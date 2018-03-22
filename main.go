package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/desmondrawls/rock-paper-scissors/web_ui"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.Parse()

	handler := &web_ui.Handler{}
	url := fmt.Sprintf("127.0.0.1:%s", port)
	fmt.Fprintf(os.Stderr, "Server: http://%s\n", url)
	log.Fatal(http.ListenAndServe(url, handler))
}
