package routes

import (
	"github.com/delaram-gholampoor-sagha/Deli-Authenticator/controllers"
	"github.com/delaram-gholampoor-sagha/Deli-Authenticator/middleware"
	"github.com/delaram-gholampoor-sagha/Deli-Authenticator/services"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(userService), rc.authController.LogoutUser)
	router.GET("/verifyemail/:verificationCode", rc.authController.VerifyEmail)
	router.POST("/forgotPassword", rc.authController.ForgotPassword)
}
