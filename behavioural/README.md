# Behavioural

## Observer
### Description
pattern นี้ให้ subject ส่งอีเวลล์หา object อื่นหลายชนิดที่มี interface ร่วมกัน
1. สร้าง interface ของ observer
2. implement ตัว observer กับ object ที่จะทำเป็น observer
3. สร้าง subject เพื่อใช้งาน observer
### Use Cases
- Email subscribe
- E-commerce : ลูกค้าเช็คว่าของมีใน stock หรือไม่

## Template
template(interface) -> compose methods()

object A(template) -> implement interface method with own logic
object B(template) -> implement interface method with own logic
