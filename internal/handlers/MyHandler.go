package handlers

import (
	bot_init "gobot/cmd/bot"
	Config "gobot/config"
	Recipient "gobot/internal/models"
	"net/http"
	"strconv"
)

type MyHandler struct {
	Config *Config.Config
	Bot    bot_init.UpgradeBot
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	x := r.Form.Get("text")
	result, existUser := h.Bot.Users.FindAll()

	for result.Next() {
		h.Bot.Users.Db.ScanRows(result, &existUser)
		texttt := &Recipient.Recipient{
			User: strconv.FormatInt(existUser.TelegramId, 10),
		}
		texttt.Recipient()
		h.Bot.Bot.Send(texttt, x)
	}

}
