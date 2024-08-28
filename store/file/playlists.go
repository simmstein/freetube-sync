package file

import (
	"encoding/json"

	config "gitnet.fr/deblan/freetube-sync/config/client"
	"gitnet.fr/deblan/freetube-sync/file"
	"gitnet.fr/deblan/freetube-sync/model"
)

func LoadPlaylists() []model.Playlist {
	lines := file.GetLines(config.GetConfig().DbPath("playlists"))
	collection := []model.Playlist{}
	added := make(map[string]bool)

	for i := len(lines) - 1; i >= 0; i-- {
		var item model.Playlist
		json.Unmarshal([]byte(lines[i]), &item)

		if !added[item.RemoteId] {
			added[item.RemoteId] = true
			collection = append(collection, item)
		}
	}

	return collection
}

func UpdatePlaylists(data []string) {
	file.WriteDatabase(config.GetConfig().DbPath("playlists"), data)
}
