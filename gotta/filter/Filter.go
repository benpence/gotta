package filter

import (
    "gotta/model"
)

// Filter is a lazy filter on the nodes in a Tree
type Filter chan model.Tree

// Walk enumerates all nodes in Tree root, depth first (determistic order)
func Walk(t model.Tree) Filter {
    ch := make(chan model.Tree)

    go func() {
        model.WalkWithDepth(t, func(child model.Tree, depth int) {
            ch <- child
        })

        close(ch)
    }()

    return Filter(ch)
}

// WhereAncestor filters any Tree and its descendants if the Tree does not
// meet a predicate 'p'
// TODO: Change input parameter to a Filter and run in O(n)
func WhereAncestor(t model.Tree, p Predicate) Filter {
    filteredCh := make(chan model.Tree)

    // For each non-compliant Tree, skip all its descendants (wait until
    // another Tree at current or less depth appears)
    go func() {
        skippingDescendants := false
        var badTree int

        model.WalkWithDepth(t, func(tree model.Tree, depth int) {
            if skippingDescendants {
                // Stop skipping descendants?
                if depth <= badTree {
                    skippingDescendants = false

                // Continue skipping
                } else {
                    return
                }
            }

            // Yield this tree?
            if p(tree) {
                filteredCh <- tree

            // Start skipping all this tree's descendants
            } else {
                skippingDescendants = true
                badTree = depth
            }
        })

        close(filteredCh)
    }()

    return Filter(filteredCh)
}

// Tasks maps a Filter channel of Trees onto their respective Tasks
func (f Filter) Tasks() chan model.Task {
    taskCh := make(chan model.Task)

    go func() {
        for tree := range (chan model.Tree)(f) {
            taskCh <- tree.Task()
        }

        close(taskCh)
    }()

    return taskCh
}

// Where filters any Tree from the stream that does not meet a Predicate
func (f Filter) Where(p Predicate) Filter {
    filteredCh := make(chan model.Tree)

    go func() {
        for tree := range chan model.Tree(f) {
            if p(tree) {
                filteredCh <- tree
            }
        }

        close(filteredCh)
    }()

    return Filter(filteredCh)
}

