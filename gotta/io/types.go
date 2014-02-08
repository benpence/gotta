package io

import (
    "gotta/model"
)

type Inputter interface {
    Input(string) (model.Tree, error)
}

type Outputter interface {
    Output(model.Tree) (string, error)
}

func NewInputter(f func(string) (model.Tree, error)) Inputter {
    return Inputter(inputter(f))
}
type inputter func(string) (model.Tree, error)
func (i inputter) Input(s string) (model.Tree, error) {
    t, err := i(s)
    return t, err
}

func NewOutputter(f func(model.Tree) (string, error)) Outputter {
    return Outputter(outputter(f))
}
type outputter func(model.Tree) (string, error)
func (o outputter) Output(t model.Tree) (string, error) {
    s, err := o(t)
    return s, err
}
