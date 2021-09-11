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

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)


	log.Fatal(http.ListenAndServe("localhost:8000",router))
}

