# How to Manage Schedules

A **Schedule** defines *when* a registered task should run and *what parameters* to run with.

## 1. Registering a Schedule

Once you have a registered task, you can add multiple schedules to it.

```go
// tsk1 is the task object from rb.RegisterTask(...)
sc, err := tsk1.RegisterSchedule(
    blueberry.TaskParams{
        "param1": "value1",
        "param2": 1,
        "param3": true,
    }, 
    "@every 1m", // See "Scheduling Syntax" for all options
)

if err != nil {
    log.Fatalf("Failed to register schedule: %v", err)
}

// You can add more schedules to the same task
sc2, err := tsk1.RegisterSchedule(
    blueberry.TaskParams{
        "param1": "another-value",
        "param2": 99,
        "param3": false,
    }, 
    blueberry.RunAtMidnight, // Use a predefined constant
)
```

## 2\. Deleting a Schedule

When you register a schedule, the returned `sc` object contains a unique `EntryID`. You must use this ID to delete the schedule from the task.

```go
// sc2 is the schedule object from tsk1.RegisterSchedule(...)
err := tsk1.DeleteSchedule(sc2.EntryID)
if err != nil {
    log.Printf("Failed to delete schedule: %v", err)
}
```

## 3\. Updating a Schedule Dynamically

To update a schedule (e.g., change its cron timing or parameters), you must:

1.  **Delete** the current schedule using its `EntryID`.
2.  **Register** a new schedule with the updated information.

\!\!\! WARNING:

If you are updating schedules dynamically while the scheduler is running, make sure to handle this in a thread-safe manner (e.g., using a mutex) to prevent race conditions.