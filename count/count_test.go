package count_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/jm96441n/gotools/count"
)

func TestLinesCountsTheLinesFromTheInput(t *testing.T) {
	t.Parallel()
	fakeInput := bytes.NewBuffer([]byte{})
	fakeInput.Write([]byte("One\nTwo\nThree\n"))
	c, err := count.NewCounter(count.WithInput(io.Reader(fakeInput)))
	if err != nil {
		t.Error(err)
	}

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
	c, err := count.NewCounter(count.WithInput(fakeInput))
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
	c, err := count.NewCounter(count.WithInputFromArgs(args))
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
	c, err := count.NewCounter(count.WithInput(inputBuf), count.WithInputFromArgs(args))
	if err != nil {
		t.Error(err)
	}

	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestWithInputFromArgsMultiple(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt", "testdata/two_lines.txt"}
	c, err := count.NewCounter(count.WithInputFromArgs(args))
	if err != nil {
		t.Error(err)
	}

	want := 5
	got := c.Lines()

	if want != got {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestWordsCountsTheWordsFromTheInput(t *testing.T) {
	t.Parallel()
	fakeInput := bytes.NewBufferString("1\n2 words\n3 this time\n")
	c, err := count.NewCounter(count.WithInput(io.Reader(fakeInput)))
	if err != nil {
		t.Error(err)
	}

	got := c.Words()
	want := 6
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
