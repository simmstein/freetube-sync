package file

import (
	"encoding/json"

	config "gitnet.fr/deblan/freetube-sync/config/client"
	"gitnet.fr/deblan/freetube-sync/file"
	"gitnet.fr/deblan/freetube-sync/model"
)

func LoadHistory() []model.WatchedVideo {
	lines := file.GetLines(config.GetConfig().DbPath("history"))
	collection := []model.WatchedVideo{}

	for _, line := range lines {
		var item model.WatchedVideo
		json.Unmarshal([]byte(line), &item)

		collection = append(collection, item)
	}

	return collection
}

func UpdateHistory(data []string) {
	file.WriteDatabase(config.GetConfig().DbPath("history"), data)
}
