package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleSnakes(snakeServer SnakeServer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		headerSetCors(w)

		snakes := snakeServer.Snakes()

		if err := json.NewEncoder(w).Encode(snakes); err != nil {
			log.Printf(
				"Failed to send snakes to the user! %v",
				err,
			)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}
