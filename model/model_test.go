package model_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/bowl/model"
)

var _ = Describe("Scoring", func() {
	It("scores a simple game", func() {
		scoredGame, err := model.Score(simple_game)
		Expect(err).NotTo(HaveOccurred())

		Expect(scoredGame.Game).To(Equal(simple_game))
		Expect(scoredGame.Scores).To(Equal([]int{
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
		}))

	})
})

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
