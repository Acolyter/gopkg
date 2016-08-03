package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("date")
	cout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is: %s\nThe Stderr is: %s\n", cout, cmd.Stderr)
}
