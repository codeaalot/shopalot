package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"shopalot-backend/models"

	"github.com/google/uuid"
)

// initialize test products
func init() {
	newProducts := []models.Product{
		{
			ID: "test-id-1",
			BaseProduct: models.BaseProduct{
				Name:        "Test Product 1",
				Description: "This is a test description for product 1",
				Price:       9.99,
			},
			InStock: true,
		},
		{
			ID: "test-id-2",
			BaseProduct: models.BaseProduct{
				Name:        "Test Product 2",
				Description: "This is a test description for product 2",
				Price:       19.99,
			},
			InStock: false,
		},
		{
			ID: "test-id-3",
			BaseProduct: models.BaseProduct{
				Name:        "Test Product 3",
				Description: "This is a test description for product 3",
				Price:       29.99,
			},
			InStock: true,
		},
	}
	products = append(products, newProducts...)
}

func TestCreateProduct(t *testing.T) {
	// Create a new product in JSON format
	newProduct := models.BaseProduct{
		Name:        "New Test Product",
		Description: "New test product description",
		Price:       15.99,
	}
	productJSON, _ := json.Marshal(newProduct)

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createProduct)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var gotProduct models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &gotProduct)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if gotProduct.Name != newProduct.Name || gotProduct.Description != newProduct.Description || gotProduct.Price != newProduct.Price {
		t.Errorf("handler returned unexpected body: got %v want %v",
			gotProduct, newProduct)
	}

	if _, err := uuid.Parse(gotProduct.ID); err != nil {
		t.Errorf("handler returned invalid UUID: got %v", gotProduct.ID)
	}
}

func TestGetProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getProducts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var gotProducts []models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &gotProducts)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	expectedProducts := products

	if len(gotProducts) != len(expectedProducts) {
		t.Fatalf("handler returned unexpected number of products: got %v want %v",
			len(gotProducts), len(expectedProducts))
	}

	for index, product := range gotProducts {
		if product != expectedProducts[index] {
			t.Errorf("handler returned unexpected body at index %d: got %v want %v",
				index, product, expectedProducts[index])
		}
	}
}

func TestGetProductByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/products/test-id-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductByID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var gotProduct models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &gotProduct)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	expectedProduct := products[0]

	if gotProduct != expectedProduct {
		t.Errorf("handler returned unexpected body: got %v want %v", gotProduct, expectedProduct)
	}
}
