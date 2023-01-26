package writer_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jm96441n/gotools/writer"
)

func TestWriteToFile(t *testing.T) {
	t.Parallel()
	path := "testdata/write_test.txt"
	_, err := os.Stat(path)
	if err == nil {
		t.Fatalf("test artifact not cleaned up: %q", path)
	}
	defer os.Remove(path)
	want := []byte{1, 2, 3}
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	perm := stat.Mode().Perm()
	if perm != 0600 {
		t.Errorf("want file mode 0600, got 0%o", perm)
	}
}

func TestWriteToFileOverwritesExistingFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "write_overwite_test.txt"

	err := os.WriteFile(path, []byte("this should not make it into the file\n"), 0600)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{1, 2, 3}
	err = writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestWriteFilePermsClosed(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/perms_test.txt"
	err := os.WriteFile(path, []byte{}, 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = writer.WriteToFile(path, []byte{})
	if err != nil {
		t.Fatal(err)
	}

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	perms := stat.Mode().Perm()

	if perms != 0600 {
		t.Errorf("expected permissions to be closed to 0600, stayed as 0%o", perms)
	}
}

func TestWriteZeroes(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		count         int
		expectedZeros []byte
	}{
		"no zeroes": {
			count:         0,
			expectedZeros: []byte{},
		},
		"one zero": {
			count:         1,
			expectedZeros: []byte{0},
		},
		"many zeroes": {
			count: 1000,
			expectedZeros: func() []byte {
				zeros := make([]byte, 0, 1000)
				for i := 0; i < 1000; i++ {
					zeros = append(zeros, byte(0))
				}
				return zeros
			}(),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			path := t.TempDir() + fmt.Sprintf("/%s.txt", strings.ReplaceAll(name, " ", "_"))
			err := writer.WriteZerosToFile(path, tc.count)
			if err != nil {
				t.Fatal(err)
			}

			got, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(tc.expectedZeros, got) {
				t.Errorf(cmp.Diff(tc.expectedZeros, got))
			}
		})
	}
}
