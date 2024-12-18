package Settings

import (
	"context"
	"fmt"
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
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// 供前端直接调用的函数都写在这里

func (s *App) Settings(a, b int) int {
	return settings(a, b)
}

func (s *App) Check_now_is_latest(currentVersion, latestVersion string) bool {
	return check_now_is_latest(currentVersion, latestVersion)
}

func (s *App) Get_latest_version_code() string {
	latestVersion := get_latest_version_code()
	fmt.Println("latestVersion", latestVersion)
	return latestVersion
}
