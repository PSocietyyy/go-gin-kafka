# API Documentation - Ticket Booking System

This document provides a comprehensive overview of all available REST API endpoints across the microservices in the Ticket Booking System.

---

## 1. Auth Service

**Base URL:** `http://localhost:8080`

### Authentication

#### `POST /auth/login`

- **Description:** Login user and receive JWT
- **Request Body:**
  ```json
  {
    "email": "...",
    "password": "..."
  }
  ```

#### `POST /auth/register`

- **Description:** Register a new user
- **Request Body:**
  ```json
  {
    "name": "...",
    "email": "...",
    "password": "..."
  }
  ```

### Users Management (Protected)

#### `GET /users`

- **Description:** Get all users

#### `POST /users`

- **Description:** Create a new user (Admin)
- **Request Body:**
  ```json
  {
    "name": "...",
    "email": "...",
    "password": "..."
  }
  ```

#### `GET /users/:id`

- **Description:** Get specific user detail

#### `PUT /users/:id`

- **Description:** Update user details
- **Request Body:**
  ```json
  {
    "name": "..."
  }
  ```

#### `DELETE /users/:id`

- **Description:** Delete a user

---

## 2. Product Service

**Base URL:** `http://localhost:8081`

### Trains

#### `GET /trains`

- **Description:** Get all trains

#### `GET /trains/:id`

- **Description:** Get specific train

#### `POST /trains`

- **Description:** Create a new train
- **Request Body:**
  ```json
  {
    "name": "...",
    "code": "..."
  }
  ```

#### `PUT /trains/:id`

- **Description:** Update a train
- **Request Body:**
  ```json
  {
    "name": "...",
    "code": "..."
  }
  ```

#### `DELETE /trains/:id`

- **Description:** Delete a train

### Train Seats (Physical Seats)

#### `GET /trains/:id/seats`

- **Description:** Get all physical seats for a specific train

#### `GET /train_seats/:id`

- **Description:** Get a specific train seat detail

#### `POST /train_seats`

- **Description:** Create a physical seat
- **Request Body:**
  ```json
  {
    "train_id": 1,
    "seat_number": "A1"
  }
  ```

#### `PUT /train_seats/:id`

- **Description:** Update physical seat
- **Request Body:**
  ```json
  {
    "seat_number": "A2"
  }
  ```

#### `DELETE /train_seats/:id`

- **Description:** Delete physical seat

### Schedules

#### `GET /schedules`

- **Description:** Get all schedules

#### `GET /schedules/:id`

- **Description:** Get a specific schedule

#### `POST /schedules`

- **Description:** Create a new schedule
- **Request Body:**
  ```json
  {
    "train_id": 1,
    "departure_time": "...",
    "arrival_time": "..."
  }
  ```

#### `PUT /schedules/:id`

- **Description:** Update a schedule
- **Request Body:**
  ```json
  {
    "departure_time": "..."
  }
  ```

#### `DELETE /schedules/:id`

- **Description:** Delete a schedule

### Schedule Seats (Booking Status)

#### `GET /schedules/:id/seats`

- **Description:** Get all seat mapped for a schedule including their status (Available/Booked)

#### `GET /schedule_seats/:id`

- **Description:** Get detail for a specific schedule seat

#### `POST /schedule_seats`

- **Description:** Map a physical train seat into a schedule
- **Request Body:**
  ```json
  {
    "schedule_id": 1,
    "train_seat_id": 1,
    "status": "AVAILABLE"
  }
  ```

#### `POST /schedule_seats/:id/book`

- **Description:** Book a schedule seat

#### `POST /schedule_seats/:id/cancel`

- **Description:** Cancel a booking for a schedule seat

---

## 3. Order Service

**Base URL:** `http://localhost:8082`

### Orders

#### `GET /orders`

- **Description:** Get all orders

#### `GET /orders/:id`

- **Description:** Get specific order detail

#### `POST /orders`

- **Description:** Create a new order
- **Request Body:**
  ```json
  {
    "user_id": 1,
    "schedule_seat_id": 1
  }
  ```

#### `PUT /orders/:id`

- **Description:** Update order status
- **Request Body:**
  ```json
  {
    "status": "PAID"
  }
  ```
