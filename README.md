# Go Bank-API REST Example
A Simple RESTful API example with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library)

## Installation & Run
```bash
# Download this project
go get github.com/viyorke/Bank-Api
```

Before running API server, you should set the database config with yours or set the your database config
```go
func GetConfig() *Config {

	db,err:=sql.Open("mysql","root:yourPassWordDBHere@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}


	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDB{db}
	}
}
```

```bash
# Build and Run
cd gBank-Api
go build
./gBank-Api

# API Endpoint : http://127.0.0.1:3000
```

## Structure
```
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── projects.go  // APIs for Project model
│   │   └── tasks.go     // APIs for Task model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /projects
* `GET` : Get all ClientsBank
* `POST` : Create a new Client

#### /projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project


#### /Curl: Get All
curl --location --request GET 'http://localhost:8000/customers'


#### /Curl: Get ById
curl --location --request GET 'http://localhost:8000/customers/2000'


## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
