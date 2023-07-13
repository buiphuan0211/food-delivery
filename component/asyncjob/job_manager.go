package asyncjob

import (
	"context"
	"fmt"
	"food-delivery/common"
	"log"
	"sync"
)

// group ...
type group struct {
	jobs         []Job
	isConcurrent bool
	wg           *sync.WaitGroup
}

// NewGroup ...
func NewGroup(isConcurrent bool, jobs ...Job) *group {
	g := &group{
		jobs:         jobs,
		isConcurrent: isConcurrent,
		wg:           new(sync.WaitGroup),
	}

	return g
}

// Run ...
func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	for i, _ := range g.jobs {
		if g.isConcurrent {
			go func(aj Job) {
				defer common.AppRecover()
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()

			}(g.jobs[i])

			continue
		}

		// nếu muốn cancel ngay khi gặp lỗi
		job := g.jobs[i]

		err := g.runJob(ctx, job)

		if err != nil {
			return err
		}

		errChan <- err
		g.wg.Done()
	}

	g.wg.Wait()

	var err error

	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			return v
		}
	}

	fmt.Println("Done group !!!")

	return err
}

// RunJob ...
func (g *group) runJob(ctx context.Context, j Job) error {
	if err := j.Execute(ctx); err != nil {
		for {
			log.Println(err)

			if j.State() == StateRetryFailed {
				return err
			}

			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}

	return nil
}
