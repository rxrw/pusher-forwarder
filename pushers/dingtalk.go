package pushers

import (
	"reprover/push-convert/parsers"

	"github.com/kataras/iris/v12"
)

const DINGTALK_URL = "https://oapi.dingtalk.com/robot/send?access_token=%s"

const PUSHOVER_URL = "https://api.pushover.net/1/messages.json"

func HandlePushover(ctx iris.Context) {
	userToken := ctx.URLParam("user")
	appToken := ctx.URLParam("token")
	body := make(map[string]interface{})
	ctx.ReadJSON(&body)
	standard := parsers.ExplainToStandard(body).ToPushover()
	standard.SetToken(userToken, appToken)
	Send(PUSHOVER_URL, standard)
}
