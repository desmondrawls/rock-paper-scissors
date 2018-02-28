package web_ui

import (
    "errors"
    "fmt"
    "net/http"
    "strings"

    "github.com/desmondrawls/rock-paper-scissors/play"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        w.Write([]byte(`<body>
        <form action="/play" method="POST">
        <label for:"player1">P1</label>
        <input name="player1" type="string"/>
        <br>
        <label for:"player2">P2</label>
        <input name="player2" type="string"/>
        <br>
        <input type="submit" value="Play" />
        </form>
        </body>`))
        return
    }
    if r.URL.Path == "/play" && r.Method == "POST" {
        p1Throw := r.FormValue("player1")
        p2Throw := r.FormValue("player2")
        inputs, err := buildInputs(p1Throw, p2Throw)
        if err != nil {
            w.Write([]byte(`invalid input: TODO: return HTML HERE`))
            return
        }

        play.Play(inputs, &web_ui{ResponseWriter: w})
        fmt.Printf("p1: %s\np2: %s\n", p1Throw, p2Throw)
    }
}

type web_ui struct {
    http.ResponseWriter
}

func (w web_ui) Winner(name string) {
    w.Write([]byte(fmt.Sprintf("<body>%s <br> WINS!</body>", name)))
}

func (w web_ui) Draw() {
    w.Write([]byte("TIE!"))
}

func parseThrow(input string) (play.Throw, bool) {
    switch strings.ToLower(input) {
    case "rock":
        return play.ROCK, true
    case "paper":
        return play.PAPER, true
    case "scissors":
        return play.SCISSORS, true
    }
    return 0, false
}

func buildInputs(p1Input, p2Input string) (play.Inputs, error) {
    p1Throw, ok := parseThrow(p1Input)
    if !ok {
        return play.Inputs{}, errors.New("unable to parse p1 input as a throw")
    }
    p2Throw, ok := parseThrow(p2Input)
    if !ok {
        return play.Inputs{}, errors.New("unable to parse p2 input as a throw")
    }
    throws := play.Inputs{
        Player1Name:  "player1",
        Player2Name:  "player2",
        Player1Throw: p1Throw,
        Player2Throw: p2Throw,
    }
    return throws, nil
}
