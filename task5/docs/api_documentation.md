# Task Management API Documentation

This RESTful API allows users to manage a collection of tasks using the **Go** programming language and the **Gin framework**. Data is persisted in a **MongoDB** database.

## Base URL
`http://localhost:8080`

---

## 📋 Data Model (Task Struct)
Each task object is mapped to a MongoDB document. Note that the `id` field in JSON maps directly to the `_id` field in the MongoDB collection.


| Field       | Type   | Description                                      |
| :---------- | :----- | :----------------------------------------------- |
| `id`        | string | Unique identifier (Maps to `_id` in MongoDB).    |
| `title`     | string | The name of the task.                            |
| `description`| string | Detailed information about the task.             |
| `due_date`  | string | ISO8601 formatted date (e.g., 2024-11-20T10:00Z).|
| `status`    | string | "Pending", "In Progress", or "Completed".        |

---

## 🚀 Endpoints

### 1. Get All Tasks
- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieves all task documents from the MongoDB collection.
- **Success Response:** `200 OK` (JSON array of task objects).

### 2. Get Task by ID
- **URL:** `/tasks/:id`
- **Method:** `GET`
- **Description:** Retrieves a specific task using its unique string ID.
- **Success Response:** `200 OK` (Single task object).
- **Error Response:** `404 Not Found` if the ID does not exist in the database.

### 3. Create a New Task
- **URL:** `/tasks`
- **Method:** `POST`
- **Description:** Inserts a new task document into MongoDB.
- **Request Body:**
  ```json
  {
    "id": "1",
    "title": "Study Go",
    "description": "Complete the Gin framework tutorial",
    "due_date": "2024-11-20T10:00:00Z",
    "status": "In Progress"
  }
### 4. Update Task by ID
- **URL:** `/tasks/:id`
- **Method:** `PUT`
- **Description:** Updates the details of an existing task in MongoDB. The `id` in the URL must match the task's `_id`.
- **Request Body:**
  ```json
  {
    "title": "Updated Task Title",
    "description": "Updated description text",
    "due_date": "2024-12-01T15:00:00Z",
    "status": "Completed"
  }
### 5. Delete Task by ID
- **URL:** `/tasks/:id`
- **Method:** `DELETE`
- **Description:** Removes a specific task from the in-memory database using its unique ID.
- **Success Response:** `200 OK`
  ```json
  {
    "message": "Task deleted successfully"
  }