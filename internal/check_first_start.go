package internal

import (
	"fmt"
	"log"
	"os"
)

func CheckIfFirstStart() bool {

	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(PathOs(config.Local.RootPath, config.Remote.RootPath, "/dll/update.old"))

	if err != nil {
		if err := os.MkdirAll(PathOs(config.Local.RootPath, config.Remote.RootPath, "/dll"), 0755); err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(PathOs(config.Local.RootPath, config.Remote.RootPath, "/dll/update.old"))
		if err != nil {
			fmt.Println(err)
			// log.Fatalln(err)
		}
		file.Close()
		return true
	}

	file.Close()
	return false
}
