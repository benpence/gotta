package io

import (
    "gotta/model"
)

var HtmlOutputter = NewOutputter(htmlOutput)
func htmlOutput(t model.Tree) (string, error) {
    s := "<html>\n<body>\n<ul>"

    prevDepth := 0
    model.WalkWithDepth(t, func(tree model.Tree, depth int) {
        switch {
        case depth > prevDepth:
            s += "<ul>\n"
        case depth < prevDepth:
            s += "</ul>\n"
        }

        s += "<li>" + tree.Task().String() + "<li>\n"
    })

    s += "</ul>\n</body>\n</html>"

    return s, nil
}
