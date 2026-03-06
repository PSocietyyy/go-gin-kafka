# Dokumentasi API

## Auth Service

Base URL:`http://localhost:8080/auth`

### Login

URL: `/login`
Method: `POST`
Request Body:

```json
{
  "email": "ferdi@example.test",
  "password": "password"
}
```

Response Body:

```json
{
  "token": "eyJhbxxxx....."
}
```

Response Error:

```json
{
  "error": "invalid password"
}
```

### Register

URL: `/register`
Method: `POST`
Request Body:

```json
{
  "name": "ferdiansyah",
  "email": "ferdia@example.test",
  "password": "password"
}
```

Response Body:

```json
{
  "message": "User created successfully"
}
```

Response Error:
**Duplicate Email**

```json
{
  "error": "Error 1062 (23000): Duplicate entry 'ferdia@example.test' for key 'users.uni_users_email'"
}
```

## User Service

Base URL:`http://localhost:8080/users`

### Get All Users

URL: `/users`
Method: `GET`

Response Body:

```json
{
  "data": [
    {
      "id": 1,
      "name": "ferdiansyah pratama",
      "email": "ferdi@example.test"
    },
    {
      "id": 5,
      "name": "ferdiansyah",
      "email": "ferdia@example.test"
    },
    {
      "id": 7,
      "name": "ferdiansyah",
      "email": "ferdian@example.test"
    }
  ]
}
```

### Create User

URL: `/users`
Method: `POST`
Request Body:

```json
{
  "name": "admin salah",
  "email": "admin@example.test",
  "password": "password"
}
```

Response Body:

```json
{
  "message": "User created successfully"
}
```

Response Error:
**Duplicate Email**

```json
{
  "error": "Error 1062 (23000): Duplicate entry 'ferdia@example.test' for key 'users.uni_users_email'"
}
```

### Get User By ID

URL: `/users/:id`
Method: `GET`

Response Body:

```json
{
  "data": {
    "id": 1,
    "name": "ferdiansyah pratama",
    "email": "ferdi@example.test"
  }
}
```

Response Error:
**User Not Found**

```json
{
  "error": "user not found"
}
```

### Update User

URL: `/users/:id`
Method: `PUT`
Request Body:

```json
{
  "name": "ferdiansyah pratama",
  "email": "ferdi@example.test",
  "password": "password"
}
```

Response Body:

```json
{
  "message": "User updated successfully"
}
```

Response Error:
**User Not Found**

```json
{
  "error": "user not found"
}
```

**Duplicate Email**
**Kecuali email milik sendiri**

```json
{
  "error": "email already exists"
}
```

### Delete User

URL: `/users/:id`
Method: `DELETE`

Response Body:

```json
{
  "message": "User deleted successfully"
}
```

## Product Service

Base URL: `http://localhost:8081`

### Trains

#### Create Train

URL: `/trains`
Method: `POST`
Request Body:

```json
{
  "name": "Argo Bromo Anggrek",
  "code": "ABA-01"
}
```

Response Body:

```json
{
  "message": "Train created successfully"
}
```

Response Error:
**Duplicate Train**

```json
{
  "error": "Error 1062 (23000): Duplicate entry 'ABA-01' for key 'trains.uni_trains_code'"
}
```

#### Get All Trains

URL: `/trains`
Method: `GET`

Response Body:

```json
{
  "data": [
    {
      "id": 1,
      "name": "Argo Bromo Anggrek",
      "code": "ABA-01"
    }
  ],
  "message": "Get All Trains Successfully"
}
```

#### Get Train By ID

URL: `/trains/:id`
Method: `GET`

Response Body:

```json
{
  "data": {
    "id": 1,
    "name": "Argo Bromo Anggrek",
    "code": "ABA-01"
  },
  "message": "Get Train By ID Successfully"
}
```

Response Error:
**Train Not Found**

```json
{
  "error": "train not found"
}
```

#### Update Train

URL: `/trains/:id`
Method: `PUT`
Request Body:

```json
{
  "name": "Argo Bromo Anggrek",
  "code": "ABA-01"
}
```

Response Body:

