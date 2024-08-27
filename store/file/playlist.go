package file

import (
	"encoding/json"

	config "gitnet.fr/deblan/freetube-sync/config/client"
	"gitnet.fr/deblan/freetube-sync/file"
	"gitnet.fr/deblan/freetube-sync/model"
)

func LoadPlaylists() []model.Playlist {
	lines := file.GetLines(config.GetConfig().Path + "/playlists.db")
	collection := []model.Playlist{}
	added := make(map[string]bool)

	for i := len(lines) - 1; i >= 0; i-- {
		var item model.Playlist
		json.Unmarshal([]byte(lines[i]), &item)

		if !added[item.Id] {
			added[item.Id] = true
			collection = append(collection, item)
		}
	}

	return collection
}
