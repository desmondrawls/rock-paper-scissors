package play_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/desmondrawls/rock-paper-scissors/play"
	"github.com/desmondrawls/rock-paper-scissors/play/playfakes"
)

var _ = Describe("GameUseCase", func() {
	It("invalid throws", func() {
		throws := play.Inputs{
			Player1Name:  "gabe",
			Player2Name:  "desmond",
			Player1Throw: "sailboat",
			Player2Throw: "paper",
		}
		uiSpy := &playfakes.UISpy{}

		play.Play(throws, uiSpy)

		Expect(uiSpy.InvalidCallCount()).To(Equal(1))
		Expect(uiSpy.InvalidArgsForCall(0)).To(Equal(throws))
		Expect(uiSpy.DrawCallCount()).To(Equal(0))
		Expect(uiSpy.WinnerCallCount()).To(Equal(0))
	})

	It("rock beats scissors", func() {
		throws := play.Inputs{
			Player1Name:  "gabe",
			Player2Name:  "desmond",
			Player1Throw: "rock",
			Player2Throw: "paper",
		}
		uiSpy := &playfakes.UISpy{}

		play.Play(throws, uiSpy)

		Expect(uiSpy.WinnerCallCount()).To(Equal(1))
		Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("desmond"))
		Expect(uiSpy.DrawCallCount()).To(Equal(0))
		Expect(uiSpy.InvalidCallCount()).To(Equal(0))
	})

	It("scissors beats paper", func() {
		throws := play.Inputs{
			Player1Name:  "gabe",
			Player2Name:  "desmond",
			Player1Throw: "scissors",
			Player2Throw: "paper",
		}
		uiSpy := &playfakes.UISpy{}

		play.Play(throws, uiSpy)

		Expect(uiSpy.WinnerCallCount()).To(Equal(1))
		Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("gabe"))
		Expect(uiSpy.DrawCallCount()).To(Equal(0))
		Expect(uiSpy.InvalidCallCount()).To(Equal(0))
	})

	It("scissors beats paper", func() {
		throws := play.Inputs{
			Player1Name:  "gabe",
			Player2Name:  "other-player",
			Player1Throw: "paper",
			Player2Throw: "scissors",
		}
		uiSpy := &playfakes.UISpy{}

		play.Play(throws, uiSpy)

		Expect(uiSpy.WinnerCallCount()).To(Equal(1))
		Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("other-player"))
		Expect(uiSpy.DrawCallCount()).To(Equal(0))
		Expect(uiSpy.InvalidCallCount()).To(Equal(0))
	})

	It("paper beats rock", func() {
		throws := play.Inputs{
			Player1Name:  "gabe",
			Player2Name:  "player2-name",
			Player1Throw: "rock",
			Player2Throw: "paper",
		}
		uiSpy := &playfakes.UISpy{}

		play.Play(throws, uiSpy)

		Expect(uiSpy.WinnerCallCount()).To(Equal(1))
		Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("player2-name"))
		Expect(uiSpy.DrawCallCount()).To(Equal(0))
		Expect(uiSpy.InvalidCallCount()).To(Equal(0))
	})

	It("paper beats rock", func() {
		throws := play.Inputs{
			Player1Name:  "gabe",
			Player2Name:  "player2-name",
			Player1Throw: "paper",
			Player2Throw: "rock",
		}
		uiSpy := &playfakes.UISpy{}

		play.Play(throws, uiSpy)

		Expect(uiSpy.WinnerCallCount()).To(Equal(1))
		Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("gabe"))
		Expect(uiSpy.DrawCallCount()).To(Equal(0))
		Expect(uiSpy.InvalidCallCount()).To(Equal(0))
	})

	It("draw", func() {
		throws := play.Inputs{
			Player1Name:  "gabe",
			Player2Name:  "desmond",
			Player1Throw: "rock",
			Player2Throw: "rock",
		}
		uiSpy := &playfakes.UISpy{}

		play.Play(throws, uiSpy)

		Expect(uiSpy.DrawCallCount()).To(Equal(1))
		Expect(uiSpy.WinnerCallCount()).To(Equal(0))
		Expect(uiSpy.InvalidCallCount()).To(Equal(0))
	})
})
