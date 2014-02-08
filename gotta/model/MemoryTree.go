package model

import (
    "fmt"
    "errors"
)

// NewMemoryTree creates a MemoryTree with 'name' as its Task
func NewMemoryTree(name string) *MemoryTree {
    return &MemoryTree{
        task: NewTask(name),
        children: make([]*MemoryTree, 0),
    }
}

// MemoryTree is an in-memory implementation of a Tree.
type MemoryTree struct {
    task Task
    children []*MemoryTree
}

func (t MemoryTree) Task() Task {
    return t.task
}

func (t MemoryTree) Children() chan Tree {
    ch := make(chan Tree)

    go func(){
        for _, tree := range t.children {
            ch <- Tree(tree)
        }

        close(ch)
    }()

    return ch
}

func (t *MemoryTree) AddChild(task Task) Tree {
    child := &MemoryTree{
        task: task,
        children: make([]*MemoryTree, 0),
    }

    t.children = append(t.children, child)

    return Tree(child)
}

func (t *MemoryTree) RemoveChild(task Task) error {
    for i, child := range t.children {
        // TODO: Better equality?

        if child.task == task {
            t.children = append(t.children[:i], t.children[i + 1:]...)
            return nil
        }
    }

    return errors.New(fmt.Sprintf("%v was not in %v", task, t))
}

func (t MemoryTree) String() string {
    return fmt.Sprintf("MemoryTree(%v)", t.Task())
}
