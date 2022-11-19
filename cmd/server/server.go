package server

import (
	"fmt"
	"gobot/internal/delivery"
	"gobot/internal/models"
	"log"
	"net/http"
)

func Server(ch chan models.Message) {
	server := delivery.New(ch)
	fmt.Print("Starting server at port 8080...\nhttp://localhost:8080/\n")

	if err := http.ListenAndServe(":8080", server.Route()); err != nil {
		log.Println(err)
		return
	}
}
