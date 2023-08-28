package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func stopServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("invalid pid in %q - %w", pidFile, err)
	}

	if pid <= 0 {
		return fmt.Errorf("bad pid in %q - %d", pidFile, pid)
	}

	defer os.Remove(pidFile)

	log.Printf("stopping server with PID %d", pid)
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("can't find process %d - %w", pid, err)
	}
	if err := proc.Kill(); err != nil {
		return fmt.Errorf("can't kill process %d - %w", pid, err)
	}

	return nil
}

func main() {
	if err := stopServer("httpd.pid"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("server not running")
		} else {
			log.Fatalf("error: %s", err)
		}
	}
}
