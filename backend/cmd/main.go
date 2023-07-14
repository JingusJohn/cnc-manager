package main

import (
	"go.uber.org/fx"

	"github.com/JingusJohn/kyle-cnc/backend/controllers"
	"github.com/JingusJohn/kyle-cnc/backend/server"
	"github.com/JingusJohn/kyle-cnc/backend/storage"
)

func main() {
	fx.New(
		fx.Provide(
			// create general handlers
			controllers.CreateGeneralController,
			controllers.CreateAuthController,
			// create storage solutions
			storage.NewMinio,
		),
		fx.Invoke(server.CreateServer),
	).Run()
}
