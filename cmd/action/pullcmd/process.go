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

func ProcessPlaylists() bool {
	log.Print("Pull of playlists")
	items, err := client.PullPlaylists()
	res := true
	_ = items

	if err != nil {
		log.Print("Error while pulling playlists: " + err.Error())
		res = false
	} else {
		lines := []string{}

		for _, item := range items {
			line, _ := json.Marshal(item)
			lines = append(lines, string(line))
		}

		file.UpdatePlaylists(lines)
	}

	return res
}

func ProcessProfiles() bool {
	log.Print("Pull of profiles")
	items, err := client.PullProfiles()
	res := true
	_ = items

	if err != nil {
		log.Print("Error while pulling profiles: " + err.Error())
		res = false
	} else {
		lines := []string{}

		for _, item := range items {
			line, _ := json.Marshal(item)
			lines = append(lines, string(line))
		}

		file.UpdateProfiles(lines)
	}

	return res
}

func Run() {
	a := ProcessHistory()
	b := ProcessPlaylists()
	c := ProcessProfiles()

	if a && b && c {
		os.Exit(0)
	}

	os.Exit(1)
}
