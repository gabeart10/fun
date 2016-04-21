package main

import (
	"bufio"
	"fmt"
	c "github.com/skilstak/go/colors"
	"math/rand"
	"os"
	"time"
)

func q1() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(c.Clear + c.Red + "What is life?" + c.Reset)
	text, _ := reader.ReadString('\n')
	if text == "47\n" {
		return true
	} else {
		return false
	}
}

func q2() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(c.Clear + c.Red + "What is my name?" + c.Reset)
	text, _ := reader.ReadString('\n')
	if text == "gabe\n" {
		return true
	} else {
		return false
	}
}

func q3() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(c.Clear + c.Red + "What is love?" + c.Reset)
	text, _ := reader.ReadString('\n')
	if text == "uh\n" {
		return true
	} else {
		return false
	}
}

var questions = []func() bool{q1, q2, q3}

func main() {
	for len(questions) > 0 {
		reader := bufio.NewReader(os.Stdin)
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		num := r1.Intn(len(questions))
		result := questions[num]()
		if result {
			fmt.Println(c.Green + "Correct" + c.Reset)
			questions = append(questions[:num], questions[num+1:]...)
			reader.ReadString('\n')
		} else {
			fmt.Println(c.Red + "Wrong" + c.Reset)
			reader.ReadString('\n')
		}
	}
}
