graph TB
    subgraph Frontend[Frontend Layer]
        A[Next.js Web App]
        B[React Components]
        C[Tailwind CSS]
    end
    
    subgraph API[API Layer]
        D[Next.js API Routes]
        E[Server Actions]
        F[Authentication]
    end
    
    subgraph External[External Services]
        G[LINE Net KYC API]
        H[Camera/Media API]
        I[SMS Gateway]
        J[Email Service]
    end
    
    subgraph Database[Database Layer]
        K[User Data]
        L[Verification Records]
        M[Document Storage]
    end
    
    subgraph Admin[Admin Dashboard]
        N[Dashboard UI]
        O[User Management]
        P[Reports & Analytics]
    end
    
    A --> D
    B --> A
    C --> B
    D --> E
    E --> F
    D --> G
    D --> H
    E --> I
    E --> J
    E --> K
    E --> L
    E --> M
    N --> D
    O --> N
    P --> N