package handler

import (
    "os"
    "fmt"
    "time"
    "strings"

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

    "bonfire/model"
)

func init() {
    // url := os.Getenv("HOST_URL")

    goth.UseProviders(
        google.New(os.Getenv("OAUTH_GOOGLE_KEY"), os.Getenv("OAUTH_GOOGLE_SECRET"), "http://192.168.197.130:8000/auth/google/callback"),
        github.New(os.Getenv("OAUTH_GITHUB_KEY"), os.Getenv("OAUTH_GITHUB_SECRET"), "http://192.168.197.130:8000/auth/github/callback"),
    )
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

func generateJwtToken(userId string) (string, error) {
    now := time.Now()
    
    claims := new(jwt.StandardClaims)
    claims.Subject   = userId
    claims.IssuedAt  = now.Unix()
    claims.ExpiresAt = now.Add(24 * time.Hour).Unix()

    // create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func parseJwtToken(tokenString string) (string, error) {
    // super user test mode
    if tokenString == "super" {
        return "", nil
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
        return "", err
    }
    if token == nil {
        return "", fmt.Errorf("token is nil")
    }

    // get claims
    claims, ok := token.Claims.(*jwt.StandardClaims)
    if !ok {
        return "", fmt.Errorf("not found claims in %s", tokenString)
    }

    // validate
    err = claims.Valid()
    if err != nil {
        return "", err
    }

    return claims.Subject, err
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
    formUsername := c.FormValue("username")
    formPassword := c.FormValue("password")

    // check user
    _, err := model.GetUser(&model.User{Mail: formMail})
    if err == nil {
        return fmt.Errorf("already exist mail %s", formMail)
    }
    _, err = model.GetUser(&model.User{Username: formUsername})
    if err == nil {
        return fmt.Errorf("already exist user %s", formUsername)
    }

    // hash password
    hash, err := encodePassword(formPassword)
    if err != nil {
        return fmt.Errorf("hash password failed")
    }

    // create user
    user := &model.User{
        Username: formUsername,
        Mail:     formMail,
    }
    user, err = model.CreateUser(user)
    if err != nil {
        return fmt.Errorf("create user failed")
    }

    // register password
    _, err = model.CreatePassword(&model.Password{
        UserId: user.UserId,
        Hash: hash,
    })
    if err != nil {
        return fmt.Errorf("register password failed")
    }

    // generate jwt token by user_id
    tokenString, err := generateJwtToken(user.UserId)
    if err != nil {
        return err
    }

    // return jwt
    return c.String(http.StatusOK, tokenString)
}

func loginEmail(c echo.Context) error {
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
        return fmt.Errorf("not exist password %s", user.UserId)
    }

    // check match
    err = comparePassword(formPassword, password.Hash)
    if err != nil {
        return fmt.Errorf("confirm password failed")
    }

    // generate jwt token by user_id
    tokenString, err := generateJwtToken(user.UserId)
    if err != nil {
        return err
    }

    // return jwt
    return c.String(http.StatusOK, tokenString)
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

    // check user acount
    if err != nil {
        // create oauth
        oauth, err = model.CreateOAuth(&model.OAuth{
            Provider: provider,
            ExternId: externId,
        })

        if err != nil {
            return fmt.Errorf("register oauth failed")
        }
    } else {
        user, err := model.GetUser(&model.User{
            UserId: oauth.UserId,
        })

        // login when exist user
        if err == nil {
            return authorizedRedirect(c, user)
        }
    }

    // redirect register
    return c.Redirect(http.StatusFound, fmt.Sprintf("/signup/%v/%v", provider, oauth.UserId))
}

func registerOAuthUser(c echo.Context) error {
    // get temp user id
    userId := c.Param("user_id")

    // get form value
    formUsername := c.FormValue("username")
    formMail := c.FormValue("mail")

    // check acount exist
    oauth, err := model.GetOAuth(&model.OAuth{
        UserId: userId,
    })
    if err != nil {
        return fmt.Errorf("not exist oauth %s", userId)
    }

    // create user
    user := &model.User{
        UserId:   oauth.UserId,
        Username: formUsername,
        Mail:     formMail,
    }
    user, err = model.CreateUser(user)
    if err != nil {
        return fmt.Errorf("create user failed")
    }

    // generate jwt token by user_id
    tokenString, err := generateJwtToken(user.UserId)
    if err != nil {
        return err
    }

    // return jwt
    return c.String(http.StatusOK, tokenString)
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
