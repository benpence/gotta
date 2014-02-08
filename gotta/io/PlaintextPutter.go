package io

import (
    "fmt"
    "strings"
    "errors"

    "gotta/model"
)

type PlaintextPutter struct {
    string
}

func NewPlaintextPutter(s string) *PlaintextPutter {
    return &PlaintextPutter{s}
}

func (p PlaintextPutter) Input(contents string) (model.Tree, error) {
    lines := strings.Split(contents, "\n")

    if len(lines) == 0 {
        return nil, errors.New("Empty file")
    }

    name, lines := lines[0], lines[1:]

    path := []model.Tree{ model.Tree(model.NewMemoryTree(name)) }

    indent := p.string
    for i, line := range lines {
        prevGeneration := len(path) - 1
        generation := 0

        // Measure the indentation
        for strings.HasPrefix(line, indent) {
            generation++
            line = line[len(indent):]
        }

        line = strings.Trim(line, " ")

        switch {

        // Duplicate root node
        case generation == 0:
            return nil, errors.New(fmt.Sprintf("Second root node at line %v: \"%v\"", i + 2, line))

        // Sibling or older 'generation' -> Add child and replace path as appropriate
        case generation <= prevGeneration:
            childTree := path[generation - 1].AddChild(model.NewTask(line))
            path[generation] = childTree

            // We could be backtracking many generations here
            path = path[:generation + 1]

        // New 'generation' -> Add child and extend path
        case generation == prevGeneration + 1:
            childTree := path[generation - 1].AddChild(model.NewTask(line))
            path = append(path, childTree)

        // No parent -> error
        default:
            return nil, errors.New(fmt.Sprintf("Invalid input at line %v: \"%v\"", i + 2, line))
        }
    }

    return path[0], nil
}

// TODO: Use some kind of Writer instead?
func (p PlaintextPutter) Output(t model.Tree) (string, error) {
    s := ""

    model.WalkWithDepth(t, func(tree model.Tree, depth int) {
        s += "\n"

        for i := 0; i < depth; i++ {
            s += p.string
        }

        s += tree.Task().String()
    })

    return s, nil
}
