package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":8080", srv))
}
