package userValidation_middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	common_constant "github.com/PARKNAMSU/pay-seperator/constant/common"
	common_model "github.com/PARKNAMSU/pay-seperator/model/common"
	"github.com/mileusna/useragent"
	"github.com/valyala/fasthttp"
)

func UserValidation(ctx *fasthttp.RequestCtx) (error, int) {
	body := map[string]interface{}{}

	var err error

	if bodyByte := ctx.Request.Body(); len(bodyByte) > 0 {
		err = json.Unmarshal(bodyByte, &body)

	}
	if err != nil {
		fmt.Printf("header parse err:%s", err.Error())
		return errors.New("invalidate body"), 400
	}

	userAgentStr := string(ctx.UserAgent())
	userAgentInfo := useragent.Parse(userAgentStr)

	var osInfor common_model.OSInformation

	if info, ok := common_constant.OSInformations[strings.ToUpper(userAgentInfo.OS)]; ok {
		osInfor = info
	} else {
		osInfor = common_constant.DefautOS
	}

	var hostname string

	if len(ctx.Request.Header.Peek("hostname")) > 0 {
		hostname = string(ctx.Request.Header.Peek("hostname"))
	} else {
		hostname = common_constant.DefaultHostname
	}

	ctx.SetUserValue("commonParams", common_model.ApiCommonParams{
		UserId:         100001,
		ClientIp:       ctx.RemoteIP().String(),
		UserAgent:      userAgentStr,
		Hostname:       hostname,
		ClientOS:       osInfor,
		ClientPlatform: userAgentInfo.Name,
		ClientDevice:   userAgentInfo.Device,
	})

	fmt.Println(ctx.UserValue("commonParams"))
	return nil, 200
}