```json
{
  "data": {
    "id": 1,
    "name": "Argo Bromo Anggrek",
    "code": "ABA-01"
  },
  "message": "Train Updated Successfully"
}
```

Response Error:
**Train Not Found**

```json
{
  "message": "Train Not Found"
}
```

**Duplicate Train**
**Kecuali train milik sendiri**

```json
{
  "message": "Failed to Update Train"
}
```

#### Delete Train

URL: `/trains/:id`
Method: `DELETE`

Response Body:

```json
{
  "message": "Train Deleted Successfully"
}
```

### Train Seats

#### Create Train Seat

URL: `/train_seats`
Method: `POST`
Request Body:

```json
{
  "train_id": 1,
  "seat_number": "A1"
}
```

Response Body:

```json
{
  "data": {
    "id": 2,
    "train_id": 2,
    "seat_number": "A2"
  },
  "message": "Train Seat Created Successfully"
}
```

Response Error:
**Train Not Found**

```json
{
  "message": "Train Not Found"
}
```

#### Get All Train Seats

URL: `/train_seats`
Method: `GET`

Response Body:

```json
{
  "data": [
    {
      "id": 1,
      "train_id": 2,
      "seat_number": "A1"
    },
    {
      "id": 2,
      "train_id": 2,
      "seat_number": "A2"
    }
  ],
  "message": "Get All Train Seats Successfully"
}
```

#### Get Train Seat By ID

URL: `/train_seats/:id`
Method: `GET`

Response Body:

```json
{
  "data": [
    {
      "id": 1,
      "train_id": 2,
      "seat_number": "A1"
    },
    {
      "id": 2,
      "train_id": 2,
      "seat_number": "A2"
    }
  ],
  "message": "Get Train Seats By Train ID Successfully"
}
```

Response Error:
**Train Seat Not Found**

```json
{
  "data": [],
  "message": "Get Train Seats By Train ID Successfully"
}
```

#### Update Train Seat

URL: `/train_seats/:id`
Method: `PUT`
Request Body:

```json
{
  "train_id": 2,
  "seat_number": "A2"
}
```

Response Body:

```json
{
  "data": {
    "id": 2,
    "train_id": 2,
    "seat_number": "A4"
  },
  "message": "Train Seat Updated Successfully"
}
```

Response Error:
**Train Seat Not Found**

```json
{
  "message": "Train Seat Not Found"
}
```

**Train Not Found**

```json
{
  "error": "train not found",
  "message": "Failed to Update Train Seat"
}
```

#### Delete Train Seat

URL: `/train_seats/:id`
Method: `DELETE`

Response Body:

```json
{
  "message": "Train Seat Deleted Successfully"
}
```

Response Error:
**Train Seat Not Found**

```json
{
  "message": "Train Seat Not Found"
}
```

### Schedules

#### Create Schedule

URL: `/schedules`
Method: `POST`
Request Body:

```json
{
  "train_id": 1,
  "departure_time": "2026-10-15T08:00:00Z",
  "arrival_time": "2026-10-15T12:00:00Z"
}
```

Response Body:

```json
{
  "data": {
    "id": 2,
    "train_id": 2,
    "train": {
      "id": 0,
      "name": "",
      "code": ""
    },
    "departure_time": "2026-10-15T08:00:00Z",
    "arrival_time": "2026-10-15T12:00:00Z"
  },
  "message": "Schedule Created Successfully"
}
```

Response Error:
**Train Not Found**

```json
{
  "message": "Failed to Create Schedule"
}
```

#### Get All Schedules

URL: `/schedules`
Method: `GET`

Response Body:

```json
{
  "data": [
    {
      "id": 2,
      "train_id": 2,
      "train": {
        "id": 2,
        "name": "Argo Bromo Anggrek",
        "code": "ABA-02"
      },
      "departure_time": "2026-10-15T15:00:00+07:00",
      "arrival_time": "2026-10-15T19:00:00+07:00"
    }
  ],
  "message": "Get All Schedules Successfully"
}
```

#### Get Schedule By ID

URL: `/schedules/:id`
Method: `GET`

