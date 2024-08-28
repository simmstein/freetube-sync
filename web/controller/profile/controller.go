package profile

import (
	"time"

	"github.com/labstack/echo/v4"
	"gitnet.fr/deblan/freetube-sync/model"
	"gitnet.fr/deblan/freetube-sync/store/database"
	"gitnet.fr/deblan/freetube-sync/web/helper"
	"gitnet.fr/deblan/freetube-sync/web/route"
	"gorm.io/gorm/clause"
)

func Init(c echo.Context) error {
	profiles := []model.Profile{}
	err := c.Bind(&profiles)
	manager := database.GetManager()

	if err != nil {
		return helper.Ko(c, err)
	}

	for _, profile := range profiles {
		manager.Db.Create(&profile)
	}

	return helper.Ok(c)
}

func Push(c echo.Context) error {
	profiles := []model.Profile{}
	err := c.Bind(&profiles)
	manager := database.GetManager()

	if err != nil {
		return helper.Ko(c, err)
	}

	hostname := c.Request().Header.Get("X-Machine")

	pull := model.Pull{
		Hostname: hostname,
		Database: "playlist",
	}

	manager.Db.Where(pull).First(&pull)

	ids := []string{}

	for _, profile := range profiles {
		if profile.Name == "" {
			continue
		}

		var existingProfile model.Profile
		manager.Db.Preload("Subscriptions").Where(model.Profile{
			RemoteId: profile.RemoteId,
		}).First(&existingProfile)

		if existingProfile.ID == 0 {
			manager.Db.Create(&profile)
			ids = append(ids, profile.Name)
		} else {
			existingProfile.Name = profile.Name
			existingProfile.BgColor = profile.BgColor
			existingProfile.TextColor = profile.TextColor

			for _, v := range existingProfile.Subscriptions {
				manager.Db.Select(clause.Associations).Delete(v)
			}

			existingProfile.Subscriptions = profile.Subscriptions
			manager.Db.Save(&existingProfile)
			ids = append(ids, existingProfile.Name)
		}
	}

	if len(ids) > 0 {
		var profilesToDelete []model.Profile

		manager.Db.Find(
			&profilesToDelete,
			"name not in (?)",
			ids,
		)

		for _, entity := range profilesToDelete {
			manager.Db.Select(clause.Associations).Delete(&entity)
		}
	}

	return helper.Ok(c)
}

func Pull(c echo.Context) error {
	profiles := []model.Profile{}
	manager := database.GetManager()

	manager.Db.Preload("Subscriptions").Find(&profiles)

	pull := model.Pull{
		Hostname: c.Request().Header.Get("X-Machine"),
		Database: "profiles",
	}

	manager.Db.Where(pull).FirstOrCreate(&pull)
	pull.PullAt = time.Now()
	manager.Db.Save(&pull)

	return c.JSON(200, profiles)
}

func Register(e *echo.Echo) {
	e.POST(route.ProfilesInit, Init)
	e.POST(route.ProfilesPush, Push)
	e.GET(route.ProfilesPull, Pull)
}
