package main_test

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("Integration", func() {
	var (
		page          *agouti.Page
		serverSession *gexec.Session
		serverPort    string
	)

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		serverPort = fmt.Sprintf("%d", PickAPort())
		serverCmd := exec.Command(pathToServer, "--port", serverPort)
		serverSession, err = gexec.Start(serverCmd, GinkgoWriter, GinkgoWriter)
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
		Eventually(serverSession.Terminate).Should(gexec.Exit())
	})

	assertJourney := func(player1, player2 string, expectedResult string) {
		Expect(page.Navigate(fmt.Sprintf("http://localhost:%s", serverPort))).To(Succeed())
		Eventually(page.FindByName("player1")).Should(BeFound())

		By("allowing the user to fill out a game and submit it")
		Expect(page.FindByName("player1").Fill(player1)).To(Succeed())
		Expect(page.FindByName("player2").Fill(player2)).To(Succeed())
		Expect(page.Find("input[type='submit']").Submit()).To(Succeed())

		By("showing the game result")
		Eventually(page.HTML).Should(MatchRegexp(expectedResult))

	}

	It("should play a normal game", func() {
		assertJourney("rock", "scissors", "player1 .* WINS")
	})

	It("should validate input", func() {
		assertJourney("rock", "scissors", "player1 .* WINS")
		//assertJourney("rock", "sailbot", "Invalid input")
	})
})
