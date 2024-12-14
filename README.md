# Crypto Dashboard Backend

A scalable microservices-based cryptocurrency dashboard backend system built with Go.

## System Architecture

The system follows a microservices architecture pattern with the following key components:

- API Gateway - Central entry point for all client requests
- Microservices:
  - User Service - Authentication and user management 
  - Market Data Service - Real-time crypto market data
  - Whale Tracking Service - Large transaction monitoring
  - Portfolio Service - User portfolio management
  - Notification Service - Real-time alerts and notifications

[See detailed architecture diagrams](docs/architecture/diagrams/architecture.md)


## Getting Started

### Prerequisites

- Go 1.20+
- Docker & Docker Compose 
- PostgreSQL 14+
- Redis 6+
- Kafka/RabbitMQ

### Development Setup

1. Clone the repository
```bash
git clone https://github.com/vietgs03/crypto-dashboard-backend
cd crypto-dashboard-backend
```

2. Copy configuration files

```bash
cp configs/development/config.example.yaml configs/development/config.yaml
```

3. Start dependencies with Docker Compose

```bash
docker-compose up -d
```


### Project Structure

```
├── cmd/                 # Service entry points
├── internal/            # Private application code
├── pkg/                # Shared libraries
├── configs/            # Configuration files
├── docs/              # Documentation
└── tests/             # Tests
```
See complete structure

### Development Guidelines

See our Development Guide for:

- Code organization

- Clean architecture principles

- SOLID principles implementation

- Testing guidelines

- Error handling

### Key Features

See Key Features for detailed feature list including:

- Real-time market data

- Whale activity tracking

- Portfolio management

- Alert system

- Technical analysis tools

### API Documentation

API documentation is available at:

- Development: http://localhost:8080/swagger/

- Production: https://api.your-domain.com/swagger/


### Deployment

1. Build Docker images

```bash
docker build -t crypto-dashboard-backend .
```

2. Deploy to Kubernetes

```bash
kubectl apply -f configs/kubernetes/
```

### Monitoring

- Metrics: Prometheus (localhost:9090)

- Dashboards: Grafana (localhost:3000)

- Logs: ELK Stack

### Contributing

1. Follow Development Guidelines

2. Create feature branch

3. Make changes

4. Submit PR with description

## Document 

### Architecture
- [System Architecture](docs//architecture/diagrams/architecture.md)
- [Database Schema](docs/architecture/diagrams/database-schema.md)
- [Service Communication](docs/architecture/diagrams/service-communication.md)
- [Security Flow](docs/architecture/diagrams/security.md)

### Development
- [Development Guidelines](docs/devguide.md)
- [Code Standards](docs/devguide.md#coding-standards)
- [Testing Strategy](docs/devguide.md#testing-guidelines)
- [Error Handling](docs/devguide.md#error-handling)

### Features
- [Key Features](docs/keyfeature.md)
- [API Documentation](docs/api/README.md)
- [Service Capabilities](docs/keyfeature.md#core-features)

### Operations
- [Scripts Guide](docs/scripts.md)
- [Deployment Guide](docs/deployment.md)
- [Monitoring Setup](docs/monitoring.md)

## License

MIT License

Contact
For support or queries, contact the team at: team@your-domain.com

```

This README provides a comprehensive overview of the project while linking to the detailed documentation available in the workspace. It follows best practices for open source projects and includes all necessary information for developers to get started.
This README provides a comprehensive overview of the project while linking to the detailed documentation available in the workspace. It follows best practices for open source projects and includes all necessary information for developers to get started.

```