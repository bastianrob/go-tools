package graph

type DFS struct {
	root   Traversable
	action Delegate
}

func NewDFS(
	root Traversable,
) *DFS {
	return &DFS{
		root: root,
	}
}

// PreOrderTraverse
// TODO: Doesn't prevent circular cycle
func (dfs *DFS) PreOrderTraverse(parent Traversable) {
	children := parent.Children()
	shouldBreak := dfs.action(parent)
	if shouldBreak {
		return
	}

	for i := 0; i < len(children); i++ {
		dfs.PreOrderTraverse(children[i])
	}
}

// PostOrderTraverse
// TODO: Doesn't prevent circular cycle
func (dfs *DFS) PostOrderTraverse(parent Traversable) {
	children := parent.Children()
	for i := 0; i < len(children); i++ {
		dfs.PostOrderTraverse(children[i])
	}

	dfs.action(parent)
}
