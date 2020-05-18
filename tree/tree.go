package tree

//TreeNode ...
type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

//Tree ...
type Tree struct {
	Root *TreeNode
}

//AddNode ...
func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}
