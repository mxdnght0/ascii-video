package output

import (
	"fmt"
	"time"
)

func PrintFrames(frames [][]string, fps int) {
	delay := time.Second / time.Duration(fps)

	for _, frame := range frames {
		fmt.Print("\033[H\033[2J")
		for _, line := range frame {
			fmt.Println(line)
		}
		time.Sleep(delay)
	}
}