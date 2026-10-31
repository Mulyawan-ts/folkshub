package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

// menuApp membangun menu aplikasi. ctx berasal dari App.ctx yang diisi startup.
func menuApp(app *App) *menu.Menu {
	m := menu.NewMenu()

	berkas := m.AddSubmenu("Berkas")
	berkas.AddText("Muat Ulang", keys.CmdOrCtrl("r"), func(_ *menu.CallbackData) {
		runtime.WindowReload(app.ctx)
	})
	berkas.AddSeparator()
	berkas.AddText("Keluar", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	bantuan := m.AddSubmenu("Bantuan")
	bantuan.AddText("Tentang Folkshub", nil, func(_ *menu.CallbackData) {
		runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Folkshub",
			Message: "Aplikasi desktop HRIS.\nData disimpan lokal dalam file SQLite.",
		})
	})

	return m
}

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:     "Folkshub",
		Width:     1024,
		Height:    768,
		MinWidth:  900,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Menu:             menuApp(app),
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
