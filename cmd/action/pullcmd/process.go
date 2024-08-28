package pullcmd

import (
	"encoding/json"
	"log"
	"os"

	"gitnet.fr/deblan/freetube-sync/store/file"
	"gitnet.fr/deblan/freetube-sync/web/client"
)

func ProcessHistory() bool {
	log.Print("Pull of history")
	items, err := client.PullHistory()
	res := true

	if err != nil {
		log.Print("Error while pulling history: " + err.Error())
		res = false
	} else {
		lines := []string{}

		for _, item := range items {
			line, _ := json.Marshal(item)
			lines = append(lines, string(line))
		}

		file.UpdateHistory(lines)
	}

	return res
}

func Run() {
	a := ProcessHistory()
	// b := Process("playlists", route.PlaylistPull)
	// c := Process("profiles", route.ProfilePull)

	// if a && b && c {
	if a {
		os.Exit(0)
	}

	os.Exit(1)
}
