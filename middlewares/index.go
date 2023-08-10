package middlewares

import (
	apivalidation_middleware "github.com/PARKNAMSU/pay-seperator/middlewares/apiValidation"
	userValidation_middleware "github.com/PARKNAMSU/pay-seperator/middlewares/userValidation"
	router_module "github.com/PARKNAMSU/pay-seperator/modules/router"
)

func DefautAPIMiddlewares() router_module.PaySeperatorRouter {
	return router_module.PaySeperatorRouter{
		Middlewares: []router_module.PaySeperatorMiddleware{
			apivalidation_middleware.ApiValidation,
			userValidation_middleware.UserValidation,
		},
	}
}
