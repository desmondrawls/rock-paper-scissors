package play_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "github.com/desmondrawls/rock-paper-scissors/play"
)

var _ = Describe("Comparison", func() {
    Specify("rock beats scissors", func() {
        Expect(play.Compare("rock", "scissors")).To(Equal("rock"))
    })

    Describe("Finding the winner", func() {
        Specify("it returns the name of the player who had the winning throw", func() {
            throws := map[string]string{
                "gabe":    "zebra",
                "desmond": "sailboat",
            }
            comparerStub := func(_ string, _ string) string {
                return "sailboat"
            }
            winFinder := &play.WinFinder{
                Comparer: comparerStub,
            }

            result, err := winFinder.GetWinner(throws)
            Expect(err).NotTo(HaveOccurred())

            Expect(result.IsDraw()).To(BeFalse())
            Expect(result.Winner).To(Equal("desmond"))
        })

        Specify("draw", func() {
            throws := map[string]string{
                "gabe":    "zebra",
                "desmond": "sailboat",
            }
            comparerStub := func(_ string, _ string) string {
                return "draw"
            }
            winFinder := &play.WinFinder{
                Comparer: comparerStub,
            }

            result, err := winFinder.GetWinner(throws)
            Expect(err).NotTo(HaveOccurred())
            Expect(result.IsDraw()).To(BeTrue())
        })
    })
})
