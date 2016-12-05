package lib

import "testing"

func TestHello(t *testing.T) {
	hello := Hello{Language: "Golang"}

	if hello.Hello() != "Hello, Golang" {
		t.Fatal("Hello() doesn't match")
	}
}