package main

import (
	"fmt"
	"log"

	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
)

func main() {

	robotgo.KeyTap("c", "control")

	text, err := clipboard.ReadAll()
	if err != nil {
		log.Println("clipboard read all error: ", err)
		return
	}

	fmt.Println(translator.Translate(text))
}
