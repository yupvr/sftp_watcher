package main

import (
	"log"
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
)

// Tool to monitor shared directory for new files
// and print path to STDOUT in host machine
// Usage:
//   $ export UPLOAD_DIR=/path/to/upload/directory/host/machine
//   $ watcher /path/to/upload/directory/in/container
func main() {
	watchDir := os.Args[1]
	uploadDir := os.Getenv("UPLOAD_DIR")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println(path.Join(uploadDir, path.Base(event.Name)))
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(watchDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
