package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleUpdates(snakeRequests chan<- SnakeRequest) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		headerSetCors(w)

		conn, err := websocketUpgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Printf(
				"Failed to upgrade a connecting user to a websocket! %v",
				err,
			)

			w.WriteHeader(http.StatusBadRequest)

			return
		}

		var (
			snakeUpdates = make(chan SnakeUpdate, 1)
			snakeCookie  = make(chan string)
		)

		snakeRequests <- SnakeRequest{
			Updates:       snakeUpdates,
			CookieRequest: snakeCookie,
		}

		cookie := <-snakeCookie

		defer conn.Close()

		for update := range snakeUpdates {

			err := conn.WriteJSON(update)

			if err != nil {
				log.Printf(
					"Failed to write an orderbook update to a connected user! %v",
					err,
				)

				// disconnect this goroutine from the orderbook server...

				snakeRequests <- SnakeRequest{
					CookieClose: cookie,
				}

				close(snakeUpdates)

				return
			}
		}
	}
}
