package LRU_Cache

type command int

const (
	// MoveToFront define
	MoveToFront command = iota
	// PushFront define
	PushFront
	// Delete define
	Delete
)

type clear struct {
	done chan struct{}
}