package helpers

import (
	"fmt"
	"os"
)

func ProcessArgs() {
	args := os.Args

	if len(args) != 2 || args[1] == "-h" || args[1] == "--help" {
		getHelp()
		os.Exit(1)
	}

	if args[1] == "." {
		ConvertAll()
	}

	if len(args) == 2 {
		Convert(args[1])
	}
}

func getHelp() {
	fmt.Println("Usage:")
	fmt.Println("=====")
	fmt.Println()
	fmt.Println("convert [mov-to-mp4 | mov-to-webm | mp4-to-webm] <filename>.[mov | mp4]")
	fmt.Println("convert .")
	fmt.Println()
	fmt.Println("-h | --help : prints this help")
	fmt.Println("mov-to-mp4  : converts <filename>.mov to <filename>.mp4")
	fmt.Println("mov-to-webm : converts <filename>.mov to <filename>.webm")
	fmt.Println("mp4-to-webm : converts <filename>.mp4 to <filename>.webm")
	fmt.Println(".           : converts all files [mov | mp4] to .webm")
	fmt.Println()
}
