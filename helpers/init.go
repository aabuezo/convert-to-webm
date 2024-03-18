package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func Init() {
	args := os.Args

	if len(args) < 2 || (len(args) == 2 && (args[1] == "-h" || args[1] == "--help")) {
		getHelp()
		os.Exit(1)
	}

	// check if ffmpeg is already installed
	if !isFfmpegPresent() {
		log.Println("Error: ffmpeg is not present. Please install ffmpeg first.")
		log.Println("Run 'brew install ffmpeg' on MacOS.")
		os.Exit(1)
	}

	// check if file exists
	if _, err := os.Stat(args[1]); err != nil {
		log.Printf("File %v does not exist.", args[1])
		os.Exit(1)
	}

	start := time.Now()
	// calling main convertion routine
	convertAll(args)
	elapsed := time.Since(start)
	// print elapsed time to process all files
	log.Printf("Took: %s\n", elapsed)
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
