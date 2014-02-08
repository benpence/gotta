package filter

import (
    "gotta/model"
)

// Predicate represents a condition to test against a Tree (or stream of Trees)
type Predicate func(model.Tree) bool

// Or will return true if any of its Predicates is true for a Tree
func Or(ps ...Predicate) Predicate {
    return func(t model.Tree) bool {
        for _, p := range ps {
            if p(t) {
                return true
            }
        }

        return false
    }
}

// And will return true if all of its Predicates is true for a Tree
func And(ps ...Predicate) Predicate {
    return func(t model.Tree) bool {
        for _, p := range ps {
            if !p(t) {
                return false
            }
        }

        return true
    }
}

// Not will return the opposite of its Predicate input
func Not(p Predicate) Predicate {
    return func(t model.Tree) bool {
        return !p(t)
    }
}
