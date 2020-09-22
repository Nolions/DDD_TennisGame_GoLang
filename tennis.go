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

/*
	初始化&建立一個Game實體
*/
func TennisGame(player1, player2 string) *Game {
	game := &Game{
		Player1:      player1,
		Player2:      player2,
		Player1Score: 0,
		Player2Score: 0,
	}

	return game
}

/*
	取得目錢比較狀況
*/
func (g *Game) Result() string {
	result := "Love - All"

	if isTied(g.Player1Score, g.Player2Score) {
		if g.Player1Score >= 3 {
			result = "Deuce"
		} else {
			result = scoreCodes[g.Player1Score] + " All"
		}
	} else if g.isScoreGreaterEqualForty() {
		advantagePlayer := g.whoAdvantage()

		if isAdvantage(g.Player1Score, g.Player2Score) {
			result = "Advantage " + advantagePlayer
		} else {
			result = "Win " + advantagePlayer
		}
	} else {
		result = scoreCodes[g.Player1Score] + " - " + scoreCodes[g.Player2Score]
	}

	return result
}

// 比分是否為平手
func isTied(score1, score2 int) bool {
	if score1 == score2 {
		return true
	}

	return false
}

// 是否Advantage狀況
func isAdvantage(score1, score2 int) bool {
	if math.Abs(float64(score1)-float64(score2)) == 1 {
		return true
	}

	return false
}

// 目前誰占優勢
func (g *Game) whoAdvantage() string {
	if g.Player1Score > g.Player2Score {
		return g.Player1
	}

	return g.Player2
}

// 目前雙方比方是否皆為大於等於Forty
func (g *Game) isScoreGreaterEqualForty() bool {
	if g.Player1Score >= 3 && g.Player2Score >= 3 {
		return true
	}

	return false
}

// Player1 得分
func (g *Game) pointPlayer1() {
	g.Player1Score++
}

// Player2 得分
func (g *Game) pointPlayer2() {
	g.Player2Score++
}
