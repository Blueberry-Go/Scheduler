# How to Define Tasks

A **Task** is the core unit of work in BlueBerry. It consists of two parts: the **Task Function** (the Go code to be executed) and the **Task Schema** (the parameters the function accepts).

## 1. The Task Function

A task function is any Go function that accepts a `*task.TaskContext` pointer. This context provides everything the task needs to run:

- `tctx.GetContext()`: A standard `context.Context` for handling cancellations and timeouts.
- `tctx.GetLogger()`: An instance of the BlueBerry logger to log status back to the GUI and database.
- `tctx.GetParams()`: A map of parameters (`blueberry.TaskParams`) passed to this specific task execution.

```go
func myTaskFunction(tctx *task.TaskContext) error {
    ctx := tctx.GetContext()
    logger := tctx.GetLogger()
    params := tctx.GetParams()

    // 1. Log the start
    _ = logger.Info("My task is starting...")

    // 2. Get parameters
    userName, _ := params.GetString("user")
    _ = logger.Info(fmt.Sprintf("Processing for user: %s", userName))

    // 3. Perform work, respecting cancellation
    select {
    case <-time.After(10 * time.Minute):
        if err := logger.Success("Task completed successfully"); err != nil {
            return err
        }
        return nil
    case <-ctx.Done(): // Triggered by graceful shutdown or manual cancellation
        if err := logger.Error("Task cancelled"); err != nil {
            return err
        }
        return ctx.Err()
    }
}
```

## 2\. The Task Schema

A **Task Schema** defines the parameters your task function expects, including their types. This is used by the GUI and API to validate input.

  - `blueberry.TypeString`
  - `blueberry.TypeInt`
  - `blueberry.TypeBool`

<!-- end list -->

```go
var (
    myTaskSchema = blueberry.NewTaskSchema(blueberry.TaskParamDefinition{
        "user":   blueberry.TypeString,
        "retries": blueberry.TypeInt,
        "force":   blueberry.TypeBool,
    })
)
```

## 3\. Registering the Task

Finally, register the function and its schema with the BlueBerry instance.

```go
// rb is your blueberry.NewBlueBerryInstance(db)
tsk, err := rb.RegisterTask(
    "my-task-name", // A unique string identifier
    myTaskFunction, 
    myTaskSchema,
)

if err != nil {
    log.Fatalf("Failed to register task: %v\n", err)
}
```