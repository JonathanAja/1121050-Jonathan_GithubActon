package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsersGorm(w http.ResponseWriter, r *http.Request) {
	db := connectGorm()
	var users []User
	db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	if len(users) < 100 {
		var response UsersResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = users
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Error Array Size Not Correct"
		json.NewEncoder(w).Encode(response)
	}

}
func InsertUserGorm(w http.ResponseWriter, r *http.Request) {
	db := connectGorm()
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")

	var users User
	users.Name = name
	users.Age = age
	users.Address = address

	db.Create(&users)

	var response UserResponse
	response.Data.Age = age
	response.Data.Name = name
	response.Data.Address = address

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"
	name := r.URL.Query()["name"]
	age := r.URL.Query()["age"]
	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name ='" + name[0] + "'"
	}
	if age != nil {
		if name != nil {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age='" + age[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		} else {
			users = append(users, user)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if len(users) < 10 {
		var response UsersResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = users
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Error Array Size Not Correct"
		json.NewEncoder(w).Encode(response)
	}

}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")

	_, errQuery := db.Exec("INSERT INTO users(name,age,address)values (?,?,?)",
		name,
		age,
		address,
	)
	var response UserResponse
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	userId := vars["user_id"]
	fmt.Println(userId)

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		userId,
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	userId := vars["user_id"]
	fmt.Println(userId)

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		userId,
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	userId, _ := strconv.Atoi(r.Form.Get("id"))
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	fmt.Println(userId)
	_, errQuery := db.Exec("UPDATE users SET name = ?, age = ?, address = ? WHERE id = ?",
		name,
		age,
		address,
		userId,
	)
	var response UserResponse

	response.Data.Age = age
	response.Data.Name = name
	response.Data.Address = address
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

func GetDetailUserTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT t.id,t.quantity,u.id,u.name,u.age,u.address,p.id,p.name,p.price FROM users u JOIN transactions t ON u.id =t.userID JOIN products p ON p.id=t.productid"
	userID := r.URL.Query()["userID"]
	if userID != nil {
		fmt.Println(userID[0])
		query += " WHERE userID ='" + userID[0] + "'"
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var td DetailTrannsaction
	var tds []DetailTrannsaction
	for rows.Next() {
		if err := rows.Scan(&td.ID, &td.Quantity, &td.User.ID, &td.User.Name, &td.User.Age, &td.User.Address,
			&td.Products.ID, &td.Products.Name, &td.Products.Price); err != nil {
			log.Println(err)
			return
		} else {
			tds = append(tds, td)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if len(tds) < 5 {
		var response DetailTransactionsResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = tds
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Error Array Size Not Correct"
		json.NewEncoder(w).Encode(response)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response UsersResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	name := r.URL.Query()["name"]
	row := db.QueryRow("SELECT * FROM users WHERE name=?", name[0])
	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
		sendErrorResponse(w, "error")
	} else {
		user.UserType = 0
		generateToken(w, user.ID, user.Name, user.UserType)
		sendSuccessResponse(w)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUserToken(w)
	var response UserResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
