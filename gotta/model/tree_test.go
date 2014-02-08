package model

import (
    "testing"
)

var TreeTest1 Tree = Tree(
    &MemoryTree{
        task: Task("root"),
        children: []*MemoryTree{

            &MemoryTree{
                task: Task("root1"),
                children: []*MemoryTree{

                    &MemoryTree{
                        task: Task("root1child1"),
                        children: []*MemoryTree{

                            &MemoryTree{
                                task: Task("root1child1child1"),
                                children: []*MemoryTree{},
                            },

                            &MemoryTree{
                                task: Task("root1child1child2"),
                                children: []*MemoryTree{},
                            },
                        },
                    },
                },
            },

            &MemoryTree{
                task: Task("root2"),
                children: []*MemoryTree{

                    &MemoryTree{
                        task: Task("root2child1"),
                        children: []*MemoryTree{},
                    },
                },
            },
        },
    },
)

var TreeTest2 Tree = Tree(
    &MemoryTree{
        task: Task("root"),
        children: []*MemoryTree{

            &MemoryTree{
                task: Task("root1"),
                children: []*MemoryTree{

                    &MemoryTree{
                        task: Task("root1child1"),
                        children: []*MemoryTree{

                            &MemoryTree{
                                task: Task("root1child1child1"),
                                children: []*MemoryTree{},
                            },

                            &MemoryTree{
                                task: Task("root1child1child2"),
                                children: []*MemoryTree{},
                            },
                        },
                    },
                },
            },

            &MemoryTree{
                task: Task("root2"),
                children: []*MemoryTree{

                    &MemoryTree{
                        task: Task("root2child1"),
                        children: []*MemoryTree{},
                    },
                },
            },
        },
    },
)

var TreeTest3 Tree = Tree(
    &MemoryTree{
        task: Task("root"),
        children: []*MemoryTree{

            &MemoryTree{
                task: Task("root1"),
                children: []*MemoryTree{},
            },

            &MemoryTree{
                task: Task("root2"),
                children: []*MemoryTree{

                    &MemoryTree{
                        task: Task("root2child1"),
                        children: []*MemoryTree{},
                    },

                    &MemoryTree{
                        task: Task("root2child2"),
                        children: []*MemoryTree{

                            &MemoryTree{
                                task: Task("root1child1child1"),
                                children: []*MemoryTree{},
                            },

                            &MemoryTree{
                                task: Task("root1child1child2"),
                                children: []*MemoryTree{},
                            },
                        },
                    },
                },
            },
        },
    },
)

var EqualTest [][]Tree = [][]Tree{
    []Tree{ TreeTest1, TreeTest2 },
    []Tree{ TreeTest1, TreeTest1 },
    []Tree{ TreeTest2, TreeTest2 },
    []Tree{ TreeTest3, TreeTest3 },
}

var UnequalTest [][]Tree = [][]Tree{
    []Tree{ TreeTest1, TreeTest3 },
    []Tree{ TreeTest2, TreeTest3 },
}

func TestEquals(t *testing.T) {
    // Equal
    for i, pair := range EqualTest {
        if !Equals(pair[0], pair[1]) {
            t.Errorf("Equals test %v: %v should equal %v", i, pair[0], pair[1])
        }
    }

    // Unequal
    for i, pair := range UnequalTest {
        if Equals(pair[0], pair[1]) {
            t.Errorf("Equals test %v: %v should equal %v", i, pair[0], pair[1])
        }
    }
}
