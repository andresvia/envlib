// Go package to manipulate the environment
package envlib

// Selects the Nameset of variables from the envlib environment
func Select(names Nameset) {
	new_env := map[string]string{}
	new_env_list := []string{}
	for _, name := range env_list {
		value := env[name]
		if _, found := names[name]; found {
			new_env[name] = value
			new_env_list = append(new_env_list, name)
		}
	}
	env = new_env
	env_list = new_env_list
}
