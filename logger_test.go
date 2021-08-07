package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestInitGlobalConsole(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SimpleGlobalConsole()
			zap.L().Debug("test")
		})
	}
}
