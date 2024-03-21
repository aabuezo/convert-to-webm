package converter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"time"
)

func Run() {
	args := os.Args

	// check if ffmpeg is already installed
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		fmt.Println("Error: ffmpeg is not present.")
		fmt.Println("Run 'brew install ffmpeg' to install on MacOS.")
		os.Exit(1)
	}

	// enable concurrency only if requested - default: false
	if slices.Contains(args, "-c") {
		concurrent = true
		idx := slices.Index(args, "-c")
		args = slices.Delete(args, idx, idx+1)
	}

	if len(args) < 2 || slices.Contains(args, "-h") || slices.Contains(args, "--help") {
		getHelp()
		os.Exit(1)
	}

	// meassure execution time
	start := time.Now()
	fmt.Println()

	// calling main convertion routine
	convertAll(args)

	elapsed := time.Since(start)
	// print elapsed time to process all files
	fmt.Println()
	log.Printf("Took: %s\n", elapsed)
	fmt.Println()
}

func getHelp() {
	fmt.Println(`
Usage:
		convert-to-webm [-h | --help]
		convert-to-webm [-c] <filename1>[.mov|.mp4] <filename2>[.mov|.mp4] ...
		convert-to-webm [-c] <path/to/>*

  -h | --help : prints this help
  -c          : convert files concurrently (faster, but VERY resource intensive!!)
   *          : converts all [.mov|.mp4] files to .webm`)
	fmt.Println()
}
