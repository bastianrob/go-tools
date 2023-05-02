package graph

type Traversable interface {
	Children() []Traversable
}

type Break = bool

type Delegate = func(t Traversable) Break
