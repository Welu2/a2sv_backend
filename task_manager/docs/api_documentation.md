# Task Management API Documentation

This RESTful API allows users to manage a collection of tasks using the **Go** programming language and the **Gin framework**. Data is currently stored in an in-memory database.

## Base URL
`http://localhost:8080`

---

## 📋 Data Model (Task Struct)
Each task object contains the following fields:
- `id` (string): Unique identifier for the task.
- `title` (string): The name of the task.
- `description` (string): Detailed information about the task.
- `due_date` (ISO8601 String): The date the task is due.
- `status` (string): Current status (e.g., "Pending", "In Progress", "Completed").

---

## 🚀 Endpoints

### 1. Get All Tasks
- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieves a list of all stored tasks.
- **Success Response:** `200 OK` (JSON array of task objects).

### 2. Get Task by ID
- **URL:** `/tasks/:id`
- **Method:** `GET`
- **Description:** Retrieves details for a specific task using its unique ID.
- **Success Response:** `200 OK` (Single task object).
- **Error Response:** `404 Not Found` if the ID does not exist.

### 3. Create a New Task
- **URL:** `/tasks`
- **Method:** `POST`
- **Description:** Adds a new task to the in-memory database.
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
- **Description:** Updates the details of an existing task. The request body should contain the fields to be updated.
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