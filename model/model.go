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
	for i, _ := range game {
		scores[i] = ScoreFrame(game[i:]) + total
		total = scores[i]
	}
	return ScoredGame{
		Game:   game,
		Scores: scores,
	}, nil
}

const MarkSpare = Throw('/')

func ScoreFrame(frames []Frame) int {
	if frames[0][1] == MarkSpare {
		return 10 + frames[1][0].AsNumber()
	}
	return frames[0][0].AsNumber() + frames[0][1].AsNumber()
}
