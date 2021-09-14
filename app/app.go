package app

import (
	"github.com/viyorkes/Bank-API/domain"
	"github.com/viyorkes/Bank-API/service"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start(){


	router := mux.NewRouter()

	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}


	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)


	log.Fatal(http.ListenAndServe("localhost:8000",router))
}

