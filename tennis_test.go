package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestLoveAll(t *testing.T) {
	game := TennisGame("A", "B")
	result := game.Result()
	assert.Equal(t, "Love All", result)
}

func TestFifteenLove(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(1, game.Player1)
	result := game.Result()
	assert.Equal(t, "Fifteen - Love", result)
}

func TestThirtyLove(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(2, game.Player1)

	result := game.Result()
	assert.Equal(t, "Thirty - Love", result)
}

func TestFortyLove(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(3, game.Player1)

	result := game.Result()
	assert.Equal(t, "Forty - Love", result)
}

func TestLoveFifteen(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(1, game.Player2)

	result := game.Result()
	assert.Equal(t, "Love - Fifteen", result)
}

func TestLoveThirty(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(2, game.Player2)

	result := game.Result()
	assert.Equal(t, "Love - Thirty", result)
}

func TestLoveForty(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(3, game.Player2)

	result := game.Result()
	assert.Equal(t, "Love - Forty", result)
}

func TestFifteenAll(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(1, game.Player1)
	game.point(1, game.Player2)

	result := game.Result()
	assert.Equal(t, "Fifteen All", result)
}

func TestThirtyAll(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(2, game.Player1)
	game.point(2, game.Player2)

	result := game.Result()
	assert.Equal(t, "Thirty All", result)
}

func TestDeuce(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(3, game.Player1)
	game.point(3, game.Player2)

	result := game.Result()
	assert.Equal(t, "Deuce", result)
}

func TestDeuceScoreOver3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	max := 99
	min := 4
	score := rand.Intn(max - min + 1 + min)

	game := TennisGame("A", "B")
	game.point(score, game.Player1)
	game.point(score, game.Player2)

	result := game.Result()
	assert.Equal(t, "Deuce", result)
}

func (g *Game) point(score int, player string) {
	for i := 0; i < score; i++ {
		if player == g.Player1 {
			g.pointPlayer1()
		} else {
			g.pointPlayer2()
		}
	}
}
