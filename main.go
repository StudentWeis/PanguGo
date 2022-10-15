package main

import (
	"fmt"
	"os"

	"github.com/vinta/pangu"
)

func main() {
	s := pangu.SpacingText(os.Args[1])
	fmt.Println(s)
}
