## Go-Health-ID
ออกแบบโครงสร้างคล้าย Hexagonal แต่จะย่อๆ บางอย่างลงให้ดูไม่ซับซ้อน

### คำแนะนำ
- HttpClient ให้เรียกใช้ http ที่อยู่ใน struct ของ repository ที่จะส่ง stats ไปที่ syslog server
- Log ให้เรียกใช้ Log ที่อยู่ใน struct ของ repository ที่จะ mark field เพื่อแยก log package และ module ได้งง่ายขึ้น
- JSON Marshal/Unmarshal ให้ใช้ Lib ที่ชื่อว่า sonic มีลักษณะการใช้งานเหมือน Lib json มาตรฐาน
- Hexagonal กึ่ง MVC จะมีด้วยกันคือ
  - Handler หรือจะเรียกอีกชื่อหนึ่งว่า Controller ไว้จัดการ Request/Response ที่เข้ามาจาก API ด้วย DTO Struct (Data Transfer Object)
  - Service เป็น Business Logic ของงานนั้นๆ
  - Repository เป็นศูนย์รวม Repo ที่ใช้เชื่อมต่อกับระบบภายนอกทั้งหมด

### Module
- Autobot สำหรับ automate รับคำสั่งจาก line bot
- Frontweb สำหรับใช้รับ request จาก frontend web
- frontweb สำหรับใช้รับ request api ที่เปิดให้คนอื่นเข้ามาใช้งาน
- B2B สำหรับใช้ Back 2 Back MOPH App

### สำหรับนักพัฒนา
- สิ่งที่ต้องใช้
  - ไฟล์ .env ที่ define user/pass ต่างๆ ใน docker-compose
  - ไฟล์ config.yml เป็น file config ของ server
- เมื่อเตรียม Dependencies ครบแล้ว ก็ใช้คำสั่ง `make dev` เพื่อ start server
- หากต้องการ build docker ก็ให้ใช้คำสั่ง `make beta` เพื่อ build docker image ภายในชื่อ tag ว่า `go-health-id:oss`
- ตัว Server ใช้ CI/CD ในการ Deploy ระบบขึ้น Production เท่านั้น โดยจะอ้างอิงจาก Tag version ใน Branch Master