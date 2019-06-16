package xlog

import "sync"

// hook represent a registered hook func
type hook struct {
	f  func(*Log)
	ch chan *Log
}

func newHook(f func(*Log)) *hook {
	h := &hook{f: f, ch: make(chan *Log)}
	go func() {
		for l := range h.ch {
			h.trigger(l)
		}
	}()
	return h
}

func (h *hook) trigger(l *Log) {
	defer func() {
		if e := recover(); e != nil {
			Error("trigger hook failed.", e)
		}
	}()
	h.f(l)
}

func (h *hook) inform(l *Log) {
	h.ch <- l
}

// hooks
type hooks struct {
	sync.Mutex
	hs []*hook
	ch chan *Log
	no bool
}

func newHooks() *hooks {
	h := &hooks{
		ch: make(chan *Log),
		no: true,
	}
	go func() {
		for l := range h.ch {
			h.handle(l)
		}
	}()
	return h
}

func (h *hooks) addHook(f func(*Log)) {
	h.Lock()
	defer h.Unlock()
	h.hs = append(h.hs, newHook(f))
	h.no = len(h.hs) == 0
}

func (h *hooks) broadcast(l *Log) {
	if h.no {
		return
	}
	h.ch <- l
}

func (h *hooks) handle(l *Log) {
	hs := h.hs
	if len(hs) == 0 {
		return
	}
	for _, hook := range hs {
		hook.inform(l)
	}
}
