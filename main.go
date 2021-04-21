package main

import (
	"net/http"

	"github.com/beevekmgr/go-fx-comparison/server"
)

func main() {
	mux := http.NewServeMux()
	server.New(mux)

	http.ListenAndServe(":8080", mux)
}
