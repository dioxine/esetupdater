package internal

import (
	"log"
	"os"
	"path/filepath"
)

func CheckNupExist(filename string) bool {

	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(filepath.Join(filepath.Clean(config.Local.RootPath), filename)); err == nil {
		return true
	}

	return false
}
