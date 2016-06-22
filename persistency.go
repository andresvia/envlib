package envlib

import (
	"os"
)

// Store the envlib environment to a file
func Persist(file string) (err error) {
	var fd *os.File
	if fd, err = os.Create(file); err == nil {
		defer fd.Close()
		fd.Write([]byte(ToString()))
	}
	return
}
