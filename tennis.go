package main

import "math"

var scoreCodes map[int]string

type Game struct {
	Player1      string
	Player2      string
	Player1Score int
	Player2Score int
}

func init() {
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

	if g.Player1Score == g.Player2Score {
		if g.Player1Score >= 3 {
			result = "Deuce"
		} else {
			result = scoreCodes[g.Player1Score] + " All"
		}
	} else if g.Player1Score > 0 || g.Player2Score > 0 {
		if g.Player1Score >= 3 && g.Player2Score >= 3 {
			player := g.Player1
			if g.Player1Score < g.Player2Score {
				player = g.Player2
			}

			if math.Abs(float64(g.Player1Score) - float64(g.Player2Score)) == 1 {
				result = "Advantage " + player
			} else {
				result = "Win " + player
			}
		} else {
			result = scoreCodes[g.Player1Score] + " - " + scoreCodes[g.Player2Score]
		}
	}

	return result
}

func (g *Game) pointPlayer1() {
	g.Player1Score++
}

func (g *Game) pointPlayer2() {
	g.Player2Score++
}
