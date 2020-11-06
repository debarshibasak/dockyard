package documentmanager

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func (d *DocumentManager) Watch() {
	fmt.Println("starting to watch...")
	watcher, _ = fsnotify.NewWatcher()

	defer watcher.Close()

	// starting at the root of the project, walk each file/directory searching for
	// directories
	if err := filepath.Walk(d.config.Location, watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	done := make(chan bool)
	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)
				fmt.Println("generating new events")
				d.Generate()
				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}
func watchDir(path string, fi os.FileInfo, err error) error {

	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}
