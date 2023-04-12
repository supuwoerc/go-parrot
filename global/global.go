package global

import "go.uber.org/zap"

var (
	Logger *zap.SugaredLogger //全局日志组件实例(单例)
)
