# Server

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
