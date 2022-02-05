package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	// DefaultListen address to use when serving the local client
	DefaultListen = ":8080"

	// MaxPrice to set in a snake update
	MaxPrice = 1_000_000
)

func main() {

	rand.Seed(time.Now().Unix())

	var listenAddress string

	if len(os.Args) > 1 {
		listenAddress = os.Args[1]
	} else {
		listenAddress = DefaultListen
	}

	snakeServer := NewSnakeServer()

	go snakeServer.Run()

	snakeMessages := snakeServer.Messages()

	var (
		snakesHandler  = handleSnakes(snakeServer)
		bidsHandler    = handleBids(snakeServer)
		updatesHandler = handleUpdates(snakeMessages)
	)

	http.HandleFunc("/api/snakes", snakesHandler)

	http.HandleFunc("/api/bids", bidsHandler)

	http.HandleFunc("/api/updates", updatesHandler)

	log.Printf("Listening on %#v!", listenAddress)

	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Fatalf(
			"Failed to listen on %#v! %v",
			listenAddress,
			err,
		)
	}
}
