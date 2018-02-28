package play_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "github.com/desmondrawls/rock-paper-scissors/play"
    "github.com/desmondrawls/rock-paper-scissors/play/playfakes"
)

var _ = Describe("GameUseCase", func() {
    It("rock beats scissors", func() {
        throws := play.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "desmond",
            Player1Throw: play.ROCK,
            Player2Throw: play.PAPER,
        }
        uiSpy := &playfakes.UISpy{}

        play.Play(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("desmond"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
    })

    It("scissors beats paper", func() {
        throws := play.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "desmond",
            Player1Throw: play.SCISSORS,
            Player2Throw: play.PAPER,
        }
        uiSpy := &playfakes.UISpy{}

        play.Play(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("gabe"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
    })

    It("scissors beats paper", func() {
        throws := play.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "other-player",
            Player1Throw: play.PAPER,
            Player2Throw: play.SCISSORS,
        }
        uiSpy := &playfakes.UISpy{}

        play.Play(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("other-player"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
    })

    It("paper beats rock", func() {
        throws := play.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "player2-name",
            Player1Throw: play.ROCK,
            Player2Throw: play.PAPER,
        }
        uiSpy := &playfakes.UISpy{}

        play.Play(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("player2-name"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
    })

    It("paper beats rock", func() {
        throws := play.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "player2-name",
            Player1Throw: play.PAPER,
            Player2Throw: play.ROCK,
        }
        uiSpy := &playfakes.UISpy{}

        play.Play(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("gabe"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
    })

    It("draw", func() {
        throws := play.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "desmond",
            Player1Throw: play.ROCK,
            Player2Throw: play.ROCK,
        }
        uiSpy := &playfakes.UISpy{}

        play.Play(throws, uiSpy)

        Expect(uiSpy.DrawCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerCallCount()).To(Equal(0))
    })
})
