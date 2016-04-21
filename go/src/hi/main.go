package main

import (
	"bufio"
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(c.Clear + c.Blue + "What is your name?\n" + c.Reset)
	text, _ := reader.ReadString('\n')
	fmt.Println(c.Red + "Hi " + text + c.Reset)
}
