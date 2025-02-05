1. หากคุณยังไม่ได้สร้าง SSH Key บน Github Account ของคุณ ให้ทำตามขั้นตอนจากตัวอย่างในบทความนี้ [Python Back-End Development for Beginners](https://blog.pjjop.org/back-end-programming/)
2. Clone Git Repo บน github.com มายังเครื่องของคุณ
    1. ไปที่ `<> Code` บน Git Repo ของ Assignment นี้ เลือก `SSH` แล้วกดปุ่ม `Copy url`
    2. ไปยังเครื่องของคุณ ใช้คำสั่ง `git clone url` ที่ Command line

    ~~~
    git clone url
    ~~~
    
    3. เข้าไปยัง Folder Project ที่ Clone มา แล้วเปิด VS Code ด้วยคำสั่ง `code .`

    ~~~
    code .
    ~~~

3. เริ่มต้น Go Module ด้วยคำสั่ง `go mod init myproject` (go จะสร้างไฟล์ go.mod ใน Prject ของคุณ)

    ~~~
    go mod init myproject
    ~~~

## Ex1 คำนวณพื้นที่สี่เหลี่ยมผืนผ้าให้ถูกต้อง (10 คะแนน)
### แก้ไข Function calculateArea() คำนวณพื้นที่สี่เหลี่ยมผืนผ้าให้ถูกต้อง โดยรับค่าความกว้างและความยาวจากตัวแปร width และ height แล้ว Return ผลลัพธ์

## Ex2 ผลรวมของ Array (10 คะแนน)
### สร้างฟังก์ชัน sumArray ที่รับ Array ของจำนวนเต็มและคืนค่าผลรวมของตัวเลขทั้งหมดใน Array

4. รัน Unit Test เพื่อทดสอบความถูกต้องของ Code ด้วยคำสั่ง `go test -v` (หากโค้ดและ Unit Test ทำงานถูกต้อง คุณจะได้ผลลัพธ์เป็น PASS)

    ~~~
    go test -v
    ~~~

5. บันทึก Version ของ Code ลง Git Repo ด้วยคำสั่ง `git add .` และ `git commit -m 'commit message'` (แก้ไข commit message เพื่ออธิบายว่าคุณทำอะไรกับ Code ที่บันทึกลง Git)

    ~~~
    git add .
    git commit -m 'edit function calculateArea'
    ~~~

6. Submit Code ขึ้น Github เพื่อตรวจ Code และให้คะแนน ด้วยคำสั่ง `git push origin main`
    
    ~~~
    git push origin main
    ~~~

7. ไปยัง Git Repo ของ Assignment นี้บน github.com กด `Actions` (Github จะใช้เวลาในการตรวจคำตอบสักพัก) เลือก Commit ล่าสุดเพื่อดูคะแนน
