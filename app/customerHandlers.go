package app

import (
	"encoding/json"
	"github.com/viyorkes/Bank-API/service"
	"github.com/gorilla/mux"
	"net/http"
)






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
		writeResponse(w,err.Code,err.AsMessage())
	}else{
		writeResponse(w,http.StatusOK, customer)
		}
	}


	func writeResponse(w http.ResponseWriter,code int,data interface{}){
		w.Header().Add("Content-Type","application/json")
		w.WriteHeader(code)
		if err:= json.NewEncoder(w).Encode(data); err !=nil{
			panic(err)
		}

	}







