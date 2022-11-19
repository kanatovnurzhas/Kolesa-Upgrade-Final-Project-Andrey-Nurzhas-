package tgbot

import (
	"fmt"
	"gobot/internal/models"
	"gopkg.in/telebot.v3"
	"log"
)

type Recipient struct {
	user string
}

func (r Recipient) Recipient() string {
	r.user = "934574456 "
	return r.user
}

func (bot *UpgradeBot) StartHandler(ctx telebot.Context) error {
	newUser := models.User{
		Name:       ctx.Sender().Username,
		TelegramId: ctx.Chat().ID,
		FirstName:  ctx.Sender().FirstName,
		LastName:   ctx.Sender().LastName,
		ChatId:     ctx.Chat().ID,
	}

	existUser, err := bot.Users.FindOne(ctx.Chat().ID)

	if err != nil {
		log.Printf("Ошибка получения пользователя %v", err)
	}

	if existUser == nil && err == nil {
		err := bot.Users.Create(newUser)

		if err != nil {
			log.Printf("Ошибка создания пользователя %v", err)
		}
	}

	return ctx.Send("Привет, " + ctx.Sender().FirstName)
}

func (bot *UpgradeBot) SendMessage() {
	msg := <-bot.Channel
	fmt.Println("Сэнд сообщение бота " + msg.Message)
	recep := &Recipient{
		user: "",
	}

	_, err := bot.Bot.Send(recep, msg.Message)
	if err != nil {
		return
	}

}
