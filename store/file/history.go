package file

import (
	"encoding/json"

	config "gitnet.fr/deblan/freetube-sync/config/client"
	"gitnet.fr/deblan/freetube-sync/file"
	"gitnet.fr/deblan/freetube-sync/model"
)

func LoadHistory() []model.Video {
	lines := file.GetLines(config.GetConfig().Path + "/history.db")
	collection := []model.Video{}

	for _, line := range lines {
		var item model.Video
		json.Unmarshal([]byte(line), &item)

		collection = append(collection, item)
	}

	return collection
}
