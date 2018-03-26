package testutil

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	defaultDelayThreadStart = time.Millisecond * 5
)

func SleepForThreadStart(t *testing.T) {
	time.Sleep(delayThreadStart(t))
}

func delayThreadStart(t *testing.T) time.Duration {
	if val := os.Getenv("TEST_DELAY_THREAD_START"); val != "" {
		d, err := time.ParseDuration(val)
		require.NoError(t, err)
		return d
	}
	return defaultDelayThreadStart
}