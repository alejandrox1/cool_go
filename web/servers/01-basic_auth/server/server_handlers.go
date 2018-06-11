package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

// Credentials models the structure of a user in the request body and the db.
type Credentials struct {
    Password string `json:"password", db:"password"`
    Username string `json:"username", db:"username"`
}

// Signup 
func Signup(w http.ResponseWriter, r *http.Request) {
    // Parse and decode the request body into a new instance of Credentials.
    creds := &Credentials{}
    err := json.NewDecoder(r.Body).Decode(creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Printf("Signup error while decoding request: %s\n", err)
        return
    }

    // Salt and hash the password using the bcrypt algorithm. The second
    // argument is the cost of hasing, which we arbitrarily set as 8.
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
    if err != nil {
        fmt.Printf("Signup error while hashing password: %s\n", err)
        return
    }

    // Insert the username along with its hashed password into the database.
     _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Printf("Signup error while storing Credentials into database: %s\n", err)
        return
    }

    // Default status 200 is sent back.
}


// Signin
func Signin(w http.ResponseWriter, r *http.Request) {
    // Parse and decode the request body into a new instance of Credentials.
    creds := &Credentials{}
    err := json.NewDecoder(r.Body).Decode(creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Printf("Signup error while decoding request: %s\n", err)
        return
    }

    // Get the existing entry in the database for the specified user.
    storedCreds := &Credentials{}
    result := db.QueryRow("select password from users where username=$1", creds.Username)
    err = result.Scan(&storedCreds.Password)
    if err != nil {
        fmt.Printf("Signin error while looking for existing user in database: %s\n", err)

        if err == sql.ErrNoRows {
            w.WriteHeader(http.StatusUnauthorized)
            return
        } else {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }

    // Compare passwords.
    err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password))
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Printf("Signin error while comparing password: %s\n", err)
        return
    }

    // Default status is sent back.
}
