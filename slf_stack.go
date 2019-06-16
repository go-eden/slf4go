package xlog

import (
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

// for better preformance, use atomic-map
var stackLock = new(sync.Mutex)
var stackCache = new(atomic.Value)

// stackInfo represent pc's stack details.
type stackInfo struct {
	pc       uintptr
	pkgName  string
	fileName string
	funcName string
	line     int
}

// cacheStack save the specified stachInfo into global atomic map.
func cacheStack(s *stackInfo) {
	stackLock.Lock()
	defer stackLock.Unlock()
	var oldMap map[uintptr]*stackInfo
	if x := stackCache.Load(); x != nil {
		oldMap = x.(map[uintptr]*stackInfo)
	}
	newMap := make(map[uintptr]*stackInfo)
	if oldMap != nil {
		for k, v := range oldMap {
			newMap[k] = v
		}
	}
	newMap[s.pc] = s
	stackCache.Store(newMap)
}

// loadStack retrieve cached stackInfo, could be nil
func loadStack(pc uintptr) *stackInfo {
	x := stackCache.Load()
	if x == nil {
		return nil
	}
	infoMap := x.(map[uintptr]*stackInfo)

	return infoMap[pc]
}

// parseStack retrieve pc's stack info
func parseStack(pc uintptr) (s *stackInfo) {
	var frame runtime.Frame
	if frames := runtime.CallersFrames([]uintptr{pc}); frames != nil {
		frame, _ = frames.Next()
	}
	s = &stackInfo{pc: pc, line: frame.Line}
	// parse package and function
	var pkgName, funcName string
	if frame.Func != nil {
		var off int
		name := frame.Func.Name()
		for i := len(name) - 1; i >= 0; i-- {
			if name[i] == '/' {
				break
			}
			if name[i] == '.' {
				off = i
			}
		}
		if off > 0 {
			pkgName = name[:off]
			if off < len(name)-1 {
				funcName = name[off+1:]
			}
		} else {
			pkgName = name
		}
	}
	s.pkgName = pkgName
	s.funcName = funcName
	// parse fileName
	var fileName = frame.File
	if off := strings.LastIndexByte(fileName, '/'); off > 0 && off < len(fileName)-1 {
		fileName = fileName[off+1:]
	}
	s.fileName = fileName
	return
}

// ParseStack retrieve pc's stack details, should cache result for performance optimization.
func ParseStack(pc uintptr) (s *stackInfo) {
	if s = loadStack(pc); s != nil {
		return
	}
	s = parseStack(pc)
	cacheStack(s)
	return
}
