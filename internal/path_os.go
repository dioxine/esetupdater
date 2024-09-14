package internal

import (
	"path/filepath"
)

func PathOs(paths ...string) string {
	for path := range paths {
		paths[path] = filepath.Clean(paths[path])
	}
	return filepath.Join(paths...)
}
