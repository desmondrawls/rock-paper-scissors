package play

import (
    "strings"

    "github.com/desmondrawls/rock-paper-scissors/models"
)

//go:generate counterfeiter --fake-name UISpy . UI
type UI interface {
    Winner(string)
    Draw()
    Invalid(models.Inputs)
}

type Throw int

const (
    ROCK Throw = iota
    PAPER
    SCISSORS
)

//go:generate counterfeiter --fake-name RepositorySpy . Repository
type Repository interface {
    Save(models.Record) error
}

type UseCase struct {
    Repository Repository
}

type SaveRecordUseCase struct {
    Repository Repository
    Inputs     models.Inputs
}

func (s *SaveRecordUseCase) Winner(name string) {
    s.Repository.Save(models.Record{
        Inputs: s.Inputs,
        Result: models.Result{
            Winner: name,
            IsDraw: false,
        },
    })
}

func (s *SaveRecordUseCase) Draw()                 {}
func (s *SaveRecordUseCase) Invalid(models.Inputs) {}

func (u *UseCase) Execute(throws models.Inputs, ui UI) {
    Play(throws, &SaveRecordUseCase{Repository: u.Repository, Inputs: models.Inputs(throws)})
    Play(throws, ui)
}

func Play(playerThrows models.Inputs, ui UI) {
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
