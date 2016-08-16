package main

import (
	"fmt"
	"os"
	"time"

	"github.com/slok/gospinner"
)

func spine(kind gospinner.AnimationKind, color gospinner.ColorAttr) {
	s, err := gospinner.NewSpinnerWithColor(kind, color)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.Start("Starting, please wait...")
	time.Sleep(3 * time.Second)
	s.FinishWithMessage("", "Loaded!")
}

func main() {
	spine(gospinner.Ball, gospinner.FgCyan)
	spine(gospinner.Column, gospinner.FgGreen)
	spine(gospinner.Slash, gospinner.FgBlue)
	spine(gospinner.Square, gospinner.FgMagenta)
	spine(gospinner.Triangle, gospinner.FgWhite)
	spine(gospinner.Dots, gospinner.FgHiBlue)
	spine(gospinner.Dots2, gospinner.FgHiGreen)
	spine(gospinner.Pipe, gospinner.FgHiRed)
	spine(gospinner.SimpleDots, gospinner.FgHiMagenta)
	spine(gospinner.SimpleDotsScrolling, gospinner.FgHiWhite)
	spine(gospinner.GrowVertical, gospinner.FgHiYellow)
	spine(gospinner.GrowHorizontal, gospinner.FgHiCyan)
	spine(gospinner.Arrow, gospinner.FgYellow)
	spine(gospinner.BouncingBar, gospinner.FgMagenta)
	spine(gospinner.BouncingBall, gospinner.FgGreen)
	spine(gospinner.Pong, gospinner.FgHiRed)
	spine(gospinner.ProgressBar, gospinner.FgHiBlue)
}
