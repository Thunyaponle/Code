[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/8eKcK-hV)
# Mini Library Management System Test (100 คะแนน)

## 📝 ภาพรวม

ให้นักศึกษาสร้าง “ระบบจัดการหนังสือในห้องสมุด” โดยมีความสามารถหลัก ๆ ดังนี้

- จัดเก็บ รายการหนังสือ ด้วย Slice ของ Book
- เพิ่มหนังสือ (AddBook) พร้อมตรวจสอบ “ID ซ้ำ”
- ค้นหาหนังสือ (GetBookByID)
- มีการคืนค่า error หากพบปัญหา เมื่อ ID ซ้ำ และค้นหาไม่พบ
- มี Interface Describable

## 📚 รายละเอียดโจทย์

## โครงสร้าง Project (ไม่รวม File Unit Test)
```bash
myproject/
├── go.mod
├── library
│   ├── library.go        // struct Book + Library + Method + Interface
│   └── library_test.go   // Unit Test
└── cmd
    └── main.go           // Code ตัวอย่างเรียกใช้งาน
```
### เริ่มต้นโมดูล Go ด้วยคำสั่ง 
```go
go mod init myproject
```

### 1. สร้าง Struct Book มี Field 3 Field ได้แก่ ID, Title และ Author
### 2. สร้าง Struct Library ที่ใช้ Slice เก็บรายการหนังสือ โดยประกาศ Method 2 ตัว ดังต่อไปนี้
    2.1 AddBook(b Book) error
    - ตรวจสอบว่า b.ID ซ้ำกับใน l.Books หรือไม่
    - ถ้าซ้ำให้ return errors.New("duplicate book ID")
    - ถ้าไม่ซ้ำให้ append ลง slice และ return nil

    2.2 GetBookByID(id int) (Book, error)
    - วนลูปค้นหาใน l.Books
    - ถ้าพบ ID == id ให้ return book, nil
    - ถ้าไม่พบให้ return Book{}, errors.New("book not found")

    2.3 ใช้ Pointer Receiver func (l *Library) ... ในการเพิ่มหนังสือ เพื่อแก้ไข Slice ตัวจริงใน Library

### 3. สร้าง Interface `Describable` ที่กำหนด Method `Describe() string`
### 4. Implement  Method `Describe()` ให้ Book ที่มีการ Return ดังต่อไปนี้

```go
fmt.Sprintf("Book[ID=%d, Title=%s, Author=%s]", b.ID, b.Title, b.Author)
```

## 📤 การส่งงาน

1. Clone Repo จาก GitHub Classroom
2. สร้าง/แก้ไข File ตามโจทย์ที่กำหนด
3. Build และรัน Program
4. ทดสอบด้วย `go test -v ./...`
5. Commit และ Push code
6. ตรวจสอบคะแนนใน Actions tab