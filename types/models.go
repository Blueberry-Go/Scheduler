package types

import "time"

type TaskRun struct {
	ID        int                    `json:"id"`
	TaskName  string                 `json:"task_name"`
	StartTime time.Time              `json:"start_time"`
	EndTime   time.Time              `json:"end_time"`
	Params    map[string]interface{} `json:"params"`
	Status    string                 `json:"status"` // "started", "completed", "failed", "cancelled"
}

// TaskRunLog represents a log entry for a task run
type TaskRunLog struct {
	ID        int
	TaskRunID int
	Timestamp time.Time
	Level     string
	Message   string
}
