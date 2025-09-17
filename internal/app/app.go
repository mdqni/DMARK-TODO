package app

import (
	"context"
	"database/sql"
	"embed"

	_ "github.com/lib/pq"
	"github.com/mdqni/dmark-todo/internal/config"
	"github.com/mdqni/dmark-todo/internal/repository/postgres"
	"github.com/mdqni/dmark-todo/internal/service"
	"github.com/mdqni/dmark-todo/internal/usecase"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type App struct {
	DB     *sql.DB
	TaskUC *usecase.TaskUseCase
}

var assets embed.FS

func NewApp(ctx context.Context, config *config.Config) (*App, error) {
	db, err := sql.Open("postgres", config.ConnString)
	if err != nil {
		return nil, err
	}

	taskRepo, err := postgres.NewPostgresTaskRepo(db)
	if err != nil {
		return nil, err
	}
	taskService := service.NewTaskService(taskRepo)
	taskUC := usecase.NewTaskUseCase(taskService)

	return &App{
		DB:     db,
		TaskUC: taskUC,
	}, nil
}

func (a *App) Run() error {
	appOptions := &options.App{
		Title:  "To-Do App",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			a.TaskUC,
		},
		OnStartup: func(ctx context.Context) {},
		OnShutdown: func(ctx context.Context) {
			_ = a.DB.Close()
		},
	}

	return wails.Run(appOptions)
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.DB.Close()
}
