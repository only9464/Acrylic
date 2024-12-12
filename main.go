package main

import (
	"embed"

	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed assets/icon.png
var appIcon []byte

func main() {
	// Create an instance of the app structure
	appManager := NewAppManager()
	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Acrylic",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        appManager.StartupHandler,
		Bind:             appManager.GetBindings(),
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			Theme:                0,
			BackdropType:         windows.Acrylic,
		},
		WindowStartState: options.Maximised,
		Linux: &linux.Options{
			WindowIsTranslucent: true,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "Acrylic",
			Icon:                appIcon,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
