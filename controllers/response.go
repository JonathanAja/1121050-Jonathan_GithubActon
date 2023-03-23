package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Req(w http.ResponseWriter, r *http.Request, req interface{}) {
	fmt.Println("Success 200")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}
