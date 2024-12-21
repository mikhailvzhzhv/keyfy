package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/mikhailvzhzhv/keyfy/pkg/mapper"
)

func main() {

	robotgo.KeyTap("c", "ctrl")
	robotgo.MilliSleep(50)

	text, err := robotgo.ReadAll()
	if err != nil {
		return
	}

	mapped_text := mapper.Mapkeys(text)
	err = robotgo.WriteAll(mapped_text)
	if err != nil {
		return
	}

	robotgo.MilliSleep(50)
	robotgo.KeyTap("v", "ctrl")
}
