package repository_test

import (
    "github.com/desmondrawls/rock-paper-scissors/models"
    "github.com/desmondrawls/rock-paper-scissors/repository"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Contract", func() {
    theContract := func(repository repository.Repository) {
        It("keeps track of played games", func() {
            By("starting off empty")
            records, err := repository.List()
            Expect(err).NotTo(HaveOccurred())
            Expect(records).To(BeEmpty())

            By("saving a game")
            someRecord := models.Record{
                Inputs: models.Inputs{
                    Player1Name:  "desmond",
                    Player2Name:  "gabe",
                    Player1Throw: "rock",
                    Player2Throw: "scissors",
                },
                Result: models.Result{
                    Winner: "desmond",
                    IsDraw: false,
                },
            }
            Expect(repository.Save(someRecord)).To(Succeed())

            By("expecting to see the game in the history")
            records, err = repository.List()
            Expect(err).NotTo(HaveOccurred())
            Expect(records).To(Equal([]models.Record{someRecord}))
        })
    }

    theContract(&repository.FakeRepository{})
})
