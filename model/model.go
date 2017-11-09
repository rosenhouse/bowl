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

func (f Frame) IsSpare() bool {
	return f[1] == MarkSpare
}

func (f Frame) IsStrike() bool {
	return f[1] == MarkStrike
}

func (f Frame) PinsDownTotal() int {
	if f.IsSpare() || f.IsStrike() {
		return 10
	}
	return f[0].AsNumber() + f[1].AsNumber()
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
const MarkStrike = Throw('X')
const MarkNoThrow = Throw(' ')

func ScoreFrame(frames []Frame) int {
	if frames[0].IsStrike() {
		if frames[1].IsSpare() {
			return 20
		}
		return 10 + frames[1].PinsDownTotal()
	}
	if frames[0].IsSpare() {
		if frames[1].IsStrike() {
			return 20
		}
		return 10 + frames[1][0].AsNumber()
	}
	return frames[0].PinsDownTotal()
}
