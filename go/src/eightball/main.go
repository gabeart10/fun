package main

import (
	"bufio"
	"fmt"
	c "github.com/skilstak/go/colors"
	"math/rand"
	"os"
	"strings"
	"time"
)

var answers = []string{
	"Yes",
	"No",
	"Maybe",
}

var easterEgg1 = []string{"kill", "Well Then"}
var easterEgg2 = []string{"love", "Love you too ;)"}
var responses = [][]string{easterEgg1, easterEgg2}

func main() {
	for true {
		egg := 0
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(c.Green + "What do you want to know?\n" + c.Reset)
		text, _ := reader.ReadString('\n')
		for _, easterEgg := range responses {
			if strings.Contains(text, easterEgg[0]) {
				fmt.Println(c.Red + easterEgg[1] + c.Reset)
				egg = 1
				reader.ReadString('\n')
				fmt.Print(c.Clear)
			}
		}
		if egg == 0 {
			num := r1.Intn(3)
			fmt.Println(c.Yellow + answers[num] + c.Reset)
			reader.ReadString('\n')
			fmt.Print(c.Clear)
		}
	}
}
