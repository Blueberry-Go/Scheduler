package blueberry

import (
	"context"
	"time"
)

type TaskContext struct {
	ctx    context.Context
	params TaskParams
	logger *Logger
}

func (t *TaskContext) GetContext() context.Context {
	return t.ctx
}

func (t *TaskContext) GetParams() TaskParams {
	return t.params
}

func (t *TaskContext) GetLogger() *Logger {
	return t.logger
}

type TaskFunc func(*TaskContext) error

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

// DB is the interface that wraps basic database operations
type DB interface {
	SaveTaskRun(ctx context.Context, taskRun *TaskRun) error
	SaveTaskRunLog(ctx context.Context, taskRunLog *TaskRunLog) error
	GetTaskRuns(ctx context.Context) ([]TaskRun, error)
	GetTaskRunByID(ctx context.Context, id int) (*TaskRun, error)
	GetTaskRunLogs(ctx context.Context, taskRunID int) ([]TaskRunLog, error)
	GetPaginatedTaskRunLogs(ctx context.Context, taskRunID int, level string, page, size int) ([]TaskRunLog, int, error)
	GetPaginatedTaskRunsForTaskName(ctx context.Context, name string, page, limit int) ([]TaskRun, error)
	GetTaskRunsCountForTaskName(ctx context.Context, name string) (int, error)
	Close() error
}
