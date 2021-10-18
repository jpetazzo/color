package main

import (
	"fmt"
	"net/http"
	"os"
        "runtime"
	"strings"
)

const template = `<!DOCTYPE html>` + "\r" +
	`<html>` + "\r" +
	`<body style="background: %s;">` + "\r" +
	`<p style="background: white;">` + "\r" +
	`This is pod %s, running on %s/%s, serving %s for %s.` + "\n" +
	`</p>` + "\r" +
	`</body>` + "\r" +
	`</html>` + "\r"

func serve(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("HOSTNAME")
	color := strings.SplitN(hostname, "-", 2)[0]
	html := fmt.Sprintf(template, color, hostname, runtime.GOOS, runtime.GOARCH, r.URL, r.RemoteAddr)
	w.Write([]byte(html))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	fmt.Printf("Starting HTTP server (%s/%s) listening on port %s.\n", runtime.GOOS, runtime.GOARCH, port)
	http.HandleFunc("/", serve)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Couldn't create server on port %s (%s).\n", port, err)
		fmt.Printf("You can change the port by setting the PORT environment variable.\n")
		os.Exit(1)
	}
}
