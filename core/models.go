package blueberry

import (
	"context"

	stype "github.com/blueberry-go/scheduler/types"
)

type TaskFunc func(*TaskContext) error

// DB is the interface that wraps basic database operations
type DB interface {
	SaveTaskRun(ctx context.Context, taskRun *stype.TaskRun) error
	SaveTaskRunLog(ctx context.Context, taskRunLog *stype.TaskRunLog) error
	GetTaskRuns(ctx context.Context) ([]stype.TaskRun, error)
	GetTaskRunByID(ctx context.Context, id int) (*stype.TaskRun, error)
	GetTaskRunLogs(ctx context.Context, taskRunID int) ([]stype.TaskRunLog, error)
	GetPaginatedTaskRunLogs(ctx context.Context, taskRunID int, level string, page, size int) ([]stype.TaskRunLog, int, error)
	GetPaginatedTaskRunsForTaskName(ctx context.Context, name string, page, limit int) ([]stype.TaskRun, error)
	GetTaskRunsCountForTaskName(ctx context.Context, name string) (int, error)
	Close() error
}
