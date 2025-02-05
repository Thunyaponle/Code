[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/VKEJ0clA)
# Go Assignment : Pointers, Structs, Methods and Interfaces (10 ข้อ)

## คำอธิบายโจทย์

ใน Assignment นี้ นักศึกษาจะได้ฝึกทักษะในการเขียนโปรแกรมภาษา Go โดยการใช้งาน Pointer, Struct, Method รวมถึง Interface ทั้งหมด 10 ข้อ (`ex1.go` - `ex10.go`) นักศึกษาต้องเติม Code ให้ถูกต้องตามเงื่อนไขที่กำหนด แต่ละข้อจะมีรายละเอียดและตัวอย่างการใช้งานเบื้องต้น เพื่อให้นักศึกษาเข้าใจแนวทาง หลังจากทำเสร็จให้รันคำสั่งทดสอบ (`go test -v`) แล้ว Push ขึ้น GitHub Classroom เพื่อให้ระบบตรวจและให้คะแนนแบบอัตโนมัติ

ระบบการให้คะแนนเป็นแบบ Autorgrading โดย GitHub Actions ซึ่งจะแสดงผลคะแนนในแท็บ **Actions** บน GitHub หลังจากนักศึกษา Push Code ขึ้นสู่ Github Classroom แล้ว

**ความรู้ที่ใช้:**
- การใช้งาน Pointer เพื่อเข้าถึงและปรับเปลี่ยนค่าของตัวแปร
- การประกาศและใช้งาน Struct และการนำ Struct มา Embed กัน
- การเขียน Method ทั้งแบบ Value Receiver และ Pointer Receiver
- การประกาศและ Implement Interface รวมถึงการใช้ Interface Embedding และ Type Assertion
- การจัดการข้อผิดพลาดด้วย error และ nil Pointer Checking

## รายละเอียดโจทย์

1. **Pointer Basics (`ex1.go`) [10 คะแนน]**
   
   ให้นักศึกษาสร้าง Function `doubleValue(x *int)` ที่รับ Pointer ของตัวแปรจำนวนเต็ม และคูณค่าที่ชี้ไปด้วย 2

   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   num := 5
   fmt.Println("Before:", num) // 5
   doubleValue(&num)
   fmt.Println("After:", num)  // 10
   ~~~

2. **Struct and Pointer Field Update (`ex2.go`) [10 คะแนน]**
   
   ให้นักศึกษาประกาศ Struct `Person` ที่มี Field `Name string, Age int` แล้วเขียน Function `incrementAge(p *Person)` เพื่อเพิ่มอายุของบุคคล 1 ปี

   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   p := Person{Name: "Alice", Age: 30}
   incrementAge(&p)
   fmt.Println(p.Age) // 31
   ~~~

3. **Method Receiver กับ Struct (`ex3.go`) [10 คะแนน]**

   ให้นักศึกษาประกาศ Struct `Rectangle` มี Field `Width, Height float64` แล้วสร้าง Method ดังต่อไปนี้
   - Method `Area()` แบบ Value Receiver คืนค่าพื้นที่ (Width * Height)
   - Method `Resize(w, h float64)` แบบ Pointer Receiver ปรับค่าขนาดของสี่เหลี่ยม

   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   r := Rectangle{Width:10, Height:5}
   fmt.Println(r.Area()) // 50
   r.Resize(2,2)
   fmt.Println(r.Area()) // 4
   ~~~

4. **Interface Implementation (`ex4.go`) [10 คะแนน]**
   
   ให้นักศึกษาประกาศ Interface `Shape` ที่มี Method `Area() float64` สร้าง Struct `Circle{Radius float64}` และ `Square{Side float64}` ให้ Implement `Shape` และ Function `printArea(s Shape)` พิมพ์ค่าพื้นที่

   **หมายเหตุ** สามารถใช้ `math.Pi` ของ Package `math` ได้

   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   c := Circle{Radius:5}
   sq := Square{Side:4}
   printArea(c)  // พิมพ์ค่าพื้นที่วงกลมรัศมี 5
   printArea(sq) // พิมพ์ค่าพื้นที่สี่เหลี่ยมจัตุรัสด้าน 4
   ~~~

5. **Interface และ Pointer Receiver (`ex5.go`) [10 คะแนน]**
   
   ให้นักศึกษาประกาศ Interface `Account` มี Method ดังนี้
   - `Deposit(amount float64)` (เป็น Pointer Receiver) สำหรับการฝากเงิน
   - `Withdraw(amount float64) error` สำหรับการถอนเงิน โดยถ้ายอดถอนมากกว่าเงินในบัญชีจะไม่สามารถถอนเงินได้ และ Return Error ที่ไม่ใช่ nil แต่ถ้ามีเงินพอที่จะถอนจะ Return `nil`
   - `Balance() float64` สำหรับดูยอดเงินในบัญชี
   
   แล้วสร้าง Struct `BankAccount{owner string, balance float64}` ให้ Implement Interface `Account`
   
   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   acc := BankAccount{owner:"Bob", balance:1000.0}
   acc.Deposit(500)
   fmt.Println(acc.Balance()) // 1500.0
   acc.Withdraw(1200)
   fmt.Println(acc.Balance()) // 300.0
   ~~~

   **หมายเหตุ** ตัวอย่างการ Return Error ด้วย Function `New()` จาก Package `errors`

   ~~~
   if amount > acc.balance {
		return errors.New("ยอดเงินไม่พอสำหรับการถอน")
   }
   ~~~

