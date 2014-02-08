package model

import (
    "testing"
)

func TestNewMemoryTree(t *testing.T) {
    tree := NewMemoryTree("test")

    if tree.task.String() != "test" {
        t.Errorf("%v should be \"test\"'test", tree.task)
    }

    if tree.children == nil || len(tree.children) != 0 {
        t.Errorf("%v should be []*Task{}", tree.children)
    }
}

func TestAddChild(t *testing.T) {
    tree := NewMemoryTree("test")

    firstTask := NewTask("Root1")
    firstResult := tree.AddChild(firstTask)
    secondTask := NewTask("Root2")
    secondResult := tree.AddChild(secondTask)

    ch := tree.Children()

    if result := <- ch; result != firstResult || result.Task() != firstTask {
        t.Errorf("%v should be %v and %v should be %v", result, firstResult, result.Task(), firstTask)
    }

    if result := <- ch; result != secondResult || result.Task() != secondTask {
        t.Errorf("%v should be %v and %v should be %v", result, secondResult, result.Task(), secondTask)
    }
}

func TestRemoveChild(t *testing.T) {
    tree := NewMemoryTree("test")

    firstTask := NewTask("Root1")
    firstResult := tree.AddChild(firstTask)
    secondTask := NewTask("Root2")
    secondResult := tree.AddChild(secondTask)

    ch := tree.Children()

    if result := <- ch; result != firstResult || result.Task() != firstTask {
        t.Errorf("%v should be %v and %v should be %v", result, firstResult, result.Task(), firstTask)
    }

    if result := <- ch; result != secondResult || result.Task() != secondTask {
        t.Errorf("%v should be %v and %v should be %v", result, secondResult, result.Task(), secondTask)
    }
}
