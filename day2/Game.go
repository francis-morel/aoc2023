package main

type Game struct {
	number int
	red    int
	green  int
	blue   int
}

func (game Game) isValid() bool {
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	return game.red <= redLimit && game.green <= greenLimit && game.blue <= blueLimit
}
