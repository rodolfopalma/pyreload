package main

import (
	"flag"
	"fmt"
	"github.com/go-fsnotify/fsnotify"
	"log"
	"os/exec"
	"time"
)

var python_file = flag.String("f", "main.py", "Name of Python file that is going to be reran.")
var watch_dir = flag.String("d", ".", "Name of folder that is going to be watched.")

func main() {
	flag.Parse()

	watcher, _ := fsnotify.NewWatcher()
	last_time := time.Now()

	defer watcher.Close()

	runFile()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if time.Now().After(last_time.Add(500 * time.Millisecond)) {
					last_time = time.Now()

					log.Println(event.Name, "modified. Reloading...")
					runFile()
				}

			case err := <-watcher.Errors:
				log.Println("error:", err)
			}

		}
	}()

	_ = watcher.Add(*python_file)
	_ = watcher.Add(*watch_dir)

	<-done
}

func runFile() {
	out, _ := exec.Command("python3", *python_file).CombinedOutput()
	fmt.Printf("%s", out)
}
