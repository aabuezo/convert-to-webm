package converter

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var concurrent bool = false

func convertAll(args []string) {

	// set working directory
	dir, _, _ := splitPath(args[1])
	if dir != "" {
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}
	}

	// create waitgroups to run concurrently
	wg := sync.WaitGroup{}

	// iterate over the list of files in CWD
	for _, f := range args[1:] {
		_, filename, ext := splitPath(f)
		src := filename + ext
		dest := filename + ".webm"

		if !isFileExists(src) {
			log.Printf("File %v does not exist.", src)
			continue
		}

		// if .webm already exists, then skip
		if isFileExists(dest) {
			continue
		}

		// convert .mp4 or .mov only
		ext = strings.ToLower(ext)
		if ext == ".mp4" || ext == ".mov" {
			// create new thread when concurrency is on
			if concurrent {
				wg.Add(1)
				go convert(src, ext, dest, &wg)
			} else {
				convert(src, ext, dest, &wg)
			}
		}
	}

	wg.Wait()
}

func convert(src string, ext string, dest string, wg *sync.WaitGroup) {
	log.Printf("Converting %v to %v\n", src, dest)

	// call ffmpeg based on src file extension
	var cmd *exec.Cmd
	switch ext {
	// add more cases accordingly
	case ".mp4":
		cmd = exec.Command("ffmpeg", "-i", src, "-c:v", "libvpx-vp9", "-crf", "10", "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", dest)
	case ".mov":
		cmd = exec.Command("ffmpeg", "-i", src, "-c:v", "libvpx-vp9", "-crf", "30", "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", dest)
	}

	// check for failures on ffmpeg
	if err := cmd.Run(); err != nil {
		log.Fatal("Error processing ffmpeg:", err)
		// do not Exit, just keep processing next file if any
	}

	// new file .webm got created successfully
	if isFileExists(dest) {
		log.Println(dest, "Done.")
	}

	// close thread
	if concurrent {
		wg.Done()
	}
}
