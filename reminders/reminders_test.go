package reminders_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/jm96441n/gotools/reminders"
)

func nopCloser(r io.ReadWriter) io.ReadWriteCloser {
	return nopcloser{r}
}

type nopcloser struct {
	io.ReadWriter
}

func (n nopcloser) Close() error { return nil }

func TestRemindWritesInputToStorageWhenArgsArePresent(t *testing.T) {
	args := []string{"buy milk"}
	fakeStorageBuf := bytes.NewBuffer([]byte{})
	fakeStorage := nopCloser(fakeStorageBuf)
	reminder, err := reminders.NewReminder(
		reminders.WithArgs(args),
		reminders.WithStorage(io.ReadWriteCloser(fakeStorage)),
	)
	if err != nil {
		t.Error(err)
	}
	reminder.Remind()
	want := fmt.Sprintf("%s\n", args[0])
	got := fakeStorageBuf.String()
	if want != got {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestRemindWritesInputToStorageWhenArgsArePresentWithMultipleWrites(t *testing.T) {
	args := []string{"buy milk", "buy cheese", "change laundry"}
	wantBuilder := strings.Builder{}
	fakeStorageBuf := bytes.NewBuffer([]byte{})
	fakeStorage := nopCloser(fakeStorageBuf)
	for _, arg := range args {
		reminder, err := reminders.NewReminder(
			reminders.WithArgs([]string{arg}),
			reminders.WithStorage(io.ReadWriteCloser(fakeStorage)),
		)
		if err != nil {
			t.Error(err)
		}
		reminder.Remind()
		wantBuilder.Write([]byte(fmt.Sprintf("%s\n", arg)))
	}
	got := fakeStorageBuf.String()
	want := wantBuilder.String()
	if want != got {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestRemindDisplaysAllRemindersWhenCalledWithEmptyArgs(t *testing.T) {
	want := "buy milk\nbuy eggs\nchange laundry\n"
	fakeStorageBuf := bytes.NewBuffer([]byte(want))
	fakeStorage := nopCloser(fakeStorageBuf)
	fakeTerminal := bytes.NewBuffer([]byte{})
	reminder, err := reminders.NewReminder(
		reminders.WithArgs([]string{}),
		reminders.WithStorage(fakeStorage),
		reminders.WithOutput(fakeTerminal),
	)
	if err != nil {
		t.Error(err)
	}
	reminder.Remind()
	got := fakeTerminal.String()
	if want != got {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
