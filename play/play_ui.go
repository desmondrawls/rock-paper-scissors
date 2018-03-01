package play

import (
    "strings"
)

//go:generate counterfeiter --fake-name UISpy . UI
type UI interface {
    Winner(string)
    Draw()
    Invalid(Inputs)
}

type Throw int

const (
    ROCK Throw = iota
    PAPER
    SCISSORS
)

type Inputs struct {
    Player1Name, Player2Name   string
    Player1Throw, Player2Throw string
}

func Play(playerThrows Inputs, ui UI) {
    player1Throw, ok := parseThrow(playerThrows.Player1Throw)
    if !ok {
        ui.Invalid(playerThrows)
        return
    }
    player2Throw, ok := parseThrow(playerThrows.Player2Throw)
    if !ok {
        ui.Invalid(playerThrows)
        return
    }
    if player1Throw == player2Throw {
        ui.Draw()
        return
    }
    if beats(player1Throw, player2Throw) {
        ui.Winner(playerThrows.Player1Name)
        return
    }
    ui.Winner(playerThrows.Player2Name)
}

func beats(a, b Throw) bool {
    switch a {
    case ROCK:
        return b == SCISSORS
    case PAPER:
        return b == ROCK
    case SCISSORS:
        return b == PAPER
    }
    return false
}

func parseThrow(input string) (Throw, bool) {
    switch strings.ToLower(input) {
    case "rock":
        return ROCK, true
    case "paper":
        return PAPER, true
    case "scissors":
        return SCISSORS, true
    }
    return 0, false
}
