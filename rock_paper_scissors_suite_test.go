package main_test

import (
    "sync"
    "testing"

    . "github.com/onsi/ginkgo"
    "github.com/onsi/ginkgo/config"
    . "github.com/onsi/gomega"
    "github.com/onsi/gomega/gexec"
    "github.com/sclevine/agouti"
)

func TestRockPaperScissors(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "RockPaperScissors Suite")
}

var agoutiDriver *agouti.WebDriver
var pathToServer string

var _ = BeforeSuite(func() {
    agoutiDriver = agouti.PhantomJS()
    Expect(agoutiDriver.Start()).To(Succeed())

    var err error
    pathToServer, err = gexec.Build("github.com/desmondrawls/rock-paper-scissors")
    Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
    Expect(agoutiDriver.Stop()).To(Succeed())
    gexec.CleanupBuildArtifacts()
})

var (
    lastPortUsed int
    mutex        sync.Mutex
    once         sync.Once
)

// PickAPort returns a port that is likely free for use in a Ginkgo test
func PickAPort() int {
    mutex.Lock()
    defer mutex.Unlock()

    if lastPortUsed == 0 {
        once.Do(func() {
            const portRangeStart = 61000
            lastPortUsed = portRangeStart + config.GinkgoConfig.ParallelNode
        })
    }

    lastPortUsed += config.GinkgoConfig.ParallelTotal
    return lastPortUsed
}
