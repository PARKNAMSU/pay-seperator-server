package user_router

import (
	user_controller "github.com/PARKNAMSU/pay-seperator/api/v1/user/controller"
	"github.com/PARKNAMSU/pay-seperator/middlewares"
	"github.com/fasthttp/router"
)

func SetUserRouterV1(group *router.Group) {
	userGroup := group.Group("/user")
	commonRouter := middlewares.DefautAPIMiddlewares()
	userGroup.GET("/test", commonRouter.Handler(user_controller.TestController))
}
