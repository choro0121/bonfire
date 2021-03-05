package handler

import (
    "os"
    "fmt"
    "time"
    "strings"
    "strconv"

    // oauth
    "github.com/markbates/goth"
    "github.com/markbates/goth/gothic"
    "github.com/markbates/goth/providers/google"
    "github.com/markbates/goth/providers/github"

    // server
    "net/http"
    "github.com/labstack/echo"

    // jwt
    "github.com/dgrijalva/jwt-go"

    // salt
    "io"
    "reflect"
    "crypto/rand"
    "encoding/base64"
    "golang.org/x/crypto/scrypt"
    "github.com/google/uuid"

    "bonfire/model"
)

func init() {
    url := os.Getenv("HOST_URL")

    goth.UseProviders(
        google.New(os.Getenv("OAUTH_GOOGLE_KEY"), os.Getenv("OAUTH_GOOGLE_SECRET"), url + "/auth/google/callback"),
        github.New(os.Getenv("OAUTH_GITHUB_KEY"), os.Getenv("OAUTH_GITHUB_SECRET"), url + "/auth/github/callback"),
    )
}

func makeUuid() string {
    return strings.Replace(uuid.New().String(), "-", "", -1)
}

const (
    PASSWORD_SALT_BYTES = 16
    PASSWORD_HASH_BYTES = 16
)
func encodePassword(password string) (string, error) {
    // random bytes
    salt := make([]byte, PASSWORD_SALT_BYTES)
    _, err := io.ReadFull(rand.Reader, salt)
    if err != nil {
        return "", err
    }

    // hash with salt
    hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PASSWORD_HASH_BYTES)
    if err != nil {
        return "", err
    }

    // concat bytes
    decode := append(salt, hash...)
    encode := base64.StdEncoding.EncodeToString(decode)

    return encode, nil
}

func comparePassword(password string, encode string) error {
    // decode string
    decode, err := base64.StdEncoding.DecodeString(encode)
    if err != nil {
        return err
    }

    // split salt
    salt := decode[:PASSWORD_SALT_BYTES]

    // hash with salt
    hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PASSWORD_HASH_BYTES)
    if err != nil {
        return err
    }

    // compare hash
    if reflect.DeepEqual(decode[PASSWORD_SALT_BYTES:PASSWORD_SALT_BYTES+PASSWORD_HASH_BYTES], hash) {
        return nil
    } else {
        return fmt.Errorf("hashed password not matched")
    }
}

func generateJwtToken(userId int) (string, error) {
    now := time.Now()
    
    claims := new(jwt.StandardClaims)
    claims.Subject   = strconv.Itoa(userId)
    claims.IssuedAt  = now.Unix()
    claims.ExpiresAt = now.Add(24 * time.Hour).Unix()

    // create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func parseJwtToken(tokenString string) (int, error) {
    // super user test mode
    if tokenString == "super" {
        return 1, nil
    }

    // trim 'Bearer'
    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    // parse token
    token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return "", fmt.Errorf("unexpected  signing method: %v", token.Header["alg"])
        }
        return os.Getenv("JWT_SECRET_KEY"), nil
    })

    // error handling
    if err != nil {
        return 0, err
    }
    if token == nil {
        return 0, fmt.Errorf("token is nil")
    }

    // get claims
    claims, ok := token.Claims.(*jwt.StandardClaims)
    if !ok {
        return 0, fmt.Errorf("not found claims in %s", tokenString)
    }

    // validate
    err = claims.Valid()
    if err != nil {
        return 0, err
    }

    // get user
    userId, err := strconv.Atoi(claims.Subject)
    if err != nil {
        return 0, fmt.Errorf("jwt subject not integer %s", claims.Subject)
    }

    return userId, err
}

func authorizedRedirect(c echo.Context, user *model.User) error {
    // generate jwt token by user_id
    tokenString, err := generateJwtToken(user.UserId)
    if err != nil {
        return err
    }

    // redirect with token
    return c.Redirect(http.StatusFound, "/authorized?token=" + tokenString)
}

func signupEmail(c echo.Context) error {
    // get form value
    formMail := c.FormValue("mail")
    
    // check user
    _, err := model.GetUser(&model.User{Mail: formMail})
    if err == nil {
        return fmt.Errorf("already exist %s", formMail)
    }

    // check temp user
    tempUser, err := model.GetTempEmailUser(&model.TempEmailUser{Mail: formMail})
    if err == nil {
        model.DeleteTempEmailUser(tempUser)
    }

    // create temp user
    tempUser, err = model.CreateTempEmailUser(&model.TempEmailUser{
        TempEmailUserId: makeUuid(),
        Mail: formMail,
    })
    if err != nil {
        return err
    }

    return c.String(http.StatusOK, tempUser.TempEmailUserId)
}

