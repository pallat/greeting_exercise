# Greeting Exercise

## ข้อมูลเพิ่มเติม

ในภาษา Go การเปลี่ยน type ของตัวแปร ที่เป็น basic type แบ่งเป็น 2 แบบ
    1. แบบที่มีโครงสร้างพื้นฐานเดียวกัน เช่น ตัวเลขจาก int -> uint หรือ int -> float64
    2. แบบที่มีโครงสร้างพื้นฐานต่างกัน เช่น string -> int หรือ int -> string

1. แบบแรก สามารถทำได้ด้วยการทำได้ด้วยการทำ type casting เช่น
   int -> float64:

   ```go
   var n int = 10
   f := float64(n)
   ```

2. แบบที่สองต้องอาศัยตัวช่วย เช่น
   int -> string

   ```go
    var n int = 10
    var s string
    s = strconv.Itoa(n)
   ```

   หรือ

   ```go
    var n int = 10
    var s string
    s = fmt.Sprintf("%d",n)
   ```
