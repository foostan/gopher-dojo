package main

import (
	"fmt"
	"github.com/foostan/gopher-dojo/imgconverter"
	"os"
)

func main(){
	if err := imgconverter.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

