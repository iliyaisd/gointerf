package asyncrunner

import(
    "github.com/iliyaisd/gointerf/scheduler"
)

type AsyncRunner struct {}

func (ar AsyncRunner) Run(w scheduler.Worker) {
    w.Run()
}