package watchcmd

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	config "gitnet.fr/deblan/freetube-sync/config/client"
)

func Run() {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Print("Error while creating the watcher: " + err.Error())
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					switch event.Name {
					case config.GetConfig().Path + "/history.db":
						fmt.Printf("%+v\n", "update history")
					case config.GetConfig().Path + "/playlists.db":
						fmt.Printf("%+v\n", "update playlists")
					case config.GetConfig().Path + "/profiles.db":
						fmt.Printf("%+v\n", "update profiles")
					}
				}
			}
		}
	}()

	watcher.Add(config.GetConfig().Path)

	<-make(chan struct{})
}
