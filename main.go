package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"gobot/cmd/bot"
	"gobot/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Config struct {
	Env      string
	BotToken string
	Dsn      string
}
type Recipient interface {
	// Must return legit Telegram chat_id or username
	Recipient() string
}
type martian struct{}

func (m martian) Recipient() string {
	return "nack nack"
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
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

	upgradeBot := bot_init.UpgradeBot{
		Bot:   bot_init.InitBot(cfg.BotToken),
		Users: &models.UserModel{Db: db},
	}

	upgradeBot.Bot.Handle("/start", upgradeBot.StartHandler)
	upgradeBot.Bot.Start()
}
func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	x := r.Form.Get("text")
	fmt.Fprint(w, x)
	//Я сделал вывод в консоль пост запроса с нашим сообщением, осталось вывести её в переменную и разослать всем с помощью метода telebot.bot.Send()
}
