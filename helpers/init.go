package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"time"
)

func Init() {
	args := os.Args

	if len(args) < 2 || args[1] == "-h" || args[1] == "--help" {
		getHelp()
		os.Exit(1)
	}

	// check if ffmpeg is already installed
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		fmt.Println("Error: ffmpeg is not present.")
		fmt.Println("Run 'brew install ffmpeg' to install on MacOS.")
		os.Exit(1)
	}

	// enable concurrency only if requested - default: false
	if args[1] == "-c" {
		concurrent = true
		args = slices.Delete(args, 1, 2)
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
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("        convert-to-webm [-h | --help]")
	fmt.Println("        convert-to-webm [-c] <filename1>[.mov|.mp4] <filename2>[.mov|.mp4] ...")
	fmt.Println("        convert-to-webm [-c] <path/to/>*")
	fmt.Println()
	fmt.Println("  -h | --help : prints this help")
	fmt.Println("  -c          : convert files concurrently (faster, but VERY resource intensive!!)")
	fmt.Println("   *          : converts all [.mov|.mp4] files to .webm")
	fmt.Println()
	fmt.Println()
}
