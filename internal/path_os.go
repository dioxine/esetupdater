package internal

import (
	"path/filepath"
)

func PathOs(paths ...string) string {
	for el := range paths {
		paths[el] = filepath.Clean(paths[el])
	}
	return filepath.Join(paths...)
}
