//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/riku0202/todo-app-go/go-app/app"
	"github.com/riku0202/todo-app-go/go-app/db"
	"github.com/riku0202/todo-app-go/go-app/infrastructure/mysql/todo"
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
