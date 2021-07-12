package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func wsHandle(w http.ResponseWriter, r *http.Request) {
	upgeader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	if conn, err = upgeader.Upgrade(w, r, nil); err != nil {
		conn.Close()
	}

	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			conn.Close()
		}

		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			conn.Close()
		}
	}

}

func main() {
	http.HandleFunc("/ws", wsHandle)
	http.ListenAndServe("0.0.0.0:8083", nil)

}
