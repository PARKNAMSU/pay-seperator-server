package root_router

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func SetRootRouterGroup(rootRouter *router.Router) {
	testGroup := rootRouter.Group("/")

	testGroup.GET("/health/check", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Server Alive!!\n\n")

		// var urlstring string = ctx.UserValue("url").(string)
		// urlparsed, _ := url.Parse(urlstring)
		// urlparsedstring := urlparsed.String()
		// fmt.Fprintf(ctx, "URL: %s\n", urlparsedstring)

		ctx.SetContentType("text/plain; charset=utf8")
		ctx.Response.SetBody([]byte("Server Alive!!"))
	})
}
