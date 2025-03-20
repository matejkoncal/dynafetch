package watch

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

type EventType int

const (
	Write EventType = iota
	Rename
)

type FileEvent struct {
	EventType EventType
	FilePath  string
}

func WatchFile(filePath string, channel chan<- FileEvent) {

	dir := filepath.Dir(filePath)

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	err = watcher.Add(dir)

	var debounceTimer *time.Timer

	for {
		select {
		case event, _ := <-watcher.Events:

			if event.Op&fsnotify.Rename != fsnotify.Rename && event.Op&fsnotify.Write != fsnotify.Write {
				continue
			}

			if debounceTimer != nil {
				debounceTimer.Stop()
			}

			debounceTimer = time.AfterFunc(1000*time.Millisecond, func() {
				if event.Name == filePath {
					if event.Op&fsnotify.Rename == fsnotify.Rename {
						channel <- FileEvent{EventType: Rename, FilePath: event.Name}
					} else {
						channel <- FileEvent{EventType: Write, FilePath: event.Name}
					}
				}
			})

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Chyba:", err)
		}
	}
}
