//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/riku0202/todo-app-go/app"
	"github.com/riku0202/todo-app-go/db"
	"github.com/riku0202/todo-app-go/infrastructure/mysql/todo"
)

func InitApp() (*app.App, error) {
	wire.Build(
		app.NewApp,
		todo.NewRepository,
		wire.Bind(new(app.Repository), new(*todo.Repository)),
		db.NewDB,
		db.NewConf,
	)

	return nil, nil
}

func InitQuery() (*todo.Query, error) {
	wire.Build(
		todo.NewQuery,
		db.NewDB,
		db.NewConf,
	)

	return nil, nil
}
