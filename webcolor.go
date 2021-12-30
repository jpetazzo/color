package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
        "runtime"
	"strings"
)

const template = `<!DOCTYPE html>` + "\r" +
	`<html>` + "\r" +
	`<body style="background: %s; text-align: center;">` + "\r" +
	`<div style="padding: 4em;"></div>` + "\r" +
	`<span style="padding: 4em; background: %s;">` + "\r" +
	`<span style="padding: 2px; background: white;">` + "\r" +
	`%sThis is %s on %s/%s, serving %s for %s.` + "\n" +
	`</span>` + "\r" +
	`</span>` + "\r" +
	`</body>` + "\r" +
	`</html>` + "\r"

func getHostname() string {
	return os.Getenv("HOSTNAME")
}

func getNamespace() string {
	{
		namespace := os.Getenv("NAMESPACE")
		if namespace != "" {
			return namespace
		}
	}
	{
		namespace, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
		if err == nil {
			return string(namespace)
		}
	}
	return ""
}

func getCircle(color string) string {
	circles := map[string]string{
		"red":    "🔴",
		"orange": "🟠",
		"yellow": "🟡",
		"green":  "🟢",
		"blue":   "🔵",
		"purple": "🟣",
		"brown":  "🟤",
		"black":  "⚫",
		"white":  "⚪",
	}
	circle, exists := circles[color]
	if exists {
		return circle
	} else {
		return ""
	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	hostname := getHostname()
	namespace := getNamespace()
	displayName := ""
	if namespace == "" {
		displayName = hostname
	} else {
		displayName = "pod " + string(namespace) + "/" + hostname
	}
	podColor := strings.SplitN(hostname, "-", 2)[0]
	circles := getCircle(namespace) + getCircle(podColor)
	html := fmt.Sprintf(template, namespace, podColor, circles, displayName, runtime.GOOS, runtime.GOARCH, r.URL, r.RemoteAddr)
	w.Write([]byte(html))
	fmt.Printf("%s %s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.Proto, r.Header["User-Agent"])
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
