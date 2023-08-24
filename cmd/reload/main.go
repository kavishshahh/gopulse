package main

import (
	"fmt"

	"kavishshahh.com/goreload/internal/hotreload"
)

func main() {
	fmt.Println("Starting the Reload...")

	watcher := hotreload.NewWatcher(".")
	defer watcher.Close()

	watcher.Start()
}
