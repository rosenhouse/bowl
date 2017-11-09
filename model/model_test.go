package model_test

import (
	. "github.com/rosenhouse/bowl/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("NonFinalFrame", func() {
	DescribeTable("PinsKnockedDown",
		func(frame Frame, expected []int) {
			Expect(frame.PinsKnockedDown()).To(Equal(expected))
		},
		Entry("zeros", NonFinalFrame{0, 0}, []int{0, 0}),
		Entry("bagel disabled", NonFinalFrame{1, 0}, []int{1, 0}),
		Entry("not bad", NonFinalFrame{8, 1}, []int{8, 1}),
		Entry("not bad", NonFinalFrame{8, ThrowSpare}, []int{8, 2}),
		Entry("strike", NonFinalFrame{ThrowStrike}, []int{10}),

		Entry("final frame: max score", FinalFrame{ThrowStrike, ThrowStrike, ThrowStrike}, []int{10, 10, 10}),
		Entry("final frame: almost max", FinalFrame{ThrowStrike, ThrowStrike, 7}, []int{10, 10, 7}),
		Entry("final frame: spare strike", FinalFrame{3, ThrowSpare, ThrowStrike}, []int{3, 7, 10}),
		Entry("final frame: spare ok", FinalFrame{3, ThrowSpare, 2}, []int{3, 7, 2}),
		Entry("final frame: spare oops", FinalFrame{3, ThrowSpare, 0}, []int{3, 7, 0}),
		Entry("final frame: meh", FinalFrame{3, 6}, []int{3, 6}),
	)

	DescribeTable("scoring a frame, given its followers",
		func(frames []Frame, expected int) {
			Expect(ScoreFrames(frames)).To(Equal(expected))
		},
		Entry("meh", []Frame{NonFinalFrame{7, 1}, NonFinalFrame{3, 1}}, 8),
		Entry("spare, meh", []Frame{NonFinalFrame{3, ThrowSpare}, NonFinalFrame{3, 1}}, 13),
		Entry("spare, strike", []Frame{NonFinalFrame{3, ThrowSpare}, NonFinalFrame{ThrowStrike}, NonFinalFrame{1, 3}}, 20),
		Entry("strike, meh", []Frame{NonFinalFrame{ThrowStrike}, NonFinalFrame{3, 1}}, 14),
		Entry("strike, spare", []Frame{NonFinalFrame{ThrowStrike}, NonFinalFrame{3, ThrowSpare}}, 20),
		Entry("strike, strike, meh", []Frame{NonFinalFrame{ThrowStrike}, NonFinalFrame{ThrowStrike}, NonFinalFrame{3, 5}}, 23),
		Entry("strike, strike, meh", []Frame{NonFinalFrame{ThrowStrike}, NonFinalFrame{ThrowStrike}, FinalFrame{3, 5}}, 23),
		Entry("strike, strike, spare", []Frame{NonFinalFrame{ThrowStrike}, NonFinalFrame{ThrowStrike}, NonFinalFrame{3, ThrowSpare}}, 23),
		Entry("strike, strike, strike", []Frame{NonFinalFrame{ThrowStrike}, NonFinalFrame{ThrowStrike}, NonFinalFrame{ThrowStrike}}, 30),
		Entry("strike, strike, strike, meh", []Frame{NonFinalFrame{ThrowStrike}, FinalFrame{ThrowStrike, ThrowStrike, 3}}, 30),
	)
})
