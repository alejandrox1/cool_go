package main

import (
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
)

var hmacSecret = []byte("secret")


func main() {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "foo": "bar",
        "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
    })
    tokenString, err := token.SignedString(hmacSecret)


    token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return hmacSecret, nil
    })


    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        fmt.Println(claims["foo"], claims["nbf"])
    } else {
        fmt.Println(err)
    }
}
