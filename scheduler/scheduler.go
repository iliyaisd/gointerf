package scheduler

type Scheduler struct {}

func NewScheduler() Scheduler {
    return Scheduler{}
}

type Runner interface {
    Run(w Worker)
}

type Worker interface {
    Do()
}

func (s Scheduler) Schedule(r Runner, w Worker) error {
    r.Run(w)
    return nil
}