6. **Struct Embedding และ Method Resolution (`ex6.go`) [10 คะแนน]**

   ให้นักศึกษาสร้าง Struct `Animal{Name string}` โดยมี Method `Breathe()` พิมพ์ว่า "<Name> กำลังหายใจ" สร้าง `Cat` ที่ Embed `Animal` และมี Method `Meow()` พิมพ์ว่า "<Name> กำลังร้อง: เมี้ยว! เมี้ยว!"
   
   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   myCat := Cat{Animal: Animal{Name:"ตุนตัง"}, Breed:"มันช์กิ้น"}
   myCat.Breathe() // "ตุนตัง กำลังหายใจ"
   myCat.Meow()    // "ตุนตัง กำลังร้อง: เมี้ยว! เมี้ยว!"
   ~~~

7. **Nil Pointer Checking (`ex7.go`) [10 คะแนน]**
   
   ให้นักศึกษาสร้าง​ Function `safeDereference(ptr *int) (int, error)`
   - ถ้า ptr เป็น nil คืนค่า (0, error ที่ไม่เป็น nil)
   - ถ้า ptr ไม่เป็น nil คืนค่า value ที่ชี้ไปและ error เป็น nil
   
   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   var p *int
   val, err := safeDereference(p)
   // val = 0, err != nil

   x := 42
   val, err = safeDereference(&x)
   // val = 42, err = nil
   ~~~

8. **Interface + Type Assertion (`ex8.go`) [10 คะแนน]**
   
   ให้นักศึกษาประกาศ Interface `Soundmaker` ที่มี Method `MakeSound() string` ประกาศ Dog และ Cow ให้ MakeSound() เป็น `"Woof!"` และ `"Moo!"` ตามลำดับ แล้วสร้าง Function `identifyAnimal(s Soundmaker) string` ที่ใช้ Type Assertion แยกกรณี ดังต่อไปนี้
   - Dog => `"This is a Dog!"`
   - Cow => `"This is a Cow!"`
   - อื่นๆ => `"Unknown Animal!"`
   
   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   fmt.Println(identifyAnimal(Dog{})) // "This is a Dog!"
   fmt.Println(identifyAnimal(Cow{})) // "This is a Cow!"
   ~~~

9. **Pointer Receiver Method vs Value Receiver Method (`ex9.go`) [10 คะแนน]**
   
   ให้นักศึกษาสร้าง Struct `Counter{Count int}` ที่มี Method `Increment()` แบบ Pointer Receiver เพิ่ม Count ขึ้น 1 และ Method `GetCount()` แบบ Value Receiver คืนค่า Count
   
   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   c := Counter{Count:10}
   c.Increment()
   c.Increment()
   fmt.Println(c.GetCount()) // 12
   ~~~

10. **Interface Embedding (`ex10.go`) [10 คะแนน]**

   ให้นักศึกษาประกาศ Interface ดังต่อไปนี้
   - Interface `Mover` มี Method `Move() string`
   - Interface `Sounder` มี Method `MakeSound() string`
   - Interface `Creature` Embed `Mover` และ `Sounder`
   
   และประกาศ Struct `Horse` ที่ Implement `Creature` โดย
   - มี Method `Move()` คืนค่า `"Horse is galloping"`
   - มี Method `MakeSound()` คืนค่า `"Neigh!"`
   แล้วสร้าง Function `describeCreature(c Creature) string` คืนค่า `"<Move()> and <MakeSound()>"`
   
   ตัวอย่างการใช้งานใน Function `main()`
   ~~~
   h := Horse{}
   fmt.Println(describeCreature(h)) // "Horse is galloping and Neigh!"
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

4. **แก้ไข Code ในไฟล์ `ex1.go`-`ex10.go` ตามโจทย์ข้อ 1 ถึง 10 ตามลำดับ**

5. **รัน Unit Test** เพื่อทดสอบความถูกต้องของ Code ด้วยคำสั่ง `go test -v` (หากโค้ดและ Unit Test ทำงานถูกต้อง คุณจะได้ผลลัพธ์เป็น PASS)

    ~~~
    go test -v
    ~~~

6. **บันทึก Version ของ Code ลง Git Repo** ด้วยคำสั่ง `git add .` และ `git commit -m 'commit message'` (แก้ไข commit message เพื่ออธิบายว่าคุณทำอะไรกับ Code ที่บันทึกลง Git)

    ~~~
    git add .
    git commit -m 'complete ex1-ex10'
    ~~~

7. **Submit Code ขึ้น Github** เพื่อตรวจ Code และให้คะแนน ด้วยคำสั่ง `git push origin main`
    
    ~~~
    git push origin main
    ~~~

8. **ไปยัง Git Repo ของ Assignment นี้บน github.com** กด `Actions` (Github จะใช้เวลาในการตรวจคำตอบสักพัก) เลือก Commit ล่าสุดเพื่อดูคะแนน

## ขอให้สนุกกับการทำ Assignment ครับ