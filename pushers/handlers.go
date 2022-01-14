package pushers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reprover/push-convert/parsers"

	"github.com/kataras/iris/v12"
)

const DINGTALK_URL = "https://oapi.dingtalk.com/robot/send?access_token=%s"

const PUSHOVER_URL = "https://api.pushover.net/1/messages.json"

func HandlePushover(ctx iris.Context) {
	userToken := ctx.URLParam("user")
	appToken := ctx.URLParam("token")
	title := ctx.URLParamDefault("title", "通知转发")
	body := make(map[string]interface{})
	ctx.ReadJSON(&body)
	var message []byte
	message, _ = json.Marshal(body)
	fmt.Println(string(message))
	standard := parsers.ExplainToStandard(body).ToPushover()
	standard.SetToken(userToken, appToken)
	standard.SetTitle(title)
	send(PUSHOVER_URL, standard)
}

func HandleDingtalk(ctx iris.Context) {
	token := ctx.URLParam("token")
	body := make(map[string]interface{})
	ctx.ReadJSON(&body)
	standard := parsers.ExplainToStandard(body).ToDingtalk()

	send(fmt.Sprintf(DINGTALK_URL, token), standard)
}

func send(url string, body interface{}) {

	var res *bytes.Buffer

	switch body := body.(type) {
	case string:
		res = bytes.NewBufferString(body)
	default:
		res1, _ := json.Marshal(body)
		res = bytes.NewBuffer(res1)
	}

	client := http.DefaultClient
	resp, err := client.Post(url, "application/json", res)
	fmt.Println(resp, err)
}
