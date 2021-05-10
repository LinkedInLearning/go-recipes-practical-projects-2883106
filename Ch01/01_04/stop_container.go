package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := killContainer("cid.txt"); err != nil {
		log.Fatal(err)
	}
}

func killContainer(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	defer os.Remove(fileName)

	cid := strings.TrimSpace(string(data))
	if !isValidID(cid) {
		return fmt.Errorf("bad container id: %q", cid)
	}

	log.Printf("stopping container %s", cid)
	if err := exec.Command("docker", "rm", "-f", cid).Run(); err != nil {
		return fmt.Errorf("failed to stop %s: %w", cid, err)
	}

	return nil
}

func isValidID(cid string) bool {
	return len(cid) == 12 || len(cid) == 64
}
