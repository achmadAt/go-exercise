package exif

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/ncruces/go-exiftool"
)

type ExifUtil interface {
	//SetupExifTool() (*exiftool.Server, error)
	RawOrientation(filepath string) int
	AddOrientation(filepath string) error
	fixMetaJPEG(filepath, destination string) error
}
type BaseUtil struct {
	server *exiftool.Server
}

func NewExifUtil(server *exiftool.Server) ExifUtil {
	return &BaseUtil{server: server}
}
func (b *BaseUtil) RawOrientation(filepath string) int {
	out, err := b.server.Command("--printConv", "-short3", "-fast2", "-Orientation", filepath)
	if err != nil {
		return 0
	}

	var orientation int
	_, err = fmt.Fscanf(bytes.NewReader(out), "%d\n", &orientation)
	if err != nil {
		return 0
	}

	return orientation
}
func (b *BaseUtil) AddOrientation(filepath string) error {
	command := "-Orientation=0"
	switch ori := b.RawOrientation(filepath); ori {
	case 1:
		command = "-Orientation=1"
	case 2:
		command = "-Orientation=2"
	case 3:
		command = "-Orientation=3"
	case 4:
		command = "-Orientation=4"
	case 5:
		command = "-Orientation=5"
	case 6:
		command = "-Orientation=6"
	case 7:
		command = "-Orientation=7"
	case 8:
		command = "-Orientation=8"
	}
	opts := []string{command, "-n", filepath}
	log.Print("exiftool (add orientation)...")
	_, err := b.server.Command(opts...)
	os.Remove(fmt.Sprint(filepath, "_original"))
	return err
}

func (b *BaseUtil) fixMetaJPEG(filepath, destination string) error {
	opts := []string{"-tagsFromFile", filepath, "-fixBase",
		"-CommonIFD0", "-ExifIFD:all", "-GPS:all",
		"-IPTC:all", "-XMP-dc:all", "-XMP-dc:Format=",
		"-fast", "-ignoreMinorErrors",
		"-overwrite_original", destination}

	log.Print("exiftool (fix jpeg)...")
	_, err := b.server.Command(opts...)
	return err
}
