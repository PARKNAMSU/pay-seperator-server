package router_module

import (
	"github.com/valyala/fasthttp"
)

type PaySeperatorMiddleware = func(ctx *fasthttp.RequestCtx) (error, int)

type PaySeperatorRouter struct {
	Middlewares []PaySeperatorMiddleware
}

func (r PaySeperatorRouter) Handler(controller fasthttp.RequestHandler, appendMiddleware ...PaySeperatorMiddleware) fasthttp.RequestHandler {
	execMiddleware := append(r.Middlewares, appendMiddleware...)

	return func(ctx *fasthttp.RequestCtx) {
		var resErr error
		var status int = 200
		for _, middleware := range execMiddleware {
			err, getStatus := middleware(ctx)
			if err != nil {
				resErr = err
				status = getStatus
				break
			}
		}
		if resErr != nil {
			ctx.Error(resErr.Error(), status)
		} else {
			controller(ctx)
		}
	}
}
