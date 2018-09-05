/*
 * curl --noproxy "*" -H "Content-Type: application/json" -x POST -d '{"description": "hi"}' http://775ab64a.ngrok.io/products/hover-shooters/feedback
 */
package main

import (
	"net/http"
	"os"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Secret for jwt.
var mySigningKey = []byte("secret-key")

// Product contains information about VR experiences.
type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

var products = []Product{
	{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Different hoverboards"},
	{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea"},
	{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Ride a T-Rex"},
	{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars"},
	{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Master the art of archery"},
	{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the world"},
}

// jwtMiddleware will handle requests and verify the jwt passed.
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func main() {
	r := mux.NewRouter()

	// Serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// Serve static assets from /static/{file} route.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/status", statusHandler).Methods("GET")

	r.HandleFunc("/get-token", getTokenHandler).Methods("GET")

	r.Handle("/products",
		jwtMiddleware.Handler(http.HandlerFunc(productsHandler))).Methods("GET")
	r.Handle("/products/{slug}/feedback",
		jwtMiddleware.Handler(http.HandlerFunc(addFeedbackHandler))).Methods("POST")

	http.ListenAndServe("0.0.0.0:3000", handlers.LoggingHandler(os.Stdout, r))
}
