package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	output := exec.Command("/root/rkdeveloptool/rkdeveloptool", "ld")
	o, err := output.Output()
	if err != nil {
		log.Fatalf("failed to output: %v\n", err)
	}
	str := string(o)

	fmt.Println(str)
	fmt.Printf("equal: %v\n", str == "not found any devices!\n")

}
