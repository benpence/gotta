package main

import (
    "fmt"

//"gotta/model"
    f "gotta/filter"
    "gotta/io"
)

const input = `root #abc
  child1 #doom 
    child1child1 "fuck all of you" #doom
    child1child2 "fuck all of you" #doom
    child1child3 "fuck all of you" #doom
  child2 #doom #abcd
  child3 #doom #abc
    child3child1 "fuck all of you" #こんにちは
    child3child2 "fuck all of you"
    child3child3 "fuck all of you"`


func main() {
    p := io.NewPlaintextPutter("  ")
    if tree, err := p.Input(input); err != nil {
        fmt.Println(err)

    } else {
        tasks := f.Walk(tree).
            Where(f.And(
                f.HasWords("#こんにちは"),
                f.HasWords("all"),
                f.IsLeaf)).
            Tasks()

        fmt.Println("Matching:")
        for task := range tasks {
            fmt.Println(task)
        }

        if s, err := p.Output(tree); err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(s)
        }
    }
}
