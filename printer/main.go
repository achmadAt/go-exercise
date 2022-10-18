package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	platform := runtime.GOOS
	command := "bash"
	argument := `lpstat -p | awk '{print $2}'`
	print := "lpr"
	if platform == "window" {
		command = "cmd"
		argument = `wmic printer list brief`
		print = "print"
	}
	cmd := exec.Command(command, "-c", argument)
	_output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := strings.TrimSpace(string(_output))
	_printers := strings.Split(output, "\n")
	var printers []string
	for _, printer := range _printers {
		if printer != "to" {
			printers = append(printers, printer)
		}
	}
	check_file, err := filepath.Abs("/home/alfazari/Pictures/Wallpapers/download1.jpeg")
	if err != nil {
		log.Fatal(err)
		return
	}
	pic := exec.Command(print, "-o", "media:5", "-o", "orientation-requested=4", "-o", "print-quality=5", check_file)
	_printed, error := pic.CombinedOutput()
	if error != nil {
		log.Fatal(error)
	}
	printed := string(_printed[:])
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(printed)
	fmt.Println(dir)
	fmt.Println(pic.Path)
	fmt.Println(check_file)
	fmt.Println(printers)
}
