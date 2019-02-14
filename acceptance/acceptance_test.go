package acceptance_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Scratchpad", func() {
	var session *gexec.Session

	BeforeEach(func() {
		scratchpadPath := buildScratchpad()
		session = runScratchpad(scratchpadPath)
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("exits with status code 0", func() {
		Eventually(session).Should(gexec.Exit(0))
	})

	It("prints 'Hello World' to stdout", func() {
		Eventually(session).Should(gbytes.Say("Hello World"))
	})
})

func buildScratchpad() string {
	scratchpadPath, err := gexec.Build("github.com/williammartin/scratchpad")
	Expect(err).NotTo(HaveOccurred())

	return scratchpadPath
}

func runScratchpad(path string) *gexec.Session {
	cmd := exec.Command(path)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
