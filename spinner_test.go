package gospinner

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestCorrectCreation(t *testing.T) {
	tests := []struct {
		kind AnimationKind
	}{
		{Ball},
		{Dots},
		{BouncingBar},
	}

	for _, test := range tests {
		_, err := NewSpinner(test.kind)
		if err != nil {
			t.Errorf("%+v\n - Creation shouldn't fail, it did: %s", test, err)
		}
	}
}

func TestStartMultipleTimes(t *testing.T) {
	s, err := NewSpinner(Ball)
	if err != nil {
		t.Errorf("\n - Creation shouldn't fail, it did: %s", err)
	}
	err = s.Start("test")
	if err != nil {
		t.Errorf("\n - First run shouldn't fail, it did: %s", err)
	}
	time.Sleep(1 * time.Millisecond)
	err = s.Start("test")
	if err == nil {
		t.Errorf("\n - Second run should fail, it didn't")
	}

}

func TestCreateFrames(t *testing.T) {
	tests := []struct {
		kind         AnimationKind
		startMessage string

		WantFrames []string
	}{
		{Ball, "This is a test", []string{"◐ This is a test", "◓ This is a test", "◑ This is a test", "◒ This is a test"}},
		{Dots, "This is another test", []string{"⠋ This is another test", "⠙ This is another test", "⠹ This is another test", "⠸ This is another test", "⠼ This is another test", "⠴ This is another test", "⠦ This is another test", "⠧ This is another test", "⠇ This is another test", "⠏ This is another test"}},
		{BouncingBar, "Это тест", []string{"[    ] Это тест", "[   =] Это тест", "[  ==] Это тест", "[ ===] Это тест", "[====] Это тест", "[=== ] Это тест", "[==  ] Это тест", "[=   ] Это тест"}},
	}

	for _, test := range tests {
		s, _ := NewSpinnerNoColor(test.kind)
		s.message = test.startMessage
		s.createFrames()

		if !reflect.DeepEqual(s.frames, test.WantFrames) {
			t.Errorf("%+v\n - Frames should be the same, got: %v, want: %v", test, s.frames, test.WantFrames)
		}
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		kind         AnimationKind
		startMessage string

		WantFrames []string
	}{
		{Ball, "This is a test", []string{"◐ This is a test", "◓ This is a test", "◑ This is a test", "◒ This is a test"}},
		{Dots, "This is another test", []string{"⠋ This is another test", "⠙ This is another test", "⠹ This is another test", "⠸ This is another test", "⠼ This is another test", "⠴ This is another test", "⠦ This is another test", "⠧ This is another test", "⠇ This is another test", "⠏ This is another test"}},
		{BouncingBar, "Это тест", []string{"[    ] Это тест", "[   =] Это тест", "[  ==] Это тест", "[ ===] Это тест", "[====] Это тест", "[=== ] Это тест", "[==  ] Это тест", "[=   ] Это тест"}},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		s, _ := NewSpinnerNoColor(test.kind)
		s.message = test.startMessage
		s.Writer = &buf

		s.createFrames()

		for _, f := range test.WantFrames {
			s.Render()
			if !strings.Contains(buf.String(), f) {
				t.Errorf("%+v\n - Wrong frame rendered, got: %v, want: %v", test, buf.String(), f)
			}
		}
	}
}

func TestRenderError(t *testing.T) {
	s, _ := NewSpinnerNoColor(Ball)
	err := s.Render()
	if err == nil {
		t.Errorf("\n - Render should return error, it didn't")
	}

}

