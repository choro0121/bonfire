package handler

import (
    "log"
    "time"
    "strings"
    "strconv"

    // server
    "net/http"
    "github.com/labstack/echo"

    "bonfire/model"
)

type (
    Api struct {
        Path        string              `json:"path"`
        Method      string              `json:"method"`
        Auth        bool                `json:"auth"`
        Handler     echo.HandlerFunc    `json:"-"`
    }
)

var apiList = map[string][]Api{
    "/v1": []Api{
        Api{
            Path: "/posts",
            Method: http.MethodGet,
            Auth: false,
            Handler: defaultHandler,
        },
        Api{
            Path: "/posts",
            Method: http.MethodPost,
            Auth: true,
            Handler: newPost,
        },
        Api{
            Path: "/posts/:post_id",
            Method: http.MethodGet,
            Auth: false,
            Handler: getPost,
        },
        Api{
            Path: "/posts/:post_id",
            Method: http.MethodPut,
            Auth: true,
            Handler: updatePost,
        },
        Api{
            Path: "/posts/:post_id",
            Method: http.MethodDelete,
            Auth: true,
            Handler: deletePost,
        },
        Api{
            Path: "/bookmark",
            Method: http.MethodGet,
            Auth: false,
            Handler: getBookmarks,
        },
        Api{
            Path: "/bookmark/:post_id",
            Method: http.MethodPost,
            Auth: true,
            Handler: newBookmark,
        },
        Api{
            Path: "/bookmark/:post_id",
            Method: http.MethodDelete,
            Auth: true,
            Handler: deleteBookmark,
        },
        Api{
            Path: "/good",
            Method: http.MethodGet,
            Auth: false,
            Handler: getGoods,
        },
        Api{
            Path: "/good/:post_id",
            Method: http.MethodPost,
            Auth: true,
            Handler: newGood,
        },
        Api{
            Path: "/good/:post_id",
            Method: http.MethodDelete,
            Auth: true,
            Handler: deleteGood,
        },
        Api{
            Path: "/comment",
            Method: http.MethodGet,
            Auth: false,
            Handler: getComments,
        },
        Api{
            Path: "/comment/:post_id",
            Method: http.MethodPost,
            Auth: true,
            Handler: newComment,
        },
        Api{
            Path: "/comment/:post_id/:comment_id",
            Method: http.MethodPut,
            Auth: true,
            Handler: updateComment,
        },
        Api{
            Path: "/comment/:post_id/:comment_id",
            Method: http.MethodDelete,
            Auth: true,
            Handler: deleteComment,
        },
        Api{
            Path: "/users",
            Method: http.MethodGet,
            Auth: false,
            Handler: getUsers,
        },
        Api{
            Path: "/users/:username",
            Method: http.MethodGet,
            Auth: false,
            Handler: getUser,
        },
        Api{
            Path: "/users/:username/follows",
            Method: http.MethodGet,
            Auth: false,
            Handler: getFollows,
        },
        Api{
            Path: "/users/:username/followers",
            Method: http.MethodGet,
            Auth: false,
            Handler: getFollowers,
        },
        Api{
            Path: "/user",
            Method: http.MethodGet,
            Auth: true,
            Handler: getSelf,
        },
        Api{
            Path: "/user",
            Method: http.MethodPut,
            Auth: true,
            Handler: defaultHandler,
        },
        Api{
            Path: "/user",
            Method: http.MethodDelete,
            Auth: true,
            Handler: defaultHandler,
        },
        Api{
            Path: "/user/follow/:username",
            Method: http.MethodPost,
            Auth: true,
            Handler: newFollow,
        },
        Api{
            Path: "/user/follow/:username",
            Method: http.MethodDelete,
            Auth: true,
            Handler: deleteFollow,
        },
        Api{
            Path: "/user/notifications",
            Method: http.MethodGet,
            Auth: true,
            Handler: defaultHandler,
        },
        Api{
            Path: "/user/notification_id",
            Method: http.MethodGet,
            Auth: true,
            Handler: defaultHandler,
        },
        Api{
            Path: "/user/notification_id",
            Method: http.MethodPut,
            Auth: true,
            Handler: defaultHandler,
        },
    },
}

