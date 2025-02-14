package main

import (
	"fmt"
	"strconv"
)

func main() {

	initMap()
	var player Player
	player = player.initPoor()

	fmt.Println("Welcome to Eat The Rich!")
	fmt.Println("You start like a poor")

	var err string

	for {
		fmt.Printf("You`re at %v (%s), please choose your next move:\n", player.currentPosition, player.getCurrentSpaceState())

		for _, i := range player.getMovement() {
			fmt.Printf("%v (%v)\n", i, resolveSpaceEnum(mapSlice[i][0]))
		}

		var input string
		fmt.Scanln(&input)
		err, player = player.move(strconv.Atoi(input))
		if err != "" {
			fmt.Println("Error:", err)
			continue
		}
	}
}
