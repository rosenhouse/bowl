package model_test

import (
	. "github.com/rosenhouse/bowl/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Frames", func() {
	DescribeTable("PinsDownInFirstThrow",
		func(frame Frame, expected int) {
			Expect(frame.PinsDownInFirstThrow()).To(Equal(expected))
		},
		Entry("frame not started", []Throw{}, -1),
		Entry("gutter ball", []Throw{0}, 0),
		Entry("zeros", []Throw{0, 0}, 0),
		Entry("bagel disabled", []Throw{1, 0}, 1),
		Entry("not bad", []Throw{8, 1}, 8),
		Entry("strike", []Throw{ThrowStrike}, 10),
	)

	DescribeTable("PinsDownTotal",
		func(frame Frame, expected int) {
			Expect(frame.PinsDownTotal()).To(Equal(expected))
		},
		Entry("frame not started", []Throw{}, -1),
		Entry("incomplete frame", []Throw{8}, -1),
		Entry("zeros", []Throw{0, 0}, 0),
		Entry("bagel disabled", []Throw{0, 1}, 1),
		Entry("not bad", []Throw{8, 1}, 9),
		Entry("spare", []Throw{4, ThrowSpare}, 10),
		Entry("strike", []Throw{ThrowStrike}, 10),
		Entry("10th frame: max score", []Throw{ThrowStrike, ThrowStrike, ThrowStrike}, 30),
		Entry("10th frame: almost max", []Throw{ThrowStrike, ThrowStrike, 7}, 27),
		Entry("10th frame: spare strike", []Throw{3, ThrowSpare, ThrowStrike}, 20),
		Entry("10th frame: spare ok", []Throw{3, ThrowSpare, 2}, 12),
		Entry("10th frame: spare oops", []Throw{3, ThrowSpare, 0}, 10),
		Entry("10th frame: meh", []Throw{3, 6}, 9),
	)

	DescribeTable("scoring a frame, given its followers",
		func(frames []Frame, expected int) {
			Expect(ScoreFrames(frames)).To(Equal(expected))
		},
		Entry("meh", []Frame{{7, 1}, {3, 1}}, 8),
		Entry("spare, meh", []Frame{{3, ThrowSpare}, {3, 1}}, 13),
		Entry("spare, strike", []Frame{{3, ThrowSpare}, {ThrowStrike}, {1, 3}}, 20),
		Entry("strike, meh", []Frame{{ThrowStrike}, {3, 1}}, 14),
		Entry("strike, spare", []Frame{{ThrowStrike}, {3, ThrowSpare}}, 20),
		Entry("strike, strike, meh", []Frame{{ThrowStrike}, {ThrowStrike}, {3, 5}}, 23),
		Entry("strike, strike, spare", []Frame{{ThrowStrike}, {ThrowStrike}, {3, ThrowSpare}}, 23),
		Entry("strike, strike, strike", []Frame{{ThrowStrike}, {ThrowStrike}, {ThrowStrike}, {5}}, 30),
	)

})