func getApiList(c echo.Context) error {
    ver := "/" + c.Param("version")

    return c.JSON(http.StatusOK, struct {
        BaseUrl     string  `json:"baseurl"`
        ApiList     []Api   `json:"apis"`
    } {
        "/api" + ver,
        apiList[ver],
    })
}

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    
    requireAuth := func(c echo.Context) bool {
        // split url
        // "/api/v1/posts" -> [0:api, 1:v1, 2~:posts]
        pathList := strings.Split(c.Path()[1:], "/")
        ver      := "/" + pathList[1]
        path     := "/" + strings.Join(pathList[2:], "/")
        method   := c.Request().Method

        for _, a := range apiList[ver] {
            log.Print(a.Path, path)
            if a.Path == path && a.Method == method {
                return a.Auth
            }
        }
        return false
    }

    return func(c echo.Context) error {
        if requireAuth(c) {
            tokenString := c.Request().Header.Get("Authorization")

            // check jwt token
            userId, err := parseJwtToken(tokenString)

            if err != nil {
                return err
            }

            // set user_id
            c.Set("user_id", userId)
        }

        return next(c)
    }
}

func defaultHandler(c echo.Context) error {
    log.Print(c.Get("user_id").(int))
    return c.String(http.StatusOK, "defaultHandler")
}


func newPost(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // body
    post := new(model.Post)
    if err := c.Bind(post); err != nil {
        return err
    }

    post.UserId = userId

    // create
    post, err := model.CreatePost(post)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, post)
}

func getPost(c echo.Context) error {
    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }

    // read
    post, err := model.GetPost(&model.Post{
        PostId: postId,
    })
    if err != nil {
        return err
    }
    
    return c.JSON(http.StatusOK, post)
}

func updatePost(c echo.Context) error {
    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }

    // body
    post := new(model.Post)
    if err = c.Bind(post); err != nil {
        return err
    }

    post.PostId = postId

    // update
    post, err = model.UpdatePost(post)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, post)
}

func deletePost(c echo.Context) error {
    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }

    // delete
    post, err := model.DeletePost(&model.Post{
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, post)
}


func getBookmarks(c echo.Context) error {
    // query
    userId, err := strconv.Atoi(c.QueryParam("user_id"))
    postId, err := strconv.Atoi(c.QueryParam("post_id"))
    offset, err := strconv.Atoi(c.QueryParam("offset"))

    // read
    bookmarks, err := model.GetBookmarks(offset, &model.Bookmark{
        UserId: userId,
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, bookmarks)
}

func newBookmark(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }

    // create
    bookmark, err := model.CreateBookmark(&model.Bookmark{
        UserId: userId,
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, bookmark)
}

func deleteBookmark(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }

    // delete
    bookmark, err := model.DeleteBookmark(&model.Bookmark{
        UserId: userId,
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, bookmark)
}


func getGoods(c echo.Context) error {
    // query
    userId, err := strconv.Atoi(c.QueryParam("user_id"))
    postId, err := strconv.Atoi(c.QueryParam("post_id"))
    offset, err := strconv.Atoi(c.QueryParam("offset"))

    // read
    goods, err := model.GetGoods(offset, &model.Good{
        UserId: userId,
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, goods)
}

func newGood(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }

    // create
    good, err := model.CreateGood(&model.Good{
        UserId: userId,
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, good)
}

func deleteGood(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }

    // delete
    good, err := model.DeleteGood(&model.Good{
        UserId: userId,
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, good)
}


func getComments(c echo.Context) error {
    // query
    userId, err := strconv.Atoi(c.QueryParam("user_id"))
    postId, err := strconv.Atoi(c.QueryParam("post_id"))
    offset, err := strconv.Atoi(c.QueryParam("offset"))

    // read
    comments, err := model.GetComments(offset, &model.Comment{
        UserId: userId,
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, comments)
}

func newComment(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }

    // body
    comment := new(model.Comment)
    if err = c.Bind(comment); err != nil {
        return err
    }

    comment.UserId = userId
    comment.PostId = postId

    // create
    comment, err = model.CreateComment(comment)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, comment)
}

func updateComment(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }
    commentId, err := strconv.Atoi(c.Param("comment_id"))
    if err != nil {
        return err
    }

    // body
    comment := new(model.Comment)
    if err = c.Bind(comment); err != nil {
        return err
    }

    comment.CommentId = commentId
    comment.UserId    = userId
    comment.PostId    = postId

    // update
    comment, err = model.UpdateComment(comment)
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, comment)
}

