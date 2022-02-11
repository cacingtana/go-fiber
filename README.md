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

    //Public endpoint
    //JSON Example payload to register user
    {
    		"email": "epankbole@gmail.com",
    		"password": "12345",
        "level": 1
    }
    
    POST Method "/login" //to access endpoint register
    
    //JSON Example payload to Login.
    //this is user dengan credential / role_id = 1
    {
    		"email": "epankbole@gmail.com",
    		"password": "12345"
    }
    
    POST Method "/login" 
    //to access endpoint login
    //also attach the token that has been generated
    
    //Private endpoint with Authentication
    //to endpoint diary, also attach the token that has been generated when access endpoint login
    GET    Method "localhost:8000/api/diary", to get all diary
    GET    Method "localhost:8000/api/diary/1", to get diary by id
    POST   Method "localhost:8000/api/diary", to create new diary
    PUT    Method "localhost:8000/api/diary/1", to update diary by id
    DELETE Method "localhost:8000/api/diary/1", to delete user by id



