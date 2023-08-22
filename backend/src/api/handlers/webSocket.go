package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"mines/src/api/errorHandler"
	"mines/src/mines"
	"mines/structs"
	"net/http"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	//Opens up the connection for websocket
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ConnectWS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	wsConnection, err := wsUpgrader.Upgrade(w, r, nil) // Upgrade connection
	if err != nil {
		errorHandler.Err(w, "Couldn't convert request to websocket", http.StatusInternalServerError)
		return
	}

	player := structs.Player{
		Ws:   wsConnection,
		Name: fmt.Sprintf("bobby%d", rand.Intn(100)),
	}

	mines.Players = append(mines.Players, &player)
	mines.SocketListener(&player)
}