Response Body:

```json
{
  "data": {
    "id": 2,
    "train_id": 2,
    "train": {
      "id": 2,
      "name": "Argo Bromo Anggrek",
      "code": "ABA-02"
    },
    "departure_time": "2026-10-15T15:00:00+07:00",
    "arrival_time": "2026-10-15T19:00:00+07:00"
  },
  "message": "Get Schedule By ID Successfully"
}
```

Response Error:
**Schedule Not Found**

```json
{
  "message": "Schedule Not Found"
}
```

#### Update Schedule

URL: `/schedules/:id`
Method: `PUT`
Request Body:

```json
{
  "train_id": 2,
  "departure_time": "2026-10-15T08:00:00Z",
  "arrival_time": "2026-10-15T12:00:00Z"
}
```

Response Body:

```json
{
  "data": {
    "id": 2,
    "train_id": 2,
    "train": {
      "id": 2,
      "name": "Argo Bromo Anggrek",
      "code": "ABA-02"
    },
    "departure_time": "2026-10-15T08:00:00Z",
    "arrival_time": "2026-10-15T12:00:00Z"
  },
  "message": "Schedule Updated Successfully"
}
```

Response Error:
**Schedule Not Found**

```json
{
  "message": "Schedule Not Found"
}
```

**Train Not Found**

```json
{
  "message": "Failed to Update Schedule"
}
```

#### Delete Schedule

URL: `/schedules/:id`
Method: `DELETE`

Response Body:

```json
{
  "message": "Schedule Deleted Successfully"
}
```

Response Error:
**Schedule Not Found**

```json
{
  "message": "Schedule Not Found"
}
```

### Schedule Seats

#### Get All Schedule Seats

URL: `/schedule_seats`
Method: `GET`

Response Body:

```json
{
  "data": [
    {
      "id": 1,
      "schedule_id": 3,
      "train_seat_id": 2,
      "status": "AVAILABLE",
      "schedule": {
        "id": 3,
        "train_id": 2,
        "train": {
          "id": 2,
          "name": "Argo Bromo Anggrek",
          "code": "ABA-02"
        },
        "departure_time": "2026-10-15T15:00:00+07:00",
        "arrival_time": "2026-10-15T19:00:00+07:00"
      },
      "train_seat": {
        "id": 2,
        "train_id": 2,
        "seat_number": "A4"
      }
    },
    {
      "id": 2,
      "schedule_id": 3,
      "train_seat_id": 2,
      "status": "AVAILABLE",
      "schedule": {
        "id": 3,
        "train_id": 2,
        "train": {
          "id": 2,
          "name": "Argo Bromo Anggrek",
          "code": "ABA-02"
        },
        "departure_time": "2026-10-15T15:00:00+07:00",
        "arrival_time": "2026-10-15T19:00:00+07:00"
      },
      "train_seat": {
        "id": 2,
        "train_id": 2,
        "seat_number": "A4"
      }
    }
  ],
  "message": "Get All Schedule Seats Successfully"
}
```

#### Create Schedule Seat

URL: `/schedule_seats`
Method: `POST`
Request Body:

```json
{
  "schedule_id": 3,
  "train_seat_id": 2,
  "status": "AVAILABLE"
}
```

Response Body:

```json
{
  "data": {
    "id": 3,
    "schedule_id": 3,
    "train_seat_id": 2,
    "status": "AVAILABLE",
    "schedule": {
      "id": 0,
      "train_id": 0,
      "train": {
        "id": 0,
        "name": "",
        "code": ""
      },
      "departure_time": "0001-01-01T00:00:00Z",
      "arrival_time": "0001-01-01T00:00:00Z"
    },
    "train_seat": {
      "id": 0,
      "train_id": 0,
      "seat_number": ""
    }
  },
  "message": "Schedule Seat Created Successfully"
}
```

#### Get Schedule Seat By Schedule ID

URL: `/schedules/:id/seats`
Method: `GET`

Response Body:

