package heap

import (
	"testing"
)

func Test_Push(t *testing.T) {
	h := Heap{
		Objects:     make([]*Object, 0),
		IndexObject: make(map[string]int),
	}
	h.Push(&Object{
		Key: "1",
		Exp: 3,
	})
	h.Push(&Object{
		Key: "2",
		Exp: -2,
	})
	if o, flag := h.Pop(); !flag || o.Key != "2" {
		t.Error("not correct")
	}
	h.Push(&Object{
		Key: "3",
		Exp: 7,
	})
	h.Push(&Object{
		Key: "8",
		Exp: 11,
	})
	if o, flag := h.GetLastItem(); !flag || o.Key != "1" {
		t.Error("not correct")
	}
	h.Push(&Object{
		Key: "9",
		Exp: -1,
	})
	if o, flag := h.Pop(); !flag || o.Key != "9" {
		t.Error("not correct")
	}
	h.Push(&Object{
		Key: "89",
		Exp: 1,
	})
	if o, flag := h.Pop(); !flag || o.Key != "89" {
		t.Error("not correct")
	}
	h.Push(&Object{
		Key: "17",
		Exp: -20,
	})
	if o, flag := h.GetLastItem(); !flag || o.Key != "17" {
		t.Error("not correct")
	}
	h.ChangeObject(&Object{
		Key: "17",
		Exp: -40,
	})
	if o, flag := h.GetLastItem(); !flag || o.Key != "17" {
		t.Error("not correct")
	}
	h.ChangeObject(&Object{
		Key: "3",
		Exp: -80,
	})
	if o, flag := h.GetLastItem(); !flag || o.Key != "3" {
		t.Error("not correct")
	}
	h.ChangeObject(&Object{
		Key: "3",
		Exp: 80,
	})
	if o, flag := h.Pop(); !flag || o.Key != "17" {
		t.Error("not correct")
	}
}
