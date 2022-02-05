package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func pickSnakeId() string {
	i := rand.Intn(len(words))
	return fmt.Sprintf("snake-%v", words[i])
}

func createSnake() snakeInternal {
	snakeId := pickSnakeId()

	return makeSnake(snakeId)
}

func makeSnake(snakeId string) snakeInternal {
	return snakeInternal{
		id:    snakeId,
		stage: 1,
		bids:  make([]int, 0),
	}
}

func newSnakes(max int) (snakes []snakeInternal) {
	snakes = make([]snakeInternal, max)

	for i := 0; i < max; i++ {
		snakes[i] = createSnake()
	}

	return
}

func headerSetCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func generateCookie() string {
	b := make([]byte, 5)

	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", string(b))
}