package internal

import (
	"fmt"
	"os"
)

func PrintConfig() error {
	b, err := os.ReadFile("config.toml")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
