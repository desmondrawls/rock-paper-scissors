package play

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Play", func() {
    It("tests stuff", func() {
        Expect(true).To(Equal(true))
    })

    It("sometimes stuff fails", func() {
        Expect(true).To(Equal(false))
    })
})
