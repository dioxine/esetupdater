package internal

import (
	"log"
	"os"
	"path/filepath"
)

func RenameUpdateVer() error {

	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(filepath.Clean(config.Local.RootPath) + filepath.Clean(config.Remote.RootPath)[1:] + filepath.Clean("/dll/update.ver")); err == nil {
		file, err := os.Open(filepath.Clean(config.Local.RootPath) + filepath.Clean(config.Remote.RootPath)[1:] + filepath.Clean("/dll/update.ver"))
		if err != nil {
			log.Fatalln(err)
		}

		// Check file size > 0
		if info, err := file.Stat(); err == nil && info.Size() > 0 {
			// Rename update.ver to update.old
			os.Remove(filepath.Clean(config.Local.RootPath) + filepath.Clean(config.Remote.RootPath)[1:] + filepath.Clean("/dll/update.old"))

			if err = os.Rename(filepath.Clean(config.Local.RootPath)+filepath.Clean(config.Remote.RootPath)[1:]+filepath.Clean("/dll/update.ver"), filepath.Clean(config.Local.RootPath)+filepath.Clean(config.Remote.RootPath)[1:]+filepath.Clean("/dll/update.old")); err != nil {
				file.Close()
				log.Fatalln(err)
			}
		}
	}
	return nil
}
