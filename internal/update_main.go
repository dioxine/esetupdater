package internal

import (
	"fmt"
)

func UpdateMain(sectionBlockArrayOld []SectionBlock, sectionBlockArrayNew []SectionBlock, fullDownload bool) {
	for idx, sectionBlock := range sectionBlockArrayNew {

		// Full download toggle
		if fullDownload {
			fmt.Println(sectionBlock["file"])
			DownloadFile(sectionBlock["file"])

			// Check if does not exist
		} else if !CheckNupExist(sectionBlock["file"]) {
			fmt.Println("Section: " + sectionBlock["sectionName"])
			fmt.Println("File does not exist on place. Downloading New file: " + sectionBlock["file"][1:])
			fmt.Println(sectionBlock["file"])
			DownloadFile(sectionBlock["file"])

			// Version differs check
		} else if sectionBlock["versionid"] != sectionBlockArrayOld[idx]["versionid"] {
			fmt.Println("Section: " + sectionBlock["sectionName"])
			fmt.Println("VersionID differs! Old: " + sectionBlockArrayOld[idx]["versionid"] + " New: " + sectionBlock["versionid"])
			fmt.Println("Downloading New file: " + sectionBlock["file"][1:])
			fmt.Println(sectionBlock["file"])
			DownloadFile(sectionBlock["file"])
		}
	}

}
