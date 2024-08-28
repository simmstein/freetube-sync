package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "gitnet.fr/deblan/freetube-sync/config/server"
	"gitnet.fr/deblan/freetube-sync/store/database"
	"gitnet.fr/deblan/freetube-sync/web/controller/history"
	"gitnet.fr/deblan/freetube-sync/web/controller/playlist"
	"gitnet.fr/deblan/freetube-sync/web/controller/profile"
)

func main() {
	config.InitConfig()

	database.GetManager().AutoMigrate()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	history.Register(e)
	playlist.Register(e)
	profile.Register(e)

	e.Logger.Fatal(e.Start(config.GetConfig().BindAddress))
}
