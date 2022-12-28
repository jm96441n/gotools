package counter_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/jm96441n/gotools/linecount/counter"
)

func TestLinesCountsTheLinesFromTheInput(t *testing.T) {
	t.Parallel()
	fakeInput := bytes.NewBuffer([]byte{})
	fakeInput.Write([]byte("One\nTwo\nThree\n"))
	c := counter.NewLineCounter()
	c.Input = io.Reader(fakeInput)

	got := c.Lines()
	want := 3
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestLinesWithFunctionalOptionsCountsLines(t *testing.T) {
	t.Parallel()
	fakeInput := bytes.NewBuffer([]byte{})
	fakeInput.Write([]byte("One\nTwo\nThree\n"))
	c, err := counter.NewCounter(counter.WithInput(fakeInput))
	if err != nil {
		t.Error(err)
	}

	got := c.Lines()
	want := 3
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
