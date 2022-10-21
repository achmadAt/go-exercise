package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ncruces/go-exiftool"
)

var exifserver *exiftool.Server

func setupExifTool() (s *exiftool.Server, err error) {
	exifserver, err = exiftool.NewServer("-ignoreMinorErrors")
	return exifserver, err
}

func getMetaHTML(path string) ([]byte, error) {
	log.Print("exiftool (get meta)...")
	return exifserver.Command("-htmlFormat", "-groupHeadings", "-long", "-fixBase", path)
}

func fixMetaDNG(orig, dest, name string) error {
	opts := []string{"-tagsFromFile", orig, "-fixBase",
		"-MakerNotes", "-OriginalRawFileName-=" + filepath.Base(orig)}
	if name != "" {
		opts = append(opts, "-OriginalRawFileName="+filepath.Base(name))
	}
	opts = append(opts, "-overwrite_original", dest)

	log.Print("exiftool (fix dng)...")
	_, err := exifserver.Command(opts...)
	return err
}
func checkOrientation(file string) (int, error) {
	opts := []string{"-Orientation", "-n", file}
	out, err := exifserver.Command(opts...)
	if err != nil {
		return 0, err
	}
	var orientation int
	a, err := fmt.Fscanf(bytes.NewReader(out), "%d\n", &orientation)
	if err != nil {
		return 0, err
	}
	fmt.Println("a", a)
	return orientation, nil
}
func fixMetaJPEG(orig, dest string) error {
	opts := []string{"-tagsFromFile", orig, "-fixBase",
		"-CommonIFD0", "-ExifIFD:all", "-GPS:all", // https://exiftool.org/forum/index.php?topic=8378.msg43043#msg43043
		"-IPTC:all", "-XMP-dc:all", "-XMP-dc:Format=",
		"-fast", "-ignoreMinorErrors",
		"-overwrite_original", dest}

	log.Print("exiftool (fix jpeg)...")
	_, err := exifserver.Command(opts...)
	return err
}

func addOrientation(file string) error {
	command := "-Orientation=0"
	switch ori := rawOrientation(file); ori {
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
	opts := []string{command, "-n", file}
	log.Print("exiftool (add orientation)...")
	_, err := exifserver.Command(opts...)
	os.Remove(fmt.Sprint(file, "_original"))
	return err
}

func dngHasEdits(path string) bool {
	log.Print("exiftool (has edits?)...")
	out, err := exifserver.Command("-XMP-photoshop:all", path)
	return err == nil && len(out) > 0
}

func rawOrientation(path string) int {
	log.Print("exiftool (get orientation)...")
	out, err := exifserver.Command("--printConv", "-short3", "-fast2", "-Orientation", path)
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

func cameraMatchingWhiteBalance(path string) string {
	log.Print("exiftool (get camera matching white balance)...")
	out, err := exifserver.Command("-duplicates", "-short3", "-fast", "-ExifIFD:WhiteBalance", "-MakerNotes:WhiteBalance", path)
	if err != nil {
		return ""
	}

	for scan := bufio.NewScanner(bytes.NewReader(out)); scan.Scan(); {
		switch wb := scan.Text(); wb {
		case "Auto", "Daylight", "Cloudy", "Shade", "Tungsten", "Fluorescent", "Flash":
			return wb
		case "Sunny":
			return "Daylight"
		case "Overcast":
			return "Cloudy"
		case "Incandescent":
			return "Tungsten"
		}
	}
	return "As Shot"
}
