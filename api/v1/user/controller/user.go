package user_controller

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func TestController(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=utf8")

	body, _ := json.Marshal(map[string]string{
		"message": "user test",
	})
	ctx.Response.SetBody(body)
}

func SignUp() {

}
