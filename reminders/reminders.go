package reminders

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Reminder struct {
	args    []string
	storage io.ReadWriteCloser
	output  io.Writer
}

type option func(*Reminder) error

var (
	ErrNilStorage      = errors.New("nil storage")
	ErrNilOutputWriter = errors.New("nil output writer")
)

func NewReminder(opts ...option) (Reminder, error) {
	storage, err := os.OpenFile("reminders.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if len(os.Args) == 1 {
		_, err := storage.Seek(0, io.SeekStart)
		if err != nil {
			return Reminder{}, err
		}
	}
	if err != nil {
		return Reminder{}, err
	}
	r := Reminder{
		storage: storage,
		args:    os.Args[1:],
		output:  os.Stdout,
	}
	for _, opt := range opts {
		err = opt(&r)
		if err != nil {
			return Reminder{}, err
		}
	}
	return r, nil
}

func WithStorage(storage io.ReadWriteCloser) option {
	return func(r *Reminder) error {
		if storage == nil {
			return ErrNilStorage
		}
		r.storage = storage
		return nil
	}
}

func WithArgs(args []string) option {
	return func(r *Reminder) error {
		r.args = args
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(r *Reminder) error {
		if output == nil {
			return ErrNilOutputWriter
		}
		r.output = output
		return nil
	}
}

func (r Reminder) Remind() {
	if len(r.args) == 0 {
		r.displayReminders()
		return
	}

	r.addReminder()
}

func (r Reminder) addReminder() {
	defer r.storage.Close()
	task := strings.Join(r.args, " ")
	r.storage.Write([]byte(fmt.Sprintf("%s\n", task)))
}

func (r Reminder) displayReminders() {
	defer r.storage.Close()
	scanner := bufio.NewScanner(r.storage)
	for scanner.Scan() {
		task := fmt.Sprintf("%s\n", scanner.Text())
		_, err := r.output.Write([]byte(task))
		if err != nil {
			panic(err)
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}
}

// func (r Reminder) storeReminder() {}
func Remind() {
	r, err := NewReminder()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	r.Remind()
}
