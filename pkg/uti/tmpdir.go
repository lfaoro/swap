package uti

import (
	"fmt"
	"os"
	"path/filepath"
)

func MakeTempDir(name string) (path string) {
	var tempDir string
	switch os.Getenv("GOOS") {
	case "windows":
		tempDir = os.Getenv("TEMP")
	case "darwin", "linux":
		tempDir = "/tmp"
	default:
		fmt.Println("unsupported operating system")
		return
	}

	trocaDir := filepath.Join(tempDir, "troca")

	err := os.Mkdir(trocaDir, 0755)
	if err != nil {
		fmt.Printf("error creating directory: %v\n", err)
		return
	}
	return trocaDir
}
