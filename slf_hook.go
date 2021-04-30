package slog

import (
	"runtime/debug"
	"sync"
	"sync/atomic"
)

const (
	hookCacheSize  = 1 << 10
	hookInitStatus = -1
)

// hook represent a registered hook func
type hook struct {
	hookFun func(*Log)
	logChan chan *Log
}

func newHook(f func(*Log)) *hook {
	h := &hook{
		hookFun: f,
		logChan: make(chan *Log, hookCacheSize),
	}
	go func() {
		for l := range h.logChan {
			h.trigger(l)
		}
	}()
	return h
}

func (t *hook) trigger(l *Log) {
	defer func() {
		if recover() != nil {
			Errorf("hook trigger error:\n", string(debug.Stack()))
		}
	}()
	t.hookFun(l)
}

// hooks
type hooks struct {
	sync.Mutex
	logChan    chan *Log
	hookStatus int32
	hookList   atomic.Value // []*hook
}

func newHooks() *hooks {
	t := &hooks{
		hookStatus: hookInitStatus,
		logChan:    make(chan *Log, hookCacheSize),
	}
	t.hookList.Store([]*hook{})
	return t
}

func (t *hooks) addHook(f func(*Log)) {
	if atomic.AddInt32(&t.hookStatus, 1) == 0 {
		go func() {
			defer func() {
				if recover() != nil {
					Errorf("BUG: async-boradcast error:\n", string(debug.Stack()))
				}
			}()
			for logInst := range t.logChan {
				for _, h := range t.hookList.Load().([]*hook) {
					h.logChan <- logInst
				}
			}
		}()
	}
	t.Lock()
	hs := t.hookList.Load().([]*hook)
	hs = append(hs, newHook(f))
	t.hookList.Store(hs)
	atomic.StoreInt32(&t.hookStatus, int32(len(hs)))
	t.Unlock()
}

func (t *hooks) broadcast(l *Log) {
	if atomic.LoadInt32(&t.hookStatus) <= 0 {
		return
	}
	select {
	case t.logChan <- l:
	default:
		Warnf("broadcast failed, log channel is full")
	}
}
