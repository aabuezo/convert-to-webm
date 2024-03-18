package helpers

import (
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

	wg := sync.WaitGroup{}

	for _, f := range args[1:] {
		_, filename, ext := splitPath(f)
		src := filename + ext
		dest := filename + ".webm"

		if _, err := os.Stat(src); err != nil {
			log.Printf("File %v does not exist.", src)
		}

		if ext == ".mp4" || ext == ".mov" {
			wg.Add(1)
			go convert(src, ext, dest, &wg)
		}
	}

	wg.Wait()
}

func convert(src string, ext string, dest string, wg *sync.WaitGroup) {
	log.Printf("Converting %v to %v\n", src, dest)

	var cmd *exec.Cmd
	switch ext {
	case ".mp4":
		cmd = exec.Command("ffmpeg", "-i", src, "-c:v", "libvpx-vp9", "-crf", "10", "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", dest)
	case ".mov":
		cmd = exec.Command("ffmpeg", "-i", src, "-c:v", "libvpx-vp9", "-crf", "30", "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", dest)
	}

	if err := cmd.Run(); err != nil {
		log.Fatal("Error processing ffmpeg:", err)
	}

	if _, err := os.Stat(dest); err == nil {
		log.Println(dest, "Done.")
	}

	wg.Done()
}
