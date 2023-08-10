package user_controller

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func TestController(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=utf8")
	// params := ctx.UserValue("commonParams").(common_model.ApiCommonParams)
	// fmt.Printf("[ip]:[%s], [os]:[%s], [platform]:[%s], [device]:[%s]", params.ClientIp, params.ClientOS, params.ClientPlatform, params.ClientDevice)
	body, _ := json.Marshal(map[string]string{
		"message": "user test",
	})
	ctx.Response.SetBody(body)
}

func SignUpController(ctx *fasthttp.RequestCtx) {

}
