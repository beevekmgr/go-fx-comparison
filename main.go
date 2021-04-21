package main

import (
	"GoFxvsNon/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server.New(mux)

	http.ListenAndServe(":8080", mux)
}
