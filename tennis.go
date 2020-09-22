package main

type Game struct {
	Player1      string
	Player2      string
	Player1Score int
	Player2Score int
}

func TennisGame(player1, player2 string) *Game {
	game := &Game{
		Player1:      player1,
		Player2:      player2,
		Player1Score: 0,
		Player2Score: 0,
	}

	return game
}

func (g *Game) Result() string {
	result := "Love - All"
	return result
}
