package history

import (
	"time"

	"github.com/labstack/echo/v4"
	"gitnet.fr/deblan/freetube-sync/model"
	"gitnet.fr/deblan/freetube-sync/store/database"
	"gitnet.fr/deblan/freetube-sync/web/helper"
	"gitnet.fr/deblan/freetube-sync/web/route"
)

func InitPush(c echo.Context) error {
	payload := []model.WatchedVideo{}
	err := c.Bind(&payload)
	manager := database.GetManager()

	if err != nil {
		return helper.Ko(c, err)
	}

	for _, item := range payload {
		manager.Db.Where(item).FirstOrCreate(&item)
	}

	return helper.Ok(c)
}

func Pull(c echo.Context) error {
	entities := []model.WatchedVideo{}
	manager := database.GetManager()

	manager.Db.Find(&entities)

	pull := model.Pull{
		Hostname: c.Request().Header.Get("X-Machine"),
		Database: "history",
	}

	manager.Db.Where(pull).FirstOrCreate(&pull)
	pull.PullAt = time.Now()
	manager.Db.Save(&pull)

	return c.JSON(200, entities)
}

func Register(e *echo.Echo) {
	e.POST(route.HistoryInit, InitPush)
	e.POST(route.HistoryPush, InitPush)
	e.GET(route.HistoryPull, Pull)
}
