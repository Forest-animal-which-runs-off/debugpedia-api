package router

import (
	"debugpedia-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, dc controller.IDebugController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4000", os.Getenv("FE_URL")},
		// ヘッダー経由でcsrfトークンを受け取れるようにする。
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		// cookieの送受信を可能にする。
		AllowCredentials: true,
	}))


	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	d := e.Group("/debugs")
	// echojwtのミドルウェアを使用。
	d.Use(echojwt.WithConfig(echojwt.Config{
		// jwtトークンを作る時と同じsecret
		SigningKey: []byte(os.Getenv("SECRET")),
		// クライアントから送られてくるjwtトークンがどこに格納されているかを指定
		TokenLookup: "cookie:token",
	}))
	d.GET("", dc.GetAllDebugs)
	d.GET("/:debugId", dc.GetDebugById)
	d.POST("", dc.CreateDebug)
	d.PUT("/:debugId", dc.UpdateDebug)
	d.DELETE("/:debugId", dc.DeleteDebug)

	return e
}
