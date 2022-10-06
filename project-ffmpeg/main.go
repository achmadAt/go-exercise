package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := "ffmpeg -i test.mp4 -vf \"fps=5,scale=320:-1:flags=lanczos\" -c:v pam -f image2pipe - | convert -delay 5 - -loop 0 -layers optimize test.gif"
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("failed to execute commdand %s", cmd)
	}
}
