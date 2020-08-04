package asyncrunner

import(
    "github.com/iliyaisd/gointerf/scheduler"
)

type AsyncRunner struct {}

func NewAsyncRunner() AsyncRunner {
    return AsyncRunner{}
}

func (ar AsyncRunner) Run(w scheduler.Worker) {
    go w.Do()
}