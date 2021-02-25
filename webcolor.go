package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const template = `<!DOCTYPE html>
<html>
  <body style="background: %s;">
    <h1 style="background: white;">This is pod %s.</h1>
  </body>
</html>`

func serve(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("HOSTNAME")
	color := strings.SplitN(hostname, "-", 2)[0]
	html := fmt.Sprintf(template, color, hostname)
	w.Write([]byte(html))
}

func main() {
	fmt.Printf("Starting webcolor listening on port 8000.\n")
	http.HandleFunc("/", serve)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
