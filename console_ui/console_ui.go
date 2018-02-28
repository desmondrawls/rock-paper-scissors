package console_ui

import (
    "fmt"
)

type UI struct {
}

func (ui UI) Winner(name string) {
    fmt.Printf("THE WINNER IS: %s", name)
}

func (ui UI) Draw() {
    fmt.Printf("DRAW!")
}
