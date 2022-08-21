package interfaces

import (
	"fmt"
)

func InterfaceInGo() {
	fmt.Println("InterFace In Go!!")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World this is Go!!"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
