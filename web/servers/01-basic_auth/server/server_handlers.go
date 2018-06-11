package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "net/http"

    _ "github.com/lib/pq"
)

// Credentials models the structure of a user in the request body and the db.
type Credentials struct {
    Password string `json:"password", db:"password"`
    Username string `json:"username", db:"username"`
}

// Signup 
func Signup(w http.ResponseWriter, r *http.Request){
    // Parse and decode the request body into a new instance of Credentials.
    creds := &Credentials{}
    err := json.NewDecoder(r.Body).Decode(creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Println(err)
        return
    }

    // Salt and hash the password using the bcrypt algorithm. The second
    // argument is the cost of hasing, which we arbitrarily set as 8.
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Insert the username along with its hashed password into the database.
     _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Println(err)
        return
    }

    // Default status 200 is sent back.
}
