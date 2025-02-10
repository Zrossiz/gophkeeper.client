package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	Version   string
	BuildTime string
)

const targetServer = "http://localhost:8080"

func handleProxy(w http.ResponseWriter, r *http.Request) {
	targetURL, err := url.Parse(targetServer)
	if err != nil {
		http.Error(w, "Bad target URL", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(w, r)
}

func main() {
	showVersion := flag.Bool("version", false, "Show version and build time")
	flag.Parse()

	if *showVersion {
		fmt.Printf("GophKeeper CLI\nVersion: %s\nBuild Time: %s\n", Version, BuildTime)
		os.Exit(0)
	}

	port := ":3000"
	http.HandleFunc("/", handleProxy)

	log.Println("Proxy server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
