package models

type Inputs struct {
    Player1Name, Player2Name   string
    Player1Throw, Player2Throw string
}

type Result struct {
    Winner string
    IsDraw bool
}

type Record struct {
    Inputs
    Result
}
