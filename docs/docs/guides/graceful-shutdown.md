# How to Handle Graceful Shutdown

To ensure all running tasks are completed or cancelled properly before your application exits, you must implement graceful shutdown handling.

This involves listening for system signals (like `SIGINT` or `SIGTERM`) and calling the `rb.Shutdown()` method.

```go
import (
    "log"
    "os"
    "os/signal"
    "syscall"
)

// rb is your blueberry.NewBlueBerryInstance(db)

// Handle system signals for graceful shutdown
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

go func() {
    sig := <-sigChan
    log.Printf("Received signal: %v. Shutting down...", sig)
    
    // This triggers the cancellation context for all running tasks
    rb.Shutdown() 
    
    os.Exit(0)
}()

// ... proceed to start your scheduler
rb.InitTaskScheduler()
rb.RunAPI("8080")
```

When `rb.Shutdown()` is called, the `context.Context` passed to all active task functions (via `tctx.GetContext()`) will be cancelled. Your task functions should use this signal to stop work and return, as shown in the "Defining Tasks" guide.
