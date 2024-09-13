package internal

import (
	"log"
	"os"
)

func CheckNupExist(filename string) bool {

	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(PathOs(config.Local.RootPath, filename)); err == nil {
		return true
	}

	return false
}
