package utils

import (
	"gmountie/pkg/utils/log"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

func Must0(t testing.TB, err error) {
	if err != nil {
		log.Log.Error("Unexpected error", zap.Error(err))
		t.FailNow()
	}
}

func Must1[T any](suite suite.TestingSuite, obj T, err error) T {
	if err != nil {
		log.Log.Error("Unexpected error", zap.Error(err))
		suite.T().FailNow()
	}
	return obj
}
