package internal

import (
	"fmt"
)

func UpdateMain(sectionBlockArrayOld []SectionBlock, sectionBlockArrayNew []SectionBlock, fullDownload bool) {
	for idx, sectionBlock := range sectionBlockArrayNew {

		// Full download toggle
		if fullDownload {
			fmt.Println("Section: " + sectionBlock["sectionName"] + " full downloading")
			DownloadFile(sectionBlock["file"])
			// Check if does not exist
		} else if !CheckNupExist(sectionBlock["file"]) {
			fmt.Println("Section: " + sectionBlock["sectionName"] + " continue downloading")
			DownloadFile(sectionBlock["file"])
			// Version differs check
		} else if sectionBlock["versionid"] != sectionBlockArrayOld[idx]["versionid"] {
			fmt.Println("Section: " + sectionBlock["sectionName"] + " downloading delta")
			fmt.Println("VersionID differs! Old: " + sectionBlockArrayOld[idx]["versionid"] + " New: " + sectionBlock["versionid"])
			DownloadFile(sectionBlock["file"])
		}
	}

}
