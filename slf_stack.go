package slog

import (
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

// for better preformance, use atomic-map
var (
	stackLock  = new(sync.Mutex)
	stackCache = new(atomic.Value)

	stackMap = sync.Map{} // for test/compare
)

// Stack represent pc's stack details.
type Stack struct {
	pc       uintptr
	Package  string `json:"package"`
	Filename string `json:"filename"`
	Function string `json:"function"`
	Line     int    `json:"line"`
}

// cacheStack save the specified stackInfo into global atomic map.
func cacheStack(s *Stack) {
	stackLock.Lock()
	defer stackLock.Unlock()
	var oldMap map[uintptr]*Stack
	if x := stackCache.Load(); x != nil {
		oldMap = x.(map[uintptr]*Stack)
	}
	// don't modify oldMap to avoid concurrency problem
	newMap := make(map[uintptr]*Stack)
	if oldMap != nil {
		for k, v := range oldMap {
			newMap[k] = v
		}
	}
	newMap[s.pc] = s
	stackCache.Store(newMap)
}

// loadStack retrieve cached Stack, could be nil
func loadStack(pc uintptr) *Stack {
	x := stackCache.Load()
	if x == nil {
		return nil
	}
	infoMap := x.(map[uintptr]*Stack)

	return infoMap[pc]
}

// parseStack retrieve pc's stack info
func parseStack(pc uintptr) (s *Stack) {
	var frame runtime.Frame
	if frames := runtime.CallersFrames([]uintptr{pc}); frames != nil {
		frame, _ = frames.Next()
	}
	s = &Stack{pc: pc, Line: frame.Line}
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
	s.Package = pkgName
	s.Function = funcName
	// parse Filename
	var fileName = frame.File
	if off := strings.LastIndexByte(fileName, '/'); off > 0 && off < len(fileName)-1 {
		fileName = fileName[off+1:]
	}
	s.Filename = fileName
	return
}

// ParseStack retrieve pc's stack details, should cache result for performance optimization.
func ParseStack(pc uintptr) (s *Stack) {
	if s = loadStack(pc); s != nil {
		return
	}
	s = parseStack(pc)
	cacheStack(s)
	return
}

// ParseStack2 this is slower, cannot use it.
func ParseStack2(pc uintptr) (s *Stack) {
	if v, ok := stackMap.Load(pc); ok {
		s = v.(*Stack)
	} else {
		s = parseStack(pc)
		stackMap.Store(pc, s)
	}
	return
}
