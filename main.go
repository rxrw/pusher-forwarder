package main

import (
	"reprover/push-convert/pushers"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	// app.Post("/dingtalk", pushers.HandleDingtalk)

	app.Post("/pushover", pushers.HandlePushover)

	app.Listen(":8080")
}
