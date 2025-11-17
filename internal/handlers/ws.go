package handlers

import (
    "net/http"
    "github.com/gorilla/websocket"
)

var up = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func WS() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        c, err := up.Upgrade(w, r, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        defer c.Close()
        for {
            _, msg, err := c.ReadMessage()
            if err != nil {
                break
            }
            c.WriteMessage(websocket.TextMessage, msg)
        }
    }
}