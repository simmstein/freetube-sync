package database

import (
	"log"

	config "gitnet.fr/deblan/freetube-sync/config/server"
	"gitnet.fr/deblan/freetube-sync/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Manager struct {
	Db *gorm.DB
}

var manager *Manager

func GetManager() *Manager {
	if manager == nil {
		manager = &Manager{}
		db, err := gorm.Open(sqlite.Open(config.GetConfig().DbPath), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}

		manager.Db = db
	}

	return manager
}

func (m *Manager) AutoMigrate() {
	m.Db.AutoMigrate(&model.Pull{})
	m.Db.AutoMigrate(&model.WatchedVideo{})
	m.Db.AutoMigrate(&model.PlaylistVideo{})
	m.Db.AutoMigrate(&model.Subscription{})
	m.Db.AutoMigrate(&model.Playlist{})
	m.Db.AutoMigrate(&model.Profile{})
}
