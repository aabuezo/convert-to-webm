package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Init() {
	args := os.Args

	if len(args) < 2 || (len(args) == 2 && (args[1] == "-h" || args[1] == "--help")) {
		getHelp()
		os.Exit(1)
	}

	if !isFfmpegPresent() {
		log.Println("Error: ffmpeg is not present. Please install ffmpeg first.")
		log.Println("Run 'brew install ffmpeg' on MacOS.")
		os.Exit(1)
	}

	convertAll(args)
}

func getHelp() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("convert-to-webm <path/to/><filename>[.mov | .mp4]")
	fmt.Println("convert-to-webm <filename>[.mov | .mp4]")
	fmt.Println("convert-to-webm <path/to/>*")
	fmt.Println("convert-to-webm *")
	fmt.Println()
	fmt.Println("-h | --help : prints this help")
	fmt.Println("*           : converts all [.mov | .mp4] files to .webm")
	fmt.Println()
}

func isFfmpegPresent() bool {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return false
	} else {
		return true
	}
}
