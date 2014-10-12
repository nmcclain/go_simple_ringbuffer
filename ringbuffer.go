package ringbuffer

import (
	"container/list"
	"sync"
)

type RingBuffer struct {
	size int
	ring *list.List
	lock sync.Mutex
}

// New creates a RingBuffer of the requested size
func New(s int) *RingBuffer {
	r := RingBuffer{size: s, ring: list.New()}
	return &r
}

// Add puts something on the RingBuffer, dropping the last entry if necessary
func (r *RingBuffer) Add(v interface{}) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.ring.Len() >= r.size {
		_ = r.ring.Remove(r.ring.Back())
	}
	_ = r.ring.PushFront(v)
}

// Dump returns the RingBuffer contents from oldest to newest
func (r *RingBuffer) Dump() []interface{} {
	r.lock.Lock()
	defer r.lock.Unlock()
	out := make([]interface{}, 0)
	for e := r.ring.Back(); e != nil; e = e.Prev() {
		out = append(out, e.Value)
	}
	return out
}

// Dump returns the RingBuffer contents from newest to oldest
func (r *RingBuffer) DumpRev() []interface{} {
	r.lock.Lock()
	defer r.lock.Unlock()
	out := make([]interface{}, 0)
	for e := r.ring.Front(); e != nil; e = e.Next() {
		out = append(out, e.Value)
	}
	return out
}
