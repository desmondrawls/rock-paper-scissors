package play_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "github.com/desmondrawls/rock-paper-scissors/models"
    "github.com/desmondrawls/rock-paper-scissors/play"
    "github.com/desmondrawls/rock-paper-scissors/play/playfakes"
)

var _ = Describe("GameUseCase", func() {
    var (
        playUseCase *play.UseCase
        repository  *playfakes.RepositorySpy
    )

    BeforeEach(func() {
        repository = &playfakes.RepositorySpy{}
        playUseCase = &play.UseCase{Repository: repository}
    })

    It("invalid throws", func() {
        throws := models.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "desmond",
            Player1Throw: "sailboat",
            Player2Throw: "paper",
        }
        uiSpy := &playfakes.UISpy{}

        playUseCase.Execute(throws, uiSpy)

        Expect(uiSpy.InvalidCallCount()).To(Equal(1))
        Expect(uiSpy.InvalidArgsForCall(0)).To(Equal(throws))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
        Expect(uiSpy.WinnerCallCount()).To(Equal(0))

        By("not saving invalid games")
        Expect(repository.SaveCallCount()).To(Equal(0))
    })

    It("rock beats scissors", func() {
        throws := models.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "desmond",
            Player1Throw: "rock",
            Player2Throw: "paper",
        }
        uiSpy := &playfakes.UISpy{}

        playUseCase.Execute(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("desmond"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
        Expect(uiSpy.InvalidCallCount()).To(Equal(0))

        By("saving the inputs and result")
        Expect(repository.SaveCallCount()).To(Equal(1))
        expectedRecord := models.Record{
            Inputs: models.Inputs(throws),
            Result: models.Result{
                Winner: "desmond",
                IsDraw: false,
            },
        }
        Expect(repository.SaveArgsForCall(0)).To(Equal(expectedRecord))
    })

    It("scissors beats paper", func() {
        throws := models.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "desmond",
            Player1Throw: "scissors",
            Player2Throw: "paper",
        }
        uiSpy := &playfakes.UISpy{}

        playUseCase.Execute(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("gabe"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
        Expect(uiSpy.InvalidCallCount()).To(Equal(0))
    })

    It("scissors beats paper", func() {
        throws := models.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "other-player",
            Player1Throw: "paper",
            Player2Throw: "scissors",
        }
        uiSpy := &playfakes.UISpy{}

        playUseCase.Execute(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("other-player"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
        Expect(uiSpy.InvalidCallCount()).To(Equal(0))
    })

    It("paper beats rock", func() {
        throws := models.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "player2-name",
            Player1Throw: "rock",
            Player2Throw: "paper",
        }
        uiSpy := &playfakes.UISpy{}

        playUseCase.Execute(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("player2-name"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
        Expect(uiSpy.InvalidCallCount()).To(Equal(0))
    })

    It("paper beats rock", func() {
        throws := models.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "player2-name",
            Player1Throw: "paper",
            Player2Throw: "rock",
        }
        uiSpy := &playfakes.UISpy{}

        playUseCase.Execute(throws, uiSpy)

        Expect(uiSpy.WinnerCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerArgsForCall(0)).To(Equal("gabe"))
        Expect(uiSpy.DrawCallCount()).To(Equal(0))
        Expect(uiSpy.InvalidCallCount()).To(Equal(0))
    })

    It("draw", func() {
        throws := models.Inputs{
            Player1Name:  "gabe",
            Player2Name:  "desmond",
            Player1Throw: "rock",
            Player2Throw: "rock",
        }
        uiSpy := &playfakes.UISpy{}

        playUseCase.Execute(throws, uiSpy)

        Expect(uiSpy.DrawCallCount()).To(Equal(1))
        Expect(uiSpy.WinnerCallCount()).To(Equal(0))
        Expect(uiSpy.InvalidCallCount()).To(Equal(0))
    })
})
