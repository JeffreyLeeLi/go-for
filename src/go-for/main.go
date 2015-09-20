package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("./go-for.bash")
	fmt.Println(c.Output())
}
