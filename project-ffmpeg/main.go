package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := "ffmpeg"
	//cmd := "ffmpeg -i test.mp4 -vf \"fps=5,scale=320:-1:flags=lanczos\" -c:v pam -f image2pipe - | convert -delay 5 - -loop 0 -layers optimize test.gif"
	Output, err := exec.Command("ffmpeg", "--help").Output()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("failed to execute commdand %s\n", cmd)
	}
	if Output != nil {
		str := Output
		str1 := string(str[:])
		fmt.Println(str1)
	}
}
