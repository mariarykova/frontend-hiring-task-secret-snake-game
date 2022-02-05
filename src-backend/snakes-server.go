package main

import (
	"math/rand"
	"time"
)

const (
	// TradeInterval between sending out a new snake update
	TradeInterval = 5 * time.Second

	// SnakeMaxSize of snakes that can be in existence at once
	SnakeMaxSize = 10

	// SnakeMaxBid that can be made
	SnakeMaxBid = 1_000_000
)

type (
	SnakeRequest struct {
		// Updates may be set to specify a channel to receive updates down.
		// Should be paired with CookieRequest to receive the cookie to
		// deregister when connected.
		Updates chan<- SnakeUpdate

		// CookieRequest is a field to be set to request that a cookie is to
		// sent down to a connected user to facilitate them disconnecting
		// with a message.
		CookieRequest chan<- string

		// CookieClose can be set with the field that should disconnect a
		// user from the server using their cookie.
		CookieClose string
	}

	SnakeUpdate struct {
		Id    string `json:"id"`
		Stage int    `json:"stage"`
		Bid   int    `json:"bid"`
	}

	snakeInternal struct {
		id    string
		stage int
		bids  []int
	}

	SnakeServer struct {
		connected     map[string]chan<- SnakeUpdate
		snakeUpdates  chan SnakeUpdate
		snakeRequests chan chan []snakeInternal
		messages      chan SnakeRequest
	}
)

func NewSnakeServer() SnakeServer {
	return SnakeServer{
		connected:    make(map[string]chan<- SnakeUpdate),
		snakeUpdates: make(chan SnakeUpdate),
		snakeRequests: make(chan chan []snakeInternal, 0),
		messages:     make(chan SnakeRequest),
	}
}

func NewSnakeUpdate(snakeId string, stage, bid int) SnakeUpdate {
	return SnakeUpdate{
		Id:    snakeId,
		Stage: stage,
		Bid:   bid,
	}
}

func (server SnakeServer) Messages() chan<- SnakeRequest {
	return server.messages
}

func (server SnakeServer) Bids(snakeId string) (bids []int) {
	snakesChan := make(chan []snakeInternal, 0)

	server.snakeRequests <- snakesChan

	snakes := <-snakesChan

	bids = make([]int, 0)

	for _, snake := range snakes {
		if snake.id == snakeId {
			bids = snake.bids
			break
		}
	}

	return bids
}

func (server SnakeServer) Snakes() (snakeUpdates []SnakeUpdate) {
	snakesChan := make(chan []snakeInternal, 0)

	server.snakeRequests <- snakesChan

	snakes := <-snakesChan

	snakeUpdates = make([]SnakeUpdate, len(snakes))

	for i, snake := range snakes {
		snakeUpdates[i] = SnakeUpdate{
			Id:    snake.id,
			Stage: snake.stage,
		}
	}

	return snakeUpdates
}

func (server SnakeServer) Run() {

	var (
		connected     = server.connected
		messages      = server.messages
		snakeRequests = server.snakeRequests
		snakeUpdates  = server.snakeUpdates
	)

	// snakeRequests should be handled outside of this goroutine so as to
	// not deadlock

	go server.createSnakesServer(snakeRequests, TradeInterval)

	for {
		select {
		case request := <-messages:

			var (
				updates       = request.Updates
				cookieRequest = request.CookieRequest
				cookieClose   = request.CookieClose
			)

			shouldSubscribeUser := updates != nil && cookieRequest != nil

			if shouldSubscribeUser {
				cookie := generateCookie()
				cookieRequest <- cookie
				connected[cookie] = updates
			}

			shouldDisconnectUser := cookieClose != ""

			if shouldDisconnectUser {
				delete(connected, cookieClose)
			}

		case snakeUpdate := <-snakeUpdates:

			for _, connected := range connected {

				// if the channel cannot be read at the time, we skip it

				select {
				case connected <- snakeUpdate:
				default:
				}
			}
		}
	}
}

func (server SnakeServer) createSnakesServer(snakeRequests <-chan chan []snakeInternal, frequency time.Duration) {

	updates := server.snakeUpdates

	ticker := time.Tick(frequency)

	snakes := newSnakes(SnakeMaxSize)

	for {
		select {

		case snakeRequest := <-snakeRequests:
			snakeRequest <- snakes

		case _ = <-ticker:
			snakeNumber := rand.Intn(len(snakes))

			snake := snakes[snakeNumber]

			var (
				snakeId    = snake.id
				snakeStage = snake.stage
				snakeBids  = snake.bids
			)

			// if shouldUpdateState is true, then we increment the underlying state of
			// the snake, if it's already 3, the cycle begins anew and a new snake is
			// chosen to replace the stage 3 snake

			shouldUpdateSnake := rand.Intn(100) < 20

			switch true {

			case snakeStage == 1 && shouldUpdateSnake:
				fallthrough

			case snakeStage == 2 && shouldUpdateSnake:

				newStage := snakeStage + 1

				snakes[snakeNumber].stage = newStage

				updates <- SnakeUpdate{
					Id:    snakeId,
					Stage: newStage,
					Bid:   0,
				}

			case snakeStage == 3 && shouldUpdateSnake:

				// the snake is at stage 3 and we're updating states, this should be
				// replaced! we're going to update that this happened

				newSnake := createSnake()

				snakes[snakeNumber] = newSnake

				updates <- SnakeUpdate{
					Id:    newSnake.id,
					Stage: newSnake.stage,
					Bid:   0,
				}

			case snakeStage == 1:

				// the snake isn't going to have its stage updated, so we simply generate
				// a new bid.

				bid := rand.Intn(SnakeMaxBid)

				newBids := append(snakeBids, bid)

				snakes[snakeNumber].bids = newBids

				updates <- SnakeUpdate{
					Id:    snakeId,
					Stage: snakeStage,
					Bid:   bid,
				}

			case snakeStage == 2:

				snakeBidsLen := len(snakeBids)

				if snakeBidsLen == 0 {
					break
				}

				bidNumber := rand.Intn(snakeBidsLen)

				snakes[snakeNumber].bids[bidNumber] = 0

			case snakeStage == 3:

				// this snake is in its final stage and is ignored!

			}
		}
	}
}
