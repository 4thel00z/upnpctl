package utils

import (
	"fmt"
	"os"
)

func Cry(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
