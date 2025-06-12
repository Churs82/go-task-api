# Go Task API

This project is a simple HTTP API for managing background tasks. It allows users to create, delete, and retrieve results of tasks, as well as check the status of each task.

## Features

- Create a new background task
- Delete an existing task
- Retrieve results of completed tasks
- Get list of tasks
- Get the status of a task, including:
  - Status (Pending, Running, Completed, Failed)
  - Date of creation
  - Duration of task execution

## Project Structure

```
go-task-api
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── handlers.go  # Handlers for HTTP requests
│   └── task
│       ├── iotask.go    # Example of an IO-bound task
│       ├── task.go      # Model for a task
│       ├── task_manager.go # Manager for tasks
│       └── task_registry.go # Registry for tasks
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

- **Endpoint:** `POST /task`
- **Request Body:** `{ "task_type": "io" }`
- **Response:** Returns the created task details.

### Delete Task

- **Endpoint:** `DELETE /task/{task_id}`
- **Response:** Confirms deletion of the task.

### Get Task

- **Endpoint:** `GET /task/{task_id}`
- **Response:** Returns the details of the specified task.

### Get Tasks List

- **Endpoint:** `GET /tasks`
- **Response:** Returns the list of all tasks.

### Get Task Status

- **Endpoint:** `GET /task/{task_id}/status`
- **Response:** Returns the status of the specified task.

### Get Task Result

- **Endpoint:** `GET /task/{task_id}/result`
- **Response:** Returns the result of the specified task.



## To implement a new task

If you want to implement a new type of background task:
- Create a new file in **internal/task** directory
- Extend package **task** and implement Run method of **Task_i** interface
- Add a string ```registry.RegisterTask([Your Task Type Name], &[Your Task Structure])``` to the init() function in **internal/task/task_registry.go** file.
