package ringbuffer

import (
	"testing"
)

// TestRB exercises our three RingBuffer methods: New, Add, and Dump
func TestRB(t *testing.T) {
	r := New(5)
	if len(r.Dump()) > 0 {
		t.Errorf("Empty RB not empty!?")
	}

	f := []string{"a", "bb", "ccc"}
	for _, s := range f {
		r.Add(s)
	}
	if len(r.Dump()) != 3 {
		t.Errorf("RB wrong length %d != %d", len(r.Dump()), 3)
	}
	for i, s := range r.Dump() {
		if f[i] != s.(string) {
			t.Errorf("RB wrong contents")
		}
	}

	r.Add("dddd")
	r.Add("eeeee")
	f = append(f, "dddd")
	f = append(f, "eeeee")
	if len(r.Dump()) != 5 {
		t.Errorf("RB wrong length2 %d != %d", len(r.Dump()), 5)
	}
	for i, s := range r.Dump() {
		if f[i] != s.(string) {
			t.Errorf("RB wrong contents2")
		}
	}

	r.Add("ffffff")
	r.Add("ggggggg")
	f = []string{"ccc", "dddd", "eeeee", "ffffff", "ggggggg"}
	if len(r.Dump()) != 5 {
		t.Errorf("RB wrong length3 %d != %d", len(r.Dump()), 5)
	}
	for i, s := range r.Dump() {
		if f[i] != s.(string) {
			t.Errorf("RB wrong contents3")
		}
	}

}
