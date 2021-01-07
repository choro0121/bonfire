package handler

import (
    "os"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

var e *echo.Echo

func New() {
    e = echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.File("/", "view/index.html")

    // auth
    e.POST("/signup", signup)
    e.File("/signup", "view/signup.html")
    e.POST("/signup/:provider", signupProvider)
    e.File("/signup/:provider", "view/signupProvider.html")
    e.POST("/login", login)
    e.File("/login", "view/login.html")
    e.GET("/auth/:provider", authProvider)
    e.GET("/auth/:provider/callback", authProviderCallback)
    e.GET("/logout/:provider", logoutProvider)

    // api
    api := e.Group("/api")
    api.GET("/:version", getApiList)
    for ver, list := range apiList {
        for _, a := range list {
            api.Add(a.Method, ver + a.Path, a.Handler, jwtMiddleware)
        }
    }

    e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func Close() error {
    return e.Close()
}
