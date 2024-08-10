package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/routing"
)

func ServeApp() {
	config.Set()
	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
