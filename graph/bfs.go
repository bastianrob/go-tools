package graph

type BFS struct {
	root   Traversable
	action Delegate
}

func NewBFS(
	root Traversable,
) *BFS {
	return &BFS{
		root: root,
	}
}

// Traverse
// TODO: Doesn't prevent circular cycle
func (bfs *BFS) Traverse() {
	queue := []Traversable{bfs.root}

	for len(queue) > 0 {
		node := queue[0]
		shouldBreak := bfs.action(node)
		if shouldBreak {
			return
		}

		children := node.Children()
		if children == nil || len(children) <= 0 {
			queue = queue[1:]
		} else {
			queue = append(queue[1:], children...)
		}
	}
}
