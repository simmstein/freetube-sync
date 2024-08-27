package main

import (
	"flag"
	"fmt"
	"os"

	"gitnet.fr/deblan/freetube-sync/cmd/action/initcmd"
	"gitnet.fr/deblan/freetube-sync/cmd/action/watchcmd"
	config "gitnet.fr/deblan/freetube-sync/config/client"
)

func main() {
	config.InitConfig()
	action := flag.Arg(0)

	switch action {
	case "init":
		initcmd.Run()
	case "watch":
		watchcmd.Run()
	default:
		fmt.Print("You must pass a sub-command: init, watch")
		os.Exit(1)
	}

	// lines := file.GetLines("/home/simon/.config/FreeTube/history.db")
	// collection := []model.Video{}
	//
	// for _, line := range lines {
	// 	var item model.Video
	// 	json.Unmarshal([]byte(line), &item)
	//
	// 	collection = append(collection, item)
	// }
	//
	// data, err := json.Marshal(collection)
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// req, err := http.NewRequest("POST", "http://localhost:1323/history/push", bytes.NewBuffer(data))
	// req.Header.Set("X-Machine", "endurance")
	// req.Header.Set("Content-Type", "application/json")
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	//
	// fmt.Printf("%+v\n", data)
	// fmt.Printf("%+v\n", collection)
}
