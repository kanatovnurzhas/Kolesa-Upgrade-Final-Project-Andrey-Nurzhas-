package handlers

import (
	"fmt"
	bot_init "gobot/cmd/bot"
	Config "gobot/config"
	Recipient "gobot/internal/models"
	"log"
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
		err := h.Bot.Users.Db.ScanRows(result, &existUser)
		if err != nil {
			fmt.Print(err)
		}
		texttt := &Recipient.Recipient{
			User: strconv.FormatInt(existUser.TelegramId, 10),
		}
		texttt.Recipient()
		log.Fatal(h.Bot.Bot.Send(texttt, x))
	}

}
