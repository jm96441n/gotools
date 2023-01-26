package findgo

import (
	"io/fs"
	"path/filepath"
)

func Files(fileSys fs.FS) int {
	var count int

	fs.WalkDir(fileSys, ".", func(filename string, d fs.DirEntry, err error) error {
		if filepath.Ext(filename) == ".go" {
			count++
		}
		return nil
	})
	return count
}
