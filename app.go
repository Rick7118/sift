package main

import (
	"context"

	"github.com/Rick7118/sift/database"
)

type App struct {
	ctx context.Context
	db  *database.Database
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	db, err := database.Open("sift.db")
	if err != nil {
		panic(err)
	}

	a.db = db
}

func (a *App) shutdown(ctx context.Context) {
	if a.db != nil {
		a.db.Close()
	}
}

func (a *App) ExecuteQuery(query string) (*database.QueryResult, error) {
	return a.db.Execute(query)
}

func (a *App) GetTables() ([]string, error) {
	return a.db.Tables()
}

func (a *App) GetSchema(table string) ([]database.Column, error) {
	return a.db.Schema(table)
}
