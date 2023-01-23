package writer_test

import (
	"testing"

	"github.com/jm96441n/goTools/writer"
)

func TestWriteToFile(t *testing.T) {
	t.Parallel()
	path := "testdata/write_test.txt"
	want := []byte{1, 2, 3}
	err := writer.WriteToFile(path, want)
}
