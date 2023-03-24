package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM transactions"
	userID := r.URL.Query()["userID"]
	productID := r.URL.Query()["productID"]
	if userID != nil {
		fmt.Println(userID[0])
		query += " WHERE userID ='" + userID[0] + "'"
	}
	if productID != nil {
		if userID[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " productID='" + productID[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var transaction Transactions
	var transactions []Transactions
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity); err != nil {
			log.Println(err)
			return
		} else {
			transactions = append(transactions, transaction)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if len(transactions) < 5 {
		var response TransactionsResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = transactions
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Error Array Size Not Correct"
		json.NewEncoder(w).Encode(response)
	}
}

func InsertTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	userID := r.Form.Get("userID")
	productID, _ := strconv.Atoi(r.Form.Get("productID"))
	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))
	var transaction Transactions
	var counter int

	errCount := db.QueryRow("Select count(*) from products where id = ?", transaction.ProductID).Scan(&counter)

	if errCount != nil {
		fmt.Println(err)
		sendErrorResponse(w, "failed")
	} else {
		if counter == 0 {
			_, errProduct := db.Exec("INSERT INTO products(Id,name,price) values (?,?,?)", productID, "", "")
			if errProduct != nil {
				sendErrorResponse(w, "failed")
				return
			}
		}
	}

	_, errQuery := db.Exec("INSERT INTO transactions(userID,productID, quantity)values (?,?,?)",
		userID,
		productID,
		quantity,
	)
	var response ProductResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Insert Failed"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteTransactions(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	transactionId := vars["transactions_id"]
	fmt.Println(transactionId)

	_, errQuery := db.Exec("DELETE FROM transactions WHERE id=?",
		transactionId,
	)
	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed!"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	id, _ := strconv.Atoi(r.Form.Get("id"))
	userId, _ := strconv.Atoi(r.Form.Get("userID"))
	productID, _ := strconv.Atoi(r.Form.Get("productID"))
	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))
	fmt.Println(userId)
	_, errQuery := db.Exec("UPDATE transactions SET userID = ?, productID = ?, quantity = ? WHERE id = ?",
		userId,
		productID,
		quantity,
		id,
	)
	var response TransactionResponse

	response.Data.UserID = userId
	response.Data.ProductID = productID
	response.Data.Quantity = quantity
	response.Data.ID = id

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Update Failed"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
