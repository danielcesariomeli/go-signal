package main

import (
	"log"
	"net/http"
	"time"

	"github.com/danielcesariomeli/go-signal/cmd/api/handler"
	"github.com/danielcesariomeli/go-signal/internal/counter"
	"github.com/gorilla/mux"
)

func main() {

	service := counter.NewService(make(chan bool), false)
	handler := handler.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/start", handler.HandleStart).Methods("POST")
	router.HandleFunc("/stop", handler.HandleStop).Methods("POST")
	runServer(router)
}

func runServer(router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
