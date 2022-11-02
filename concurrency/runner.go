package concurrency

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(task ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
		fmt.Println("complete")
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		fmt.Println("timeout")
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) goInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
