// oberver-and-subject-interface.go

package main

// Subject interface
type Subject interface {
	Attach(o Observer) (bool, error)
	Detach(o Observer) (bool, error)
	Notify() (bool, error)
}

// Observer Interface
type Observer interface {
	Update(string)
}
