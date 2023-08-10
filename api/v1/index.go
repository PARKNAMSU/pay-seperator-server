package api_v1_router

import (
	user_router "github.com/PARKNAMSU/pay-seperator/api/v1/user/router"
	"github.com/fasthttp/router"
)

func SetAPIRouterV1(rootRouter *router.Router) {
	v1Group := rootRouter.Group("/api/v1")
	user_router.SetUserRouterV1(v1Group)
}
