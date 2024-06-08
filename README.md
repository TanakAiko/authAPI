# authAPI
An API for the authentification of a user (client)

## Before starting
import some package: 
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) : `go get golang.org/x/crypto/bcrypt`
- [sqlite3](github.com/mattn/go-sqlite3) : `github.com/mattn/go-sqlite3`
- [gorilla websocket](https://pkg.go.dev/github.com/gorilla/websocket) : `go get github.com/gorilla/websocket`
- [UUID](https://github.com/gofrs/uuid) : `go get github.com/google/uuid`

## Note
The stucture of the body of the request is the following :
- ### register
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
```
{
    action: "login"
    body: {
	    identifier  string
	    password    string
    }
}
```

- ### authorized
```
{
    action: "authorized"
    body: {
	    sessionID  string
    }
}
```


## ToDo list
- Have to handle the not allowed methode ❌
- test a register with the same email and nickname ✅
- Hanve to make a script for the download of the dependencies ❌ (maybe the docker compose will handle it)
- 



## Testing
### register
- #### request
Execute the following command :
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
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusCreated (201)
- body    : "New user created"
```

### login
- #### request
Execute the following command :
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
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusOK (200)
- body    : the user data
```

### authorized
- #### request
Execute the following command :
```
curl -X POST http://localhost:8080/ -d '{
  "action":"authorized", 
  "body": 
    { 
      "sessionID": "6a09a3da-26ee-4b35-870c-d7a4f22f939c"
    }
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusAccepted (202)
- body    : "The session is valid"
```