func TestRenderColor(t *testing.T) {
	tests := []struct {
		kind         AnimationKind
		startMessage string
		color        ColorAttr

		WantFrames []string
	}{
		{Ball, "This is a test", FgHiCyan, []string{"\x1b[96m◐\x1b[0m This is a test", "\x1b[96m◓\x1b[0m This is a test", "\x1b[96m◑\x1b[0m This is a test", "\x1b[96m◒\x1b[0m This is a test"}},
		{Dots, "This is another test", FgMagenta, []string{"\x1b[35m⠋\x1b[0m This is another test", "\x1b[35m⠙\x1b[0m This is another test", "\x1b[35m⠹\x1b[0m This is another test", "\x1b[35m⠸\x1b[0m This is another test", "\x1b[35m⠼\x1b[0m This is another test", "\x1b[35m⠴\x1b[0m This is another test", "\x1b[35m⠦\x1b[0m This is another test", "\x1b[35m⠧\x1b[0m This is another test", "\x1b[35m⠇\x1b[0m This is another test", "\x1b[35m⠏\x1b[0m This is another test"}},
		{BouncingBar, "Это тест", FgHiGreen, []string{"\x1b[92m[    ]\x1b[0m Это тест", "\x1b[92m[   =]\x1b[0m Это тест", "\x1b[92m[  ==]\x1b[0m Это тест", "\x1b[92m[ ===]\x1b[0m Это тест", "\x1b[92m[====]\x1b[0m Это тест", "\x1b[92m[=== ]\x1b[0m Это тест", "\x1b[92m[==  ]\x1b[0m Это тест", "\x1b[92m[=   ]\x1b[0m Это тест"}},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		s, _ := NewSpinnerWithColor(test.kind, test.color)
		s.message = test.startMessage
		s.Writer = &buf

		s.createFrames()

		for _, f := range test.WantFrames {
			s.Render()
			if !strings.Contains(buf.String(), f) {
				t.Errorf("%+v\n - Wrong frame rendered, got: %v, want: %v", test, buf.String(), f)
			}
		}
	}
}

func TestSetMessage(t *testing.T) {
	tests := []struct {
		kind          AnimationKind
		startMessage  string
		secondMessage string

		WantFrames []string
	}{
		{Ball, "1st", "2nd", []string{"◐ 1st", "◓ 2nd"}},
		{Dots, "first one", "second one", []string{"⠋ first one", "⠙ second one"}},
		{BouncingBar, "один", "दुई", []string{"[    ] один", "[   =] दुई"}},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		s, _ := NewSpinnerNoColor(test.kind)
		s.message = test.startMessage
		s.Writer = &buf
		s.createFrames()

		s.Render()
		if !strings.Contains(buf.String(), test.WantFrames[0]) {
			t.Errorf("%+v\n - Wrong frame rendered, got: %v, want: %v", test, buf.String(), test.WantFrames[0])
		}
		s.SetMessage(test.secondMessage)
		s.Render()
		if !strings.Contains(buf.String(), test.WantFrames[1]) {
			t.Errorf("%+v\n - Wrong frame rendered, got: %v, want: %v", test, buf.String(), test.WantFrames[1])
		}
	}
}
func TestStop(t *testing.T) {
	var buf bytes.Buffer
	s, _ := NewSpinnerNoColor(Ball)
	s.Writer = &buf
	s.Start("test")

	time.Sleep(10 * time.Millisecond)

	if !s.running {
		t.Errorf("- Spinner should be running, it isn't")
	}

	s.Stop()

	if s.running {
		t.Errorf("- Spinner should be stopped, it isn't")
	}
}

func TestStopError(t *testing.T) {
	s, _ := NewSpinnerNoColor(Ball)
	err := s.Stop()
	if err == nil {
		t.Errorf("\n - Stop should return error, it didn't")
	}
}

func TestReset(t *testing.T) {
	var buf bytes.Buffer
	frames := []string{"◐ This is a test", "◓ This is a test", "◑ This is a test", "◒ This is a test"}

	s, _ := NewSpinnerNoColor(Ball)
	s.Writer = &buf
	s.message = "This is a test"
	s.step = 5

	s.Reset()

	if s.step != 0 {
		t.Errorf("- Spinner step should be 0, got: %d", s.step)
	}

	if !reflect.DeepEqual(s.frames, frames) {
		t.Errorf("- Frames should be the same, got: %v, want: %v", s.frames, frames)
	}
}

