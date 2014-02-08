package terminal

import (
    "gotta/view"
    "gotta/filter"
)

type Terminal struct {
    root view.Item
    layout []Task
    layout map[Task]Item
    view []Task
}

func NewTerminal() *Terminal {
    
}

func (t Terminal) FoldDescendants(generation view.Generation) {
    switch {
    case generation == view.Nobody:
        for tree := range filter.Walk(t.root) {
            item.Unfold()
            item.Show()
        }

    default:
        foldDepth := int(generation)

        model.WalkWithDepth(func(tree model.Tree, depth int) {
            // TODO: Convert Tree to 

            switch {
            case depth <= foldDepth:
                item.Show()
                continue

            // 
            case depth < foldDepth:  item.Unfold()
            // Last shown descendants
            case depth == foldDepth: item.Fold()
            }
        }

        for tree := range filter.Walk(t.root) {
            if item.IsFolded() {
                item.Unfold()
            }

            if !item.IsShown() {
                item.Show()
            }
        }

}
