package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/slok/gospinner"
)

var steps = []string{
	"Generating src files",
	"Compiling binary",
	"Setting environment",
	"Running unit tests",
	"Running integration tests",
}

func main() {
	s, err := gospinner.NewSpinner(gospinner.Dots2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, p := range steps {
		total := 0
		s.Start(p)
		for i := 0; i < 15; i++ {
			seed := rand.NewSource(time.Now().UnixNano())
			r := rand.New(seed)

			// Increment % randomly
			total = total + r.Intn(15)
			if total > 100 {
				total = 100
				break
			}

			msg := fmt.Sprintf("%s (%d%%)", p, total)
			s.SetMessage(msg)

			ms := r.Intn(1000)
			time.Sleep(time.Duration(ms) * time.Millisecond)

		}
		s.SetMessage(fmt.Sprintf(p))
		s.Succeed()
	}

}
