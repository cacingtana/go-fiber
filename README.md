# Implement CRUD Data Go and Mysql using Authentication & Authorization

Microservices with Go using fiber Framework. The code implementation was inspired by port and adapter pattern or known as hexagonal:
The hexagonal architecture is based on three principles and techniques:

- Explicitly separate User-Side, Business Logic, and Server-Side
- Dependencies are going from User-Side and Server-Side to the Business Logic
- We isolate the boundaries by using Ports and Adapters

<br>

## Data initialization
To describe about how port and adapter interaction, this example will have databases supported. There are MySQL using gorm as library.

## How To Consume The API

    //Exoprt DDL sql file first

    //Example public endpoint
    //JSON Example payload to register user
    {
    	"email": "epankbole@gmail.com",
    	"password": "12345",
        "level": 1
    }
    
    POST Method "http://localhost:8000/login" //to access endpoint register
    
    //JSON Example payload to Login.
    //this is user dengan credential
    {
    		"email": "epankbole@gmail.com",
    		"password": "12345"
    }
    
    POST Method "http://localhost:8000/login" 
    //to access endpoint login
    //also attach the token that has been generated
    
    //Example private endpoint with Authentication
    //to endpoint diary, also attach the token that has been generated when access endpoint login
    GET    Method "http://localhost:8000/api/diary", to get all diary
    GET    Method "http://localhost:8000/api/diary/1", to get diary by id
    PUT    Method "http://localhost:8000/api/diary/1", to update diary by id
    DELETE Method "http://localhost:8000//api/diary/1", to delete user by id
    
    POST   Method "http://localhost:8000/api/diary", to create new diary
    //this is Example payload to crete diary
    {
    	"title": "Golang",
    	"body": "Belajar Pemrograman Golang"
    }


## Unit Testing
Unit testing with testify library



