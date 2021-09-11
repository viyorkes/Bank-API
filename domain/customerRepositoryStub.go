package domain




type CustomerRepositoryStub struct {

	customers []Customer
}


func (s CustomerRepositoryStub) FindAll() ([]Customer,error){

	return s.customers, nil
}

func NewCustomerRepositoryStub()  CustomerRepositoryStub{
	customers := []Customer{

		{"100","victor","Sao Paulo","00000999","18/07/1992","ok"},
	}

	return CustomerRepositoryStub{customers}
}
