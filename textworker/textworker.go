package textworker

import "fmt"

type TextWorker struct {}

func NewTextWorker() TextWorker {
    return TextWorker{}
}

func (tw TextWorker) Do() {
    fmt.Println("I am working")
}
