package internal

import (
	"bufio"
	"log"
	"slices"
	"strings"
)

type SectionBlock map[string]string

func ConvertToArray(scanner *bufio.Scanner) []SectionBlock {
	var sectionBlockArray []SectionBlock
	var sectionBlock SectionBlock

	config, err := ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	filter := config.Remote.Filter

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '[' {
			if len(sectionBlock) > 0 && !slices.Contains(filter, sectionBlock["platform"]) {
				sectionBlockArray = append(sectionBlockArray, sectionBlock)
			}
			sectionBlock = make(SectionBlock)
			sectionBlock["sectionName"] = line[1 : len(line)-1]
		} else if line[0] != ';' {
			field := strings.Split(line, "=")
			sectionBlock[field[0]] = field[1]
		}
	}

	if !slices.Contains(filter, sectionBlock["platform"]) {
		sectionBlockArray = append(sectionBlockArray, sectionBlock)
	}

	return sectionBlockArray
}
