package server

import (
	"fmt"
	"gobot/internal/delivery"
	"log"
	"net/http"
)

func Server() {
	server := delivery.New()
	fmt.Print("Starting server at port 8080...\nhttp://localhost:8080/\n")

	if err := http.ListenAndServe(":8080", server.Route()); err != nil {
		log.Println(err)
		return
	}
}
