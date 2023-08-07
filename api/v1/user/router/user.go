package user_router

import (
	user_controller "github.com/PARKNAMSU/pay-seperator/api/v1/user/controller"
	userValidation_middleware "github.com/PARKNAMSU/pay-seperator/middlewares/userValidation"
	router_module "github.com/PARKNAMSU/pay-seperator/modules/router"
	"github.com/fasthttp/router"
)

func SetUserRouterV1(group *router.Group) {
	userGroup := group.Group("/user")

	// userGroup.GET("/test", middlewares.ExecMiddlewares(user_controller.TestController, userValidation_middleware.UserValidation))
	userGroup.GET("/test", router_module.PaySeperatorRouter{
		Middlewares: []router_module.PaySeperatorMiddleware{
			userValidation_middleware.UserValidation,
		},
		Controller: user_controller.TestController,
	}.Handler())
}
