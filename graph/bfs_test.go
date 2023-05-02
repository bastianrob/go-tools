package graph

import (
	"reflect"
	"testing"
)

type traversable_mock struct {
	value    int
	children []Traversable
}

func (t *traversable_mock) Children() []Traversable {
	return t.children
}

func TestBFS_Traverse(t *testing.T) {
	result := []int{}
	bfs := &BFS{
		root: &traversable_mock{
			value: 1,
			children: []Traversable{
				&traversable_mock{value: 2, children: []Traversable{
					&traversable_mock{value: 4, children: nil},
					&traversable_mock{value: 5, children: nil},
				}},
				&traversable_mock{value: 3, children: []Traversable{
					&traversable_mock{value: 6, children: nil},
					&traversable_mock{value: 7, children: nil},
				}},
			},
		},
		action: func(n Traversable) Break {
			node, _ := n.(*traversable_mock)
			result = append(result, node.value)

			return false
		},
	}

	bfs.Traverse()

	if !reflect.DeepEqual(result, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("traversal order should be 1 2 3 4 5 6 7, instead got: %v", result)
	}
}
