# Task Management API Documentation (v2 - JWT & MongoDB)

This RESTful API allows users to manage tasks using **Go**, **Gin**, and **MongoDB**. Access is restricted using **JWT Authentication** and **Role-Based Access Control (RBAC)**.

## 🔑 Authentication Flow
1. **Register**: Create an account. (The first user registered becomes an `admin` automatically).
2. **Login**: Provide credentials to receive a **JWT Token**.
3. **Authorize**: Include the token in the header of all protected requests:  
   `Authorization: Bearer <your_token_here>`

---

## 📋 Data Models

### User Struct

| Field      | Type   | Description                          |
| :--------- | :----- | :----------------------------------- |
| `username` | string | Unique login name.                   |
| `password` | string | Plaintext (stored as Bcrypt hash).   |
| `role`     | string | `admin` or `user`.                   |

### Task Struct

| Field      | Type   | Description                          |
| :--------- | :----- | :----------------------------------- |
| `id`       | string | Unique identifier (MongoDB `_id`).   |
| `title`    | string | Task name.                           |
| `status`   | string | `Pending`, `In Progress`, `Completed`.|

---

## 🛡️ User & Auth Endpoints

### 1. Register User
- **URL:** `/register`
- **Method:** `POST`
- **Payload:** `{"username": "johndoe", "password": "securepassword"}`
- **Note:** The first user in the DB is granted the `admin` role.

### 2. Login
- **URL:** `/login`
- **Method:** `POST`
- **Payload:** `{"username": "johndoe", "password": "securepassword"}`
- **Success Response:** `200 OK` with `{"token": "JWT_STRING_HERE"}`

### 3. Promote User (Admin Only)
- **URL:** `/promote/:username`
- **Method:** `POST`
- **Description:** Upgrades a regular `user` to `admin`.

---

## 🚀 Task Endpoints


| Endpoint | Method | Access Level | Description |
| :--- | :--- | :--- | :--- |
| `/tasks` | **GET** | Authenticated | View all tasks. |
| `/tasks/:id` | **GET** | Authenticated | View a specific task. |
| `/tasks` | **POST** | **Admin Only** | Create a new task. |
| `/tasks/:id` | **PUT** | **Admin Only** | Update an existing task. |
| `/tasks/:id` | **DELETE** | **Admin Only** | Remove a task. |

### Example Protected Request (POST /tasks)
**Header:** `Authorization: Bearer <token>`  
**Body:**
```json
{
  "id": "101",
  "title": "Deploy to Production",
  "description": "Finalize MongoDB Atlas connection",
  "due_date": "2024-11-25T10:00:00Z",
  "status": "Pending"
}

