package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs the timestamp, method, path, and response time for each request
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
