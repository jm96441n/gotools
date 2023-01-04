package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// Using functional options

type Counter struct {
	input      []io.ReadCloser
	output     io.Writer
	matcher    func(string, string) bool
	needle     string
	countWords bool
}

type option func(*Counter) error

func NewCounter(opts ...option) (Counter, error) {
	c := Counter{
		input:   []io.ReadCloser{io.NopCloser(os.Stdin)},
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
		var inputCloser io.ReadCloser
		var ok bool
		if inputCloser, ok = input.(io.ReadCloser); !ok {
			inputCloser = io.NopCloser(input)
		}
		c.input = []io.ReadCloser{inputCloser}
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

func FromArgs(args []string) option {
	return func(c *Counter) error {
		if len(args) == 0 {
			return nil
		}
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		wordCount := fset.Bool("w", false, "Count words instead of lines")
		fset.SetOutput(c.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		c.countWords = *wordCount
		args = fset.Args()
		inputs := make([]io.ReadCloser, len(args))
		for idx, fileName := range args {
			f, err := os.Open(fileName)
			if err != nil {
				return err
			}
			inputs[idx] = f
		}
		c.input = inputs
		return nil
	}
}

func (c Counter) Lines() int {
	lines := 0
	for _, in := range c.input {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			if c.matcher(scanner.Text(), c.needle) {
				lines += 1
			}
		}
		in.Close()
	}
	return lines
}

func Lines() int {
	c, err := NewCounter(FromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Lines()
}

func (c Counter) Words() int {
	words := 0
	scanner := bufio.NewScanner(c.input[0])
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words += 1
	}
	return words
}

func Words() int {
	c, err := NewCounter(FromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

	}

	return c.Words()
}

func RunCLI() int {
	counter, err := NewCounter(FromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if counter.countWords {
		return counter.Words()
	}
	return counter.Lines()
}

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
