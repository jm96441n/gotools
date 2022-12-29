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

func TestWithInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := counter.NewCounter(counter.WithInputFromArgs(args))
	if err != nil {
		t.Error(err)
	}

	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestWithInputFromArgsEmpty(t *testing.T) {
	t.Parallel()
	args := []string{}
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := counter.NewCounter(counter.WithInput(inputBuf), counter.WithInputFromArgs(args))
	if err != nil {
		t.Error(err)
	}

	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("Got %d, want %d", got, want)
	}
}
