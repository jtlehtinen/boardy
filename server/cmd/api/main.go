package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.1"

type config struct {
	port int
}

type application struct {
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 5001, "server port")
	flag.Parse()

	app := &application{}

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", cfg.port),
		Handler: app.routes(),
	}

	fmt.Fprintf(os.Stdout, "Server starting at %s\n", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
