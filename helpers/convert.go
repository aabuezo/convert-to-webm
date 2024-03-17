package helpers

import (
	"log"
)

func ConvertAll() {
	log.Println("Processing all files in current directory...")
}

func Convert(from string) {
	dir, filename, srcExt := SplitPath(from)
	to := dir + filename + ".webm"

	if srcExt != ".mov" && srcExt != ".mp4" {
		log.Printf("Error: invalid file extension (%v)\n", srcExt)
		return
	}

	if srcExt == ".mov" {
		convertMovToWebm(from, to)
	}
	if srcExt == ".mp4" {
		convertMp4ToWebm(from, to)
	}
}

func convertMovToWebm(from string, to string) {
	log.Printf("Converting %v to %v\n", from, to)
}

func convertMp4ToWebm(from string, to string) {
	log.Printf("Converting %v to %v\n", from, to)
}
