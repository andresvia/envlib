package envlib

import (
	"strings"
)

// Converts the envlib environment to the typical representation for files
func ToString() (s string) {
	s = strings.Join(Environ(), "\n")
	return
}
