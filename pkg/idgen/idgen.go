package idgen

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const (
	counterFile = "task_counter"
	idFormat    = "TASK-%05d"
)

var mu sync.Mutex

// Next generates the next sequential task ID (TASK-00001, TASK-00002, ...).
func Next(dataDir string) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	counterPath := filepath.Join(dataDir, counterFile)

	current := 0
	data, err := os.ReadFile(counterPath)
	if err == nil {
		if n, err := strconv.Atoi(string(data)); err == nil {
			current = n
		}
	}

	current++
	if err := os.WriteFile(counterPath, []byte(strconv.Itoa(current)), 0644); err != nil {
		return "", fmt.Errorf("failed to write counter: %w", err)
	}

	return fmt.Sprintf(idFormat, current), nil
}
