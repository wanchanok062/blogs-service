package util

import "sync"

func RunConcurrentTasks(tasks ...func()) {
	var wg sync.WaitGroup

	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(t func()) {
			defer wg.Done()
			t()
		}(task)
	}

	wg.Wait()
}
