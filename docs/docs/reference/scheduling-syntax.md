# Reference: Scheduling Syntax

BlueBerry supports standard cron expressions as well as a set of predefined constants for common intervals.

## 1. Custom Cron Expressions

You can use standard 5-part cron expressions.

```go
// "0 0 * * *" - Run every day at midnight
sc, err := tsk1.RegisterSchedule(params, "0 0 * * *")

// "@every 1h30m" - Go-specific cron syntax
sc, err := tsk1.RegisterSchedule(params, "@every 1h30m")
```

## 2\. Predefined Run Configurations

For convenience, you can use the following built-in constants.

### Common Cron Intervals

  - **RunEveryMinute**: Executes the task every minute.
  - **RunEvery5Minutes**: Executes the task every 5 minutes.
  - **RunEvery10Minutes**: Executes the task every 10 minutes.
  - **RunEvery15Minutes**: Executes the task every 15 minutes.
  - **RunEvery30Minutes**: Executes the task every 30 minutes.
  - **RunEveryHour**: Executes the task every hour.
  - **RunEvery2Hours**: Executes the task every 2 hours.
  - **RunEvery3Hours**: Executes the task every 3 hours.
  - **RunEvery4Hours**: Executes the task every 4 hours.
  - **RunEvery6Hours**: Executes the task every 6 hours.
  - **RunEvery12Hours**: Executes the task every 12 hours.
  - **RunEveryDay**: Executes the task every 24 hours.
  - **RunEveryWeek**: Executes the task every 7 days (168 hours).

### Specific Times of Day

  - **RunAtMidnight**: Executes the task at midnight (00:00) every day.
  - **RunAtNoon**: Executes the task at noon (12:00) every day.
  - **RunAt6AM**: Executes the task at 6:00 AM every day.
  - **RunAt6PM**: Executes the task at 6:00 PM every day.

### Specific Days of the Week

  - **RunEveryMondayAtNoon**: Executes the task every Monday at noon (12:00).
  - **RunEveryFridayAtNoon**: Executes the task every Friday at noon (12:00).
  - **RunEverySundayAtMidnight**: Executes the task every Sunday at midnight (00:00).

### Example

```go
// Register and schedule Task 1
sc, err := tsk1.RegisterSchedule(
    blueberry.TaskParams{"param1": "value1"}, 
    blueberry.RunEveryMinute,
)
```