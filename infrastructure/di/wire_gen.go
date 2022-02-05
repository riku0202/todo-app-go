// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/riku0202/todo-app-go/app"
	"github.com/riku0202/todo-app-go/db"
	"github.com/riku0202/todo-app-go/infrastructure/mysql/todo"
)

// Injectors from wire.go:

func InitApp() (*app.App, error) {
	config, err := db.NewConf()
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.NewDB(config)
	if err != nil {
		return nil, err
	}
	repository, err := todo.NewRepository(sqlDB)
	if err != nil {
		return nil, err
	}
	appApp := app.NewApp(repository)
	return appApp, nil
}

func InitQuery() (*todo.Query, error) {
	config, err := db.NewConf()
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.NewDB(config)
	if err != nil {
		return nil, err
	}
	query, err := todo.NewQuery(sqlDB)
	if err != nil {
		return nil, err
	}
	return query, nil
}