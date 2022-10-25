package exif

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/ncruces/go-exiftool"
)

func SetupPaths() (err error) {
	BaseDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(BaseDir)

	switch runtime.GOOS {
	case "windows":
		exiftool.Exec = BaseDir + "/excutable/exiftool_windows/exiftool/exiftool.exe"
		exiftool.Arg1 = strings.TrimSuffix(exiftool.Exec, ".exe")
	case "darwin", "linux":
		exiftool.Exec = BaseDir + "/excutable/unix/exiftool/exiftool"
	}

	return nil
}
