package app

import (
	"github.com/riku0202/todo-app-go/domain/model"
)

type Repository interface {
	Create(todo *model.Todo) error
	Delete(id *model.ID) error
	// UpdateIsFinished(id *model.ID) error
	// FindByUserID(id *model.UserId) (*model.Todo, error)
}

type App struct {
	repo Repository
}

func NewApp(repo Repository) *App {
	app := &App{}
	app.repo = repo

	return app
}
