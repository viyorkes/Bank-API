package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDB struct{
db *sql.DB
}


func (d CustomerRepositoryDB) FindAll() ([]Customer, error){

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.db.Query(findAllSql)

	if err != nil{
		log.Println("error while querying customer table" + err.Error())

	}

	customers := make([]Customer,0)

	for rows.Next(){
		var c  Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateOfBirth, &c.Zipcode, &c.status)

		if err != nil{
			log.Println("error while scanning customer table" + err.Error())

		}

		customers = append(customers, c)

	}
	return customers,nil;

}


func NewCustomerRepositoryDB()CustomerRepositoryDB{

	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}


	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDB{db}


}

