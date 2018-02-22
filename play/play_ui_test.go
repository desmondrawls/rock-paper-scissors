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
