package log

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	// Create a logger
	Log, _ = zap.NewDevelopment()
	defer Log.Sync() // Flushes buffer, if any
}
