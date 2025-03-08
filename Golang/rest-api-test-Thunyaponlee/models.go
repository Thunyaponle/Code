package main

// เมนูกาแฟ
type Coffee struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"` // เช่น "Espresso", "Latte", "Cappuccino"
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// Order
type Order struct {
	OrderID           string `json:"order_id"`
	CoffeeID          string `json:"coffee_id"`
	Quantity          int    `json:"quantity"`
	CreatedAt         string `json:"created_at"`
	EstimatedDelivery string `json:"estimated_delivery"`
	Status            string `json:"status"` // "Pending", "In Progress", "Completed"
}

// OrderRequest represents the request payload for creating a new order
type OrderRequest struct {
	CoffeeID string `json:"coffee_id" binding:"required"`
	Quantity int    `json:"quantity"`
}

// CoffeesResponse represents the response for a list of coffees
type CoffeesResponse struct {
	Coffees []Coffee `json:"coffees"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Global variables to store data (in a real application, this would be in a database)
var coffees = []Coffee{
	{ID: "c001", Name: "Espresso", Type: "Espresso", Price: 60, Description: "เข้มข้น กลมกล่อม"},
	{ID: "c002", Name: "Americano", Type: "Espresso", Price: 65, Description: "เอสเพรสโซ่ผสมน้ำร้อน"},
	{ID: "c003", Name: "Latte", Type: "Latte", Price: 75, Description: "เอสเพรสโซ่ผสมนมร้อน"},
	{ID: "c004", Name: "Cappuccino", Type: "Cappuccino", Price: 75, Description: "เอสเพรสโซ่ผสมนมร้อนและฟองนม"},
	{ID: "c005", Name: "Mocha", Type: "Mocha", Price: 80, Description: "เอสเพรสโซ่ผสมนมร้อนและช็อกโกแลต"},
}

// In-memory storage for orders
var orders = make(map[string]Order)
