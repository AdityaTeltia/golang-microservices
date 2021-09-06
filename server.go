package main

import (
	"context"
	"log"
	"microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api\t", log.LstdFlags)

	ph := handlers.NewProducts(l)

	// Serve Mux is a multiplex handler unlike other handlers

	
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}",ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/",ph.AddProducts)
	postRouter.Use(ph.MiddlewareValidateProduct)



	// sm.Handle("/products", ph).Methods("GET")

	// Created our own server
	// Why ?
	// For handling Timeouts
	// What ?
	// Basically if the user is not using our server for a while
	// User will automatically get disconnected , and this will be
	// helpful for saving resources

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Goroutine Function (syntax ~ go func )
	// Does not block the code

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Channels
	//https://gobyexample.com/channels
	// For graceful Termination
	// Means before terminating anwersing the already processed
	// requests and not take any further request so that user
	// do not face any problems
	// Remember the syntax
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved Terminate , graceful Shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)

}
