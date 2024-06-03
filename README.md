# authAPI
An API for the authentification of a user (client)

## Before starting
import some package: 
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) : `go get golang.org/x/crypto/bcrypt`
- [sqlite3](github.com/mattn/go-sqlite3) : `github.com/mattn/go-sqlite3`

## Note
The stucture for the body of the request is the following :
```
{
    action: string
    user: {
        Id        int
	    Nickname  string
	    Age       int   
	    Gender    string
	    FirstName string
	    LastName  string
	    Email     string
	    Password  string
    }
}
```

## ToDo list
- Have to handle the not allowed methode
- test a register with the same email and nickname
- 



## Testing
- ### register
```
curl -X POST http://localhost:8080/ -d '{
  "action":"register", 
  "user": 
    { 
      "Nickname": "exampleNickname",
      "Age": 30,
      "Gender": "male",
      "FirstName": "John",
      "LastName": "Doe",
      "Email": "john.doe@example.com",
      "Password": "hashedPassword"
    }
}' -H "Content-Type: application/json"
```
