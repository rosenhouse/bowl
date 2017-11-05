package ui

import (
	"strings"

	"github.com/rosenhouse/bowl/model"
)

func extractMarks(inputText string) []byte {
	mask := strings.TrimSpace(Template)
	trimmedInput := strings.TrimSpace(inputText)

	var output []byte
	for i, c := range mask {
		if c == ThrowMark {
			output = append(output, trimmedInput[i])
		}
	}

	return output
}

func ParseScorecard(inputText string) (model.Game, error) {
	marks := extractMarks(inputText)
	game := model.Game{}
	for i := 0; i < 20; i += 2 {
		game = append(game, model.Frame{
			model.Throw(marks[i]), model.Throw(marks[i+1]),
		})
	}

	return game, nil
}
