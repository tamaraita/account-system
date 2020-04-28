package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type responseObject struct {
	Message string `json:"message"`
}

type transaction struct {
	ID            string  `json:"id,omitempty"`
	Type          *string `json:"type"`
	Amount        *int    `json:"amount"`
	EffectiveDate string  `json:"effectiveDate,omitempty"`
}

type transactions []transaction

var balance int
var transactionHistory transactions

func updateBalance(newTransaction transaction) (string, int) {
	dt := time.Now()
	ttype, amount := *newTransaction.Type, *newTransaction.Amount
	var response = "Unknown error"
	var header = http.StatusInternalServerError
	newTransaction.ID = uuid.New().String()
	newTransaction.EffectiveDate = dt.Format("2006-01-02T15:04:05-0700")
	switch ttype {
	case "credit":
		if amount <= balance {
			balance -= amount
			response = "New balance $" + strconv.Itoa(balance)
			header = http.StatusOK
			transactionHistory = append(transactionHistory, newTransaction)
		} else {
			response = "Denied"
			header = http.StatusBadRequest
		}
	case "debit":
		balance += amount
		response = "New balance $" + strconv.Itoa(balance)
		header = http.StatusOK
		transactionHistory = append(transactionHistory, newTransaction)
	default:
		log.Fatal("Unknown transaction type")
	}
	return response, header
}

func home(w http.ResponseWriter, r *http.Request) {
	var response responseObject
	response.Message = "Account balance $" + strconv.Itoa(balance)
	json.NewEncoder(w).Encode(response)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactionHistory)
}

func getTransactionByID(w http.ResponseWriter, r *http.Request) {
	found := false
	transactionID := mux.Vars(r)["id"]
	for _, singleTransaction := range transactionHistory {
		if singleTransaction.ID == transactionID {
			json.NewEncoder(w).Encode(singleTransaction)
			found = true
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		var response responseObject
		response.Message = "No transaction found"
		json.NewEncoder(w).Encode(response)
	}
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction transaction
	var response responseObject
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = "Malformed request."
	}
	json.Unmarshal(reqBody, &newTransaction)
	// Validations
	if newTransaction.Type == nil || newTransaction.Amount == nil {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = "Please, fill the required information in the body of the request."
	} else if *newTransaction.Type != "credit" && *newTransaction.Type != "debit" {
		w.WriteHeader(http.StatusBadRequest)
		response.Message = "Transaction type is not correct."
	} else {
		msg, header := updateBalance(newTransaction)
		w.WriteHeader(header)
		response.Message = msg
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// initEvents()
	router := mux.NewRouter().StrictSlash(true)
	balance = 0
	transactionHistory = nil
	router.HandleFunc("/", home)
	router.HandleFunc("/transactions", getTransactions).Methods("GET")
	router.HandleFunc("/transactions", createTransaction).Methods("POST")
	router.HandleFunc("/transactions/{id}", getTransactionByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
