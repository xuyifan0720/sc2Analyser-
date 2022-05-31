package main

import (
	"fmt"
	"io/ioutil"

	"github.com/SC2Analyser/Analyser"
)

func main() {
	files, err := ioutil.ReadDir("../../../../Library")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error)
	}
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
	err = Analyser.Analyse()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error)
	}
}
