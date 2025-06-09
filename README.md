# Go Task API

This project is a simple HTTP API for managing background tasks. It allows users to create, delete, and retrieve results of tasks, as well as check the status of each task.

## Features

- Create a new background task
- Delete an existing task
- Retrieve results of completed tasks
- Get the status of a task, including:
  - Task ID
  - Status (Pending, Running, Completed, Failed)
  - Date of creation
  - Duration of task execution

## Project Structure

```
go-task-api
├── cmd
│   └── main.go          # Entry point of the application
└── README.md             # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd go-task-api
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## API Usage

### Create Task

- **Endpoint:** `POST /tasks`
- **Request Body:** `{ "task_name": "example_task" }`
- **Response:** Returns the created task details.

### Delete Task

- **Endpoint:** `DELETE /tasks/{task_id}`
- **Response:** Confirms deletion of the task.

### Get Task

- **Endpoint:** `GET /tasks/{task_id}`
- **Response:** Returns the details of the specified task.

### Get Task Status

- **Endpoint:** `GET /tasks/{task_id}/status`
- **Response:** Returns the current status of the specified task.
