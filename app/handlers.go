package app

import (
	"encoding/json"
	"fmt"
	"github.com/viyorkes/Bank-API/service"
	"github.com/gorilla/mux"
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


	customers, _ :=ch.service.GetAllCustomer()
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(customers)

}



func (ch *CustomerHandlers)getCustomer(w http.ResponseWriter, r *http.Request){


	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {

		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
	}else{

		w.Header().Add("Content-Type","application/json")
		json.NewEncoder(w).Encode(customer)
		}
	}







