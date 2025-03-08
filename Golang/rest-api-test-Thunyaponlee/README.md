[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/VdFbzCfE)
# Pair Programming : ระบบบริหารจัดการร้านกาแฟ Online (100 คะแนน)

## 📝 ภาพรวม

ให้นักศึกษาพัฒนา REST API สำหรับระบบบริหารจัดการร้านกาแฟ Online โดยใช้ Gin Framework โดยมีความต้องการของระบบ ดังนี้

- ระบบสามารถแสดงรายการเมนูกาแฟทั้งหมด
- ระบบสามารถค้นหาเมนูกาแฟด้วยชื่อหรือประเภท
- ระบบสามารถแสดงรายละเอียดของเมนูกาแฟจากรหัสเมนู
- ระบบสามารถรับ Order ใหม่และคำนวณเวลาที่คาดว่าจะเสร็จ
- ระบบต้องมี Middleware สำหรับบันทึกเวลาการเข้าถึงและ Endpoint ที่ถูกเรียกใช้งาน
- Order จะอยู่ในสถานะ Pending ตอนสร้าง และเปลี่ยนเป็น In Progress เมื่อผ่านไปครึ่งทางของเวลาเตรียม และเป็น Completed เมื่อถึงเวลา estimated delivery

## 📚 รายละเอียดโจทย์

## โครงสร้างข้อมูล 
```bash
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
```
## ข้อมูลตัวอย่าง
```go
var coffees = []Coffee{
	{ID: "c001", Name: "Espresso", Type: "Espresso", Price: 60, Description: "เข้มข้น กลมกล่อม"},
	{ID: "c002", Name: "Americano", Type: "Espresso", Price: 65, Description: "เอสเพรสโซ่ผสมน้ำร้อน"},
	{ID: "c003", Name: "Latte", Type: "Latte", Price: 75, Description: "เอสเพรสโซ่ผสมนมร้อน"},
	{ID: "c004", Name: "Cappuccino", Type: "Cappuccino", Price: 75, Description: "เอสเพรสโซ่ผสมนมร้อนและฟองนม"},
	{ID: "c005", Name: "Mocha", Type: "Mocha", Price: 80, Description: "เอสเพรสโซ่ผสมนมร้อนและช็อกโกแลต"},
}
```
## Endpoints ที่ต้องพัฒนา
```bash
GET /coffees แสดงรายการเมนูกาแฟทั้งหมด
GET /coffees/:id แสดงรายละเอียดเมนูกาแฟจากรหัส
GET /coffees/search ค้นหาเมนูกาแฟด้วย query parameters (name และ type)
POST /orders สร้าง Order ใหม่ (ใช้เวลาเตรียม 5 นาทีต่อ 1 แก้ว)
GET /orders/:order_id - ดูสถานะของ Order
```
1. GET /coffees - แสดงรายการเมนูกาแฟทั้งหมด
Request
```bash
GET /coffees HTTP/1.1
Host: localhost:8080
```
Response (HTTP Status 200 OK)
```bash
{
  "coffees": [
    {
      "id": "c001",
      "name": "Espresso",
      "type": "Espresso",
      "price": 60,
      "description": "เข้มข้น กลมกล่อม"
    },
    {
      "id": "c002",
      "name": "Americano",
      "type": "Espresso",
      "price": 65,
      "description": "เอสเพรสโซ่ผสมน้ำร้อน"
    },
    {
      "id": "c003",
      "name": "Latte",
      "type": "Latte",
      "price": 75,
      "description": "เอสเพรสโซ่ผสมนมร้อน"
    },
    {
      "id": "c004",
      "name": "Cappuccino",
      "type": "Cappuccino",
      "price": 75,
      "description": "เอสเพรสโซ่ผสมนมร้อนและฟองนม"
    },
    {
      "id": "c005",
      "name": "Mocha",
      "type": "Mocha",
      "price": 80,
      "description": "เอสเพรสโซ่ผสมนมร้อนและช็อกโกแลต"
    }
  ]
}
```
2. GET /coffees/:id - แสดงรายละเอียดเมนูกาแฟจากรหัส
Request (กรณีปกติ)
```bash
GET /coffees/c003 HTTP/1.1
Host: localhost:8080
```
Response (กรณีปกติ)
```bash
{
  "id": "c003",
  "name": "Latte",
  "type": "Latte",
  "price": 75,
  "description": "เอสเพรสโซ่ผสมนมร้อน"
}
```
Request (กรณีไม่พบ)
```bash
GET /coffees/c999 HTTP/1.1
Host: localhost:8080
```
Response (กรณีไม่พบ)
```bash
{
  "error": "Coffee not found"
}
```
3. GET /coffees/search - ค้นหาเมนูกาแฟด้วย query parameters
Request (ค้นหาด้วยชื่อ)
```bash
GET /coffees/search?name=lat HTTP/1.1
Host: localhost:8080
```
Response (ค้นหาด้วยชื่อ)
```bash
{
  "coffees": [
    {
      "id": "c003",
      "name": "Latte",
      "type": "Latte",
      "price": 75,
      "description": "เอสเพรสโซ่ผสมนมร้อน"
    }
  ]
}
```
Request (ค้นหาด้วยประเภท)
```bash
GET /coffees/search?type=Espresso HTTP/1.1
Host: localhost:8080
```
Response (ค้นหาด้วยประเภท)
```bash
{
  "coffees": [
    {
      "id": "c001",
      "name": "Espresso",
      "type": "Espresso",
      "price": 60,
      "description": "เข้มข้น กลมกล่อม"
    },
    {
      "id": "c002",
      "name": "Americano",
      "type": "Espresso",
      "price": 65,
      "description": "เอสเพรสโซ่ผสมน้ำร้อน"
    }
  ]
}
```
Request (ไม่มีผลลัพธ์)
```bash
GET /coffees/search?name=frappe HTTP/1.1
Host: localhost:8080
```
Response (ไม่มีผลลัพธ์)
```bash
{
  "coffees": []
}
```
ตัวอย่างการค้นหาคำแบบไม่คำนึงถึงตัวพิมพ์ใหญ่เล็ก
```go
strings.Contains(strings.ToLower(coffee.Name), strings.ToLower(name))
```
ตัวอย่างการเปรียบเทียบข้อความสองชุดว่าเหมือนกันหรือไม่ โดยไม่คำนึงถึงตัวพิมพ์ใหญ่เล็ก
```go
strings.EqualFold(coffee.Type, coffeeType)
```
4. POST /orders - สร้างออร์เดอร์ใหม่

