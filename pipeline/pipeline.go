package pipeline

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Pipeline struct {
	Output io.Writer
	Reader io.Reader
	Error  error
}

func New() *Pipeline {
	return &Pipeline{
		Output: os.Stdout,
	}
}

// Input functions

func FromString(in string) *Pipeline {
	p := New()
	p.Reader = strings.NewReader(in)
	return p
}

func FromFile(path string) *Pipeline {
	f, err := os.Open(path)
	if err != nil {
		return &Pipeline{Error: err}
	}
	p := New()
	p.Reader = f
	return p
}

// Filter Functions

func (p *Pipeline) Column(colNum int) *Pipeline {
	if p.Error != nil {
		p.Reader = strings.NewReader("")
		return p
	}

	if colNum < 1 {
		p.Error = fmt.Errorf("column number must be a positive number greater than zero, you provided %d", colNum)
		p.Reader = strings.NewReader("")
		return p
	}

	scanner := bufio.NewScanner(p.Reader)
	result := bytes.NewBuffer([]byte{})

	for scanner.Scan() {
		row := scanner.Text()
		cols := strings.Fields(row)
		if len(cols) < colNum {
			continue
		}
		_, err := result.Write([]byte(fmt.Sprintf("%s\n", cols[colNum-1])))
		if err != nil {
			p.Error = err
			break
		}
	}
	p.Reader = result
	return p
}

// Output functions

func (p *Pipeline) String() (string, error) {
	if p.Error != nil {
		return "", p.Error
	}
	data, err := io.ReadAll(p.Reader)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func (p *Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Reader)
}
