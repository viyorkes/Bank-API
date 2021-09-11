package app

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start(){


	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}",getCustomerById).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000",router))
}

func getCustomerById(w http.ResponseWriter,r*http.Request){

	vars:=mux.Vars(r)
	fmt.Fprint(w,vars["customer_id"])

}