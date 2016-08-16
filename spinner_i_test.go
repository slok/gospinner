package gospinner

import (
	"bytes"
	"testing"
	"time"
)

func TestIntegrationStartNoColor(t *testing.T) {
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
		s, _ := NewSpinnerNoColor(test.kind)
		s.Writer = &buf
		s.separator = "|"

		s.Start(test.startMessage)
		time.Sleep(s.animation.interval*time.Duration(len(s.animation.frames)) + 5*time.Millisecond)
		s.Stop()

		got := buf.String()
		if got != test.want {
			t.Errorf("%+v\n - Wrong result, got: %v, want: %v", test, got, test.want)
		}
	}
}
func TestIntegrationStartDefault(t *testing.T) {
	tests := []struct {
		kind         AnimationKind
		startMessage string
		color        ColorAttr

		want string
	}{
		{Ball, "This is a test", FgHiCyan, "|\x1b[96m◐\x1b[0m This is a test|\x1b[96m◓\x1b[0m This is a test|\x1b[96m◑\x1b[0m This is a test|\x1b[96m◒\x1b[0m This is a test"},
		{Dots, "This is another test", FgMagenta, "|\x1b[96m⠋\x1b[0m This is another test|\x1b[96m⠙\x1b[0m This is another test|\x1b[96m⠹\x1b[0m This is another test|\x1b[96m⠸\x1b[0m This is another test|\x1b[96m⠼\x1b[0m This is another test|\x1b[96m⠴\x1b[0m This is another test|\x1b[96m⠦\x1b[0m This is another test|\x1b[96m⠧\x1b[0m This is another test|\x1b[96m⠇\x1b[0m This is another test|\x1b[96m⠏\x1b[0m This is another test"},
		{BouncingBar, "Это тест", FgHiGreen, "|\x1b[96m[    ]\x1b[0m Это тест|\x1b[96m[   =]\x1b[0m Это тест|\x1b[96m[  ==]\x1b[0m Это тест|\x1b[96m[ ===]\x1b[0m Это тест|\x1b[96m[====]\x1b[0m Это тест|\x1b[96m[=== ]\x1b[0m Это тест|\x1b[96m[==  ]\x1b[0m Это тест|\x1b[96m[=   ]\x1b[0m Это тест"},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		s, _ := NewSpinner(test.kind)
		s.Writer = &buf
		s.separator = "|"

		s.Start(test.startMessage)
		time.Sleep(s.animation.interval*time.Duration(len(s.animation.frames)) + 5*time.Millisecond)
		s.Stop()

		got := buf.String()
		if got != test.want {
			t.Errorf("%+v\n - Wrong result, got: %s, want: %s", test, got, test.want)
		}
	}
}

func TestIntegrationStartColor(t *testing.T) {
	tests := []struct {
		kind         AnimationKind
		startMessage string
		color        ColorAttr

		want string
	}{
		{Ball, "This is a test", FgHiCyan, "|\x1b[96m◐\x1b[0m This is a test|\x1b[96m◓\x1b[0m This is a test|\x1b[96m◑\x1b[0m This is a test|\x1b[96m◒\x1b[0m This is a test"},
		{Dots, "This is another test", FgMagenta, "|\x1b[35m⠋\x1b[0m This is another test|\x1b[35m⠙\x1b[0m This is another test|\x1b[35m⠹\x1b[0m This is another test|\x1b[35m⠸\x1b[0m This is another test|\x1b[35m⠼\x1b[0m This is another test|\x1b[35m⠴\x1b[0m This is another test|\x1b[35m⠦\x1b[0m This is another test|\x1b[35m⠧\x1b[0m This is another test|\x1b[35m⠇\x1b[0m This is another test|\x1b[35m⠏\x1b[0m This is another test"},
		{BouncingBar, "Это тест", FgHiGreen, "|\x1b[92m[    ]\x1b[0m Это тест|\x1b[92m[   =]\x1b[0m Это тест|\x1b[92m[  ==]\x1b[0m Это тест|\x1b[92m[ ===]\x1b[0m Это тест|\x1b[92m[====]\x1b[0m Это тест|\x1b[92m[=== ]\x1b[0m Это тест|\x1b[92m[==  ]\x1b[0m Это тест|\x1b[92m[=   ]\x1b[0m Это тест"},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		s, _ := NewSpinnerWithColor(test.kind, test.color)
		s.Writer = &buf
		s.separator = "|"

		s.Start(test.startMessage)
		time.Sleep(s.animation.interval*time.Duration(len(s.animation.frames)) + 5*time.Millisecond)
		s.Stop()

		got := buf.String()
		if got != test.want {
			t.Errorf("%+v\n - Wrong result, got: %s, want: %s", test, got, test.want)
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
		s, _ := NewSpinnerNoColor(test.kind)
		s.Writer = &buf
		s.separator = "|"
		speed := 20 * time.Millisecond

		s.StartWithSpeed(test.startMessage, speed)
		time.Sleep(speed*time.Duration(len(s.animation.frames)) + 5*time.Millisecond)
		s.Stop()

		got := buf.String()
		if got != test.want {
			t.Errorf("%+v\n - Wrong result, got: %v, want: %v", test, got, test.want)
		}
	}
}
