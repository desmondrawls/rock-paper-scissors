package play

//go:generate counterfeiter --fake-name UISpy . UI
type UI interface {
    Winner(string)
    Draw()
}

type Throw int

const (
    ROCK Throw = iota
    PAPER
    SCISSORS
)

type Inputs struct {
    Player1Name, Player2Name   string
    Player1Throw, Player2Throw Throw
}

func Play(playerThrows Inputs, ui UI) {
    if playerThrows.Player1Throw == playerThrows.Player2Throw {
        ui.Draw()
    } else {
        ui.Winner("desmond")
    }
}
