package main

import (
    console_ui "github.com/desmondrawls/rock-paper-scissors/console_ui"
    play "github.com/desmondrawls/rock-paper-scissors/play"
)

func main() {
    throws := play.Inputs{
        Player1Name:  "gabe",
        Player2Name:  "player2-name",
        Player1Throw: play.PAPER,
        Player2Throw: play.ROCK,
    }
    play.Play(throws, &console_ui.UI{})
}
