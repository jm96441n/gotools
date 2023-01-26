package findgo_test

import (
	"testing"
	"testing/fstest"

	"github.com/jm96441n/gotools/findgo"
)

func TestCountGoFiles(t *testing.T) {
	t.Parallel()
	fileSys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := 4
	got := findgo.Files(fileSys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
