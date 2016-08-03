package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	_, err = stdin.Write([]byte("write sth,then cat will echo you"))
	if err != nil {
		fmt.Println(err)
	}
	stdin.Close()
	cmd.Stdout = os.Stdout
	cmd.Start()
}
