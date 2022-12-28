package counter

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// Original
type LineCounter struct {
	Input io.Reader
}

func NewLineCounter() *LineCounter {
	return &LineCounter{
		Input: os.Stdin,
	}
}

func (l *LineCounter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(l.Input)
	for scanner.Scan() {
		lines += 1
	}
	return lines
}

func OldLines() int {
	return NewLineCounter().Lines()
}

// Using functional options

type Counter struct {
	input   io.Reader
	output  io.Writer
	matcher func(string, string) bool
	needle  string
}

type option func(*Counter) error

func NewCounter(opts ...option) (Counter, error) {
	c := Counter{
		input:   os.Stdin,
		output:  os.Stdout,
		matcher: func(_, _ string) bool { return true },
	}
	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return Counter{}, err
		}
	}
	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *Counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *Counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}

		c.output = output
		return nil
	}
}

func WithMatcher(matcher func(string, string) bool, needle string) option {
	return func(c *Counter) error {
		c.matcher = matcher
		c.needle = needle
		return nil
	}
}

func (c Counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		if c.matcher(scanner.Text(), "hello") {
			lines += 1
		}
	}
	return lines
}

func Lines() int {
	c, err := NewCounter()
	if err != nil {
		panic("internal error")
	}
	return c.Lines()
}
