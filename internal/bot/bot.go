package tgbot

import (
	"gobot/internal/models"
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

type UpgradeBot struct {
	Bot   *telebot.Bot
	Users *models.UserModel
}

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}

func Init(token string) *telebot.Bot {
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
