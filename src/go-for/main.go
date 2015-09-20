package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("./go-for.bash")
	bs, _ := c.Output()
	fmt.Println(string(bs))
}
