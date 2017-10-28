package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestBowl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bowl Suite")
}

var binPath string

var _ = SynchronizedBeforeSuite(func() []byte {
	fmt.Fprintf(GinkgoWriter, "building binary...")
	path, err := gexec.Build("github.com/rosenhouse/bowl")
	fmt.Fprintf(GinkgoWriter, "done")
	Expect(err).NotTo(HaveOccurred())
	return []byte(path)
}, func(data []byte) {
	binPath = string(data)
	rand.Seed(config.GinkgoConfig.RandomSeed + int64(GinkgoParallelNode()))
})

var _ = SynchronizedAfterSuite(func() {}, func() {
	gexec.CleanupBuildArtifacts()
})

func getFixture(name string) string {
	inBytes, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s.txt", name))
	Expect(err).NotTo(HaveOccurred())
	return strings.TrimSpace(string(inBytes))
}
