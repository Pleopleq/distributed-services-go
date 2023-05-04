package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(addr string) *http.Server {
	httpServer := NewHTTPServer()
	r := mux.NewRouter()

	r.HandleFunc("/", httpServer.handleProduce).Methods("POST")
	r.HandleFunc("/", httpServer.handleConsume).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
