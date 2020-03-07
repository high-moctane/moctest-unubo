package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	err := exec.Command("./mocbinary-unibo", "").Run()
	if err != nil {
		log.Fatal(err)
	}
	for {
		time.Sleep(time.Hour)
	}
}
