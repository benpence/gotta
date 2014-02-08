package view

import (
    "gotta/model"
)

type View interface {
    Search(query string) filter.Filter
    FoldDescendants(generation Generation)
}

type Generation int
const (
    Nobody Generation = Generation(iota)
    OneGeneration
    TwoGeneration
    ThreeGeneration
    FourGeneration
    FiveGeneration
    SixGeneration
    SevenGeneration
    EightGeneration
    NineGeneration
)

type Item interface {
    model.Tree

    // Fold hides every descendant of this Item and visibly marks this Item folded
    // Note: If the item is already folded, this will have no effect
    Fold()
    // Fold shows every descendant of this Item that is not folded by another of its descendants
    // Note: If the item not folded, this will have no effect
    Unfold()
    IsFolded() bool

    Hide()
    Show()
    IsShown() bool

    Remove() error
}
