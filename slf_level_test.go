package slog

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevel(t *testing.T) {
	t.Log(FatalLevel.String())
	t.Log(TraceLevel)
}

func TestLevelSetting(t *testing.T) {
	var setting LevelSetting

	setting.setRootLevel(WarnLevel)
	setting.setLoggerLevel(map[string]Level{
		TraceLevel.String(): TraceLevel,
		DebugLevel.String(): DebugLevel,
		InfoLevel.String():  InfoLevel,
	})
	setting.setLoggerLevel(map[string]Level{
		WarnLevel.String():  WarnLevel,
		ErrorLevel.String(): ErrorLevel,
		PanicLevel.String(): PanicLevel,
		FatalLevel.String(): FatalLevel,
	})

	assert.True(t, setting.getLoggerLevel(rootLoggerName) == WarnLevel)
	assert.True(t, setting.getLoggerLevel(WarnLevel.String()) == WarnLevel)
	assert.True(t, setting.getLoggerLevel(ErrorLevel.String()) == ErrorLevel)
	assert.True(t, setting.getLoggerLevel("xxxx") == WarnLevel)
}

// BenchmarkLevelSetting-12    	59575768	        17.50 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLevelSetting(b *testing.B) {
	var setting LevelSetting
	for i := 0; i < 100; i++ {
		setting.setLoggerLevel(map[string]Level{
			fmt.Sprintf("logger_%v", i): WarnLevel,
		})
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		setting.getLoggerLevel("logger_0")
	}
}