func TestSucceed(t *testing.T) {
	want := "✔ test"

	var buf bytes.Buffer
	s, _ := NewSpinnerNoColor(Ball)
	s.Writer = &buf
	s.Start("test")
	time.Sleep(10 * time.Millisecond)

	if !s.running {
		t.Errorf("- Spinner should be running, it isn't")
	}

	s.Succeed()

	if s.running {
		t.Errorf("- Spinner should be stopped, it isn't")
	}
	if !strings.Contains(buf.String(), want) {
		t.Errorf("Wrong frame rendered, got: %v, want: %v", buf.String(), want)

	}
}

func TestFail(t *testing.T) {
	want := "✖ test"

	var buf bytes.Buffer
	s, _ := NewSpinnerNoColor(Ball)
	s.Writer = &buf
	s.Start("test")
	time.Sleep(10 * time.Millisecond)

	if !s.running {
		t.Errorf("- Spinner should be running, it isn't")
	}

	s.Fail()

	if s.running {
		t.Errorf("- Spinner should be stopped, it isn't")
	}
	if !strings.Contains(buf.String(), want) {
		t.Errorf("Wrong frame rendered, got: %v, want: %v", buf.String(), want)

	}
}

func TestWarn(t *testing.T) {
	want := "⚠ test"

	var buf bytes.Buffer
	s, _ := NewSpinnerNoColor(Ball)
	s.Writer = &buf
	s.Start("test")
	time.Sleep(10 * time.Millisecond)

	if !s.running {
		t.Errorf("- Spinner should be running, it isn't")
	}

	s.Warn()

	if s.running {
		t.Errorf("- Spinner should be stopped, it isn't")
	}
	if !strings.Contains(buf.String(), want) {
		t.Errorf("Wrong frame rendered, got: %v, want: %v", buf.String(), want)

	}
}

func TestFinish(t *testing.T) {
	frames := []string{"◐ test", "◓ test", "◑ test", "◒ test"}

	var buf bytes.Buffer
	s, _ := NewSpinnerNoColor(Ball)
	s.Writer = &buf
	s.Start("test")
	s.step = 5
	time.Sleep(10 * time.Millisecond)

	if !s.running {
		t.Errorf("- Spinner should be running, it isn't")
	}

	s.Warn()

	if s.running {
		t.Errorf("- Spinner should be stopped, it isn't")
	}
	if s.step != 0 {
		t.Errorf("- Spinner step should be 0, got: %d", s.step)
	}

	if !reflect.DeepEqual(s.frames, frames) {
		t.Errorf("- Frames should be the same, got: %v, want: %v", s.frames, frames)
	}

}

func TestFinishWithSymbol(t *testing.T) {
	symbol := "ℹ"
	want := "ℹ test"

	var buf bytes.Buffer
	s, _ := NewSpinnerNoColor(Ball)
	s.Writer = &buf
	s.Start("test")
	s.step = 5
	time.Sleep(10 * time.Millisecond)

	if !s.running {
		t.Errorf("- Spinner should be running, it isn't")
	}

	s.FinishWithSymbol(symbol)

	if s.running {
		t.Errorf("- Spinner should be stopped, it isn't")
	}
	if !strings.Contains(buf.String(), want) {
		t.Errorf("Wrong frame rendered, got: %v, want: %v", buf.String(), want)

	}
}

func TestFinishWithMessage(t *testing.T) {
	symbol := "ℹ"
	message := "test2"
	want := "ℹ test2"

	var buf bytes.Buffer
	s, _ := NewSpinnerNoColor(Ball)
	s.Writer = &buf
	s.Start(message)
	s.step = 5
	time.Sleep(10 * time.Millisecond)

	if !s.running {
		t.Errorf("- Spinner should be running, it isn't")
	}

	s.FinishWithMessage(symbol, message)

	if s.running {
		t.Errorf("- Spinner should be stopped, it isn't")
	}
	if !strings.Contains(buf.String(), want) {
		t.Errorf("Wrong frame rendered, got: %v, want: %v", buf.String(), want)

	}
}
