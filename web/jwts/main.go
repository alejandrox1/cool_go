/*
 * curl --noproxy "*" -H "Content-Type: application/json" -x POST -d '{"description": "hi"}' http://775ab64a.ngrok.io/products/hover-shooters/feedback
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/auth0-community/auth0"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	jose "gopkg.in/square/go-jose.v2"
)

var mySigningKey = []byte("secret-key")

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// Product contains information about VR experiences.
type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

var products = []Product{
	{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented")
}

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims.
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims.
	claims["admin"] = true
	claims["name"] = "alejandrox1"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Printf("getTokenHandler error: %s\n", err)
	}

	fmt.Fprintln(w, tokenString)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API is up and running")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := json.MarshalIndent(products, "", "\t")
	if err != nil {
		log.Printf("productHandler error: %s\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(payload))
}

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

func addFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	var found bool

	vars := mux.Vars(r)
	slug := vars["slug"]

	w.Header().Set("Content-Type", "application/json")
	for i, p := range products {
		if p.Slug == slug {
			if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
				log.Printf("addFeedbackHandler error: %s\n", err)
			}
			r.Body.Close()

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

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := []byte("{ATH0_API_SECRET}")
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{"{AUTH0_API_AUDIENCE}"}

		configuration := auth0.NewConfiguration(secretProvider, audience, "https://{AUTH0_DOMAIN}.auth0.com/", jose.HS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(r)
		if err != nil {
			log.Printf("Token is not valid: %s\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {
	r := mux.NewRouter()

	// Serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// Serve static assets from/static/{file} route.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/status", statusHandler).Methods("GET")

	r.HandleFunc("/get-token", getTokenHandler).Methods("GET")

	//r.Handle("/products", jwtMiddleware.Handler(http.HandlerFunc(productsHandler))).Methods("GET")
	//r.Handle("/products/{slug}/feedback", jwtMiddleware.Handler(http.HandlerFunc(addFeedbackHandler))).Methods("POST")
	r.Handle("/products", authMiddleware(http.HandlerFunc(productsHandler))).Methods("GET")
	r.Handle("/products/{slug}/feedback", authMiddleware(http.HandlerFunc(addFeedbackHandler))).Methods("POST")

	http.ListenAndServe("0.0.0.0:3000", handlers.LoggingHandler(os.Stdout, r))
}
