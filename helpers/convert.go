package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func convertAll(args []string) {

	dir, _, _ := splitPath(args[1])
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current working directory:", pwd)

	var files []string
	for _, f := range args[1:] {
		_, filename, ext := splitPath(f)
		src := filename + ext
		log.Println(src)
		if ext == ".mov" || ext == ".mp4" {
			files = append(files, src)
		}
	}

	for _, f := range files {
		convert(f)
	}
}

func convert(src string) {
	_, filename, ext := splitPath(src)
	dest := filename + ".webm"

	if ext != ".mov" && ext != ".mp4" {
		log.Printf("Error: invalid file extension (%v)\n", ext)
		return
	}

	if ext == ".mov" {
		convertMovToWebm(src, dest)
	}
	if ext == ".mp4" {
		convertMp4ToWebm(src, dest)
	}
}

func convertMovToWebm(src string, dest string) {
	log.Printf("Converting %v to %v\n", src, dest)
	cmd := exec.Command("ffmpeg", "-i", src, "-c:v", "libvpx-vp9", "-crf", "30", "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", dest)

	if err := cmd.Run(); err != nil {
		log.Fatal("ffmpeg error:", err)
	}

	if _, err := os.Stat(dest); err == nil {
		log.Println(dest, "... Done.")
	}
}

func convertMp4ToWebm(src string, dest string) {
	log.Printf("Converting %v to %v\n", src, dest)
	cmd := exec.Command("ffmpeg", "-i", src, "-c:v", "libvpx-vp9", "-crf", "10", "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", dest)

	if err := cmd.Run(); err != nil {
		log.Fatal("ffmpeg error:", err)
	}

	if _, err := os.Stat(dest); err == nil {
		log.Println(dest, "... Done.")
	}
}
