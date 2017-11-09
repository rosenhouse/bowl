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

	Describe("scoring a single frame, based on the following ones", func() {
		var frames []model.Frame

		Context("when the frame is not a spare or strike", func() {
			BeforeEach(func() {
				frames = []model.Frame{
					{'8', '1'},
					{'3', '3'},
				}
			})

			It("scores the frame by itself", func() {
				score := model.ScoreFrame(frames)
				Expect(score).To(Equal(9))
			})
		})

		Context("when the frame is a spare", func() {
			BeforeEach(func() {
				frames = []model.Frame{
					{'8', '/'},
					{'3', '3'},
				}
			})

			It("scores the frame as 10 plus the first throw of the following frame", func() {
				score := model.ScoreFrame(frames)
				Expect(score).To(Equal(13))
			})
		})

		Context("when there's a strike followed by a simple frame", func() {
			BeforeEach(func() {
				frames = []model.Frame{
					{' ', 'X'},
					{'5', '2'},
				}
			})

			It("scores the first frame as 10 plus the next two throws", func() {
				score := model.ScoreFrame(frames)
				Expect(score).To(Equal(17))
			})
		})

		Context("when there's a strike followed by a spare", func() {
			BeforeEach(func() {
				frames = []model.Frame{
					{' ', 'X'},
					{'5', '/'},
				}
			})

			It("scores the first frame as 20", func() {
				score := model.ScoreFrame(frames)
				Expect(score).To(Equal(20))
			})
		})

		Context("when there's a strike followed by a simple frame", func() {
			BeforeEach(func() {
				frames = []model.Frame{
					{' ', 'X'},
					{'5', '2'},
				}
			})

			It("scores the first frame as 10 plus the next two throws", func() {
				score := model.ScoreFrame(frames)
				Expect(score).To(Equal(17))
			})
		})

		Context("when there's a spare followed by a strike", func() {
			BeforeEach(func() {
				frames = []model.Frame{
					{'5', '/'},
					{' ', 'X'},
				}
			})

			It("scores the first frame as 20", func() {
				score := model.ScoreFrame(frames)
				Expect(score).To(Equal(20))
			})
		})

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
