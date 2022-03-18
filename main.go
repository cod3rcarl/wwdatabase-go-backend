package main

import (
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/app"
)

func main() {
	a := app.NewApp().WithLogger().WithClient().WithServer().WithGRPCServer().WithGracefulShutdown()
	a.Start()
}
