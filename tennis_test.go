package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoveAll(t *testing.T) {
	game := TennisGame("A", "B")
	result := game.Result()
	assert.Equal(t, "Love - All", result)
}

func TestFifteenLove(t *testing.T) {
	game := TennisGame("A", "B")
	game.point(1)
	result := game.Result()
	assert.Equal(t, "Fifteen - Love", result)
}

func TestThirtyLove(t *testing.T) {
	game := TennisGame("A", "B")

	game.point(2)

	result := game.Result()
	assert.Equal(t, "Thirty - Love", result)
}

func (g *Game) point(score int) {
	for i := 0; i < score; i++ {
		g.pointPlayer1()
	}
}
