package main

import (
    "log"
    "net/http"

    "github.com/desmondrawls/rock-paper-scissors/web_ui"
)

func main() {
    handler := &web_ui.Handler{}
    log.Fatal(http.ListenAndServe("127.0.0.1:8080", handler))
}
