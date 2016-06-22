// Go package to manipulate the environment
package envlib

import (
	"strings"
)

var (
	env      map[string]string = map[string]string{} // the environment as a map
	env_list []string          = []string{}          // the sorted variable list
)

// Gets the named variable from the envlib environment
func Getenv(name string) string {
	return env[name]
}

// Initializes the envlib with the supplied environ
func Init(environ []string) {
	env = map[string]string{}
	env_list = []string{}
	for _, e := range environ {
		nameval := strings.SplitN(e, "=", 2)
		if len(nameval) == 2 {
			name := nameval[0]
			val := nameval[1]
			env[name] = val
			env_list = append(env_list, name)
		} else {
			panic("Passed environ is not valid, error found on '" + e + "'")
		}
	}
}

// Returns the envlib environment in standard representation, removing duplicates, ignoring empty variables
func Environ() (environ []string) {
	names := map[string]int{}
	rev_env_list := []string{}
	for i := len(env_list) - 1; i >= 0; i-- {
		name := env_list[i]
		if _, found := names[name]; !found {
			names[name] = 0
			rev_env_list = append(rev_env_list, name)
		}
	}
	for i := len(rev_env_list) - 1; i >= 0; i-- {
		name := rev_env_list[i]
		value := env[name]
		if value != "" {
			environ = append(environ, name+"="+env[name])
		}
	}
	return
}
