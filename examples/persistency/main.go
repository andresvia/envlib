package main

import (
	"fmt"
	"github.com/andresvia/envlib"
	"log"
	"os"
)

//go:generate go run main.go ../../testdata/etc_environment_example ../../testdata/generated/persistency_example
func main() {
	fmt.Println("persistency example")
	if len(os.Args) != 3 {
		log.Fatal("incorrect number of arguments")
	}
	envlib.Init([]string{"abc=def"})
	envlib.MergeWith(os.Args[1])
	envlib.Select(envlib.NewNameset("ENVLIB_TEST1", "abc"))
	envlib.Persist(os.Args[2])
	fmt.Println("end")
}
