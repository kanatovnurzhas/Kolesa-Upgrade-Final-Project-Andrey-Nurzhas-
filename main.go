package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	bot_init "gobot/cmd/bot"
	"gobot/cmd/server"
	Config "gobot/config"
	"gobot/internal/models"
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

	upgradeBot := bot_init.UpgradeBot{
		Bot:   bot_init.InitBot(cfg.BotToken),
		Users: &models.UserModel{Db: db},
	}
	srv := server.Server(upgradeBot, cfg)
	go log.Fatal(srv.ListenAndServe())
	upgradeBot.Bot.Handle("/start", upgradeBot.StartHandler)
	upgradeBot.Bot.Start()

}
