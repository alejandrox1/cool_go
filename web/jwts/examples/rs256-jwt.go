/*
    openssl genrsa -out private.pem 4096
    openssl rsa -in private.pem -outform PEM -pubout -out public.pem
*/
package main

import (
    "crypto/rsa"
    "fmt"
    "io/ioutil"
    "log"
    "time"

    "github.com/dgrijalva/jwt-go"
)

const (
    privKeyPath = "./private.pem"
    pubKeyPath = "./public.pem"
)

var (
    signKey *rsa.PrivateKey
    verifyKey *rsa.PublicKey
)

func init() {
    signBytes, err := ioutil.ReadFile(privKeyPath);
    if err != nil {
        log.Fatal(err)
    }

    verifyBytes, err := ioutil.ReadFile(pubKeyPath);
    if err != nil {
        log.Fatal(err)
    }

    signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes);
    if err != nil {
        log.Fatal(err)
    }

    verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes);
    if err != nil {
        log.Fatal(err)
    }
}


func main() {
    t := jwt.New(jwt.SigningMethodRS256)
    claims := make(jwt.MapClaims)

    claims["foo"] = "bar"
    claims["nbf"] = time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()
    claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
    claims["http://wso2.org/claims/role"] = "Internal/public_keys_admin"

    t.Claims = claims
    tokenString, err := t.SignedString(signKey)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\n%s\n", tokenString)

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return verifyKey, nil
    })
    if err != nil {
        log.Fatal(err)
    }


    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        fmt.Println(claims)
        fmt.Println(claims["foo"], claims["nbf"])
        fmt.Println(claims["http://wso2.org/claims/role"])
    } else {
        fmt.Println(err)
    }
}
