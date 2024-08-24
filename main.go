package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	runner()
}

func runner() string {

	var fileName string
	var read string
	var content string

goBack:
	fmt.Printf("Do you want to list, create, show or delete a file? (list/create/show/delete): ")
	fmt.Scanln(&read)

	// displays all files in the folder
	if read == "list" || read == "List" {
		list()
		goto goBack
	}

	fmt.Printf("Enter the file name: ")
	fmt.Scanln(&fileName)
	// Creates new file
	if read == "create" || read == "Create" {
		fmt.Print("Add file content: ")
		fmt.Scanln(&content)
		fmt.Printf(create(fileName+".txt", content))
	}
	// shows exiting file by the name provided
	if read == "show" || read == "Show" {
		_, err := os.ReadFile(fileName + ".txt")
		if err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Printf(show(fileName + ".txt"))

		goto goBack
	}

	if read == "delete" || read == "Delete" {
	}

	return "No command provided"
}
