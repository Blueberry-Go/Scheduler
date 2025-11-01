# Reference: REST API

The BlueBerry API server provides endpoints to manage and monitor tasks and schedules.

When the server is running, a full **Swagger API documentation** UI is available at `/swagger/index.html`.

## Starting the API Server

To start the API server, use the `RunAPI` method.

```go
// rb is your blueberry.NewBlueBerryInstance(db)
rb.InitTaskScheduler()
rb.RunAPI("8080")
```

## Endpoints

  - **GET /api/tasks**: Get all registered tasks and their schedules.
  - **GET /api/task/:name/executions**: Get all executions for a specific task.
  - **GET /api/task\_run/:id/logs**: Get all logs for a specific task run.
  - **POST /api/execution/:id/cancel**: Cancel a specific task execution by ID.
  - **POST /api/task/:name/execute**: Execute a task by name immediately (on-demand).
