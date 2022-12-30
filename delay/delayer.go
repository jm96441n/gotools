package delay

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Delayer struct {
	input  []io.ReadCloser
	output io.Writer
	delay  time.Duration
}

var (
	ErrNilInputReader  = errors.New("nil input reader")
	ErrNilOutputWriter = errors.New("nil output writer")
)

type option func(*Delayer) error

func NewDelayer(opts ...option) (Delayer, error) {
	d := Delayer{
		input:  []io.ReadCloser{io.NopCloser(os.Stdin)},
		output: os.Stdout,
		delay:  time.Duration(500 * time.Millisecond),
	}
	for _, opt := range opts {
		err := opt(&d)
		if err != nil {
			return Delayer{}, err
		}
	}
	return d, nil
}

func WithInput(input io.Reader) option {
	return func(d *Delayer) error {
		if input == nil {
			return ErrNilInputReader
		}
		if closerInput, ok := input.(io.ReadCloser); ok {
			d.input = []io.ReadCloser{closerInput}
		} else {
			d.input = []io.ReadCloser{io.NopCloser(input)}
		}
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(d *Delayer) error {
		if output == nil {
			return ErrNilOutputWriter
		}
		d.output = output
		return nil
	}
}

func WithArgs(args []string) option {
	return func(d *Delayer) error {
		if len(args) == 0 {
			return nil
		}
		input := make([]io.ReadCloser, len(args))
		for idx, filename := range args {
			f, err := os.Open(filename)
			if err != nil {
				return err
			}
			input[idx] = f
		}
		d.input = input
		return nil
	}
}

func WithDelay(delay time.Duration) option {
	return func(d *Delayer) error {
		d.delay = delay
		return nil
	}
}

func (d Delayer) Print() {
	builder := strings.Builder{}
	for _, in := range d.input {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				break
			}
			builder.WriteString(fmt.Sprintf("%s\n", text))

		}
		in.Close()
	}
	text := builder.String()
	for _, c := range text {
		fmt.Fprintf(d.output, "%c", c)
		time.Sleep(d.delay)
	}
}

func Print() {
	d, err := NewDelayer(
		WithArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	d.Print()
}
