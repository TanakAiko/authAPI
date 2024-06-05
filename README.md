# authAPI
An API for the authentification of a user (client)

## Before starting
import some package: 
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) : `go get golang.org/x/crypto/bcrypt`
- [sqlite3](github.com/mattn/go-sqlite3) : `github.com/mattn/go-sqlite3`
- [gorilla websocket](https://pkg.go.dev/github.com/gorilla/websocket) : `go get github.com/gorilla/websocket`
- [UUID](https://github.com/gofrs/uuid) : `go get github.com/google/uuid`

## Note
- ### register
The stucture for the body of the request (Register) is the following :
```
{
    action: "register"
    body: {
	    nickname  string
	    age       int   
	    gender    string
	    firstName string
	    lastName  string
	    email     string
	    password  string
    }
}
```

- ### login
The stucture for the body of the request (Login) is the following :
```
{
    action: "login"
    body: {
	    identifier  string
	    password    string
    }
}
```

## ToDo list
- Have to handle the not allowed methode
- test a register with the same email and nickname
- Hanve to make a script for the download of the dependencies
- 



## Testing
- ### register
```
curl -X POST http://localhost:8080/ -d '{
  "action":"register", 
  "body": 
    { 
      "nickname": "exampleNickname",
      "age": 30,
      "gender": "male",
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@example.com",
      "password": "hashedPassword"
    }
}' -H "Content-Type: application/json"
```

- ### login
```
curl -X POST http://localhost:8080/ -d '{
  "action":"login", 
  "body": 
    { 
      "identifier": "john.doe@example.com",
      "password": "hashedPassword"
    }
}' -H "Content-Type: application/json"
```