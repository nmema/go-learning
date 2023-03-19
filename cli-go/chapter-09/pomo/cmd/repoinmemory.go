package cmd

import (
	"cli-go/chapter-09/pomo/pomodoro"
	"cli-go/chapter-09/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
