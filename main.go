package main

import (
    "fmt"
    "log"
    "os"

    play "github.com/desmondrawls/rock-paper-scissors/play"
)

func main() {
    args := os.Args
    throws := map[string]string{args[1]: args[2],
        args[3]: args[4]}
    winFinder := &play.WinFinder{
        Comparer: play.Compare,
    }
    result, err := winFinder.GetWinner(throws)
    if err != nil {
        log.Fatal(err)
    }
    if result.IsDraw() {
        fmt.Printf("ITS A DRAW\n")
    }
    fmt.Printf("WINNER: %s\n", result.Winner)
}
