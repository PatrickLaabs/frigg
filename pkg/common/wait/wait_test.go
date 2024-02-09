package wait_test

import (
	"github.com/PatrickLaabs/frigg/pkg/common/wait"
	"testing"
	"time"
)

func TestWaitDuration(t *testing.T) {
	tests := []struct {
		name      string
		duration  time.Duration
		expected  time.Duration
		tolerance time.Duration // How much difference in elapsed time is acceptable
	}{
		{
			name:      "Wait for 1 second",
			duration:  time.Second,
			expected:  time.Second,
			tolerance: 100 * time.Millisecond, // Allow for some overhead
		},
		{
			name:      "Wait for 500 milliseconds",
			duration:  500 * time.Millisecond,
			expected:  500 * time.Millisecond,
			tolerance: 50 * time.Millisecond, // Allow for smaller overhead
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			wait.Wait(tt.duration)
			elapsed := time.Since(startTime)

			if elapsed < tt.expected-tt.tolerance || elapsed > tt.expected+tt.tolerance {
				t.Errorf("Waited for %v, but elapsed time was %v", tt.duration, elapsed)
			}
		})
	}
}
