package helpers

import (
	"fmt"
	"log"
	"os"
)

func Init() {
	args := os.Args

	if len(args) < 2 || (len(args) == 2 && (args[1] == "-h" || args[1] == "--help")) {
		getHelp()
		os.Exit(1)
	}

	for _, f := range args[1:] {
		log.Println(f)
		convert(f)
	}
}

func getHelp() {
	fmt.Println("Usage:")
	fmt.Println("-----")
	fmt.Println()
	fmt.Println("convert-to-webm <filename>[.mov | .mp4]")
	fmt.Println("convert-to-webm *")
	fmt.Println()
	fmt.Println("-h | --help : prints this help")
	fmt.Println("*           : converts all [.mov | .mp4] files to .webm")
	fmt.Println()
}
