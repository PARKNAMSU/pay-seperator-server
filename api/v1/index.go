package api_v1_router

import (
	user_router "github.com/PARKNAMSU/pay-seperator/api/v1/user/router"
	userValidation_middleware "github.com/PARKNAMSU/pay-seperator/middlewares/userValidation"
	"github.com/fasthttp/router"
)

func SetAPIRouterV1(rootRouter *router.Router) {

	v1Group := rootRouter.Group("/api/v1")
	v1Group.ANY("/", userValidation_middleware.UserValidationMiddleware)
	user_router.SetUserRouterV1(v1Group)
}
