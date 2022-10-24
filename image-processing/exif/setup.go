package exif

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ncruces/go-exiftool"
)

var (
	ServerMode                bool
	BaseDir, DataDir, TempDir string
	DngConverter              string
)

func SetupPaths() (err error) {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	if exe, err := filepath.EvalSymlinks(exe); err != nil {
		return err
	} else {
		BaseDir = filepath.Dir(exe)
	}

	DataDir = filepath.Join(BaseDir, "data")
	TempDir = filepath.Join(os.TempDir(), "somethink")

	switch runtime.GOOS {
	case "windows":
		exiftool.Exec = "/home/alfazari/go-exercise/image-processing/excutable/exiftool_windows/exiftool/exiftool.exe"
		exiftool.Arg1 = strings.TrimSuffix(exiftool.Exec, ".exe")
	case "linux":
		exiftool.Exec = "/home/alfazari/go-exercise/image-processing/excutable/unix/exiftool/exiftool"
		//exiftool.Config = BaseDir + "/utils/exiftool_config.pl"
	}

	if testDataDir() == nil {
		return nil
	}
	if data, err := os.UserConfigDir(); err != nil {
		return err
	} else {
		DataDir = filepath.Join(data, "somethink")
	}
	return testDataDir()
}

func testDataDir() error {
	if err := os.MkdirAll(DataDir, 0700); err != nil {
		return err
	}
	if f, err := os.Create(filepath.Join(DataDir, "lastrun")); err != nil {
		return err
	} else {
		return f.Close()
	}
}
