package main

var scoreCodes map[int]string

type Game struct {
	Player1      string
	Player2      string
	Player1Score int
	Player2Score int
}

func init()  {
	scoreCodes = make(map[int]string)
	scoreCodes[0] = "Love"
	scoreCodes[1] = "Fifteen"
	scoreCodes[2] = "Thirty"
	scoreCodes[3] = "Forty"
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

	if g.Player1Score > 0 {
		result = scoreCodes[g.Player1Score] + " - Love"
	}

	return result
}

func (g *Game) pointPlayer1() {
	g.Player1Score++
}
