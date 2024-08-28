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
	watchedVideos := []model.WatchedVideo{}
	err := c.Bind(&watchedVideos)
	manager := database.GetManager()

	if err != nil {
		return helper.Ko(c, err)
	}

	for _, watchedVideo := range watchedVideos {
		manager.Db.Where(watchedVideo).FirstOrCreate(&watchedVideo)
	}

	return helper.Ok(c)
}

func Pull(c echo.Context) error {
	watchedVideos := []model.WatchedVideo{}
	manager := database.GetManager()

	manager.Db.Find(&watchedVideos)

	pull := model.Pull{
		Hostname: c.Request().Header.Get("X-Machine"),
		Database: "history",
	}

	manager.Db.Where(pull).FirstOrCreate(&pull)
	pull.PullAt = time.Now()
	manager.Db.Save(&pull)

	return c.JSON(200, watchedVideos)
}

func Register(e *echo.Echo) {
	e.POST(route.HistoryInit, InitPush)
	e.POST(route.HistoryPush, InitPush)
	e.GET(route.HistoryPull, Pull)
}
