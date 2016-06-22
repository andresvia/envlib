package envlib

import (
	"github.com/joho/godotenv"
)

// Merges the envlib environment with a file, the envlib environment have preference over the file
func MergeWith(file string) (err error) {
	current_env := env
	current_env_list := env_list
	env, err = godotenv.Read(file)
	env_list = []string{}
	for name, _ := range env {
		env_list = append(env_list, name)
	}
	for _, name := range current_env_list {
		env_list = append(env_list, name)
	}
	for name, value := range current_env {
		env[name] = value
	}
	return
}
