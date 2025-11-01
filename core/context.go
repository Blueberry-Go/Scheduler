package blueberry

import "context"

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
