package application_test

import (
	"bytes"
	"errors"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/bowl/application"
	"github.com/rosenhouse/bowl/application/fakes"
)

var _ = Describe("Application", func() {
	var (
		app          *application.App
		outBuffer    *bytes.Buffer
		commandNew   *fakes.CommandNew
		commandScore *fakes.CommandScore
	)

	BeforeEach(func() {
		outBuffer = &bytes.Buffer{}
		commandNew = &fakes.CommandNew{}
		commandScore = &fakes.CommandScore{}
		app = &application.App{
			Input:        strings.NewReader("some game input"),
			Output:       outBuffer,
			Arguments:    []string{"bowl", "some-command"},
			CommandNew:   commandNew,
			CommandScore: commandScore,
		}
	})

	Context("when the 1st argument is 'score'", func() {
		BeforeEach(func() {
			app.Arguments[1] = "score"
			commandScore.RunReturns("some score response", nil)
		})

		It("returns the result of the Score command", func() {
			Expect(app.Run()).To(Succeed())
			Expect(outBuffer.String()).To(Equal("some score response"))
		})

		It("passes the Input to the Score command", func() {
			Expect(app.Run()).To(Succeed())
			Expect(commandScore.RunArgsForCall(0)).To(Equal("some game input"))
		})

		Context("when the Score command errors", func() {
			BeforeEach(func() {
				commandScore.RunReturns("", errors.New("banana"))
			})
			It("returns the error", func() {
				Expect(app.Run()).To(MatchError("banana"))
			})
		})
	})

	Context("when the 1st argument is 'new'", func() {
		BeforeEach(func() {
			app.Arguments[1] = "new"
			commandNew.RunReturns("some empty scoreboard")
		})

		It("returns the result of the New command", func() {
			Expect(app.Run()).To(Succeed())
			Expect(outBuffer.String()).To(Equal("some empty scoreboard"))
		})
	})

	Context("when the 1st argument is nonsense", func() {
		BeforeEach(func() {
			app.Arguments = []string{"bowl", "potato"}
		})

		It("returns a usage error", func() {
			Expect(app.Run()).To(Equal(application.ErrUsage))
		})
	})

	Context("when the 1st argument is missing", func() {
		BeforeEach(func() {
			app.Arguments = []string{"bowl"}
		})

		It("returns a usage error", func() {
			Expect(app.Run()).To(Equal(application.ErrUsage))
		})
	})
})
