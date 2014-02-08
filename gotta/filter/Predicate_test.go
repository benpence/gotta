package filter

import (
    "testing"

    "gotta/model"
)

func testFalse(t model.Tree) bool {
    return false
}

func testTrue(t model.Tree) bool {
    return true
}

var testOrTrues []Predicate = []Predicate{
    Or(testTrue),
    Or(testFalse, testTrue),
    Or(testFalse, testTrue),
}

var testOrFalses []Predicate = []Predicate{
    Or(testFalse),
    Or(testFalse, testFalse),
}

func TestOr(t *testing.T) {
    for i, test := range testOrTrues {
        if !test(model.NewMemoryTree("test")) {
            t.Errorf("Test %v should be true", i)
        }
    }

    for i, test := range testOrFalses {
        if test(model.NewMemoryTree("test")) {
            t.Errorf("Test %v should be false", i)
        }
    }
}

var testAndTrues []Predicate = []Predicate{
    And(testTrue),
    And(testTrue, testTrue),
}

var testAndFalses []Predicate = []Predicate{
    And(testFalse),
    And(testFalse, testTrue),
    And(testFalse, testFalse),
}

func TestAnd(t *testing.T) {
    for i, test := range testAndTrues {
        if !test(model.NewMemoryTree("test")) {
            t.Errorf("Test %v should be true", i)
        }
    }

    for i, test := range testAndFalses {
        if test(model.NewMemoryTree("test")) {
            t.Errorf("Test %v should be false", i)
        }
    }
}

var testNotTrues []Predicate = []Predicate{
    Not(testFalse),
    Not(Not(Not(testFalse))),
}

var testNotFalses []Predicate = []Predicate{
    Not(testTrue),
    Not(Not(testFalse)),
}

func TestNot(t *testing.T) {
    for i, test := range testNotTrues {
        if !test(model.NewMemoryTree("test")) {
            t.Errorf("Test %v should be true", i)
        }
    }

    for i, test := range testNotFalses {
        if test(model.NewMemoryTree("test")) {
            t.Errorf("Test %v should be false", i)
        }
    }
}
