package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM products"
	name := r.URL.Query()["name"]
	price := r.URL.Query()["price"]
	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name ='" + name[0] + "'"
	}
	if price != nil {
		if name[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " price='" + price[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}
	var product Products
	var products []Products
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		} else {
			products = append(products, product)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if len(products) < 5 {
		var response ProductsResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = products
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Error Array Size Not Correct"
		json.NewEncoder(w).Encode(response)
	}
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	price, _ := strconv.Atoi(r.Form.Get("price"))

	_, errQuery := db.Exec("INSERT INTO products(name,price)values (?,?)",
		name,
		price,
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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	productId := vars["products_id"]
	fmt.Println(productId)
	db.Exec("DELETE FROM transactions WHERE productid=?",
		productId,
	)
	_, errQuery := db.Exec("DELETE FROM products WHERE id=?",
		productId,
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

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	userId, _ := strconv.Atoi(r.Form.Get("id"))
	name := r.Form.Get("name")
	price, _ := strconv.Atoi(r.Form.Get("price"))
	fmt.Println(userId)
	_, errQuery := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?",
		name,
		price,
		userId,
	)
	var response ProductResponse
	response.Data.Price = price
	response.Data.Name = name
	response.Data.ID = userId

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
