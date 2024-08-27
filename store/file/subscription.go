package file

import (
	"encoding/json"

	config "gitnet.fr/deblan/freetube-sync/config/client"
	"gitnet.fr/deblan/freetube-sync/file"
	"gitnet.fr/deblan/freetube-sync/model"
)

func LoadProfiles() []model.Profile {
	lines := file.GetLines(config.GetConfig().Path + "/profiles.db")
	collection := []model.Profile{}
	added := make(map[string]bool)

	for i := len(lines) - 1; i >= 0; i-- {
		var item model.Profile
		json.Unmarshal([]byte(lines[i]), &item)

		if !added[item.Name] {
			added[item.Name] = true
			collection = append(collection, item)
		}
	}

	return collection
}
