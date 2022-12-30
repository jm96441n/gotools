package delay_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/jm96441n/gotools/delay"
)

func TestPrintEchosBackAllLetters(t *testing.T) {
	t.Parallel()
	want := "hello this is a test\n"
	fakeInput := bytes.NewBufferString(want)
	fakeTerminal := bytes.NewBuffer([]byte{})
	delayer, err := delay.NewDelayer(
		delay.WithInput(fakeInput),
		delay.WithOutput(fakeTerminal),
		delay.WithDelay(time.Duration(0)),
	)
	if err != nil {
		t.Error(err)
	}
	delayer.Print()
	got := fakeTerminal.String()

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestPrintEchosBackAllLettersWithADelay(t *testing.T) {
	t.Parallel()
	want := "he\n"
	fakeInput := bytes.NewBufferString(want)
	fakeTerminal := bytes.NewBuffer([]byte{})
	delayer, err := delay.NewDelayer(
		delay.WithInput(fakeInput),
		delay.WithOutput(fakeTerminal),
		delay.WithDelay(time.Duration(1*time.Millisecond)),
	)
	if err != nil {
		t.Error(err)
	}
	startTime := time.Now()
	delayer.Print()
	endTime := time.Now()
	got := fakeTerminal.String()

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
	numWaits := time.Duration(len(want) * 1)
	timeDiff := endTime.Sub(startTime)
	minTimeDiff := time.Duration(time.Millisecond * numWaits)
	if timeDiff < minTimeDiff {
		t.Errorf("Expected time to be more than %d, got %d", minTimeDiff, timeDiff)
	}
}

func TestPrintEchosBackAllLettersAndPreservesNewLines(t *testing.T) {
	want := "hello this is a test\non another line\n"
	fakeInput := bytes.NewBufferString(want)
	fakeTerminal := bytes.NewBuffer([]byte{})
	delayer, err := delay.NewDelayer(
		delay.WithInput(fakeInput),
		delay.WithOutput(fakeTerminal),
		delay.WithDelay(time.Duration(0)),
	)
	if err != nil {
		t.Error(err)
	}
	delayer.Print()
	got := fakeTerminal.String()

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestPrintEchosBackAllLettersWithArgs(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "one file",
			args: []string{"testdata/three_lines.txt"},
			want: "hello\nthis is some\ninput from three lines\n",
		},
		{
			name: "multiple files",
			args: []string{"testdata/three_lines.txt", "testdata/one_line.txt"},
			want: "hello\nthis is some\ninput from three lines\ninput from one line\n",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fakeTerminal := bytes.NewBuffer([]byte{})
			delayer, err := delay.NewDelayer(
				delay.WithArgs(tc.args),
				delay.WithOutput(fakeTerminal),
				delay.WithDelay(time.Duration(0)),
			)
			if err != nil {
				t.Error(err)
			}
			delayer.Print()
			got := fakeTerminal.String()

			if got != tc.want {
				t.Errorf("Got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestPrintEchosBackAllLettersWithZeroArgs(t *testing.T) {
	t.Parallel()
	want := "hello\nthis is some\ninput from three lines\n"
	args := []string{"testdata/three_lines.txt"}
	fakeInput := bytes.NewBufferString(want)
	fakeTerminal := bytes.NewBuffer([]byte{})
	delayer, err := delay.NewDelayer(
		delay.WithArgs(args),
		delay.WithInput(fakeInput),
		delay.WithOutput(fakeTerminal),
		delay.WithDelay(time.Duration(0)),
	)
	if err != nil {
		t.Error(err)
	}
	delayer.Print()
	got := fakeTerminal.String()

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
