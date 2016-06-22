package main

import (
	"fmt"
	"github.com/andresvia/envlib"
	"log"
	"os"
)

//go:generate go run main.go ../../testdata/etc_environment_example
func main() {
	fmt.Println("simple example")
	if len(os.Args) != 2 {
		log.Fatal("incorrect number of arguments")
	}
	envlib.Init(os.Environ())
	envlib.Select(envlib.NewNameset("GOFILE", "GOPACKAGE"))
	envlib.MergeWith(os.Args[1])
	fmt.Println(envlib.ToString())
	fmt.Println("end")
}
