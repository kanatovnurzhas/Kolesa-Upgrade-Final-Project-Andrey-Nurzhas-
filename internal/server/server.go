package server

import (
	Config "gobot/config"
	bot_init "gobot/internal/bot"
	MyHandler "gobot/internal/handlers"
	"net/http"
	"time"
)

func Server(bot bot_init.UpgradeBot, config *Config.Config) *http.Server {
	handler := &MyHandler.MyHandler{
		Config: config,
		Bot:    bot,
	}
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server
}
