# How to Configure Databases

BlueBerry supports multiple database backends for storing task runs, schedules, and execution logs. The database interface is located in the `github.com/blueberry-go/scheduler/store` package.

## 1. SQLite

SQLite is the simplest way to get started, storing the database in a single file.

```go
import "github.com/blueberry-go/scheduler/store/sqlite"

db, err := store.NewSQLiteDB("task_scheduler.db")
if err != nil {
    log.Fatalf("Failed to initialize SQLite: %v", err)
}
defer db.Close()

rb := blueberry.NewBlueBerryInstance(db)
```

## 2\. PostgreSQL

For production environments, PostgreSQL is a robust choice.

```go
import "github.com/blueberry-go/scheduler/store/postgres"

const connStr = "postgres://username:password@localhost:5432/mydatabase?sslmode=disable"
db, err := store.NewPostgresDB(connStr)
if err != nil {
    log.Fatalf("Failed to initialize PostgreSQL: %v", err)
}
defer db.Close()

rb := blueberry.NewBlueBerryInstance(db)
```

## 3\. MongoDB

BlueBerry also supports MongoDB for NoSQL-based storage.

```go
import "github.com/blueberry-go/scheduler/store/mongodb"

const mongoURI = "mongodb://localhost:27017"
const dbName = "task_scheduler"
db, err := store.NewMongoDB(mongoURI, dbName)
if err != nil {
    log.Fatalf("Failed to initialize MongoDB: %v", err)
}
defer db.Close()

rb := blueberry.NewBlueBerryInstance(db)
```