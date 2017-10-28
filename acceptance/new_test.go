package acceptance_test

import (
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Generating a new game card", func() {
	It("prints a blank scorecard", func() {
		var err error
		cmd := exec.Command(binPath, "new")
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		outString := strings.TrimSpace(string(session.Out.Contents()))
		Expect(outString).To(Equal(getFixture("empty")))
	})
})
