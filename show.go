package main

import (
	"log"
	"os"
)

func show(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return string(data), nil
}
