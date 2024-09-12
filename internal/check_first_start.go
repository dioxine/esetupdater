package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CheckIfFirstStart() bool {

	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Open(filepath.Clean(config.Local.RootPath) + filepath.Clean(config.Remote.RootPath) + "/dll/update.old"); err != nil {

		if err := os.MkdirAll(filepath.Clean(config.Local.RootPath)+filepath.Clean(config.Remote.RootPath)+"/dll", 0755); err != nil {
			log.Fatal(err)
		}

		out, err := os.Create(filepath.Clean(config.Local.RootPath) + filepath.Clean(config.Remote.RootPath) + "/dll/update.old")
		if err != nil {
			fmt.Println(err)
			// log.Fatalln(err)
		}
		out.Close()
		return true
	}
	return false
}
