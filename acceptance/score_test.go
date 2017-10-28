package acceptance_test

import (
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Scoring", func() {
	It("scores a simple game (no spares or strikes)", func() {
		var err error
		cmd := exec.Command(binPath, "score")

		cmd.Stdin = strings.NewReader(getFixture("simple-input"))

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		outString := strings.TrimSpace(string(session.Out.Contents()))
		Expect(outString).To(Equal(getFixture("simple-output")))
	})
})
