package wanderer

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

type Point struct{ X, Y int }

type Size struct{ Width, Height int }

type Area struct {
	Base Point
	Size
}

func (a Area) Wander(stop chan bool) {
	for {
		x := rand.Intn(a.Width) + a.Base.X
		y := rand.Intn(a.Height) + a.Base.Y

		robotgo.MoveSmooth(x, y, 0.3, 0.7)

		pause := time.Duration(0.5e9 + rand.Float64()*10e9)
		timer := time.NewTimer(pause)

		select {
		case <-stop:
			return
		case <-timer.C:
		}
	}
}
