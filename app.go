package main

import (
	"flag"
	"fmt"

	api_v1_router "github.com/PARKNAMSU/pay-seperator/api/v1"
	root_router "github.com/PARKNAMSU/pay-seperator/root"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("addr", ":80", "Server Addr for Listen")
)

func main() {
	appRouter := router.New()

	api_v1_router.SetAPIRouterV1(appRouter)
	root_router.SetRootRouterGroup(appRouter)

	if err := fasthttp.ListenAndServe(*addr, appRouter.Handler); err != nil {
		fmt.Printf("server not listened [cause by]:[%s]", err.Error())
	}
}

func serverListenHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	ctx.SetContentType("application/json; charset=utf8")

	// Set arbitrary headers
	// ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	// var c fasthttp.Cookie
	// c.SetKey("cookie-name")
	// c.SetValue("cookie-value")
	// ctx.Response.Header.SetCookie(&c)
}
