package ui_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/bowl/model"
	"github.com/rosenhouse/bowl/ui"
)

var _ = Describe("Output", func() {
	It("formats a scored game card", func() {
		outText := ui.Output(scored_simple_game)
		Expect(outText).To(Equal(strings.TrimSpace(expected_score_card)))
	})
})

var scored_simple_game = model.ScoredGame{
	Game: simple_game,
	Scores: []int{
		9,
		17,
		23,
		30,
		38,
		45,
		49,
		50,
		59,
		68,
	},
}

const expected_score_card = `
╔═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╗
║     1     ║     2     ║     3     ║     4     ║     5     ║     6     ║     7     ║     8     ║     9     ║    10     ║
╟─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫───┬───┬───╢
║  9  │  0  ║  8  │  0  ║  3  │  3  ║  5  │  2  ║  4  │  4  ║  4  │  3  ║  3  │  1  ║  1  │  0  ║  0  │  9  ║ 9 │ 0 │   ║
║     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢   └───┴───╢
║      9    ║     17    ║     23    ║     30    ║     38    ║     45    ║     49    ║     50    ║     59    ║     68    ║
╚═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╝
`
