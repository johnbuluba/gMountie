package log

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	// Create a logger

	// Set logger to warn level for tests
	zapConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	Log = zap.Must(zapConfig.Build())
	Log = Log.Named("gMountie")
	zap.ReplaceGlobals(Log)

	defer Log.Sync() // Flushes buffer, if any
}
