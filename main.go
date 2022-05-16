package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/itsfhz/flutter-golang-server/api"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/test", api.TestLink).Methods("GET")

	//expenses
	myRouter.HandleFunc("/expenses", api.GetAllExpenses).Methods("GET")
	myRouter.HandleFunc("/expense/{id}", api.GetExpense).Methods("GET")
	myRouter.HandleFunc("/expense", api.AddExpenses).Methods("POST")
	myRouter.HandleFunc("/expense/{id}", api.DeleteExpense).Methods("DELETE")
	myRouter.HandleFunc("/expense/{id}", api.UpdateExpense).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	f, err := os.OpenFile("flutter_golang_server_log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println("Server Started!")

	handleRequests()

	log.Println("Server Started. Ready to serve!")

}
