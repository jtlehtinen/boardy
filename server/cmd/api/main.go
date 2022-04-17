package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server hello"))
}

type config struct {
	port int
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 5001, "server port")
	flag.Parse()

	router := http.NewServeMux()
	router.HandleFunc("/", index)

	addr := fmt.Sprintf("localhost:%d", cfg.port)

	fmt.Printf("server starting at %s\n", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
