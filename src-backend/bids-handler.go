package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleBids(snakeServer SnakeServer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		headerSetCors(w)

		snakeId := r.URL.Query().Get("snake-id")

		bids := snakeServer.Bids(snakeId)

		if err := json.NewEncoder(w).Encode(bids); err != nil {
			log.Printf(
				"Failed to send bids to the user! %v",
				err,
			)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}
