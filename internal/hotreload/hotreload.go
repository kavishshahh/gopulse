package hotreload

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	dir     string
	watcher *fsnotify.Watcher
	cmd     *exec.Cmd
	cmdMu   sync.Mutex
}

func NewWatcher(dir string) *Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		return nil
	}

	return &Watcher{
		dir:     dir,
		watcher: watcher,
	}
}

func (w *Watcher) Close() {
	w.watcher.Close()
}

func (w *Watcher) Start() {
	err := filepath.Walk(w.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return w.watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error watching directory:", err)
		return
	}

	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("File modified:", event.Name)

				w.cmdMu.Lock()
				if w.cmd != nil && w.cmd.Process != nil {
					_ = w.cmd.Process.Signal(os.Interrupt)
					w.cmd.Process.Wait()
				}
				w.cmd = exec.Command("go", "run", "main.go")
				w.cmd.Dir = w.dir
				w.cmd.Stdout = os.Stdout
				w.cmd.Stderr = os.Stderr
				err := w.cmd.Start()
				w.cmdMu.Unlock()

				if err != nil {
					fmt.Println("Error starting process:", err)
				}
			}
		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Watcher error:", err)
		}
	}
}
