package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"shopalot-backend/models"

	"github.com/google/uuid"
)

var (
	products []models.Product
	mutex    sync.Mutex
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProducts(w, r)
	case http.MethodPost:
		createProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	defer mutex.Unlock()
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")

	// Assuming the ID is the last segment of the URL path
	if len(segments) < 3 {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	productID := segments[len(segments)-1]

	mutex.Lock()
	defer mutex.Unlock()

	for _, product := range products {
		if product.ID == productID {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newBaseProduct models.BaseProduct
	if err := json.NewDecoder(r.Body).Decode(&newBaseProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add validation for product fields as needed

	mutex.Lock()
	defer mutex.Unlock()
	newProduct := models.Product{
		ID:          uuid.New().String(),
		BaseProduct: newBaseProduct,
		InStock:     true, // Default to true or add logic to determine stock
	}
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
