package main

import (
	"github.com/Marwane97/fx-eventbus/app"
	"github.com/Marwane97/fx-eventbus/service"
	"github.com/Marwane97/fx-eventbus/subscriber"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		subscriber.Module,
		service.Module,
		app.Module,
	).Run()
}
