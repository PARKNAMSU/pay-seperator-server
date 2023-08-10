package apivalidation_middleware

import (
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

var (
	_ = godotenv.Load()
)

func ApiValidation(ctx *fasthttp.RequestCtx) (error, int) {
	// apiKey := ctx.Request.Header.Peek("x-api-key")
	// if os.Getenv("API_KEY") != string(apiKey) {
	// 	return errors.New("invalidation api key"), 400
	// }
	return nil, 200
}
