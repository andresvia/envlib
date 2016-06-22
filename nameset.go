package envlib

// A map that is used as a set
type Nameset map[string]int

// Create a new Nameset
func NewNameset(names ...string) (ns Nameset) {
	ns = map[string]int{}
	for _, name := range names {
		ns[name] = 0
	}
	return
}
