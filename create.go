package main

import (
	"log"
	"os"
)

func create(pattern string, content string) (string, error) {
	f, err := os.OpenFile(pattern, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(content)

	if err != nil {
		log.Fatal(err.Error())
	}

	return "successfully Created", nil
}
