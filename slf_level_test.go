package slog

import "testing"

func TestLevel(t *testing.T) {
	t.Log(FatalLevel.String())
	t.Log(TraceLevel)
}
