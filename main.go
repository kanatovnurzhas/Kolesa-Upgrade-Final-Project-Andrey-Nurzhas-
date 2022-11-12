package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Env      string
	BotToken string
}

func main() {
	configPath := flag.String("config", "", "Path the config file")
	flag.Parse()
	conf := &Config{}
	_, err := toml.DecodeFile(*configPath, conf)
	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигурации %v", err)
	}
}
