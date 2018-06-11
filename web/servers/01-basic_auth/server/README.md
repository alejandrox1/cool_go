# Server

Server based off [Password authentication and storage in Go
(Golang)](https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/).

In order to sign up and create a user account we will make a handler that
accepts a `POST` request with a json body of the form:
```
{
    "username": "username",
    "password": "userpassword"
}
```
The handler should return a status code `200` is the request was succesfully
caried out.

To test it out,
```
curl -v -X POST -H "Content-Type: application/json" -d '{"username": "me", "password": "mypass"}' http://localhost:8000/signup
```
