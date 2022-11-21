package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	Config "gobot/config"
	bot "gobot/internal/bot"
	"gobot/internal/models"
	"gobot/internal/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()
	cfg := &Config.Config{}
	_, err := toml.DecodeFile(*configPath, cfg)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}
	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}
	upgradeBot := bot.UpgradeBot{
		Bot:   bot.Init(cfg.BotToken),
		Users: &models.UserModel{Db: db},
	}
	srv := server.Server(upgradeBot, cfg)
	go srv.ListenAndServe()
	upgradeBot.Bot.Handle("/start", upgradeBot.StartHandler)
	upgradeBot.Bot.Start()
}
