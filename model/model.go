package model

import "strconv"

type Game []Frame

type Frame []Throw

type Throw rune

func (t Throw) AsNumber() int {
	n, err := strconv.Atoi(string(t))
	if err != nil {
		panic("not yet supporting non-numeric throws")
	}
	return n
}

type ScoredGame struct {
	Game
	Scores []int
}

func Score(game Game) (ScoredGame, error) {
	scores := make([]int, 10)
	total := 0
	for i, frame := range game {
		scores[i] = frame[0].AsNumber() + frame[1].AsNumber() + total
		total = scores[i]
	}
	return ScoredGame{
		Game:   game,
		Scores: scores,
	}, nil
}
