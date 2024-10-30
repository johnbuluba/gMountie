package frontend

import "go.uber.org/zap"

// Logger is a struct that holds the logger
type Logger struct {
	logger *zap.Logger
}

// NewLogger creates a new logger
func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Print(message string) {
	l.logger.Sugar().Info(message)
}

func (l *Logger) Trace(message string) {
	l.logger.Sugar().Debug(message)
}

func (l *Logger) Debug(message string) {
	l.logger.Sugar().Debug(message)
}

func (l *Logger) Info(message string) {
	l.logger.Sugar().Info(message)
}

func (l *Logger) Warning(message string) {
	l.logger.Sugar().Warn(message)
}

func (l *Logger) Error(message string) {
	l.logger.Sugar().Error(message)
}

func (l *Logger) Fatal(message string) {
	l.logger.Sugar().Fatal(message)
}
