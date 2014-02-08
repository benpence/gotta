package model

type Task string

func NewTask(s string) Task {
    return Task(s)
}

func (t Task) String() string {
    return string(t)
}
