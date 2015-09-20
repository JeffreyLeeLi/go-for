package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	flag.Parse()

	c := exec.Command("./go-for.bash")
	bs, _ := c.Output()
	fmt.Println(string(bs))
}
