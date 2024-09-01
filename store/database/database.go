package database

import (
	"log"
	"os"
	"time"

	config "gitnet.fr/deblan/freetube-sync/config/server"
	"gitnet.fr/deblan/freetube-sync/logger"
	"gitnet.fr/deblan/freetube-sync/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	lg "gorm.io/gorm/logger"
)

type Manager struct {
	Db *gorm.DB
}

var manager *Manager

func GetManager() *Manager {
	if manager == nil {
		manager = &Manager{}
		db, err := gorm.Open(sqlite.Open(config.GetConfig().DbPath), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				lg.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  lg.LogLevel(config.GetConfig().LogLevel),
					IgnoreRecordNotFoundError: true,
					ParameterizedQueries:      true,
					Colorful:                  true,
				},
			),
		})

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
