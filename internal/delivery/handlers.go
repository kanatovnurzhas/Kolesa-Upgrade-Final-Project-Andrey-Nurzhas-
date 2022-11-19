package delivery

import (
	"encoding/json"
	"fmt"
	"gobot/internal/models"
	"io/ioutil"
	"log"
	"net/http"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("не метод POST")
		return
	}

	msg := models.Message{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Ошибка при чтении тела запроса")
	}

	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Printf("Ошибка при декодировании json %v", err)
	}

	fmt.Println("Message:" + msg.Message)

	//x := r.Form.Get("text")
	//fmt.Println(x)
}
