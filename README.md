# ChatApp - Real-time Chat Application

This project serves as a practical implementation of building real-time applications using modern technologies. It was developed to demonstrate the understanding of web socket implementation, API development, and database management while serving as a portfolio piece. The application showcases the integration of Golang's powerful backend capabilities with MongoDB's flexible data storage and WebSocket's real-time communication features.

## Tech Stack
- **Backend**: Go (Fiber Framework)
- **Database**: 
  - MongoDB (Chat Messages)
  - PostgreSQL (User Data)
- **Real-time Communication**: WebSocket
- **Authentication**: JWT (JSON Web Tokens)

## API Endpoints

### Authentication
```
POST /user/v1/register   - User registration
POST /user/v1/login      - User login
POST /user/v1/logout     - User logout
POST /user/v1/refresh-token    - Refresh access token
```

### Chat Operations
```
WS   /ws/            - WebSocket connection endpoint
GET  /message/v1/history       - Retrieve message history
POST /message/v1/send       - Send new message
```
