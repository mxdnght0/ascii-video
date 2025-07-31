package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"ascii-video/convert"
	"ascii-video/extract"
	"ascii-video/output"
)

func loadPNG(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func naturalSort(files []os.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		extractNum := func(name string) int {
			base := strings.TrimSuffix(name, filepath.Ext(name))
			parts := strings.Split(base, "_")
			if len(parts) < 2 {
				return 0
			}
			n, _ := strconv.Atoi(parts[1])
			return n
		}
		return extractNum(files[i].Name()) < extractNum(files[j].Name())
	})
}

func main() {
	videoPath := os.Args[1]
	err := extract.ExtractFrames(videoPath)
	if err != nil {
		fmt.Println("Ошибка при извлечении кадров:", err)
		return
	}

	files, err := os.ReadDir("frames")
	if err != nil {
		fmt.Println("Ошибка чтения папки frames:", err)
		return
	}

	naturalSort(files)

	var asciiFrames [][]string
	const asciiWidth = 160
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join("frames", file.Name())
		img, err := loadPNG(path)
		if err != nil {
			fmt.Println("Ошибка загрузки изображения:", path, err)
			continue
		}

		asciiFrame := ascii.ImageToASCII(img, asciiWidth)
		asciiFrames = append(asciiFrames, asciiFrame)
	}

	if len(asciiFrames) == 0 {
		fmt.Println("Нет кадров для отображения")
		return
	}

	output.PrintFrames(asciiFrames, 10)
}
