package main

import (
	"fmt"
	"log"
	"os"
)

func list() {
	files, err := os.ReadDir("./")

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
