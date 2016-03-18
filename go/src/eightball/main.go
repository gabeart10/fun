package main

import (
	"fmt"
	"bufio"
	"os"
	c"github.com/skilstak/go/colors"
	"time"
	"math/rand"
)

var eastereggs 

func main() {
	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)

