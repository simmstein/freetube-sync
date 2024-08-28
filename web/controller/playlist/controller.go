package playlist

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
	playlists := []model.Playlist{}
	err := c.Bind(&playlists)
	manager := database.GetManager()

	if err != nil {
		return helper.Ko(c, err)
	}

	for _, playlist := range playlists {
		manager.Db.Create(&playlist)
	}

	return helper.Ok(c)
}

func Push(c echo.Context) error {
	playlists := []model.Playlist{}
	err := c.Bind(&playlists)
	manager := database.GetManager()

	if err != nil {
		return helper.Ko(c, err)
	}

	hostname := c.Request().Header.Get("X-Machine")

	pull := model.Pull{
		Hostname: hostname,
		Database: "playlists",
	}

	manager.Db.Where(pull).First(&pull)

	ids := []string{}

	for _, playlist := range playlists {
		if playlist.PlaylistName == "" {
			continue
		}

		var existingPlaylist model.Playlist
		manager.Db.Preload("Videos").Where(model.Playlist{
			RemoteId: playlist.RemoteId,
		}).First(&existingPlaylist)

		if existingPlaylist.ID == 0 {
			playlist.Hostname = hostname
			manager.Db.Create(&playlist)
			ids = append(ids, playlist.RemoteId)
		} else {
			existingPlaylist.Description = playlist.Description
			existingPlaylist.LastUpdatedAt = playlist.LastUpdatedAt
			existingPlaylist.PlaylistName = playlist.PlaylistName
			existingPlaylist.Protected = playlist.Protected

			for _, v := range existingPlaylist.Videos {
				manager.Db.Delete(v)
			}

			existingPlaylist.Videos = playlist.Videos
			manager.Db.Save(&existingPlaylist)
			ids = append(ids, existingPlaylist.RemoteId)
		}
	}

	if len(ids) > 0 {
		var playlistsToDelete []model.Playlist

		manager.Db.Find(
			&playlistsToDelete,
			"remote_id not in (?) and (created_at < ? or hostname = ?)",
			ids,
			pull.PullAt,
			hostname,
		)

		for _, entity := range playlistsToDelete {
			manager.Db.Select(clause.Associations).Delete(&entity)
		}
	}

	return helper.Ok(c)
}

func Pull(c echo.Context) error {
	playlists := []model.Playlist{}
	manager := database.GetManager()

	manager.Db.Preload("Videos").Find(&playlists)

	pull := model.Pull{
		Hostname: c.Request().Header.Get("X-Machine"),
		Database: "playlist",
	}

	manager.Db.Where(pull).FirstOrCreate(&pull)
	pull.PullAt = time.Now()
	manager.Db.Save(&pull)

	return c.JSON(200, playlists)
}

func Register(e *echo.Echo) {
	e.POST(route.PlaylistsInit, Init)
	e.POST(route.PlaylistsPush, Push)
	e.GET(route.PlaylistsPull, Pull)
}
