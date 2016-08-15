package gospinner

import (
	"bytes"
	"testing"
	"time"
)

func TestIntegrationStart(t *testing.T) {
	tests := []struct {
		kind         AnimationKind
		startMessage string

		want string
	}{
		{Ball, "This is a test", "|◐ This is a test|◓ This is a test|◑ This is a test|◒ This is a test"},
		{Dots, "This is another test", "|⠋ This is another test|⠙ This is another test|⠹ This is another test|⠸ This is another test|⠼ This is another test|⠴ This is another test|⠦ This is another test|⠧ This is another test|⠇ This is another test|⠏ This is another test"},
		{BouncingBar, "Это тест", "|[    ] Это тест|[   =] Это тест|[  ==] Это тест|[ ===] Это тест|[====] Это тест|[=== ] Это тест|[==  ] Это тест|[=   ] Это тест"},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		s := &Spinner{
			animation:    animations[test.kind],
			message:      test.startMessage,
			disableColor: true,
			separator:    "|",
			Writer:       &buf,
		}
		s.createFrames()
		s.Start()
		time.Sleep(s.animation.interval*time.Duration(len(s.animation.frames)) + 5*time.Millisecond)
		s.Stop()
		got := buf.String()
		if got != test.want {
			t.Errorf("%+v\n - Wrong result, got: %v, want: %v", test, got, test.want)
		}
	}
}

func TestIntegrationStartWithMessage(t *testing.T) {
	tests := []struct {
		kind         AnimationKind
		startMessage string

		want string
	}{
		{Ball, "This is a test", "|◐ This is a test|◓ This is a test|◑ This is a test|◒ This is a test"},
		{Dots, "This is another test", "|⠋ This is another test|⠙ This is another test|⠹ This is another test|⠸ This is another test|⠼ This is another test|⠴ This is another test|⠦ This is another test|⠧ This is another test|⠇ This is another test|⠏ This is another test"},
		{BouncingBar, "Это тест", "|[    ] Это тест|[   =] Это тест|[  ==] Это тест|[ ===] Это тест|[====] Это тест|[=== ] Это тест|[==  ] Это тест|[=   ] Это тест"},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		s := &Spinner{
			animation:    animations[test.kind],
			message:      "wrooooooooooooooooooong",
			disableColor: true,
			separator:    "|",
			Writer:       &buf,
		}
		s.createFrames()
		s.StartWithMessage(test.startMessage)
		time.Sleep(s.animation.interval*time.Duration(len(s.animation.frames)) + 5*time.Millisecond)
		s.Stop()
		got := buf.String()
		if got != test.want {
			t.Errorf("%+v\n - Wrong result, got: %v, want: %v", test, got, test.want)
		}
	}
}

func TestIntegrationStartWithSpeed(t *testing.T) {
	tests := []struct {
		kind         AnimationKind
		startMessage string

		want string
	}{
		{Ball, "This is a test", "|◐ This is a test|◓ This is a test|◑ This is a test|◒ This is a test"},
		{Dots, "This is another test", "|⠋ This is another test|⠙ This is another test|⠹ This is another test|⠸ This is another test|⠼ This is another test|⠴ This is another test|⠦ This is another test|⠧ This is another test|⠇ This is another test|⠏ This is another test"},
		{BouncingBar, "Это тест", "|[    ] Это тест|[   =] Это тест|[  ==] Это тест|[ ===] Это тест|[====] Это тест|[=== ] Это тест|[==  ] Это тест|[=   ] Это тест"},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		s := &Spinner{
			animation:    animations[test.kind],
			message:      test.startMessage,
			disableColor: true,
			separator:    "|",
			Writer:       &buf,
		}
		s.createFrames()
		speed := 20 * time.Millisecond
		s.StartWithSpeed(speed)

		time.Sleep(speed*time.Duration(len(s.animation.frames)) + 5*time.Millisecond)
		s.Stop()
		got := buf.String()
		if got != test.want {
			t.Errorf("%+v\n - Wrong result, got: %v, want: %v", test, got, test.want)
		}
	}
}
