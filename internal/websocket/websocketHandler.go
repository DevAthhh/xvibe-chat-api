package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DevAthhh/xvibe-chat/internal/repo"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	chatRepo := repo.ChatRepo{}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		messages, err := chatRepo.GetMessagesByChatID(string(msg))
		if err != nil {
			continue
		}
		jsonMsgs, err := json.Marshal(messages)
		if err != nil {
			continue
		}
		if err := conn.WriteJSON(jsonMsgs); err != nil {
			continue
		}
	}
}
