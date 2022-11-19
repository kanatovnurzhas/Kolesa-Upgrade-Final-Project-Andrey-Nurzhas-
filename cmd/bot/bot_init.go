package tgbot

import (
	"flag"
	"github.com/BurntSushi/toml"
	"gobot/internal/models"
	"gopkg.in/telebot.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type UpgradeBot struct {
	Bot     *telebot.Bot
	Users   *models.UserModel
	Channel chan models.Message
}

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}

func InitBot(token string) *telebot.Bot {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatalf("Ошибка при инициализации бота %v", err)
	}

	return b
}

func RunBot(ch chan models.Message) {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}

	upgradeBot := UpgradeBot{
		Bot:     InitBot(cfg.BotToken),
		Users:   &models.UserModel{Db: db},
		Channel: ch,
	}

	upgradeBot.Bot.Handle("/start", upgradeBot.StartHandler)
	//upgradeBot.Bot.Handle("/", upgradeBot.SendMessage)
	upgradeBot.SendMessage()
	//for {
	//	upgradeBot.Bot.Send("")
	//}

	upgradeBot.Bot.Start()
}