func registerEmailUser(c echo.Context) error {
    // get temp user id
    tempUserId := c.Param("temp_id")

    // get form value
    formUsername := c.FormValue("username")
    formPassword := c.FormValue("password")

    // get temp user
    tempUser, err := model.GetTempEmailUser(&model.TempEmailUser{TempEmailUserId: tempUserId})
    if err != nil {
        return err
    }

    // hash password
    hash, err := encodePassword(formPassword)
    if err != nil {
        return fmt.Errorf("hash password failed")
    }

    // create user
    user := &model.User{
        Username: formUsername,
        Mail:     tempUser.Mail,
    }
    user, _, err = model.CreateUserWithPassword(user, &model.Password{
        Hash: hash,
    })
    if err != nil {
        return fmt.Errorf("create user failed")
    }

    // remove temp user
    _, err = model.DeleteTempEmailUser(tempUser)
    if err != nil {
        return err
    }

    // redirect with token
    return authorizedRedirect(c, user)
}

func signinEmail(c echo.Context) error {
    // get form value
    formUsername := c.FormValue("username")
    formPassword := c.FormValue("password")

    // get user from database
    user, err := model.GetUser(&model.User{Username: formUsername})
    if err != nil {
        return fmt.Errorf("not exist username %s", formUsername)
    }

    // confirm password
    password, err := model.GetPassword(&model.Password{UserId: user.UserId})
    if err != nil {
        return fmt.Errorf("not exist password %d", user.UserId)
    }

    // check match
    err = comparePassword(formPassword, password.Hash)
    if err != nil {
        return fmt.Errorf("confirm password failed")
    }

    // redirect with token
    return authorizedRedirect(c, user)
}

func authProvider(c echo.Context) error {
    provider := c.Param("provider")
    req := gothic.GetContextWithProvider(c.Request(), provider)
    res := c.Response().Writer

    // begin
    gothic.BeginAuthHandler(res, req)

    return nil
}

func authProviderCallback(c echo.Context) error {
    provider := c.Param("provider")
    req := c.Request()
    res := c.Response().Writer

    // get user
    providerAuth, err := gothic.CompleteUserAuth(res, req)
    if err != nil {
        return err
    }
    externId := providerAuth.UserID

    // check acount exist
    oauth, err := model.GetOAuth(&model.OAuth{
        Provider: provider,
        ExternId: externId,
    })
    if err != nil {
        // not exist acount, create acount

        // check temp user
        tempUser, err := model.GetTempOAuthUser(&model.TempOAuthUser{
            Provider: provider,
            ExternId: externId,
        })
        if err == nil {
            model.DeleteTempOAuthUser(tempUser)
        }

        // create temp user
        tempUser, err = model.CreateTempOAuthUser(&model.TempOAuthUser{
            TempOAuthUserId: makeUuid(),
            Provider:        provider,
            ExternId:        externId,
        })
        if err != nil {
            return err
        }

        // set default user param to query param
        url := fmt.Sprintf("/signup/%v", tempUser.TempOAuthUserId)

        return c.Redirect(http.StatusFound, url)
    } else {
        // exist user in oauth table, login
        user, err := model.GetUser(&model.User{UserId: oauth.UserId})
        if err != nil {
            return err
        }

        // redirect with token
        return authorizedRedirect(c, user)
    }
}

func registerOAuthUser(c echo.Context) error {
    // get temp user id
    tempUserId := c.Param("temp_id")

    // get form value
    formUsername := c.FormValue("username")
    formMail := c.FormValue("mail")

    // get temp user
    tempUser, err := model.GetTempOAuthUser(&model.TempOAuthUser{TempOAuthUserId: tempUserId})
    if err != nil {
        return err
    }

    // create user
    user := &model.User{
        Username: formUsername,
        Mail:     formMail,
    }
    user, _, err = model.CreateUserWithOAuth(user, &model.OAuth{
        Provider: tempUser.Provider,
        ExternId: tempUser.ExternId,
    })
    if err != nil {
        return err
    }

    // remove temp user
    _, err = model.DeleteTempOAuthUser(tempUser)
    if err != nil {
        return err
    }

    // redirect with token
    return authorizedRedirect(c, user)
}

func logoutProvider(c echo.Context) error {
    provider := c.Param("provider")
    req := gothic.GetContextWithProvider(c.Request(), provider)
    res := c.Response().Writer

    // logout
    gothic.Logout(res, req)
    
    // redirect top page
    return c.Redirect(http.StatusFound, "/")
}
