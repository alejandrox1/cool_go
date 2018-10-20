package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App creates a connection with a database and instantiates the web server.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize creates a connection to the database and register the app's
// rotues.
func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	tries := 5
	delay := time.Duration(500) * time.Millisecond
	for ; tries >= 0; tries, delay = tries-1, delay*2 {
		if err = a.DB.Ping(); err != nil {
			break
		} else if err != nil && tries == 0 {
			log.Fatal(err)
		}

		time.Sleep(delay)
	}

	// Initialize router.
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

// InitializeRoutes defines the apps endpoint's.
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/products", a.createProduct).Methods("POST")
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.getProduct).Methods("GET")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.updateProduct).Methods("PUT")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}

// Run runs the application.
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		msg := fmt.Sprintf("Invalid request payload: %s\n", err)
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}
	defer r.Body.Close()

	if err := p.createProduct(a.DB); err != nil {
		msg := fmt.Sprintf("Error creating product: %s\n", err)
		respondWithError(w, http.StatusInternalServerError, msg)
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := product{ID: id}
	if err := p.getProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	queryCount := r.URL.Query().Get("count")
	if queryCount == "" {
		queryCount = "0"
	}
	queryStart := r.URL.Query().Get("start")
	if queryStart == "" {
		queryStart = "0"
	}

	count, err := strconv.Atoi(queryCount)
	if err != nil {
		msg := fmt.Sprintf("Error parsing query string (count = '%d'): %s\n", count, err)
		respondWithError(w, http.StatusInternalServerError, msg)
		return
	}
	start, err := strconv.Atoi(queryStart)
	if err != nil {
		msg := fmt.Sprintf("Error parsing query string (start = '%d'): %s\n", count, err)
		respondWithError(w, http.StatusInternalServerError, msg)
		return
	}

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getProducts(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := fmt.Sprintf("Invalid product ID: %s\n", err)
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}

	var p product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		msg := fmt.Sprintf("Invalid request payload: %s\n", err)
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.updateProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := fmt.Sprintf("Invalid product ID: %s\n", err)
		respondWithError(w, http.StatusBadRequest, msg)
		return
	}

	p := product{ID: id}
	if err := p.deleteProduct(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
