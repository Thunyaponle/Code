package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *http.Request, expectedCode int) *httptest.ResponseRecorder {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the request
	router := SetupRouter()
	router.ServeHTTP(w, r)

	// Assert expected status code
	assert.Equal(t, expectedCode, w.Code)

	return w
}

// Test GET /coffees endpoint
func TestGetAllCoffees(t *testing.T) {
	r, _ := http.NewRequest("GET", "/coffees", nil)
	w := testHTTPResponse(t, r, http.StatusOK)

	var response CoffeesResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Len(t, response.Coffees, 5) // There should be 5 coffees in our sample data

	// Check first coffee details
	assert.Equal(t, "c001", response.Coffees[0].ID)
	assert.Equal(t, "Espresso", response.Coffees[0].Name)
}

// Test GET /coffees/:id endpoint - success case
func TestGetCoffeeByIDSuccess(t *testing.T) {
	r, _ := http.NewRequest("GET", "/coffees/c003", nil)
	w := testHTTPResponse(t, r, http.StatusOK)

	var coffee Coffee
	err := json.Unmarshal(w.Body.Bytes(), &coffee)

	assert.Nil(t, err)
	assert.Equal(t, "c003", coffee.ID)
	assert.Equal(t, "Latte", coffee.Name)
	assert.Equal(t, "Latte", coffee.Type)
	assert.Equal(t, 75.0, coffee.Price)
}

// Test GET /coffees/:id endpoint - not found case
func TestGetCoffeeByIDNotFound(t *testing.T) {
	r, _ := http.NewRequest("GET", "/coffees/c999", nil)
	w := testHTTPResponse(t, r, http.StatusNotFound)

	var errorResponse ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.Nil(t, err)
	assert.Equal(t, "Coffee not found", errorResponse.Error)
}

// Test GET /coffees/search endpoint - search by name
func TestSearchCoffeesByName(t *testing.T) {
	r, _ := http.NewRequest("GET", "/coffees/search?name=lat", nil)
	w := testHTTPResponse(t, r, http.StatusOK)

	var response CoffeesResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Len(t, response.Coffees, 1)
	assert.Equal(t, "Latte", response.Coffees[0].Name)
}

// Test GET /coffees/search endpoint - search by type
func TestSearchCoffeesByType(t *testing.T) {
	r, _ := http.NewRequest("GET", "/coffees/search?type=Espresso", nil)
	w := testHTTPResponse(t, r, http.StatusOK)

	var response CoffeesResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Len(t, response.Coffees, 2) // Espresso and Americano

	// Check that both coffees have type "Espresso"
	for _, coffee := range response.Coffees {
		assert.Equal(t, "Espresso", coffee.Type)
	}
}

// Test GET /coffees/search endpoint - no results
func TestSearchCoffeesNoResults(t *testing.T) {
	r, _ := http.NewRequest("GET", "/coffees/search?name=frappe", nil)
	w := testHTTPResponse(t, r, http.StatusOK)

	var response CoffeesResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Len(t, response.Coffees, 0) // No coffees should match
}

// Test POST /orders endpoint - success case
func TestCreateOrderSuccess(t *testing.T) {
	payload := `{"coffee_id": "c003", "quantity": 2}`
	r, _ := http.NewRequest("POST", "/orders", bytes.NewBufferString(payload))
	r.Header.Set("Content-Type", "application/json")

	w := testHTTPResponse(t, r, http.StatusCreated)

	var order Order
	err := json.Unmarshal(w.Body.Bytes(), &order)

	assert.Nil(t, err)
	assert.Equal(t, "c003", order.CoffeeID)
	assert.Equal(t, 2, order.Quantity)
	assert.Equal(t, "Pending", order.Status)

	// Verify the order ID format
	assert.True(t, strings.HasPrefix(order.OrderID, "o"))

	// Verify timestamps
	_, err = time.Parse(time.RFC3339, order.CreatedAt)
	assert.Nil(t, err)

	estimatedTime, err := time.Parse(time.RFC3339, order.EstimatedDelivery)
	assert.Nil(t, err)

	// Check that the estimated delivery is approximately (5 minutes * quantity) later
	// With a small buffer for test execution time
	createdTime, _ := time.Parse(time.RFC3339, order.CreatedAt)
	expectedDuration := time.Duration(order.Quantity*5) * time.Minute
	timeDiff := estimatedTime.Sub(createdTime)

	// Allow for 1 second of execution time
	assert.True(t, timeDiff >= expectedDuration-time.Second && timeDiff <= expectedDuration+time.Second)
}

// Test POST /orders endpoint - invalid coffee ID
func TestCreateOrderInvalidCoffeeID(t *testing.T) {
	payload := `{"coffee_id": "c999", "quantity": 1}`
	r, _ := http.NewRequest("POST", "/orders", bytes.NewBufferString(payload))
	r.Header.Set("Content-Type", "application/json")

	w := testHTTPResponse(t, r, http.StatusNotFound)

	var errorResponse ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.Nil(t, err)
	assert.Equal(t, "Coffee not found", errorResponse.Error)
}

// Test POST /orders endpoint - invalid quantity
func TestCreateOrderInvalidQuantity(t *testing.T) {
	payload := `{"coffee_id": "c003", "quantity": 0}`
	r, _ := http.NewRequest("POST", "/orders", bytes.NewBufferString(payload))
	r.Header.Set("Content-Type", "application/json")

	w := testHTTPResponse(t, r, http.StatusBadRequest)

	var errorResponse ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.Nil(t, err)
	assert.Equal(t, "Invalid quantity, must be greater than 0", errorResponse.Error)
}

// Test GET /orders/:order_id endpoint
func TestGetOrderStatus(t *testing.T) {
	// First, create an order
	payload := `{"coffee_id": "c001", "quantity": 1}`
	rCreate, _ := http.NewRequest("POST", "/orders", bytes.NewBufferString(payload))
	rCreate.Header.Set("Content-Type", "application/json")

	wCreate := httptest.NewRecorder()
	router := SetupRouter()
	router.ServeHTTP(wCreate, rCreate)

	var order Order
	err := json.Unmarshal(wCreate.Body.Bytes(), &order)
	assert.Nil(t, err)

	// Now fetch the order status
	rGet, _ := http.NewRequest("GET", "/orders/"+order.OrderID, nil)
	wGet := testHTTPResponse(t, rGet, http.StatusOK)

	var fetchedOrder Order
	err = json.Unmarshal(wGet.Body.Bytes(), &fetchedOrder)

	assert.Nil(t, err)
	assert.Equal(t, order.OrderID, fetchedOrder.OrderID)
	assert.Equal(t, order.CoffeeID, fetchedOrder.CoffeeID)
	assert.Equal(t, order.Quantity, fetchedOrder.Quantity)
}

// Test GET /orders/:order_id endpoint - order not found
func TestGetOrderStatusNotFound(t *testing.T) {
	r, _ := http.NewRequest("GET", "/orders/o99999", nil)
	w := testHTTPResponse(t, r, http.StatusNotFound)

	var errorResponse ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.Nil(t, err)
	assert.Equal(t, "Order not found", errorResponse.Error)
}
