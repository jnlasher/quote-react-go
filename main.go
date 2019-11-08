package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// JSONMessage Struct for formatting error responses back to the client
type JSONMessage struct {
	Message string `json:"message"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/quotes", getRandomQuote).Methods("GET")
	router.HandleFunc("/quotes", createQuote).Methods("POST")

	fmt.Println("Server started on PORT 8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic("ERR: " + err.Error())
	}
}

func getRandomQuote(writer http.ResponseWriter, request *http.Request) {
	quoteStruct, err := FetchRandomQuote()
	if err != nil {
		fmt.Println("Got error from database")
		writeJSON(writer, JSONMessage{err.Error()}, http.StatusUnprocessableEntity)
		return
	}
	writeJSON(writer, quoteStruct, http.StatusOK)
}

func createQuote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err := WriteNewQuote(request)
	if err != nil {
		writeJSON(writer, JSONMessage{err.Error()}, http.StatusBadRequest)
		return
	}
	writeJSON(writer, JSONMessage{"New quote added"}, http.StatusCreated)
}

func writeJSON(writer http.ResponseWriter, data interface{}, status int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(status)

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}

	_, err = fmt.Fprint(writer, string(jsonBytes))
	if err != nil {
		log.Panic(err)
	}
}
