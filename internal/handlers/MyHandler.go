package handlers

import (
	"fmt"
	Config "gobot/config"
	bot_init "gobot/internal/bot"
	Recipient "gobot/internal/models"
	"net/http"
	"strconv"
)

type MyHandler struct {
	Config *Config.Config
	Bot    bot_init.UpgradeBot
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	x := r.Form.Get("text")
	result, existUser := h.Bot.Users.FindAll()
	for result.Next() {
		h.Bot.Users.Db.ScanRows(result, &existUser)
		texttt := &Recipient.Recipient{
			User: strconv.FormatInt(existUser.TelegramId, 10),
		}
		texttt.Recipient()
		_, err := h.Bot.Bot.Send(texttt, x)
		if err != nil {
			h.Bot.Bot.Send(texttt, "Ошибка при отправке сообщения "+err.Error())
		}
	}
}
