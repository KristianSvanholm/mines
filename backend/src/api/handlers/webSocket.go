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
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var adjectives = readFromFile("./src/api/handlers/adjectives.txt", 1532)

//var nouns = readFromFile("./src/api/handlers/nouns.txt", 1323)

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

	//Find random words from slices
	rand.Seed(time.Now().UnixNano())
	ranAdj := rand.Intn(1333)
	//ranNoun := rand.Intn(1524)
	adj := adjectives[ranAdj]
	//noun := nouns[ranNoun]

	player := structs.Player{
		Ws: wsConnection,
		//Name: fmt.Sprintf(("%s%s"), strings.TrimSpace(strings.Title(adj)), strings.Title(noun)),
		Name: fmt.Sprintf(("%sBobby"), strings.TrimSpace(strings.Title(adj))),
	}

	mines.Players = append(mines.Players, &player)
	mines.SocketListener(&player)
}

//Read file
func readFromFile(filename string, lines int) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	readWord, _ := ioutil.ReadFile(filename)
	allWord := string(readWord)
	listWord := strings.Split(allWord, "\n")

	return listWord
}
