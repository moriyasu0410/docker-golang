package main

import (
	"fmt"
	"github.com/mediba-moritake/docker-golang/hello/lib"
)

func main() {
	hello := lib.Hello{Language: "Golang"}
	fmt.Println(hello.Hello())
}
