package ui_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/bowl/model"
	"github.com/rosenhouse/bowl/ui"
)

var simple_scoresheet = `
╔═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╦═══════════╗
║     1     ║     2     ║     3     ║     4     ║     5     ║     6     ║     7     ║     8     ║     9     ║    10     ║
╟─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫─────┬─────╫───┬───┬───╢
║  9  │  0  ║  8  │  0  ║  3  │  3  ║  5  │  2  ║  4  │  4  ║  4  │  3  ║  3  │  1  ║  1  │  0  ║  0  │  9  ║ 9 │ 0 │   ║
║     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢     └─────╢   └───┴───╢
║    ___    ║    ___    ║    ___    ║    ___    ║    ___    ║    ___    ║    ___    ║    ___    ║    ___    ║    ___    ║
╚═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╩═══════════╝
`

var simple_game = model.Game{
	model.Frame{'9', '0'},
	model.Frame{'8', '0'},
	model.Frame{'3', '3'},
	model.Frame{'5', '2'},
	model.Frame{'4', '4'},
	model.Frame{'4', '3'},
	model.Frame{'3', '1'},
	model.Frame{'1', '0'},
	model.Frame{'0', '9'},
	model.Frame{'9', '0'},
}

var _ = Describe("Parse", func() {
	It("extracts a Game from input text", func() {
		game, err := ui.ParseScorecard(simple_scoresheet)
		Expect(err).NotTo(HaveOccurred())

		Expect(game).To(HaveLen(10))
		Expect(game).To(Equal(simple_game))
	})
})
