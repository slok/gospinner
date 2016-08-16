package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/slok/gospinner"
)

// Shuttle launch steps taken from: http://space.stackexchange.com/a/7996
var shuttleLaunchSteps = []string{
	"Start automatic ground launch sequencer",
	"Retract orbiter access arm",
	"Start auxiliary power units",
	"Arm solid rocket booster range safety safe and arm devices",
	"Start orbiter aerosurface profile test, followed by main engine gimbal profile test",
	"Retract gaseous oxygen vent arm, or 'beanie cap'",
	"Crew members close and lock their visors",
	"Orbiter transfers from ground to internal power",
	"Ground launch sequencer is go for auto sequence start",
	"Activate launch pad sound suppression system",
	"Activate main engine hydrogen burnoff system",
	"Main engine start",
}

func main() {
	s, err := gospinner.NewSpinner(gospinner.Dots2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, p := range shuttleLaunchSteps {
		s.Start(p)

		// Random sleep simulates doing stuff...
		seed := rand.NewSource(time.Now().UnixNano())
		r := rand.New(seed)
		ms := r.Intn(3000)
		time.Sleep(time.Duration(ms) * time.Millisecond)

		switch ms % 3 {
		case 0:
			s.Succeed()
		case 1:
			s.Fail()
		case 2:
			s.Warn()
		}
	}
}
