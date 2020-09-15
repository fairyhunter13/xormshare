package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

var (
	osSignal = make(chan os.Signal, 1)
)

func initSignal() {
	signal.Notify(osSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
}

func init() {
	initEngine()
	initHandler()
	initSignal()
}

func main() {
	server := &http.Server{
		Handler: mux,
		Addr:    ":6969",
	}
	go func() {
		<-osSignal
		server.Shutdown(context.Background())
		log.Println("Server has been stopped!")
	}()
	log.Println("Server has been started!")
	server.ListenAndServe()
}
