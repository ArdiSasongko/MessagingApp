# ChatApp - Real-time Chat Application

This project serves as a practical implementation of building real-time applications using modern technologies. It was developed to demonstrate the understanding of web socket implementation, API development, and database management while serving as a portfolio piece. The application showcases the integration of Golang's powerful backend capabilities with MongoDB's flexible data storage and WebSocket's real-time communication features.

## Tech Stack
- **Backend**: Go (Fiber Framework)
- **Database**: 
  - MongoDB (Chat Messages)
  - PostgreSQL (User Data)
- **Real-time Communication**: WebSocket
- **Authentication**: JWT (JSON Web Tokens)

## Features

### Authentication System
- **User Registration**: Secure signup process with email verification
  - Password hashing using bcrypt
  - Email validation
  - Username uniqueness check
  
- **Login System**: JWT-based authentication
  - Generates access token (24h validity)
  - Generates refresh token (7 days validity)
  - Secure password comparison
  
- **Token Management**:
  - Refresh token rotation
  - Access token renewal
  - Token blacklisting for logout
  
- **Logout Mechanism**:
  - Invalidates active tokens
  - Clears session data
  - Handles multi-device logout

### Real-time Chat Features
- **WebSocket Implementation**:
  - Persistent connection management
  - Real-time message delivery
  - Online status tracking
  - Typing indicators
  
- **Message Management**:
  - Real-time message sending and receiving
  - Message history retrieval
  - Read receipts
  - Message pagination

### Database Architecture
- **PostgreSQL (User Data)**:
  - User profiles
  - Authentication records
  - Account settings
  - User relationships
  
- **MongoDB (Messages)**:
  - Chat messages
  - Media attachments
  - Message metadata
  - Chat history

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
