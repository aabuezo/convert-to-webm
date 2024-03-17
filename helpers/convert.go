package helpers

import (
	"log"
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
}

func convertMp4ToWebm(from string, to string) {
	log.Printf("Converting %v to %v\n", from, to)
}