โดยให้สร้าง orderID ด้วยคำสั่งต่อไปนี้
```go
orderID := fmt.Sprintf("o%s", uuid.New().String()[:5])
```
Request (กรณีปกติ)
```bash
POST /orders HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "coffee_id": "c003",
  "quantity": 2
}
```
Response (กรณีปกติ)
```bash
{
  "order_id": "o12345",
  "coffee_id": "c003",
  "quantity": 2,
  "created_at": "2025-03-03T15:30:45+07:00",
  "estimated_delivery": "2025-03-03T15:40:45+07:00",
  "status": "Pending"
}
```
Request (รหัสกาแฟไม่ถูกต้อง)
```bash
POST /orders HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "coffee_id": "c999",
  "quantity": 1
}
```
Response (รหัสกาแฟไม่ถูกต้อง)
```bash
{
  "error": "Coffee not found"
}
```
Request (จำนวนไม่ถูกต้อง)
```bash
POST /orders HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "coffee_id": "c003",
  "quantity": 0
}
```
Response (จำนวนไม่ถูกต้อง)
```bash
{
  "error": "Invalid quantity, must be greater than 0"
}
```
5. GET /orders/:order_id - ดูสถานะของ Order
Request (กรณีปกติ)
```bash
GET /orders/o12345 HTTP/1.1
Host: localhost:8080
```
Response (กรณีปกติ)
```bash
{
  "order_id": "o12345",
  "coffee_id": "c003",
  "quantity": 2,
  "created_at": "2025-03-03T15:30:45+07:00",
  "estimated_delivery": "2025-03-03T15:40:45+07:00",
  "status": "In Progress"
}
```
Request (กรณีไม่พบออร์เดอร์)
```bash
GET /orders/o99999 HTTP/1.1
Host: localhost:8080
```
Response (กรณีไม่พบออร์เดอร์)
```bash
{
  "error": "Order not found"
}
```
## ตัวอย่างโครงสร้าง Middleware
```go
func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // บันทึกเวลาเริ่มต้น
        startTime := time.Now()
        
        // บันทึก Endpoint ที่ถูกเรียก
        path := c.Request.URL.Path
        
        // ดำเนินการต่อไปยังส่วนถัดไป (handler หรือ middleware ถัดไป)
        c.Next()
        
        // บันทึกเวลาสิ้นสุดและคำนวณเวลาที่ใช้
        endTime := time.Now()
        latency := endTime.Sub(startTime)
        
        // แสดงข้อมูลการเข้าถึง
        log.Printf("| %s | %s | %v |", c.Request.Method, path, latency)
    }
}
```
## HTTP Status Code ที่ควรใช้
```bash
200 OK สำหรับการเรียกดูข้อมูลสำเร็จ
201 Created สำหรับการสร้างข้อมูลใหม่สำเร็จ
400 Bad Request สำหรับข้อมูลที่ส่งมาไม่ถูกต้อง
404 Not Found สำหรับไม่พบข้อมูลที่ร้องขอ
500 Internal Server Error สำหรับข้อผิดพลาดภายในเซิร์ฟเวอร์
```
## โครงสร้าง Project (ไม่รวม File Unit Test, go.mod, go.sum และ handlers.go)
```bash
.
├── README.md
├── main.go
├── main_test.go
├── middleware.go
├── models.go
└── routes.go
```
### เริ่มต้นโมดูล Go ด้วยคำสั่ง 
```go
go mod init restapi
```
## 📤 การส่งงาน
1. Clone Repo จาก GitHub Classroom
2. สร้าง/แก้ไข File handlers.go ตามโจทย์ที่กำหนด
3. Build และรัน Program
4. ทดสอบด้วย `go test -v ./...`
5. Commit และ Push code
6. ตรวจสอบคะแนนใน Actions tab