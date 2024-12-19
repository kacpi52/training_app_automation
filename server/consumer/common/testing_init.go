package common

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(filepath.Dir(filename), "../../../../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}