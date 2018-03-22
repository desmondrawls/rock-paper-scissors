package main_test

import (
	"fmt"
	"net/http"
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
		baseURL       string
	)

	getHTTPStatus := func() (int, error) {
		resp, err := http.Get(baseURL)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		return resp.StatusCode, nil
	}

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		serverPort := fmt.Sprintf("%d", PickAPort())
		serverCmd := exec.Command(pathToServer, "--port", serverPort)
		serverSession, err = gexec.Start(serverCmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		baseURL = fmt.Sprintf("http://127.0.0.1:%s", serverPort)
		Eventually(getHTTPStatus).Should(Equal(http.StatusOK))
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
		Eventually(serverSession.Terminate).Should(gexec.Exit())
	})

	assertJourney := func(player1, player2 string, expectedResult string) {
		Expect(page.Navigate(baseURL)).To(Succeed())
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
		assertJourney("rock", "sailboat", "Invalid input")
	})
})
