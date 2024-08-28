package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	config "gitnet.fr/deblan/freetube-sync/config/client"
	"gitnet.fr/deblan/freetube-sync/model"
	"gitnet.fr/deblan/freetube-sync/web/route"
)

type Data any

type PostResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Request(method, route string, data Data) ([]byte, error) {
	value, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", config.GetConfig().Server, route)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(value))
	request.Header.Set("X-Machine", config.GetConfig().Hostname)
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func Post(route string, data Data) ([]byte, error) {
	return Request("POST", route, data)
}

func Get(route string) ([]byte, error) {
	return Request("GET", route, nil)
}

func InitPush(route string, data Data) (PostResponse, error) {
	var value PostResponse

	body, err := Post(route, data)
	json.Unmarshal(body, &value)

	return value, err
}

func PullHistory() ([]model.WatchedVideo, error) {
	var items []model.WatchedVideo

	body, err := Get(route.HistoryPull)
	json.Unmarshal(body, &items)

	return items, err
}
