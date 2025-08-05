graph LR
    subgraph Input[User Input]
        A[ID Card Photos]
        B[Personal Info]
        C[Biometric Data]
    end
    
    subgraph Processing[Processing]
        D[Document Validation]
        E[Data Extraction]
        F[Biometric Analysis]
        G[LINE Net KYC]
    end
    
    subgraph Storage[Storage]
        H[Encrypted Database]
        I[Secure File Storage]
        J[Audit Logs]
    end
    
    subgraph Output[Output]
        K[Verification Result]
        L[Digital Certificate]
        M[Notifications]
    end
    
    A --> D
    B --> E
    C --> F
    D --> G
    E --> G
    F --> G
    G --> H
    G --> I
    G --> J
    H --> K
    I --> L
    J --> M