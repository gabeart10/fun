package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"time"
)

func main() {
	for i := 0; i < 6; i++ {
		fmt.Println(c.Blue + "Badger" + c.Reset)
		time.Sleep(1 * time.Second)
	}

	for i := 0; i < 2; i++ {
		fmt.Println(c.Red + "Mushroom" + c.Reset)
		time.Sleep(1 * time.Second)
	}
}
