package main_test

import (
    "errors"

    "github.com/desmondrawls/rock-paper-scissors/history"
    "github.com/desmondrawls/rock-paper-scissors/history/historyfakes"
    "github.com/desmondrawls/rock-paper-scissors/models"
    "github.com/desmondrawls/rock-paper-scissors/play"
    "github.com/desmondrawls/rock-paper-scissors/play/playfakes"
    "github.com/desmondrawls/rock-paper-scissors/repository"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("History", func() {
    var (
        repositoryStub    *historyfakes.RepositoryStub
        historyUseCase    *history.UseCase
        historyVisualizer *historyfakes.VisualizerSpy
    )

    BeforeEach(func() {
        repositoryStub = &historyfakes.RepositoryStub{}
        historyUseCase = &history.UseCase{Repository: repositoryStub}
        historyVisualizer = &historyfakes.VisualizerSpy{}
    })

    Context("when no rounds have been played", func() {
        It("returns empty history when no rounds have been played", func() {
            historyUseCase.Execute(historyVisualizer)

            Expect(historyVisualizer.EmptyCallCount()).To(Equal(1))
        })
    })

    Context("when the repository returns an error", func() {
        BeforeEach(func() {
            repositoryStub.ListReturns(nil, errors.New("potato"))
        })

        It("calls error on the visualizer", func() {
            historyUseCase.Execute(historyVisualizer)

            Expect(historyVisualizer.ErrorCallCount()).To(Equal(1))
        })
    })

    Context("when rounds have been played", func() {
        BeforeEach(func() {
            repositoryFake := &repository.FakeRepository{}
            historyUseCase = &history.UseCase{Repository: repositoryFake}
            playUseCase := &play.UseCase{Repository: repositoryFake}
            throws := models.Inputs{
                Player1Name:  "desmond",
                Player2Name:  "gabe",
                Player1Throw: "rock",
                Player2Throw: "scissors",
            }

            uiSpy := &playfakes.UISpy{}
            playUseCase.Execute(throws, uiSpy)
        })
        It("shows all the rounds in the visualizer", func() {
            historyUseCase.Execute(historyVisualizer)

            Expect(historyVisualizer.RecordsCallCount()).To(Equal(1))
            expectedInputs := models.Inputs{
                Player1Name:  "desmond",
                Player2Name:  "gabe",
                Player1Throw: "rock",
                Player2Throw: "scissors",
            }
            expectedResult := models.Result{
                Winner: "desmond",
                IsDraw: false,
            }
            expectedRecords := []models.Record{models.Record{expectedInputs, expectedResult}}
            Expect(historyVisualizer.RecordsArgsForCall(0)).To(Equal(expectedRecords))
            Expect(historyVisualizer.EmptyCallCount()).To(Equal(0))
            Expect(historyVisualizer.ErrorCallCount()).To(Equal(0))

        })
    })
})
