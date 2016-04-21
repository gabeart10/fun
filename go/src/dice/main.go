package main

import (
	"bufio"
	"fmt"
	c "github.com/skilstak/go/colors"
	"math/rand"
	"os"
	"time"
)

var die = []string{}

func main() {

	die = append(die, `
	 -------
	|       |
	|   .   |
	|       |
	 -------      
	`)
	die = append(die, `
	 -------
	|   .   |
	|       |
	|   .   |
	 -------      
	`)
	die = append(die, `
	 -------
	|   .   |
	|   .   |
	|   .   |
	 -------      
	`)
	die = append(die, `
	 -------
	| .   . |
	|       |
	| .   . |
	 -------      
	`)
	die = append(die, `
	 -------
	| .   . |
	|   .   |
	| .   . |
	 -------      
	`)
	die = append(die, `
	 -------
	| .   . |
	| .   . |
	| .   . |
	 -------      
	`)

	reader := bufio.NewReader(os.Stdin)
	for true {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		num := r1.Intn(6)
		fmt.Println(c.Clear + c.Rc() + die[num] + c.Reset)
		reader.ReadString('\n')
	}
}
