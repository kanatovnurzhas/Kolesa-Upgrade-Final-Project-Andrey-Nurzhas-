package main

import (
	"gobot/cmd/bot"
	"gobot/cmd/server"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go server.Server()
	go tgbot.RunBot()
	wg.Wait()
}
