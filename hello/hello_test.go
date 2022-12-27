package hello_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/jm96441n/gotools/hello"
)

func TestPrintsHelloMessageToWriterWithName(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	fakeReader := &bytes.Buffer{}
	fakeReader.Write([]byte("John\n"))
	p := hello.NewPrinter()
	p.Output = fakeTerminal
	p.Input = fakeReader

	p.PrintGreeting()
	want := "hello John\n"
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestPrintTimePrintsTheTimeInTheRightFormat(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		curTime string
		want    string
	}{
		{
			name:    "Time afternoon",
			curTime: "01/02 04:05:05PM '06 -0700",
			want:    "It's 5 minutes past 4\n",
		},
		{
			name:    "Time morning",
			curTime: "01/02 04:05:05AM '06 -0700",
			want:    "It's 5 minutes past 4\n",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockTimeFn := func() time.Time {
				t, _ := time.Parse(time.Layout, tc.curTime)
				return t
			}
			fakeTerminal := bytes.NewBuffer([]byte{})
			p := hello.NewPrinter()
			p.Output = fakeTerminal
			p.TimeFn = mockTimeFn

			p.PrintTime()
			got := fakeTerminal.String()

			if tc.want != got {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}
