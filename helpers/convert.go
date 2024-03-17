package helpers

import (
	"log"
	"os"
	"os/exec"
)

func convert(from string) {
	dir, filename, ext := splitPath(from)
	to := dir + filename + ".webm"

	if ext != ".mov" && ext != ".mp4" {
		log.Printf("Error: invalid file extension (%v)\n", ext)
		return
	}

	if ext == ".mov" {
		convertMovToWebm(from, to)
	}
	if ext == ".mp4" {
		convertMp4ToWebm(from, to)
	}
}

func convertMovToWebm(from string, to string) {
	log.Printf("Converting %v to %v\n", from, to)
	cmd := exec.Command("ffmpeg -i " + from + " libvpx-vp9 -crf 30 -b:v 0 -b:a 128k -c:a libopus " + to)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(to); err == nil {
		log.Println(to, "... Done.")
	}
}

func convertMp4ToWebm(from string, to string) {
	log.Printf("Converting %v to %v\n", from, to)
	cmd := exec.Command("ffmpeg -i " + from + " -vcodec copy " + to)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(to); err == nil {
		log.Println(to, "... Done.")
	}
}
