package main

import (
	"context"
	"encoding/json"
	"fmt"
	"next/common"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetAppInfo 获取程序基础信息
func (a *App) GetAppInfo() string {
	b, err := json.Marshal(common.App)
	if err != nil {
		return ""
	}
	return string(b)
}
