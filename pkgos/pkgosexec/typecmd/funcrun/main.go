package main

import (
	"log"
	"os/exec"
)

// 对比 funcstart/main.go
func main() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
