package main

import (
	"log"
	"pocketbase-lzcapp/ui"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()
	// Initialize app configuration before bootstrapping
	if err := app.Bootstrap(); err != nil {
		log.Fatal(err)
	}

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serve the index.html file
		e.Router.GET(
			"*",
			echo.StaticDirectoryHandler(ui.DistDirFS, false),
			middleware.Gzip(),
		)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
