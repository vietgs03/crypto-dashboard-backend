# Crypto Dashboard Architecture

## System Architecture Diagram

```mermaid
graph TB
    %% External Data Sources
    subgraph External["🌐 External Data Sources"]
        CGAPI["🚀 CG\nCoinGecko\nMarket Data"]
        BNAPI["💱 BN\nBinance\nTrading Info"]
        WHAPI["🐳 WA\nWhaleAlert\nTransaction Tracking"]
    end

    %% Frontend Application
    subgraph FE["💻 User Interface"]
        UI["🖥️ Dashboard\nCrypto Insights"]
        WS["📡 WebSocket\nLive Streaming"]
        Redux["📊 State Mgmt\nData Handling"]
    end

    %% API Gateway
    API["🔐 API Gateway\nRequest Routing"]

    %% Microservices
    subgraph Services["🔧 Microservices"]
        AUTH["👤 User Service\nAuth & Management"]
        MARKET["💹 Market Service\nPrice Analysis"]
        WHALE["🐬 Whale Tracker\nTransaction Monitor"]
        PORT["💼 Portfolio\nInvestment Tracking"]
        NOTIF["🔔 Notification\nReal-time Alerts"]
        ANAL["📈 Analytics\nMarket Insights"]
    end

    %% Data Layer
    subgraph Data["💾 Data Infrastructure"]
        REDIS["🚄 Redis\nFast Caching"]
        PGSQL["🗃️ PostgreSQL\nPersistent Storage"]
        KAFKA["📨 Kafka\nEvent Streaming"]
    end

    %% Monitoring
    subgraph Monitor["🔍 System Monitoring"]
        PROM["📊 Prometheus\nMetrics Collection"]
        GRAF["📉 Grafana\nVisualization"]
        LOG["📋 ELK Stack\nLog Management"]
    end

    %% Connections with custom styling
    linkStyle default stroke:#666,stroke-width:2px
    CGAPI --> |Market Data| MARKET
    BNAPI --> |Trading Info| MARKET
    WHAPI --> |Transaction Data| WHALE
    
    UI --> |User Requests| API
    WS --> |Real-time Updates| API
    
    API --> AUTH
    API --> MARKET
    API --> WHALE
    API --> PORT
    API --> NOTIF
    API --> ANAL
    
    AUTH --> REDIS
    AUTH --> PGSQL
    MARKET --> REDIS
    MARKET --> KAFKA
    WHALE --> REDIS
    WHALE --> KAFKA
    PORT --> PGSQL
    NOTIF --> KAFKA
    ANAL --> PGSQL
    
    KAFKA --> NOTIF
    
    Services --> PROM
    PROM --> GRAF
    Services --> LOG

    %% Color Theme
    classDef external fill:#FFE5B4,stroke:#FF9800
    classDef frontend fill:#E6F2FF,stroke:#2196F3
    classDef services fill:#E8F5E9,stroke:#4CAF50
    classDef data fill:#FFF3E0,stroke:#FF5722
    classDef monitor fill:#F3E5F5,stroke:#9C27B0

    class External external
    class FE frontend
    class Services services
    class Data data
    class Monitor monitor
```

## Data Flow Diagram

