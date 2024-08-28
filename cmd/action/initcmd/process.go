package initcmd

import (
	"log"
	"os"

	filestore "gitnet.fr/deblan/freetube-sync/store/file"
	"gitnet.fr/deblan/freetube-sync/web/client"
	"gitnet.fr/deblan/freetube-sync/web/route"
)

func Process(name, route string, data any) bool {
	log.Print("Init of " + name)
	response, err := client.InitPush(route, data)
	res := true

	if err != nil {
		log.Print("Error while initializing " + name + ": " + err.Error())
		res = false
	} else {
		if response.Code == 201 {
			log.Print(name + " initialized!")
		} else {
			log.Print("Error while initializing " + name + ": " + response.Message)
			res = false
		}
	}

	return res
}

func Run() {
	a := Process("history", route.HistoryInit, filestore.LoadHistory())
	b := Process("playlists", route.PlaylistsInit, filestore.LoadPlaylists())
	c := Process("profiles", route.ProfilesInit, filestore.LoadProfiles())

	if a && b && c {
		os.Exit(0)
	}

	os.Exit(1)
}
