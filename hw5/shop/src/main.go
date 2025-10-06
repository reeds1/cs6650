package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "shop/docs" 
)

// Product represents a shop product
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var products = []Product{
	{ID: 1, Name: "Phone", Price: 699.99, Stock: 10},
	{ID: 2, Name: "Laptop", Price: 1299.00, Stock: 5},
}

// @title Simple Shop API
// @version 1.0
// @description This is a simple API for managing products
// @host localhost:8080
// @BasePath /

func main() {
	// REST endpoints
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productByIDHandler)

	// Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", nil)
}

// productsHandler godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]
// @Router /products [post]
func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getProducts(w, r)
	} else if r.Method == http.MethodPost {
		createProduct(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// productByIDHandler godoc
// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func productByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"Invalid product ID"}`, http.StatusBadRequest)
		return
	}

	for _, p := range products {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, `{"error":"Product not found"}`, http.StatusNotFound)
}

// getProducts returns all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// createProduct adds a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}
