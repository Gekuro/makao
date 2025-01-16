package main

import (
	"log"
	"makao/round"
)

func main() {
	closeFile := setupLogger()
	defer closeFile()

	rnd := round.NewRound(3, 7)
	log.Println(string(rnd.GetRoundJSON()))
}
