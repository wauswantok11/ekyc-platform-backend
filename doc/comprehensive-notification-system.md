graph TD
    subgraph "Trigger Events"
        A["Application Submitted"]
        B["Document Uploaded"]
        C["Biometric Completed"]
        D["Verification Approved"]
        E["Verification Rejected"]
        F["Manual Review Required"]
        G["Expiry Warning"]
        H["System Maintenance"]
    end
    
    subgraph "Event Processing"
        A --> I["Event Queue"]
        B --> I
        C --> I
        D --> I
        E --> I
        F --> I
        G --> I
        H --> I
        
        I --> J["Event Processor"]
        J --> K["Determine Recipients"]
        K --> L["Select Notification Channels"]
        L --> M["Generate Message Content"]
    end
    
    subgraph "Channel Selection Logic"
        M --> N{"Notification Type?"}
        N -->|"Critical"| O["SMS #43; Email #43; Push"]
        N -->|"Important"| P["Email #43; Push"]
        N -->|"Informational"| Q["Email Only"]
        N -->|"Marketing"| R["Email (if opted in)"]
    end
    
    subgraph "Message Personalization"
        O --> S["Load User Preferences"]
        P --> S
        Q --> S
        R --> S
        
        S --> T["Apply Language Settings"]
        T --> U["Insert Dynamic Content"]
        U --> V["Apply Branding"]
        V --> W["Validate Message Format"]
    end
    
    subgraph "SMS Channel"
        W --> X["SMS Gateway"]
        X --> Y["Format for SMS"]
        Y --> Z["Character Limit Check"]
        Z --> AA{"Length OK?"}
        AA -->|"No"| BB["Truncate Message"]
        AA -->|"Yes"| CC["Send SMS"]
        BB --> CC
        
        CC --> DD{"SMS Sent?"}
        DD -->|"Yes"| EE["Update Delivery Status"]
        DD -->|"No"| FF["Retry Logic"]
        
        FF --> GG{"Retry Count < 3?"}
        GG -->|"Yes"| HH["Wait #43; Retry SMS"]
        GG -->|"No"| II["Mark SMS Failed"]
        
        HH --> CC
    end
    
    subgraph "Email Channel"
        W --> JJ["Email Service"]
        JJ --> KK["Load Email Template"]
        KK --> LL["Render HTML Content"]
        LL --> MM["Add Attachments"]
        MM --> NN["Send Email"]
        
        NN --> OO{"Email Sent?"}
        OO -->|"Yes"| PP["Track Email Opens"]
        OO -->|"No"| QQ["Retry Logic"]
        
        QQ --> RR{"Retry Count < 3?"}
        RR -->|"Yes"| SS["Wait #43; Retry Email"]
        RR -->|"No"| TT["Mark Email Failed"]
        
        SS --> NN
        PP --> UU["Track Link Clicks"]
    end
    
    subgraph "Push Notification Channel"
        W --> VV["Push Service"]
        VV --> WW["Check Device Registration"]
        WW --> XX{"Device Registered?"}
        XX -->|"No"| YY["Skip Push Notification"]
        XX -->|"Yes"| ZZ["Send Push Notification"]
        
        ZZ --> AAA{"Push Sent?"}
        AAA -->|"Yes"| BBB["Track Push Delivery"]
        AAA -->|"No"| CCC["Retry Push"]
        
        CCC --> DDD{"Retry Count < 2?"}
        DDD -->|"Yes"| EEE["Wait #43; Retry Push"]
        DDD -->|"No"| FFF["Mark Push Failed"]
        
        EEE --> ZZ
    end
    
    subgraph "Delivery Tracking & Analytics"
        EE --> GGG["Notification Analytics"]
        UU --> GGG
        BBB --> GGG
        II --> GGG
        TT --> GGG
        FFF --> GGG
        YY --> GGG
        
        GGG --> HHH["Update Delivery Metrics"]
        HHH --> III["Generate Reports"]
        III --> JJJ["Alert on Failures"]
        JJJ --> KKK["Optimize Delivery Times"]
    end
    
    subgraph "User Preferences Management"
        KKK --> LLL["User Notification Settings"]
        LLL --> MMM["Channel Preferences"]
        MMM --> NNN["Frequency Settings"]
        NNN --> OOO["Content Preferences"]
        OOO --> PPP["Opt-out Management"]
    end