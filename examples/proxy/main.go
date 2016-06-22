package main

import (
	"fmt"
	"github.com/andresvia/envlib"
	"log"
	"os"
)

//go:generate go run main.go ../../testdata/etc_environment_proxy_example
func main() {
	fmt.Println("proxy example")
	if len(os.Args) != 2 {
		log.Fatal("incorrect number of arguments")
	}
	proxy_vars := []string{"HTTP_PROXY", "http_proxy", "HTTPS_PROXY", "https_proxy", "NO_PROXY", "no_proxy"}
	environ := []string{}
	for _, proxy_var := range proxy_vars {
		environ = append(environ, proxy_var+"="+os.Getenv(proxy_var))
	}
	envlib.Init(environ)
	envlib.MergeWith(os.Args[1])
	fmt.Println(envlib.ToString())
	fmt.Println("end")
}
