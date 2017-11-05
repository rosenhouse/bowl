package ui

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rosenhouse/bowl/model"
)

func Output(game model.ScoredGame) string {
	out := []byte(strings.TrimSpace(Template))
	markAllScores(out, game.Scores)
	markAllThrows(out, game.Game)
	return string(out)
}

func markAllScores(out []byte, scores []int) {
	cursor := 0
	for _, score := range scores {
		cursor = bytes.Index(out, ScoreMark)
		if cursor == -1 {
			panic("got more scores than slots in the template")
		}
		markNextScore(out[cursor:], score)
	}
}

func markNextScore(out []byte, score int) {
	scoreString := fmt.Sprintf("%3d", score)
	for i, c := range []byte(scoreString) {
		out[i] = c
	}
}

func markAllThrows(out []byte, frames []model.Frame) {
	cursor := 0
	for _, frame := range frames {
		for _, throw := range frame {
			cursor = bytes.Index(out, []byte{ThrowMark})
			if cursor == -1 {
				panic("got more throws than slots in the template")
			}
			markNextThrow(out[cursor:], throw)
		}
	}
}

func markNextThrow(out []byte, throw model.Throw) {
	out[0] = byte(throw)
}
