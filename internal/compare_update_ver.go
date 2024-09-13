package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CompareUpdateVer() ([]SectionBlock, []SectionBlock, error) {

	// Read old update.ver
	file, err := OpenUpdateVer("/dll/update.old")
	if err != nil {
		return nil, nil, err
	}
	defer CloseUpdateVer(file)
	scanner := bufio.NewScanner(file)
	sectionBlockArrayOld := ConvertToArray(scanner)

	// Read new update.ver
	file, err = OpenUpdateVer("/dll/update.ver")
	if err != nil {
		return nil, nil, err
	}
	defer CloseUpdateVer(file)
	scanner = bufio.NewScanner(file)
	sectionBlockArrayNew := ConvertToArray(scanner)

	// Print length of old and new arrays

	fmt.Printf("Length of old array: %d\n", len(sectionBlockArrayOld))
	fmt.Printf("Length of new array: %d\n", len(sectionBlockArrayNew))

	return sectionBlockArrayOld, sectionBlockArrayNew, nil
}

func OpenUpdateVer(filenameWithRestOfPath string) (*os.File, error) {
	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(filepath.Join(filepath.Clean(config.Local.RootPath), filepath.Clean(config.Remote.RootPath), filepath.Clean(filenameWithRestOfPath)))
	if err != nil {
		log.Fatalln(err)
	}
	return file, nil
}

func CloseUpdateVer(file *os.File) {
	file.Close()
}
