# Getting Started

This guide will walk you through building your first BlueBerry scheduler from scratch. We'll write a complete, runnable Go program that defines, schedules, and runs a task.

## 1\. Installation

First, you'll need the core BlueBerry module and the Logger store module; (SQLite is choosen down below but you can choose anything as per your preference) for storing the logs. Make sure you have Go installed, then run:

```bash
go get github.com/blueberry-go/scheduler/core
go get github.com/blueberry-go/scheduler/store/sqlite
```

## 2\. Project Setup

Let's create a new file named `main.go`. We'll start with the package declaration and import all the libraries we'll need.

```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// Import the BlueBerry core
	blueberry "github.com/blueberry-go/scheduler/core"
	"github.com/blueberry-go/scheduler/core/task" 

    // Import the SQLite store
	store "github.com/blueberry-go/scheduler/store/sqlite" 
)
```

Above, we've imported the `core` BlueBerry library, the `task` package which contains the `TaskContext`, and the `sqlite` store. The other imports are standard Go libraries for logging, time, and system signals.

## 3\. Define a Task and Schema

A **Task** is your application's business logic—it's just a Go function. To make it dynamic, you also define a **Schema** that tells BlueBerry what parameters your task expects.

```go
// 1. Define a Task Schema (required)
var (
	task1Schema = blueberry.NewTaskSchema(blueberry.TaskParamDefinition{
		"param1": blueberry.TypeString,
		"param2": blueberry.TypeInt,
	})
)
```

This schema tells BlueBerry that our task accepts a `string` named `param1` and an `int` named `param2`. The Web GUI and API will use this schema to validate user input.

Now, let's write the **Task Function** itself. It must accept a `*task.TaskContext` as its only parameter.

```go
// 2. Define a Task Function
func task1(tctx *task.TaskContext) error {
	// Get context, logger, and params from the TaskContext
	ctx := tctx.GetContext()
	logger := tctx.GetLogger()
	params := tctx.GetParams()

	_ = logger.Info(fmt.Sprintf("The params are: %v", params))

	if err := logger.Info("Starting Task 1"); err != nil {
		return err
	}

	// This select block allows the task to do work
	// while also listening for a cancellation signal (ctx.Done())
	select {
	case <-time.After(30 * time.Second):
		if err := logger.Success("Task 1 completed successfully"); err != nil {
			return err
		}
		return nil
	case <-ctx.Done():
		if err := logger.Error("Task 1 cancelled"); err != nil {
			return err
		}
		return ctx.Err()
	}
}
```

The `TaskContext` is your gateway to BlueBerry's features:

  * `tctx.GetLogger()`: Returns a logger module; which can be used to add logs and store it in above given store.
  * `tctx.GetParams()`: Gets the parameters for this specific run.
  * `tctx.GetContext()`: A standard Go `context.Context` that handles graceful shutdowns.

## 4\. Configure and Run the Scheduler

With our task defined, we can now write the `main` function to bring it all to life.

```go
func main() {
	// 3. Initialize the Database
	// This creates a 'task_scheduler.db' file in the same directory.
	db, err := store.NewSQLiteDB("task_scheduler.db")
	if err != nil {
		log.Fatalf("Failed to initialize SQLite: %v", err)
	}
	defer db.Close()

	// 4. Initialize the BlueBerry Instance
	// Pass the database store to the new instance.
	rb := blueberry.NewBlueBerryInstance(db)
```

We first set up our database (a simple SQLite file) and then create the main BlueBerry instance, passing the database connection to it.

Next, we **Register** the task and **Schedule** it to run.
This links the identifier "task_1" to your task1 function and schema, which will be used later throught the code. Do note that this needs to be unique per task.

```go
	// 5. Register the Task
	tsk1, err := rb.RegisterTask("task_1", task1, task1Schema)
	if err != nil {
		fmt.Printf("Failed to register task: %v\n", err)
		return
	}

	// 6. Register a Schedule for the Task
	// This tells "task_1" to run "@every 1m" (every 1 minute)
	// with the provided parameters.
	_, err = tsk1.RegisterSchedule(blueberry.TaskParams{
		"param1": "value1",
		"param2": 1,
	}, "@every 1m")

	if err != nil {
		log.Fatalf("Failed to register schedule: %v", err)
	}
```

Registering the task makes it available to the GUI and API. Registering the schedule tells the scheduler to run it automatically.

Finally, we'll set up graceful shutdown handling and start the server.

```go
    // 7. Handle Graceful Shutdown
    // This is critical for production. It listens for OS signals
    // and calls rb.Shutdown(), which triggers ctx.Done() in all
    // running tasks.
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <-sigChan
        log.Printf("Received signal: %v. Shutting down...", sig)
        rb.Shutdown()
        os.Exit(0)
    }()

    // 8. Start the Scheduler and API
    rb.InitTaskScheduler() // Starts the scheduler engine
    log.Println("Starting BlueBerry API server at :8080")
    rb.RunAPI("8080") // Starts the Web GUI and REST API
}
```

## 5\. Run Your Scheduler

You're all set\! Your `main.go` file is complete. Open your terminal in the same directory and run:

```bash
# This downloads the dependencies and cleans up your go.mod file
go mod tidy

# This runs your scheduler
go run main.go
```

You should see the log output: `Starting BlueBerry API server at :8080`.

Your scheduler is now running\!

  * **Web GUI**: `http://localhost:8080`
  * **Swagger API Docs**: `http://localhost:8080/swagger/index.html`

Open the Web GUI to see your `task_1` and watch it run. You can also use the GUI or API to trigger the task manually with different parameters.