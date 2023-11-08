package main

//my attempt to write vanilla router in go, which technically should be easy

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Msg string `json:"msg"`
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hi", hiHandler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Server is starting on port 8080...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed to start: %s", err)
		}
	}()

	// Wait for an interrupt
	<-make(chan struct{})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling hello request...")
	time.Sleep(5 * time.Second)
	w.Write([]byte("Hello, World!"))
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling hi request...")
	response := Message{Msg: "hi"}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding JSON: %s", err)
	}
}
