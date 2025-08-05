graph TD
    A[เลือกประเภทการเข้าสู่ระบบ] --> B{ประเภทผู้ใช้}
    
    B -->|OneID| C[กรอกเบอร์โทร + เลขบัตรประชาชน]
    B -->|Service| D[กรอก Client ID + Secret Key]
    B -->|Business| E[กรอก Client ID + Secret Key + Shared Token]
    
    C --> F[ตรวจสอบข้อมูล OneID]
    D --> G[ตรวจสอบ Service Credentials]
    E --> H[ตรวจสอบ Business Credentials]
    
    F --> I{การตรวจสอบ}
    G --> I
    H --> I
    
    I -->|สำเร็จ| J[สร้าง Session]
    I -->|ไม่สำเร็จ| K[แสดงข้อผิดพลาด]
    
    J --> L[เข้าสู่ระบบสำเร็จ]
    K --> A