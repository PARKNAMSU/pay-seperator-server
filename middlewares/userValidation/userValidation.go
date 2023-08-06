package userValidation_middleware

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

func UserValidationMiddleware(ctx *fasthttp.RequestCtx) {
	headers := make(map[string]string)
	headerByte := ctx.Request.Header.Header()

	err := json.Unmarshal(headerByte, &headers)

	if err != nil {
		fmt.Printf("header parse err:%s", err.Error())
	}

	body := map[string]interface{}{}

	if bodyByte := ctx.Request.Body(); len(bodyByte) > 0 {
		err = json.Unmarshal(bodyByte, &body)

	}
	if err != nil {
		fmt.Printf("header parse err:%s", err.Error())
	}
	fmt.Println("user Validation finish")
}

func UserValidation(ctx *fasthttp.RequestCtx) (error, int) {
	headers := make(map[string]string)
	headerByte := ctx.Request.Header.Header()

	err := json.Unmarshal(headerByte, &headers)

	if err != nil {
		fmt.Printf("header parse err:%s", err.Error())
		return errors.New("invalidate header"), 400
	}

	body := map[string]interface{}{}

	if bodyByte := ctx.Request.Body(); len(bodyByte) > 0 {
		err = json.Unmarshal(bodyByte, &body)

	}
	if err != nil {
		fmt.Printf("header parse err:%s", err.Error())
		return errors.New("invalidate body"), 400
	}

	fmt.Printf("[headers]:[%v], [body]:[%v]", headers, body)
	fmt.Println("user Validation finish")
	return nil, 200
}
