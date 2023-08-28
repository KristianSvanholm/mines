package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mines/src/api/errorHandler"
	"mines/src/mines"
	"mines/structs"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var adjectives = readFromFile("./src/api/handlers/adjectives.txt", 1532)

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

	//Random seed
	rand.Seed(time.Now().UnixNano())
	//Find random adjective from file
	ranAdj := rand.Intn(1333)
	adj := adjectives[ranAdj]

	player := structs.Player{
		Ws:   wsConnection,
		Name: fmt.Sprintf(("%sBobby"), strings.TrimSpace(strings.Title(adj))),
	}

	mines.Players = append(mines.Players, &player)
	mines.SocketListener(&player)
}

//Read file
func readFromFile(filename string, lines int) []string {
	//Open and read file
	readWord, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	//Set file into string slice
	stringWord := string(readWord)
	listWord := strings.Split(stringWord, "\n")

	return listWord
}
