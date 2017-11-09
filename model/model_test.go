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
})
