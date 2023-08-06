package user_router

import (
	user_controller "github.com/PARKNAMSU/pay-seperator/api/v1/user/controller"
	"github.com/PARKNAMSU/pay-seperator/middlewares"
	userValidation_middleware "github.com/PARKNAMSU/pay-seperator/middlewares/userValidation"
	"github.com/fasthttp/router"
)

func SetUserRouterV1(group *router.Group) {
	userGroup := group.Group("/user")

	userGroup.GET("/test", middlewares.ExecMiddlewares(user_controller.TestController, userValidation_middleware.UserValidation))
}
