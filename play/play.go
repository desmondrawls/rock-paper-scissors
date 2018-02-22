package play

func Compare(throw1, throw2 string) string {
    return "rock"
}

type WinFinder struct {
    Comparer func(string, string) string
}

type Result struct {
    Winner string
}

func (r Result) IsDraw() bool {
    return r.Winner == ""
}

func (f *WinFinder) GetWinner(throws map[string]string) (Result, error) {
    throwValues := []string{}
    for _, v := range throws {
        throwValues = append(throwValues, v)
    }
    winningThrowValue := f.Comparer(throwValues[0], throwValues[1])
    var winner string
    for k, v := range throws {
        if v == winningThrowValue {
            winner = k
        }
    }

    return Result{Winner: winner}, nil
}
