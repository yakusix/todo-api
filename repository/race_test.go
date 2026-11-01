package repository

import (
	"sync"
	"testing"
)

func TestCreateConcurrent(t *testing.T) {
	repo := NewInMemoryRepo()
	const n = 100

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			repo.Create(Task{Title: "ТЕСТ"})
		}()

	}

	wg.Wait()

	tasks, _ := repo.List()
	if len(tasks) != n {
		t.Fatalf("expected %d tasks, got %d", n, len(tasks))
	}
}
