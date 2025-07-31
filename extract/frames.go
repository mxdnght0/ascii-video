package extract

import (
	"os/exec"
	"os"
)

func ExtractFrames(videoPath string) error {
	err := os.MkdirAll("frames", 0755)
    if err != nil {
        return err
    }

	cmd := exec.Command("ffmpeg", "-y", "-i", videoPath, "-vf", "fps=10", "frames/frame_%04d.png")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
} 