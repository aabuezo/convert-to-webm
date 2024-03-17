package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
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
	wg := sync.WaitGroup{}

	for _, f := range args[1:] {
		_, filename, ext := splitPath(f)
		src := filename + ext
		dest := filename + ".webm"

		if ext == ".mov" {
			wg.Add(1)
			go convertMovToWebm(src, dest, &wg)
		}
		if ext == ".mp4" {
			wg.Add(1)
			go convertMp4ToWebm(src, dest, &wg)
		}
	}
	wg.Wait()
}

func convertMovToWebm(src string, dest string, wg *sync.WaitGroup) {
	log.Printf("Converting %v to %v\n", src, dest)
	cmd := exec.Command("ffmpeg", "-i", src, "-c:v", "libvpx-vp9", "-crf", "30", "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", dest)

	if err := cmd.Run(); err != nil {
		log.Fatal("ffmpeg error:", err)
	}

	if _, err := os.Stat(dest); err == nil {
		log.Println(dest, "... Done.")
	}
	wg.Done()
}

func convertMp4ToWebm(src string, dest string, wg *sync.WaitGroup) {
	log.Printf("Converting %v to %v\n", src, dest)
	cmd := exec.Command("ffmpeg", "-i", src, "-c:v", "libvpx-vp9", "-crf", "10", "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", dest)

	if err := cmd.Run(); err != nil {
		log.Fatal("ffmpeg error:", err)
	}

	if _, err := os.Stat(dest); err == nil {
		log.Println(dest, "... Done.")
	}
	wg.Done()
}
