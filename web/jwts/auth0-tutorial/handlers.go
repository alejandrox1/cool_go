package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// statusHandler responds with a string advertising the server is running.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API is up and running")
}

// getTokenHandler
func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Create a token.
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims.
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims.
	claims["admin"] = true
	claims["name"] = "alejandrox1"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign token with our secret.
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Printf("getTokenHandler error: %s\n", err)
	}

	// Send token as response.
	fmt.Fprintln(w, tokenString)
}

// productsHandler formats the products into json and sends them as response.
func productsHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := json.MarshalIndent(products, "", "\t")
	if err != nil {
		log.Printf("productHandler error: %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(payload))
}

// updateProduct updates the products slice given a product struct.
func updateProduct(key int, product Product) {
	if product.Id != 0 {
		products[key].Id = product.Id
	}
	if product.Name != "" {
		products[key].Name = product.Name
	}
	if product.Description != "" {
		products[key].Description = product.Description
	}
}

// addFeedbackHandler
func addFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	var found bool

	// Get product (/products/{slug}/feedback).
	vars := mux.Vars(r)
	slug := vars["slug"]

	w.Header().Set("Content-Type", "application/json")
	// Get the struct corresponding to the product.
	for i, p := range products {
		if p.Slug == slug {
			// Save the request's data into to a Product struct.
			if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
				log.Printf("addFeedbackHandler error: %s\n", err)
			}
			r.Body.Close()

			// Update the products slice.
			updateProduct(i, product)

			payload, err := json.MarshalIndent(products[i], "", "\t")
			if err != nil {
				log.Printf("addFeedbackHandler error: %s\n", err)
			}
			found = true
			fmt.Fprint(w, string(payload))
		}
	}

	if found != true {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Product not found")
	}
}
