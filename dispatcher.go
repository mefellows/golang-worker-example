package main

import "fmt"

type WorkerPoolType chan chan WorkRequest

var WorkerPool WorkerPoolType

func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerPool = make(WorkerPoolType, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerPool)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-WorkerPool

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