```mermaid
graph LR
    %% User Interactions
    subgraph Users["👥 User Actions"]
        USR["User"]
        direction TB
        USR -->|View Dashboard| REQ1[/"Request Data"/]
        USR -->|Set Alerts| REQ2[/"Configure Alerts"/]
        USR -->|Track Portfolio| REQ3[/"Portfolio Actions"/]
    end

    %% API Processing
    subgraph Processing["⚡ Request Processing"]
        GATE["API Gateway"]
        CACHE["Redis Cache"]
        direction TB
        REQ1 --> GATE
        REQ2 --> GATE
        REQ3 --> GATE
        GATE -->|Check| CACHE
    end

    %% Services Layer
    subgraph Backend["🔧 Service Layer"]
        MKT["Market Service"]
        ALERT["Alert Service"]
        PORT["Portfolio Service"]
        direction TB
        GATE -->|Market Data| MKT
        GATE -->|Alerts| ALERT
        GATE -->|Portfolio| PORT
    end

    %% Data Storage
    subgraph Storage["💾 Data Layer"]
        DB[(PostgreSQL)]
        MQ{"Message Queue"}
        direction TB
        MKT --> DB
        ALERT --> MQ
        PORT --> DB
        MQ -->|Notifications| ALERT
    end

    %% External Data
    subgraph External["🌐 External Sources"]
        API1["CoinGecko"]
        API2["Binance"]
        API3["WhaleAlert"]
        direction TB
        API1 -->|Market Data| MKT
        API2 -->|Trade Data| MKT
        API3 -->|Whale Activity| ALERT
    end

    %% Styling
    classDef users fill:#E3F2FD,stroke:#1976D2
    classDef process fill:#E8F5E9,stroke:#388E3C
    classDef backend fill:#FFF3E0,stroke:#F57C00
    classDef storage fill:#FCE4EC,stroke:#C2185B
    classDef external fill:#F3E5F5,stroke:#7B1FA2

    class Users users
    class Processing process
    class Backend backend
    class Storage storage
    class External external
```

## Plainning Deployment Architecture

```mermaid
graph TB
    subgraph Cloud["☁️ Cloud Infrastructure"]
        direction TB
        subgraph LB["Load Balancers"]
            ALB["Application LB"]
            NLB["Network LB"]
        end

        subgraph K8S["Kubernetes Clusters"]
            direction TB
            subgraph Services["Service Pods"]
                API["API Gateway"]
                MS["Microservices"]
                DB["Database"]
            end
            
            subgraph Scaling["Auto Scaling"]
                HPA["Horizontal Pod Autoscaler"]
                VPA["Vertical Pod Autoscaler"]
            end
        end

        subgraph Storage["Persistent Storage"]
            EBS["Block Storage"]
            S3["Object Storage"]
        end

        subgraph Network["Network"]
            VPC["Virtual Private Cloud"]
            SG["Security Groups"]
        end
    end

    Internet((Internet)) --> ALB
    ALB --> API
    API --> MS
    MS --> DB
    DB --> EBS
    
    classDef cloud fill:#E1F5FE,stroke:#0288D1
    classDef network fill:#E8F5E9,stroke:#388E3C
    classDef storage fill:#FFF3E0,stroke:#F57C00
    
    class Cloud cloud
    class Network network
    class Storage storage
```


## Plan Communication Service

```mermaid
graph TB
    %% Service Discovery Layer
    subgraph Discovery["🔍 Service Discovery"]
        CONSUL["Consul Registry"]
        LB["Load Balancer"]
    end

    %% Communication Patterns
    subgraph Sync["↔️ Synchronous"]
        REST["REST APIs"]
        GRPC["gRPC Calls"]
    end

    subgraph Async["📨 Asynchronous"]
        KAFKA["Kafka Events"]
        REDIS["Redis Pub/Sub"]
    end

    %% Circuit Breakers
    subgraph Resilience["🔄 Resilience"]
        CB["Circuit Breaker"]
        RT["Retry Logic"]
        TO["Timeout Handler"]
    end

    %% Services
    subgraph Services["🔧 Services"]
        AUTH["Auth Service"]
        MARKET["Market Service"]
        WHALE["Whale Service"]
        PORT["Portfolio Service"]
        NOTIF["Notification Service"]
    end

    %% Connections
    Services -->|Register| CONSUL
    CONSUL -->|Resolve| LB
    LB -->|Route| Services

    Services -->|API Calls| REST
    Services -->|Stream| GRPC
    Services -->|Events| KAFKA
    Services -->|Real-time| REDIS

    REST --> CB
    GRPC --> CB
    CB --> RT
    RT --> TO

    %% Styling
    classDef discovery fill:#E3F2FD,stroke:#1976D2
    classDef sync fill:#E8F5E9,stroke:#388E3C
    classDef async fill:#FFF3E0,stroke:#F57C00
    classDef resilience fill:#FCE4EC,stroke:#C2185B
    classDef services fill:#F3E5F5,stroke:#7B1FA2

    class Discovery discovery
    class Sync sync
    class Async async
    class Resilience resilience
    class Services services

```

