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

	if _, err := os.Stat(filepath.Join(filepath.Clean(config.Local.RootPath), filepath.Clean(config.Remote.RootPath), filepath.Clean("/dll/update.ver"))); err == nil {
		file, err := os.Open(filepath.Join(filepath.Clean(config.Local.RootPath), filepath.Clean(config.Remote.RootPath), filepath.Clean("/dll/update.ver")))
		if err != nil {
			log.Fatalln(err)
		}
		file.Close()

		os.Remove(filepath.Join(filepath.Clean(config.Local.RootPath), filepath.Clean(config.Remote.RootPath), filepath.Clean("/dll/update.old")))

		// Rename update.ver to update.old
		if err = os.Rename(filepath.Join(filepath.Clean(config.Local.RootPath), filepath.Clean(config.Remote.RootPath), filepath.Clean("/dll/update.ver")), filepath.Join(filepath.Clean(config.Local.RootPath), filepath.Clean(config.Remote.RootPath), filepath.Clean("/dll/update.old"))); err != nil {
			log.Fatalln(err)
		}

	}
	return nil
}
