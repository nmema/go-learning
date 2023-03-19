package pomodoro_test

import (
	"testing"

	"cli-go/chapter-09/pomo/pomodoro"
	"cli-go/chapter-09/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
