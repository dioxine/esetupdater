package internal

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type result bool

func DownloadFile(filenameWithRestOfPath string) (result, error) {

	// filenameWithRestOfPath can be "dll/update.ver" or full path of NUP-file
	// like "endpointwindowslatest-rel-sta/mod_002_engine_61806/em002_64_l0.dll.nup"

	//Getting config data
	config, err := ReadConfig()
	if err != nil {
		return false, err
	}

	client := http.Client{}

	var req *http.Request

	if filenameWithRestOfPath == "/dll/update.ver" {
		//Getting update.ver if true or NUP-file if false
		fmt.Println("Downloading update.ver: " + config.Remote.Url + filepath.Clean(config.Remote.RootPath) + filenameWithRestOfPath)
		req, err = http.NewRequest("GET", config.Remote.Url+filepath.Clean(config.Remote.RootPath)+filenameWithRestOfPath, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Add("User-Agent", config.Remote.UserAgent)
		if config.Remote.Username != "" && config.Remote.Password != "" {
			req.SetBasicAuth(config.Remote.Username, config.Remote.Password)
		}
	} else {
		fmt.Println("Downloading NUP-file: " + config.Remote.Url + filenameWithRestOfPath)
		req, err = http.NewRequest("GET", config.Remote.Url+filenameWithRestOfPath, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Add("User-Agent", config.Remote.UserAgent)
		if config.Remote.Username != "" && config.Remote.Password != "" {
			req.SetBasicAuth(config.Remote.Username, config.Remote.Password)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if response.StatusCode != http.StatusOK {
		return false, errors.New("Unexpected status: " + response.Status)
	}

	defer response.Body.Close()

	var out *os.File

	//create folder if not exists and write a file

	if filenameWithRestOfPath == "/dll/update.ver" {
		// for update.ver
		if err := os.MkdirAll(filepath.Dir(PathOs(config.Local.RootPath, config.Local.CustomDllPath, filenameWithRestOfPath)), 0755); err != nil {
			log.Fatalln(err)
		}

		out, err = os.Create(PathOs(config.Local.RootPath, config.Local.CustomDllPath, filenameWithRestOfPath))
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		// for NUP-file
		if err := os.MkdirAll(filepath.Dir(PathOs(config.Local.RootPath, filenameWithRestOfPath)), 0755); err != nil {
			log.Fatalln(err)
		}
		out, err = os.Create(PathOs(config.Local.RootPath, filenameWithRestOfPath))
		if err != nil {
			log.Fatalln(err)
		}
	}
	_, err = io.Copy(out, response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	out.Close()

	return true, nil
}
