package watchcmd

import (
	"log"

	"github.com/fsnotify/fsnotify"
	config "gitnet.fr/deblan/freetube-sync/config/client"
	filestore "gitnet.fr/deblan/freetube-sync/store/file"
	"gitnet.fr/deblan/freetube-sync/web/client"
	"gitnet.fr/deblan/freetube-sync/web/route"
)

func Process(name, route string, data any) bool {
	log.Print("Push of " + name)
	response, err := client.InitPush(route, data)
	res := true

	if err != nil {
		log.Print("Error while pushing " + name + ": " + err.Error())
		res = false
	} else {
		if response.Code == 201 {
			log.Print(name + " pushed!")
		} else {
			log.Print("Error while pushing " + name + ": " + response.Message)
			res = false
		}
	}

	return res
}

func Run() {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Print("Error while creating the watcher: " + err.Error())
	}

	defer watcher.Close()
	c := config.GetConfig()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					switch event.Name {
					case c.DbPath("history"):
						Process("history", route.HistoryPush, filestore.LoadHistory())
					case c.DbPath("playlists"):
						Process("playlists", route.PlaylistPush, filestore.LoadPlaylists())
					case c.DbPath("profiles"):
						Process("profiles", route.ProfilePush, filestore.LoadProfiles())
					}
				}
			}
		}
	}()

	watcher.Add(config.GetConfig().Path)

	<-make(chan struct{})
}