func deleteComment(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    postId, err := strconv.Atoi(c.Param("post_id"))
    if err != nil {
        return err
    }
    commentId, err := strconv.Atoi(c.Param("comment_id"))
    if err != nil {
        return err
    }

    // delete
    comment, err := model.DeleteComment(&model.Comment{
        CommentId: commentId,
        UserId: userId,
        PostId: postId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, comment)
}


func getUsers(c echo.Context) error {
    log.Print(c.Get("user_id").(int))
    return c.String(http.StatusOK, "defaultHandler")
}

func getFollows(c echo.Context) error {
    // path
    username := c.Param("username")

    // query
    offset, err := strconv.Atoi(c.QueryParam("offset"))

    // username to user
    user, err := model.GetUser(&model.User{
        Username: username,
    })
    if err != nil {
        return err
    }

    // read
    follows, err := model.GetFollows(offset, &model.Follow{
        UserId: user.UserId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, follows)
}

func getFollowers(c echo.Context) error {
    // path
    username := c.Param("username")

    // query
    offset, err := strconv.Atoi(c.QueryParam("offset"))

    // username to user
    user, err := model.GetUser(&model.User{
        Username: username,
    })
    if err != nil {
        return err
    }

    // read
    followers, err := model.GetFollows(offset, &model.Follow{
        FollowUserId: user.UserId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, followers)
}


func getUser(c echo.Context) error {
    // path
    username := c.Param("username")

    // read
    user, err := model.GetUser(&model.User{
        Username: username,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, struct {
        UserId      int         `json:"user_id"`
        Username    string      `json:"username"`
        Bio         string      `json:"bio"`
        Company     string      `json:"company"`
        Location    string      `json:"location"`
        WebSite     string      `json:"website"`
        CreatedAt    time.Time   `json:"create_at"`
        UpdatedAt    time.Time   `json:"update_at"`
        // Posts       int         `json:"posts"`
        // Follows     int         `json:"follows"`
        // Followers   int         `json:"followers"`
        // Bookmarks   int         `json:"bookmarks"`
    } {
        user.UserId,
        user.Username,
        user.Bio,
        user.Company,
        user.Location,
        user.WebSite,
        user.CreatedAt,
        user.UpdatedAt,
    })
}

func newFollow(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    username := c.Param("username")

    // username to user
    user, err := model.GetUser(&model.User{
        Username: username,
    })
    if err != nil {
        return err
    }

    // create
    follow, err := model.CreateFollow(&model.Follow{
        UserId: userId,
        FollowUserId: user.UserId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, follow)
}

func deleteFollow(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // path
    username := c.Param("username")

    // username to user
    user, err := model.GetUser(&model.User{
        Username: username,
    })
    if err != nil {
        return err
    }

    // delete
    follow, err := model.DeleteFollow(&model.Follow{
        UserId: userId,
        FollowUserId: user.UserId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, follow)
}


func getSelf(c echo.Context) error {
    // token
    userId := c.Get("user_id").(int)

    // read
    user, err := model.GetUser(&model.User{
        UserId: userId,
    })
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, user)
}
