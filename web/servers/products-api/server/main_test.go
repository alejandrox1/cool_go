package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var app App

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func ensureTablesExist() {
	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE FROM products")
	app.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		app.DB.Exec("INSERT INTO products(name, price) VALUES($1, $2)",
			strconv.Itoa(i), (i+1.0)*10)
	}
}

func executeRequest(r *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, r)

	return recorder
}

func TestMain(m *testing.M) {
	app = App{}
	app.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTablesExist()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Error(err)
	}

	response := executeRequest(req)
	if response.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s\n", body)
	}
}

func TestGetNonExistentProduct(t *testing.T) {
	clearTable()

	req, err := http.NewRequest("GET", "/products/11", nil)
	if err != nil {
		t.Error(err)
	}

	response := executeRequest(req)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Product not found" {
		t.Errorf(
			"Expected the 'error' key of the response to be 'Product not found'. Got '%s'\n",
			m["error"])
	}
}

func TestCreateProduct(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"test product", "price":11.22}`)
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
	if err != nil {
		t.Error(err)
	}

	response := executeRequest(req)

	if response.Code != http.StatusCreated {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusCreated, response.Code)
	}

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test product" {
		t.Errorf("Expected product name to be `test product`. Got '%v'\n", m["name"])
	}
	if m["price"] != 11.22 {
		t.Errorf("Expected product price to be `11.22`. Got `%v`\n", m["price"])
	}
	// json unmarshalling converts numbers to floats.
	if m["id"] != 1.0 {
		t.Errorf("Expected product id to be '1'. Got '%v'\n", m["id"])
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Error(err)
	}

	response := executeRequest(req)
	if response.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}
}

func TestUpdateProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Error(err)
	}

	response := executeRequest(req)

	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	payload := []byte(`{"name":"test product - updated name", "price":11.22}`)
	req, err = http.NewRequest("PUT", "/products/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Error(err)
	}

	response = executeRequest(req)
	if response.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalProduct["id"] {
		t.Errorf("Expected id to remain the same (%v). Got %v\n", originalProduct["id"], m["id"])
	}
	if m["name"] != "test product - updated name" {
		t.Errorf("Expected '%s' name. Got '%v'\n", originalProduct["name"], m["name"])
	}
	if m["price"] != 11.22 {
		t.Errorf("Expected price '%v'. Got '%v'\n", originalProduct["price"], m["price"])
	}
}

func TestDeleteProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Error(err)
	}

	response := executeRequest(req)
	if response.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}

	req, err = http.NewRequest("DELETE", "/products/1", nil)
	if err != nil {
		t.Error(err)
	}

	response = executeRequest(req)
	if response.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
	}

	req, err = http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Error(err)
	}

	response = executeRequest(req)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.Code)
	}
}
