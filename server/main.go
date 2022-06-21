package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	database, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(database)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	srv := NewPlayerServer(store)
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