```json
{
  "data": [
    {
      "id": 1,
      "schedule_id": 3,
      "train_seat_id": 2,
      "status": "AVAILABLE",
      "schedule": {
        "id": 3,
        "train_id": 2,
        "train": {
          "id": 2,
          "name": "Argo Bromo Anggrek",
          "code": "ABA-02"
        },
        "departure_time": "2026-10-15T15:00:00+07:00",
        "arrival_time": "2026-10-15T19:00:00+07:00"
      },
      "train_seat": {
        "id": 2,
        "train_id": 2,
        "seat_number": "A4"
      }
    },
    {
      "id": 2,
      "schedule_id": 3,
      "train_seat_id": 2,
      "status": "AVAILABLE",
      "schedule": {
        "id": 3,
        "train_id": 2,
        "train": {
          "id": 2,
          "name": "Argo Bromo Anggrek",
          "code": "ABA-02"
        },
        "departure_time": "2026-10-15T15:00:00+07:00",
        "arrival_time": "2026-10-15T19:00:00+07:00"
      },
      "train_seat": {
        "id": 2,
        "train_id": 2,
        "seat_number": "A4"
      }
    },
    {
      "id": 3,
      "schedule_id": 3,
      "train_seat_id": 2,
      "status": "AVAILABLE",
      "schedule": {
        "id": 3,
        "train_id": 2,
        "train": {
          "id": 2,
          "name": "Argo Bromo Anggrek",
          "code": "ABA-02"
        },
        "departure_time": "2026-10-15T15:00:00+07:00",
        "arrival_time": "2026-10-15T19:00:00+07:00"
      },
      "train_seat": {
        "id": 2,
        "train_id": 2,
        "seat_number": "A4"
      }
    }
  ],
  "message": "Get Schedule Seats By Schedule ID Successfully"
}
```

#### Book a Seat

URL: `/schedule_seats/:id/book`
Method: `POST`

Response Body:

```json
{
  "message": "Seat Booked Successfully"
}
```

Response Error:
**Seat Not Found**

```json
{
  "message": "Schedule Seat Not Found"
}
```

#### Cancel a Seat Booking

URL: `/schedule_seats/:id/cancel`
Method: `POST`

Response Body:

```json
{
  "message": "Seat Booking Cancelled Successfully"
}
```

Response Error:
**Seat Not Found**

```json
{
  "message": "Schedule Seat Not Found"
}
```

## Order Service

BASE URL: `http://localhost:8082`

### Create an Order

URL: `/orders`
Method: `POST`

Request Body:

```json
{
  "user_id": 1,
  "schedule_seat_id": 1
}
```

Response Body:

```json
{
  "message": "Order created successfully"
}
```

### Get Order by ID

URL: `/orders/:id`
Method: `GET`

Response Body:

```json
{
  "ID": 1,
  "UserID": 1,
  "ScheduleSeatID": 1,
  "Status": "PAID",
  "CreatedAt": "2026-03-05T23:09:38.841+07:00",
  "UpdatedAt": "2026-03-05T23:09:50.98+07:00"
}
```

### Update Status

URL: `/orders/:id`
Method: `PUT`

Request Body:

```json
{
  "status": "PAID"
}
```

Response Body:

```json
{
  "message": "Order status updated successfully"
}
```

### Get All Orders

URL: `/orders`
Method: `GET`

Response Body:

```json
[
  {
    "ID": 1,
    "UserID": 1,
    "ScheduleSeatID": 1,
    "Status": "PAID",
    "CreatedAt": "2026-03-05T23:09:38.841+07:00",
    "UpdatedAt": "2026-03-06T07:02:08.568+07:00"
  },
  {
    "ID": 2,
    "UserID": 1,
    "ScheduleSeatID": 1,
    "Status": "PENDING",
    "CreatedAt": "2026-03-06T06:59:51.704+07:00",
    "UpdatedAt": "2026-03-06T06:59:51.704+07:00"
  },
  {
    "ID": 3,
    "UserID": 1,
    "ScheduleSeatID": 1,
    "Status": "PENDING",
    "CreatedAt": "2026-03-06T07:01:04.389+07:00",
    "UpdatedAt": "2026-03-06T07:01:04.389+07:00"
  }
]
```
