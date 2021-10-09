package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/giucal/wanderer"
	"github.com/go-vgo/robotgo"
)

type SizeArg wanderer.Size

func (s *SizeArg) Set(value string) error {
	size := strings.Split(value, "x")
	if len(size) != 2 {
		return fmt.Errorf("bad size format: %s", value)
	}

	width, err := strconv.Atoi(size[0])
	if err != nil || width < 0 {
		return fmt.Errorf("bad width: %d", width)
	}
	height, err := strconv.Atoi(size[1])
	if err != nil || width < 0 {
		return fmt.Errorf("bad height: %d", height)
	}

	s.Width = width
	s.Height = height

	return nil
}

func (s *SizeArg) String() string { return "1024x768" }

func main() {
	size := SizeArg{Width: 1024, Height: 768}

	flag.Var(&size, "area", "set mouse area to `<width>x<height> pixels`")
	flag.Parse()

	area := wanderer.Area{
		Size: wanderer.Size(size),
	}

	stop := make(chan bool)
	go area.Wander(stop)
	// A left click terminates the program.
	robotgo.AddMouse("left")
	stop <- true
}
