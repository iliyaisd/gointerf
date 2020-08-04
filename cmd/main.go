package main

import (
    "github.com/iliyaisd/gointerf/asyncrunner"
    "github.com/iliyaisd/gointerf/delayedrunner"
    "github.com/iliyaisd/gointerf/scheduler"
    "github.com/iliyaisd/gointerf/textworker"
)

const defaultDelayTime = 4 //seconds

func main() {
    asyncRunner := asyncrunner.NewAsyncRunner()
    delayedRunner := delayedrunner.NewDelayedRunner(defaultDelayTime)
    textWorker := textworker.NewTextWorker()
    defaultScheduler := scheduler.NewScheduler()

    defaultScheduler.Schedule(asyncRunner, textWorker)
    defaultScheduler.Schedule(delayedRunner, textWorker)
}