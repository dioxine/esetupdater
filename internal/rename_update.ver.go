package internal

import (
	"log"
	"os"
)

func RenameUpdateVer() error {

	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(PathOs(config.Local.RootPath, config.Remote.RootPath, "/dll/update.ver")); err == nil {
		file, err := os.Open(PathOs(config.Local.RootPath, config.Remote.RootPath, "/dll/update.ver"))
		if err != nil {
			log.Fatalln(err)
		}
		file.Close()

		os.Remove(PathOs(config.Local.RootPath, config.Remote.RootPath, "/dll/update.old"))

		// Rename update.ver to update.old
		if err = os.Rename(PathOs(config.Local.RootPath, config.Remote.RootPath, "/dll/update.ver"), PathOs(config.Local.RootPath, config.Remote.RootPath, "/dll/update.old")); err != nil {
			log.Fatalln(err)
		}

	}
	return nil
}
