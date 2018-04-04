package web_ui

import (
    "fmt"
    "net/http"

    "github.com/desmondrawls/rock-paper-scissors/models"
    "github.com/desmondrawls/rock-paper-scissors/play"
)

const (
    HomePageText = `<body>
        <form action="/play" method="POST">
        <label for:"player1">P1</label>
        <input name="player1" type="string"/>
        <br>
        <label for:"player2">P2</label>
        <input name="player2" type="string"/>
        <br>
        <input type="submit" value="Play" />
        </form>
        </body>`

    InvalidInputPageTemplate = `<body>
        <h1>Invalid input</h1>
        <form action="/play" method="POST">
        <label for:"player1">P1</label>
        <input name="player1" type="string" value=%q/>
        <br>
        <label for:"player2">P2</label>
        <input name="player2" type="string" value=%q/>
        <br>
        <input type="submit" value="Play" />
        </form>
        </body>`
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        w.Write([]byte(HomePageText))
        return
    }
    if r.URL.Path == "/play" && r.Method == "POST" {
        throws := models.Inputs{
            Player1Name:  "player1",
            Player2Name:  "player2",
            Player1Throw: r.FormValue("player1"),
            Player2Throw: r.FormValue("player2"),
        }
        play.Play(throws, &web_ui{
            ResponseWriter: w,
        })
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

func (w web_ui) Invalid(throws models.Inputs) {
    w.Write([]byte(fmt.Sprintf(InvalidInputPageTemplate, throws.Player1Throw, throws.Player2Throw)))
}
