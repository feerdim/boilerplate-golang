package application

import (
	"github.com/feerdim/boilerplate-golang/src/domain/auth/feature/auth/service"
	"github.com/feerdim/boilerplate-golang/src/middleware"
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo, t *toolkit.Toolkit) {
	svc := service.NewService(t.GetDB(), t.GetMail())
	mdw := middleware.NewAuthMiddleware(t.GetDB())

	e.POST("/login", loginApp(svc))
	e.POST("/refresh-token", refreshTokenApp(svc))
	e.POST("/register", registerApp(svc))
	e.POST("/verify", verifyUserApp(svc))

	validateTokenRoutes := e.Group("/", mdw.ValidateToken)
	validateTokenRoutes.POST("logout", logoutApp(svc))
	validateTokenRoutes.GET("profile", readProfileApp(svc))
	validateTokenRoutes.PUT("profile", updateProfileApp(svc))
	validateTokenRoutes.GET("verify", sendUserVerificationApp(svc))

	forgotPasswordRoutes := e.Group("/forgot-password")
	forgotPasswordRoutes.POST("", sendForgotPasswordLinkApp(svc))
	forgotPasswordRoutes.GET("/reset", validateForgotPasswordTokenApp(svc))
	forgotPasswordRoutes.POST("/reset", resetPasswordApp(svc))

	oauthRoutes := e.Group("/oauth/:provider")
	oauthRoutes.GET("", loginSSOApp(svc))
	oauthRoutes.POST("/redirect", loginSSORedirectApp(svc))
}
