package delayedrunner

import(
    "github.com/iliyaisd/gointerf/scheduler"
    "time"
)

type DelayedRunner struct {
    delayTime int
}

func NewDelayedRunner(delayTime int) DelayedRunner {
    return DelayedRunner{
        delayTime: delayTime,
    }
}

func (dr DelayedRunner) Run(w scheduler.Worker) {
    time.Sleep(time.Duration(dr.delayTime) * time.Second)
    w.Do()
}
