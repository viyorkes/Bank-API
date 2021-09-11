package app

import (
	"encoding/json"
	"github.com/viyorkes/Bank-API/service"
	"net/http"
)


type Customer struct{

	Name string `json:"full_name"`
	City string `json:"city"`
	Zipcode  string `json:"zip_code"`
}

type  CustomerHandlers struct{

	service service.CustomerService
}



func (ch* CustomerHandlers)getAllCustomers(w http.ResponseWriter, r *http.Request){

	//customers :=[]Customer{
	//	{"teste1", "teste-city","00000009"},
	//	{"teste2", "teste-city2","00000008"},
	//
	//}

	customers, _ :=ch.service.GetAllCustomer()
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(customers)

}





