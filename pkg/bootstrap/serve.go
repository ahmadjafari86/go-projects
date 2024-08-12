package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/database"
	"blog/pkg/html"
	"blog/pkg/routing"
	"blog/pkg/static"
)

func ServeApp() {
	config.Set()
	database.Connect()
	routing.Init()
	static.LoadStatic(routing.GetRouter())
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRoutes()
	routing.Serve()
}
