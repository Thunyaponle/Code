/*package main

import (
	"fmt"
	"myproject/calc"  // "ชื่อ module / folder"
	"myproject/utils" //ชื่อ module ต้องตรงกับที่ import เข้ามา
	"net/http"

	"github.com/gin-gonic/gin"
)*/

/*
func main() {
	result := calc.Add(1, 2)
	fmt.Println(result)

	st := utils.NumberToThai(result)
	fmt.Print(st)
}*/
/*
func main() {
	result := calc.Add(5, 3)
	fmt.Println("5 + 3 =", result)

	thaiNumber := utils.NumberToThai(result)
	fmt.Println("ผลลัพธ์เป็นภาษาไทย:", thaiNumber)

	// สร้าง Gin engine
	r := gin.Default()

	// กำหนด route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// กำหนด route ที่รับพารามิเตอร์
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// รัน server ที่ port 8080
	r.Run(":8080")
}
*/

// pair programming

package main

import (
	"fmt"
)

type Student struct {
	name  string
	score float64
}

func calculateGrade(sum float64) string {
	if sum >= 80 {
		return "A"
	} else if sum >= 70 {
		return "B"
	} else if sum >= 60 {
		return "C"
	} else if sum >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func (s Student) getGrade() string {
	return calculateGrade(s.score)
}

func main() {
	silceStu := []string{"Ann", "Dog", "Cat", "Boss"}
	sliceScore := []float64{85.3, 55, 88, 69}

	for i := 0; i < 4; i++ {
		student := Student{name: silceStu[i], score: sliceScore[i]}
		fmt.Printf("Name: %s, Grade: %s\n", student.name, student.getGrade())
	}

}
