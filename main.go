package main

import (
	"gobot/cmd/bot"
	"gobot/cmd/server"
)

func main() {
	server.Server()
	tgbot.RunBot()

}
