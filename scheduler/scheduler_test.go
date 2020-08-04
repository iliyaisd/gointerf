package scheduler_test

import (
	"log"
	"testing"

	"github.com/iliyaisd/gointerf/scheduler"
	"github.com/stretchr/testify/require"
)

func TestScheduler_Schedule(t *testing.T) {
	testScheduler := scheduler.NewScheduler()

	err := testScheduler.Schedule(mockRunner{}, mockWorker{})
	require.NoError(t, err)
}

type mockRunner struct{}

func (mr mockRunner) Run(_ scheduler.Worker) {
	log.Println("Runner mock works")
}

type mockWorker struct{}

func (mw mockWorker) Do() {
	log.Println("Worker mock works")
}
