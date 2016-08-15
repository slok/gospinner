package gospinner

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
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
}

// NewSpinner returns an animation type
func NewSpinner(kind AnimationKind, startMessage string) (*Spinner, error) {
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

	// Set the initial frames
	s.createFrames()

	return s, nil
}

func (s *Spinner) createFrames() {
	f := make([]string, len(s.animation.frames))

	for i, c := range s.animation.frames {
		f[i] = fmt.Sprintf("%s %s", c, s.message)
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
	s.FinishWithSymbol(successSymbol)
}

// Fail will stop the animation with a failure symbol where the spinner is
func (s *Spinner) Fail() {
	s.FinishWithSymbol(failureSymbol)
}

// Warn will stop the animation with a warning symbol where the spinner is
func (s *Spinner) Warn() {
	s.FinishWithSymbol(WarningSymbol)
}

// Finish will stop an write to the next line
func (s *Spinner) Finish() {
	s.Stop()
	s.Reset()
	fmt.Fprint(s.Writer, "\n")
}

// FinishWithSymbol will finish the animation with a symbol where the spinner is
func (s *Spinner) FinishWithSymbol(symbol string) {
	s.Stop()
	s.Reset()
	msg := fmt.Sprintf("%s%s %s\n", s.separator, symbol, s.message)
	fmt.Fprint(s.Writer, msg)
	// should maintian the spinner
}

// FinishWithMessage will finish animation setting a message and a symbol where the spinner was
func (s *Spinner) FinishWithMessage(symbol, closingMessage string) {
	s.Stop()
	s.Reset()
	fmt.Fprintf(s.Writer, "%s%s %s\n", s.separator, symbol, closingMessage)
}
