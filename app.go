package main

import (
	"context"
	"file-organizer-desktop/application"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type app struct {
	ctx context.Context
}

func NewApp() *app {
	return &app{}
}

func (app *app) startup(ctx context.Context) {
	app.ctx = ctx
}

func (app *app) OpenDirectory() (string, error) {
	return runtime.OpenDirectoryDialog(app.ctx, runtime.OpenDialogOptions{
		CanCreateDirectories: false,
	})
}

func (*app) Organize(sortBy string, dirPath string) error {
	if sortBy == "-n" {
		err := application.OrganizeFilesByName(dirPath)

		if err != nil {
			return err
		}
	} else if sortBy == "-d" {
		err := application.OrganizeFilesByDate(dirPath)

		if err != nil {
			return err
		}
	}

	return nil
}
