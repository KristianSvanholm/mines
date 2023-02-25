package main

import (
	"log"
	"mines/src/api/handlers"
	"mines/src/mines"
	"net/http"
	"os"
)

func main() {

	mines.InitField(20)
	//mines.PlantMines(mines.Field)
	//mines.CalculateCells(mines.Field)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Handle frontend
	fs := http.FileServer(http.Dir("./../frontend/dist"))
	http.Handle("/", fs)

	// API handlers
	http.HandleFunc("/api/getField", handlers.GetField) // Lets a user login with an account
	http.HandleFunc("/api/join", handlers.ConnectWS)    // Creates WS connection

	// Start server
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
