package app

import (
	"fmt"

	"github.com/riku0202/todo-app-go/go-app/domain/model"
)

func (app App) CreateTodo(userId string, title string, description string) error {

	u, err := model.NewUserId(description)
	if err != nil {
		return fmt.Errorf("userIdを作成できません:%v", err)
	}

	t, err := model.NewTitle(description)
	if err != nil {
		return fmt.Errorf("titleを作成できません:%v", err)
	}

	d, err := model.NewDescription(description)
	if err != nil {
		return fmt.Errorf("descriptionを作成できません:%v", err)
	}

	todo, err := model.NewTodo(u, t, d)
	if err != nil {
		return fmt.Errorf("todoを作成できません:%v", err)
	}

	if err = app.repo.Create(todo); err != nil {
		return fmt.Errorf("todoを作成できません:%v", err)
	}

	return nil
}
