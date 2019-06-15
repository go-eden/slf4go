package slf4go

import "strconv"

type Level int

const (
	LEVEL_TRACE Level = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
)

// Retrieve Level's name
func (l Level) String() string {
	switch l {
	case LEVEL_TRACE:
		return "TRACE"
	case LEVEL_DEBUG:
		return "DEBUG"
	case LEVEL_INFO:
		return "INFO"
	case LEVEL_WARN:
		return "WARN"
	case LEVEL_ERROR:
		return "ERROR"
	case LEVEL_FATAL:
		return "FATAL"
	default:
		return strconv.Itoa(int(l))
	}
}
