[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/pxqHs-ZM)
# Go Assignment: First-Class Functions, Arrays, Slices and Maps

## คำอธิบายโจทย์

ใน Assignment นี้ นักศึกษาต้องเขียน Code ภาษา Go โดยใช้ความรู้เกี่ยวกับ First-Class Citizen Functions, Arrays, Slices และ Maps ที่ได้ศึกษาจากบทเรียน ผมได้เตรียมไฟล์ `ex1.go` ถึง `ex4.go` โดยนักศึกษาต้องทำการเติม Code ใน File ที่กำหนดให้ถูกต้องตามเงื่อนไข เมื่อทำเสร็จให้รันคำสั่งทดสอบ (`go test -v`) แล้ว Push ขึ้น Github Classroom

ระบบการให้คะแนนเป็นแบบ Autorgrading โดย GitHub Actions ซึ่งจะแสดงผลคะแนนในแท็บ **Actions** บน GitHub หลังจากนักศึกษา Push Code ขึ้นสู่ Github Classroom แล้ว

**ความรู้ที่ใช้:**
- First-Class Citizen Functions ในภาษา Go
- การใช้งานและการคืนค่าจาก Array (Fixed-size)
- การสร้างและจัดการ Slice
- การใช้งาน Map

## รายละเอียดโจทย์

1. **First-Class Citizen Functions (10 คะแนน)**  
   ให้นักศึกษาสร้าง Function `applyOperation`  
   โดย `applyOperation` จะรับ Function `op` ที่เป็น Function การดำเนินการทางคณิตศาสตร์ (ได้แก่ add, subtract, multiply และ divide) กับเลขจำนวนจริง `a` และ `b` จากนั้นให้ return ผลลัพธ์ของ `op(a,b)`

   ตัวอย่างการใช้งานใน Function main
   ~~~
   add := func(x, y float64) float64 {
      return x + y
   }

   result := applyOperation(add, 5, 3)
   fmt.Println(result) // ควรได้ผลลัพธ์ 8
   ~~~

2. **Array (10 คะแนน)**  
   ให้นักศึกษาสร้าง Function `sumArray`  
   Function นี้จะรับ Array ความยาว 5 ของจำนวนเต็ม และ return ผลรวมของตัวเลขทั้งหมดใน Array นั้น

   ตัวอย่างการใช้งานใน Function main
   ~~~
   arr := [5]int{1,2,3,4,5}
   sum := sumArray(sum)
   fmt.Println(sum) // ควรได้ผลลัพธ์ 15
   ~~~

3. **Slice (10 คะแนน)**  
   ให้นักศึกษาสร้าง Function `filterEven`  
   Function นี้จะรับ Slice ของจำนวนเต็ม และ return Slice ใหม่ที่มีเฉพาะตัวเลขคู่เท่านั้น

   ตัวอย่างการใช้งานใน Function main
   ~~~
   s := []int{1,2,3,4,5,6}
   evens := filterEven(s)
   fmt.Println(evens) // ควรได้ผลลัพธ์ [2,4,6]
   ~~~

4. **Map (10 คะแนน)**  
   ให้นักศึกษาสร้าง Function `countOccurrences`  
   Function นี้จะรับ Slice ของ string และ return Map ที่มี key เป็น string และ value เป็นจำนวนครั้งที่ string นั้นปรากฏใน Slice

   ตัวอย่างการใช้งานใน Function main
   ~~~
   words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
   counts := countOccurrences(words) 
   fmt.Println(counts) // ควรได้ผลลัพธ์ {"apple":3, "banana":2, "orange":1}

   ~~~


## วิธีการทำงานและส่งงาน

1. **หากคุณยังไม่ได้สร้าง SSH Key บน Github Account ของคุณ ให้ทำตามขั้นตอนจากตัวอย่างในบทความนี้ [Python Back-End Development for Beginners](https://blog.pjjop.org/back-end-programming/)**

2. **Clone Repo จาก GitHub Classroom**  
   ไปที่หน้า Assignment บน GitHub Classroom แล้วกดปุ่ม **"Accept Assignment"** จากนั้นจะมี Repo ของนักศึกษาปรากฏขึ้น  
   ให้คัดลอก SSH URL ของ Repo นั้น โดย
    1. ไปที่ `<> Code` บน Git Repo ของ Assignment เลือก `SSH` แล้วกดปุ่ม `Copy url`
    2. ไปยังเครื่องของคุณ ใช้คำสั่ง `git clone url` ที่ Command line

    ~~~
    git clone url
    ~~~
    
    3. เข้าไปยัง Folder Project ที่ Clone มา แล้วเปิด VS Code ด้วยคำสั่ง `code .`

    ~~~
    code .
    ~~~

3. **เริ่มต้น Go Module** ด้วยคำสั่ง `go mod init myproject` (go จะสร้างไฟล์ go.mod ใน Prject ของคุณ)

    ~~~
    go mod init myproject
    ~~~

4. **แก้ไข Code ในไฟล์ ex1.go-ex4.go ตามโจทย์ข้อ 1 ถึง 4 ตามลำดับ**

5. **รัน Unit Test** เพื่อทดสอบความถูกต้องของ Code ด้วยคำสั่ง `go test -v` (หากโค้ดและ Unit Test ทำงานถูกต้อง คุณจะได้ผลลัพธ์เป็น PASS)

    ~~~
    go test -v
    ~~~

6. **บันทึก Version ของ Code ลง Git Repo** ด้วยคำสั่ง `git add .` และ `git commit -m 'commit message'` (แก้ไข commit message เพื่ออธิบายว่าคุณทำอะไรกับ Code ที่บันทึกลง Git)

    ~~~
    git add .
    git commit -m 'edit ex1-ex4'
    ~~~

7. **Submit Code ขึ้น Github** เพื่อตรวจ Code และให้คะแนน ด้วยคำสั่ง `git push origin main`
    
    ~~~
    git push origin main
    ~~~

8. **ไปยัง Git Repo ของ Assignment นี้บน github.com** กด `Actions` (Github จะใช้เวลาในการตรวจคำตอบสักพัก) เลือก Commit ล่าสุดเพื่อดูคะแนน

