package router

import (
	"debugpedia-api/controller"
	"net/http"
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
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		// cookieの送受信を可能にする。
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// postmanで動作確認するにはセキュア属性をfalseにしないといけないが、smaeSiteNoneModeだと勝手にtrueになる
		// CookieSameSite: http.SameSiteDefaultMode,
		CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf",uc.CsrfToken)
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
