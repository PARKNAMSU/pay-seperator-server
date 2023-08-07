package router_module

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type PaySeperatorMiddleware = func(ctx *fasthttp.RequestCtx) (error, int)

type PaySeperatorRouter struct {
	Middlewares []PaySeperatorMiddleware
	Controller  fasthttp.RequestHandler
}

func (r PaySeperatorRouter) Handler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var resErr error
		var status int = 200
		for _, middleware := range r.Middlewares {
			err, getStatus := middleware(ctx)
			if err != nil {
				fmt.Println("here")
				resErr = err
				status = getStatus
				break
			}
		}
		if resErr != nil {
			ctx.Error(resErr.Error(), status)
		} else {
			r.Controller(ctx)
		}
	}
}
