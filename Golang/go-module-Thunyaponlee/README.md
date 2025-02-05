[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/zhHfbYiN)
# Go Assignment : Package และ Module (7 ข้อ 100 คะแนน)

## คำอธิบายโจทย์

ใน Assignment นี้ นักศึกษาจะได้ฝึกทักษะในการเขียนโปรแกรมภาษา Go โดยการสร้าง Package และ Module ทั้งหมด 7 ข้อ (Ex1-Ex7) ซึ่งนักศึกษาต้องเติม Code ให้ถูกต้องตามเงื่อนไขที่กำหนด แต่ละข้อจะมีรายละเอียดและตัวอย่างการใช้งานเบื้องต้น

ระบบการให้คะแนนเป็นแบบ Unit Test โดยมีการทดสอบแต่ละ Function ว่าทำงานได้ถูกต้องตามที่โจทย์กำหนดหรือไม่

**ความรู้ที่ใช้**
- การสร้างและใช้งาน Package ใน Go
- การสร้าง Function แบบเรียกใช้งานได้จาก Package อื่น (Exported) และไม่สามารถเรียกใช้งานได้จาก Package อื่น (Unexported)
- การจัดการโครงสร้าง Project
- การใช้งาน External Package
- การจัดการ Dependencies ด้วย go mod

## รายละเอียดโจทย์

1. **Package Greetings Basic (`Ex1`) [10 คะแนน]**
   
   สร้าง Package ชื่อ greetings (อยู่ใน Folder `greetings`) โดยมี Function แบบ Exported ชื่อ `Hello(name string) string` ทำหน้าที่รับชื่อแล้วคืนค่าเป็นสตริง `"Hello, <name>"` แต่ถ้าชื่อเป็นค่าว่างให้คืนค่า `"Hello, World"`

   ```go
   result := greetings.Hello("Alice")
   fmt.Println(result) // "Hello, Alice"
   
   result = greetings.Hello("")
   fmt.Println(result) // "Hello, World"
   ```

2. **Multiple Files in Package (`Ex2`) [10 คะแนน]**
   
   สร้าง Function ใน Package `greetings` ชื่อ `Goodbye(name string) string` ที่คืนค่า `"Goodbye, <name>"` โดยให้ Function นี้อยู่ใน File `farewell.go` เพื่อทดสอบว่าหลาย File ใน Package เดียวกันก็ทำงานได้

   ```go
   result := greetings.Goodbye("Bob")
   fmt.Println(result) // "Goodbye, Bob"
   ```

3. **Unexported Function (`Ex3`) [15 คะแนน]**
   
   สร้าง Function แบบ unexported ชื่อ `formatMessage(name string) string` (ขึ้นต้นด้วยตัวพิมพ์เล็ก) ภายใน Folder `greetings` ให้คืนค่า `"*** <name> ***"`

   แต่เมื่อส่งค่าว่างไป `formatMessage` จะแทนที่ค่าว่างด้วย "World" ก่อน ทำให้ได้ผลลัพธ์ คือ `"Hello, *** World ***"`

   และสร้าง Function แบบ exported ชื่อ `FormatHello(name string) string` ที่ภายในเรียก `formatMessage(name)` และเติม "Hello, " ไว้ด้านหน้า

   ```go
   result := greetings.FormatHello("Alice")
   fmt.Println(result) // "Hello, *** Alice ***"
   
   result = greetings.FormatHello("")
   fmt.Println(result) // "Hello, *** World ***"
   ```

4. **Calculator Package (`Ex4`) [15 คะแนน]**
   
   สร้าง File `calculator.go` ภายใน Folder `calc` โดยประกาศเป็น package `calc`

   สร้าง Function `Add(a, b int) int` และ `Multiply(a, b int) int` เป็นแบบ Exported

   ```go
   sum := calc.Add(5, 3)
   product := calc.Multiply(4, 2)
   fmt.Println(sum, product) // 8, 8
   ```

5. **Thai Number Converter (`Ex5`) [15 คะแนน]**
   
   สร้าง Folder `utils` และ File `converter.go` (package utils)

   และ Function `NumberToThai(num int) string` ที่แปลงตัวเลขหลักเดียว (0-9) เป็นคำอ่านภาษาไทย

   โดยตรวจสอบว่าตัวเลขอยู่ในช่วง 0-9 หรือไม่ ถ้าไม่ใช่ Return `""`

   ```go
   result := utils.NumberToThai(5)
   fmt.Println(result) // "ห้า"
   
   result = utils.NumberToThai(10)
   fmt.Println(result) // ""
   ```

6. **UUID Generator (`Ex6`) [15 คะแนน]**
   
   ติดตั้ง Module `github.com/google/uuid` เพื่อสร้าง UUID แบบสุ่ม

   ใน File `utils/tools.go` (package `utils`) ให้สร้าง Function `GenerateUUID() string` ที่เรียก `uuid.NewString()` โดยคืนค่าเป็น String

   ```go
   uuid := utils.GenerateUUID()
   fmt.Println(uuid) // จะได้ UUID แบบสุ่ม
   ```

7. **Dependency Management (`Ex7`) [20 คะแนน]**
   
   จัดการ Dependencies ด้วย `go mod tidy`

## การส่งงาน

1. Clone Repository จาก GitHub Classroom
2. สร้าง Module ชื่อ `go_module`
3. สร้างและแก้ไข File ตามที่โจทย์กำหนด
4. รัน Unit Test ด้วยคำสั่ง `go test ./... -v` เพื่อตรวจสอบความถูกต้องทุกข้อ หรือรัน Unit Test ทีละข้อด้วยคำสั่ง ดังต่อไปนี้
- Ex1 Unit Test : `go test ./greetings -run TestHello -v`
- Ex2 Unit Test : `go test ./greetings -run "Goodbye" -v`
- Ex3 Unit Test : `go test ./greetings -run "Format" -v`
- Ex4 Unit Test : `go test ./calc -v'
- Ex5 Unit Test : `go test ./utils -run "^(TestConverterFileStructure|TestNumberToThai)$" -v`
- Ex6 Unit Test : `go test ./utils -run "^(TestToolsFileStructure|TestGenerateUUID|TestUUIDImplementation)$" -v`
- Ex7 Unit Test : `go test -run "^(TestGoModTidy|TestProjectStructure|TestGoModRequirements)$" -v`

5. Commit และ Push code ขึ้น GitHub Classroom

## โครงสร้าง Project (ไม่รวม File Unit Test)
```
go_module/
├── greetings/
│   ├── greetings.go
│   └── farewell.go
├── calc/
│   └── calculator.go
├── utils/
│   ├── converter.go
│   └── tools.go
├── go.mod
└── go.sum
```

## เกณฑ์การให้คะแนน
- Function ทำงานถูกต้องตามที่โจทย์กำหนด
- มีการจัดโครงสร้าง Project ถูกต้อง
- Package และ Dependencies ถูกจัดการอย่างเหมาะสม
- ผ่าน Unit Test ตามที่กำหนด

## ขอให้สนุกกับการทำ Assignment ครับ