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

	file, err := os.Open(filepath.Clean(config.Local.RootPath) + filepath.Clean(config.Remote.RootPath)[1:] + filepath.Clean("/dll/update.old"))

	if err != nil {
		if err := os.MkdirAll(filepath.Clean(config.Local.RootPath)+filepath.Clean(config.Remote.RootPath)[1:]+filepath.Clean("/dll"), 0755); err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(filepath.Clean(config.Local.RootPath) + filepath.Clean(config.Remote.RootPath)[1:] + filepath.Clean("/dll/update.old"))
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
