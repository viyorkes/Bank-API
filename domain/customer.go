package domain


type Customer struct{
    Id string
	Name string `json:"full_name"`
	City string `json:"city"`
	Zipcode  string `json:"zip_code"`
	DateOfBirth string
	status string
}


type CustomerRepository interface{

	FindAll() ([]Customer, error)
}

