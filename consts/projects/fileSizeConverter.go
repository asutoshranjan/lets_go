package projects

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func FileSizeConverter() {
	filesize := 4000000000
	fmt.Printf("%.2f GB \n", float32(filesize)/GB)
}
