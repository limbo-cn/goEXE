package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd.exe", "/C", "cd /d D:/ditu/DebugClient & start Jovian.ClientMap")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
