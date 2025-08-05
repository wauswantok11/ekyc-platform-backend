graph TD
    A[เหตุการณ์เกิดขึ้น] --> B{ประเภทการแจ้งเตือน}
    
    B -->|Critical| C[SMS + Email + Push]
    B -->|Important| D[Email + Push]
    B -->|Info| E[Email Only]
    
    C --> F[ส่ง SMS]
    C --> G[ส่ง Email]
    C --> H[ส่ง Push Notification]
    
    D --> G
    D --> H
    
    E --> G
    
    F --> I{SMS สำเร็จ?}
    G --> J{Email สำเร็จ?}
    H --> K{Push สำเร็จ?}
    
    I -->|ไม่| L[ลองส่ง SMS ใหม่]
    J -->|ไม่| M[ลองส่ง Email ใหม่]
    K -->|ไม่| N[ลองส่ง Push ใหม่]
    
    I -->|ใช่| O[บันทึกสถานะ SMS]
    J -->|ใช่| P[บันทึกสถานะ Email]
    K -->|ใช่| Q[บันทึกสถานะ Push]
    
    L --> R{ลองครั้งที่ 3?}
    M --> S{ลองครั้งที่ 3?}
    N --> T{ลองครั้งที่ 2?}
    
    R -->|ใช่| U[หยุดส่ง SMS]
    S -->|ใช่| V[หยุดส่ง Email]
    T -->|ใช่| W[หยุดส่ง Push]
    
    R -->|ไม่| F
    S -->|ไม่| G
    T -->|ไม่| H