package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	DryRun bool
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer) *Session {
	return &Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
}

func (s *Session) Run() {
	input := bufio.NewReader(s.Stdin)
	for {
		fmt.Fprint(s.Stdout, "> ")
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprint(s.Stdout, "\nBe seeing you!\n")
			break
		}

		cmd, err := CmdFromString(line)
		if err != nil {
			continue
		}
		if s.DryRun {
			fmt.Fprint(s.Stdout, line)
			continue
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprint(s.Stdout, err)
		}
		fmt.Fprintf(s.Stdout, "%s", out)

	}
}

func RunCLI() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
}

func CmdFromString(input string) (*exec.Cmd, error) {
	cmds := strings.Fields(input)
	if len(cmds) == 0 {
		return nil, errors.New("expected command to be non empty, it was empty")
	}
	base := cmds[0]
	return exec.Command(base, cmds[1:]...), nil
}
