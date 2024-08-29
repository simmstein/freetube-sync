package main

import (
	"flag"
	"fmt"
	"os"

	"gitnet.fr/deblan/freetube-sync/cmd/action/initcmd"
	"gitnet.fr/deblan/freetube-sync/cmd/action/pullcmd"
	"gitnet.fr/deblan/freetube-sync/cmd/action/watchcmd"
	config "gitnet.fr/deblan/freetube-sync/config/client"
)

func main() {
	config.InitConfig()

	switch flag.Arg(0) {
	case "init":
		initcmd.Run()
	case "watch":
		watchcmd.Run()
	case "pull":
		pullcmd.Run()
	default:
		fmt.Print("You must pass a sub-command: init, watch, pull")
		os.Exit(1)
	}
}
