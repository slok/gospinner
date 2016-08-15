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
		kind         AnimationKind
		startMessage string

		WantFrames []string
	}{
		{Ball, "This is a test", []string{"◐ This is a test", "◓ This is a test", "◑ This is a test", "◒ This is a test"}},
		{Dots, "This is another test", []string{"⠋ This is another test", "⠙ This is another test", "⠹ This is another test", "⠸ This is another test", "⠼ This is another test", "⠴ This is another test", "⠦ This is another test", "⠧ This is another test", "⠇ This is another test", "⠏ This is another test"}},
		{BouncingBar, "Это тест", []string{"[    ] Это тест", "[   =] Это тест", "[  ==] Это тест", "[ ===] Это тест", "[====] Это тест", "[=== ] Это тест", "[==  ] Это тест", "[=   ] Это тест"}},
	}

	for _, test := range tests {
		s, err := NewSpinner(test.kind, test.startMessage)
		if err != nil {
			t.Errorf("%+v\n - Creation shouldn't fail, it did: %s", test, err)
		}

		if !reflect.DeepEqual(s.frames, test.WantFrames) {
			t.Errorf("%+v\n - Frames should be the same, got: %v, want: %v", test, s.frames, test.WantFrames)
		}
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
		s := &Spinner{
			animation:    animations[test.kind],
			message:      test.startMessage,
			disableColor: true,
		}

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

		s := &Spinner{
			animation:    animations[test.kind],
			message:      test.startMessage,
			disableColor: true,
			Writer:       &buf,
		}
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

		s := &Spinner{
			animation:    animations[test.kind],
			message:      test.startMessage,
			disableColor: true,
			Writer:       &buf,
		}
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

	s := &Spinner{
		animation:    animations[Ball],
		message:      "test",
		disableColor: true,
		Writer:       &buf,
	}
	s.createFrames()
	s.Start()
	time.Sleep(10 * time.Millisecond)

	if !s.running {
		t.Errorf("- Spinner should be running, it isn't")
	}

	s.Stop()

	if s.running {
		t.Errorf("- Spinner should be stopped, it isn't")
	}
}

func TestReset(t *testing.T) {
	var buf bytes.Buffer
	frames := []string{"◐ This is a test", "◓ This is a test", "◑ This is a test", "◒ This is a test"}

	s := &Spinner{
		animation:    animations[Ball],
		message:      "This is a test",
		disableColor: true,
		Writer:       &buf,
		step:         5,
	}
	s.Reset()

	if s.step != 0 {
		t.Errorf("- Spinner step should be 0, got: %d", s.step)
	}

	if !reflect.DeepEqual(s.frames, frames) {
		t.Errorf("- Frames should be the same, got: %v, want: %v", s.frames, frames)
	}
}

func TestSucceed(t *testing.T) {
	var buf bytes.Buffer

	want := "✔ test"
	s := &Spinner{
		animation:    animations[Ball],
		message:      "test",
		disableColor: true,
		Writer:       &buf,
	}
	s.createFrames()
	s.Start()
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
	var buf bytes.Buffer

	want := "✖ test"
	s := &Spinner{
		animation:    animations[Ball],
		message:      "test",
		disableColor: true,
		Writer:       &buf,
	}
	s.createFrames()
	s.Start()
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
	var buf bytes.Buffer

	want := "⚠ test"
	s := &Spinner{
		animation:    animations[Ball],
		message:      "test",
		disableColor: true,
		Writer:       &buf,
	}
	s.createFrames()
	s.Start()
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
	var buf bytes.Buffer
	frames := []string{"◐ test", "◓ test", "◑ test", "◒ test"}

	s := &Spinner{
		animation:    animations[Ball],
		message:      "test",
		disableColor: true,
		Writer:       &buf,
		step:         5,
	}
	s.createFrames()
	s.Start()
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
	var buf bytes.Buffer
	symbol := "ℹ"
	want := "ℹ test"
	s := &Spinner{
		animation:    animations[Ball],
		message:      "test",
		disableColor: true,
		Writer:       &buf,
	}
	s.createFrames()
	s.Start()
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
	var buf bytes.Buffer
	symbol := "ℹ"
	message := "test2"
	want := "ℹ test2"
	s := &Spinner{
		animation:    animations[Ball],
		message:      "test",
		disableColor: true,
		Writer:       &buf,
	}
	s.createFrames()
	s.Start()
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
