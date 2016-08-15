package gospinner

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	// default colors for the application
	defaultColor        = newColor(FgHiCyan)
	defaultSuccessColor = newColor(FgHiGreen)
	defaultFailColor    = newColor(FgHiRed)
	defaultWarnColor    = newColor(FgHiYellow)
)

// Spinner is a representation of the animation itself
type Spinner struct {

	// Writer will be the target of the printing
	Writer io.Writer

	// frames are the frames that will be showed on screen, they are a representation of animation+runes
	frames []string

	// message is the content wanted to show with the loading animation
	message string

	// chars are the animation characters
	animation Animation

	// step tracks the current step
	step int

	// ticker is the animation ticker, will set the pace
	ticker *time.Ticker

	// Previous frame is used to clean the screen
	previousFrame string // TODO use  bytes so we dont allocate new strings always

	// running shows the state of the animation
	running bool

	// Separator will separate the messages each other, by default this should be carriage return
	separator string

	// color
	color *Color

	// disableColor
	disableColor bool
}

//create is a helper function for all the creators
func create(kind AnimationKind, startMessage string) (*Spinner, error) {
	an, ok := animations[kind]
	if !ok {
		return nil, errors.New("Wrong kind of animation")
	}
	s := &Spinner{
		animation: an,
		message:   startMessage,
		Writer:    os.Stdout,
		separator: "\r",
	}
	return s, nil
}

// NewSpinner is the defalult spinner for easy ussage with default colors
func NewSpinner(kind AnimationKind, startMessage string) (*Spinner, error) {

	s, err := create(kind, startMessage)
	if err != nil {
		return nil, err
	}

	s.color = defaultColor
	s.createFrames()

	return s, nil
}

func NewSpinnerNoColor(kind AnimationKind, startMessage string) (*Spinner, error) {
	s, err := create(kind, startMessage)
	if err != nil {
		return nil, err
	}

	s.color = defaultColor
	s.disableColor = true
	s.createFrames()

	return s, nil
}

// NewSpinnerWithColor returns a new spinner with color
func NewSpinnerWithColor(kind AnimationKind, startMessage string, color ColorAttr) (*Spinner, error) {
	s, err := create(kind, startMessage)
	if err != nil {
		return nil, err
	}

	s.color = newColor(color)
	s.createFrames()
	return s, nil
}

func (s *Spinner) createFrames() {
	f := make([]string, len(s.animation.frames))
	for i, c := range s.animation.frames {
		var symbol = c
		if !s.disableColor || s.color != nil {
			symbol = s.color.SprintfFunc()(c)
		}
		f[i] = fmt.Sprintf("%s %s", symbol, s.message)
	}

	// Set the new animation
	s.frames = f

}

// Start will animate with the recomended speed
func (s *Spinner) Start() {
	s.StartWithSpeed(s.animation.interval)
}

// StartWithMessage will animate with the recommended speend and a new message
func (s *Spinner) StartWithMessage(message string) {
	s.StartWithSpeed(s.animation.interval)
	s.message = message
	s.createFrames()
}

// StartWithSpeed will start animation the spinner on the
func (s *Spinner) StartWithSpeed(speed time.Duration) {
	// Start the animation in background
	go func() {
		s.running = true
		s.ticker = time.NewTicker(speed)

		for range s.ticker.C {
			s.Render()
		}
	}()
}

// Render will render manually an step
func (s *Spinner) Render() {
	s.step = s.step % len(s.frames)
	previousLen := len(s.previousFrame)
	s.previousFrame = fmt.Sprintf("%s%s", s.separator, s.frames[s.step])
	newLen := len(s.previousFrame)

	// We need to clean the previous message
	if previousLen > newLen {
		r := previousLen - newLen
		suffix := strings.Repeat(" ", r)
		s.previousFrame = s.previousFrame + suffix
	}

	fmt.Fprint(s.Writer, s.previousFrame)
	s.step++
}

// SetMessage will set new message on the animation without stoping it
func (s *Spinner) SetMessage(message string) {
	s.message = message
	s.createFrames()
}

// Stop will stop the animation
func (s *Spinner) Stop() {
	s.ticker.Stop()
	s.running = false
}

// Reset will set the spinner to its initial frame
func (s *Spinner) Reset() {
	s.step = 0
	s.createFrames()
}

// Succeed will stop the animation with a success symbol where the spinner is
func (s *Spinner) Succeed() {
	symbol := successSymbol
	if !s.disableColor || s.color != nil {
		symbol = defaultSuccessColor.SprintfFunc()(successSymbol)
	}
	s.FinishWithSymbol(symbol)
}

// Fail will stop the animation with a failure symbol where the spinner is
func (s *Spinner) Fail() {
	symbol := failureSymbol
	if !s.disableColor || s.color != nil {
		symbol = defaultFailColor.SprintfFunc()(failureSymbol)
	}
	s.FinishWithSymbol(symbol)
}

// Warn will stop the animation with a warning symbol where the spinner is
func (s *Spinner) Warn() {
	symbol := WarningSymbol
	if !s.disableColor || s.color != nil {
		symbol = defaultWarnColor.SprintfFunc()(WarningSymbol)
	}
	s.FinishWithSymbol(symbol)
}

// Finish will stop an write to the next line
func (s *Spinner) Finish() {
	s.Stop()
	s.Reset()
	fmt.Fprint(s.Writer, "\n")
}

// FinishWithSymbol will finish the animation with a symbol where the spinner is
func (s *Spinner) FinishWithSymbol(symbol string) {
	s.FinishWithMessage(symbol, s.message)
}

// FinishWithMessage will finish animation setting a message and a symbol where the spinner was
func (s *Spinner) FinishWithMessage(symbol, closingMessage string) {
	s.Stop()
	s.Reset()
	previousLen := len(s.previousFrame)
	finalMsg := fmt.Sprintf("%s %s", symbol, closingMessage)
	newLen := len(finalMsg)
	if previousLen > newLen {
		r := previousLen - newLen
		suffix := strings.Repeat(" ", r)
		finalMsg = finalMsg + suffix
	}
	fmt.Fprintf(s.Writer, "%s%s\n", s.separator, finalMsg)
}
