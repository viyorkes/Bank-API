package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type Customer struct{

	Name string `json:"full_name"`
	City string `json:"city"`
	Zipcode  string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request){

	fmt.Fprint(w,"hello")



}



func getAllCustomers(w http.ResponseWriter, r *http.Request){

	customers :=[]Customer{
		{"teste1", "teste-city","00000009"},
		{"teste2", "teste-city2","00000008"},

	}

	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(customers)

}
