package model

// Tree is a tree data structure that holds Tasks.
// Each Tree may have 0 or more child Trees.
// Note: In this documentation, a Tree's "child" has that Tree as a parent,
//   whereas a Tree's "descendant" has that Tree as an ancestor (parent's
//   parent, etc.).
// Note: A Tree data structure may not enforce concurrency.
type Tree interface {
    // Task returns the Task stored at the root node of this Tree
    Task() Task

    // Children will iterate through the MemoryTree's children.
    // The order will be consistent, assuming no children are added/removed.
    Children() chan Tree

    // AddChild will append to its children a Tree containting storing 'task'.
    AddChild(task Task) Tree

    // RemoveChild will remove 'task' from its children.
    // An error will be returned if 'task' is not stored by one of 
    RemoveChild(task Task) error
}

// Equals performs a deep equality check of two Trees.
// Note: Equals is recursive and may cause a stack overflow on very deep Trees
func Equals(first, second Tree) bool {
    if first.Task().String() != second.Task().String() {
        return false
    }

    firstChildren, secondChildren := first.Children(), second.Children()

    // TODO: This may take forever if there is a loop in the data structure(s)
    for {
        firstChild, firstOk := <- firstChildren
        secondChild, secondOk := <- secondChildren

        switch {
        // Different number of children -> unequal
        case firstOk != secondOk: return false

        // All children were equal -> equal
        case firstOk == false: return true

        // Is the current child the same?
        case !Equals(firstChild, secondChild): return false
        }
    }
}

// WalkWithDepth runs a function on all nodes in Tree root, depth first.
// 'f' should expect as parameters the current Tree node and the current depth
//   (0 being the depth of the root node)
func WalkWithDepth(t Tree, f func(Tree, int)) {
    stack := []treeAndInt{ treeAndInt{t, 0} }
    var pair treeAndInt

    // Use stack for depth-first enumeration
    for len(stack) != 0 {
        // Pop
        pair, stack = stack[0], stack[1:]
        tree, depth := pair.Tree, pair.int

        // Yield the current task
        f(tree, depth)

        // Add current task's children to stack
        for childTree := range tree.Children() {
            stack = append(stack, treeAndInt{childTree, depth + 1})
        }
    }
}

type treeAndInt struct {
    Tree
    int
}
