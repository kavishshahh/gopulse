package main

import (
	"fmt"

	"github.com/kavishshahh/GoPulse/internal/hotreload"
)

func main() {
	fmt.Println("Starting the Reload...")

	watcher := hotreload.NewWatcher(".")
	defer watcher.Close()

	watcher.Start()
}
