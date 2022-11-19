package main

import (
	"gobot/cmd/bot"
	"gobot/cmd/server"
	"gobot/internal/models"
	"sync"
)

func main() {
	msg := make(chan models.Message)
	var wg sync.WaitGroup
	wg.Add(2)
	go server.Server(msg)
	go tgbot.RunBot(msg)
	wg.Wait()
}
