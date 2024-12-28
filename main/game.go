package main

func startGame() {
	var players []Player

	for i := 0; i < 4; i++ {
		players[i] = *new(Player)
	}

	players[0].initRich()
	players[1].initPoor()
	players[2].initPoor()
	players[3].initPoor()

	initMap()

}
