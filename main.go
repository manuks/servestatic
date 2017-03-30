package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
)

func main() {

	var sourceVar string
	flag.StringVar(&sourceVar, "s", "/var/www/static", "Absolute path to folder")
	flag.Parse()
	fmt.Printf("Serving static files from %s\n", sourceVar)

	goji.Handle("/*", http.FileServer(http.Dir(sourceVar)))

	goji.Serve()
}