## DB digram

```mermaid
erDiagram
    USERS {
        uuid user_id PK
        string email UK
        string password_hash
        timestamp created_at
        timestamp updated_at
    }

    PORTFOLIOS {
        uuid portfolio_id PK
        uuid user_id FK
        string name
        decimal total_value
        timestamp updated_at
    }

    ASSETS {
        uuid asset_id PK
        uuid portfolio_id FK
        string symbol
        decimal quantity
        decimal avg_buy_price
        timestamp updated_at
    }

    TRANSACTIONS {
        uuid tx_id PK
        uuid asset_id FK
        string type
        decimal amount
        decimal price
        timestamp created_at
    }

    ALERTS {
        uuid alert_id PK
        uuid user_id FK
        string crypto_pair
        decimal trigger_price
        string condition
        boolean is_active
        timestamp created_at
    }

    MARKET_DATA {
        string symbol PK
        decimal price
        decimal volume_24h
        decimal market_cap
        timestamp updated_at
    }

    WHALE_TRANSACTIONS {
        uuid whale_tx_id PK
        string from_address
        string to_address
        decimal amount
        timestamp created_at
    }

    USERS ||--o{ PORTFOLIOS : has
    PORTFOLIOS ||--o{ ASSETS : contains
    ASSETS ||--o{ TRANSACTIONS : records
    USERS ||--o{ ALERTS : sets
    WHALE_TRANSACTIONS }|--|| MARKET_DATA : references

```

## Security Architecture Planning


```mermaid
graph TB
    %% User Authentication Flow
    subgraph Auth["🔐 Authentication"]
        LOGIN["Login Request"]
        MFA["2FA Verification"]
        JWT["JWT Generation"]
        REFRESH["Token Refresh"]
    end

    %% Authorization Layers
    subgraph Access["🛡️ Authorization"]
        RBAC["Role-Based Access"]
        PERMS["Permissions"]
        POLICY["Policy Engine"]
    end

    %% Data Security
    subgraph Data["🔒 Data Protection"]
        ENCRYPT["Encryption at Rest"]
        TLS["TLS/SSL"]
        MASK["Data Masking"]
    end

    %% Network Security
    subgraph Network["🌐 Network Security"]
        WAF["Web Application Firewall"]
        DDOS["DDoS Protection"]
        VPN["VPN Access"]
    end

    %% Security Monitoring
    subgraph Monitor["👁️ Security Monitoring"]
        AUDIT["Audit Logs"]
        IDS["Intrusion Detection"]
        ALERT["Security Alerts"]
    end

    %% Flow Connections
    LOGIN --> MFA
    MFA --> JWT
    JWT --> RBAC
    RBAC --> PERMS
    PERMS --> POLICY

    LOGIN --> WAF
    WAF --> DDOS
    
    POLICY --> ENCRYPT
    POLICY --> MASK
    
    WAF --> AUDIT
    DDOS --> IDS
    IDS --> ALERT

    %% Styling
    classDef auth fill:#E8EAF6,stroke:#3F51B5
    classDef access fill:#E8F5E9,stroke:#4CAF50
    classDef data fill:#FFF3E0,stroke:#FF9800
    classDef network fill:#F3E5F5,stroke:#9C27B0
    classDef monitor fill:#FFEBEE,stroke:#F44336

    class Auth auth
    class Access access
    class Data data
    class Network network
    class Monitor monitor
```