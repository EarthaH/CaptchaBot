package main

import (
	"text-captcha/bot/pkg/api"
	"text-captcha/bot/pkg/logger"
)

func main() {
	logger.Init()
	api.HandleRequest()
}
