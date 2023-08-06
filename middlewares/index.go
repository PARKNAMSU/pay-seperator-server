package middlewares

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func ExecMiddlewares(handler fasthttp.RequestHandler, middlewares ...func(ctx *fasthttp.RequestCtx) (error, int)) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var resErr error
		var status int = 200
		for _, middleware := range middlewares {
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
			handler(ctx)
		}
	}
}
