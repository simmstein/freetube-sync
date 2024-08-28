package file

import (
	"encoding/json"

	config "gitnet.fr/deblan/freetube-sync/config/client"
	"gitnet.fr/deblan/freetube-sync/file"
	"gitnet.fr/deblan/freetube-sync/model"
)

func LoadProfiles() []model.Profile {
	lines := file.GetLines(config.GetConfig().DbPath("profiles"))
	collection := []model.Profile{}
	added := make(map[string]bool)

	for i := len(lines) - 1; i >= 0; i-- {
		var item model.Profile
		json.Unmarshal([]byte(lines[i]), &item)

		if !added[item.RemoteId] {
			added[item.RemoteId] = true
			collection = append(collection, item)
		}
	}

	return collection
}

func UpdateProfiles(data []string) {
	file.WriteDatabase(config.GetConfig().DbPath("profiles"), data)
}
