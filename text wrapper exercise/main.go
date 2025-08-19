package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const fileName string = "story.txt"
	limit := 40
	story, _ := ioutil.ReadFile(fileName)
	index := 0
	var output []byte
	for _, value := range string(story) {
		if value == '\n' {
			index = 0
		}
		if index >= limit && value == ' ' {
			output = append(output, '\n')
			index = 0
		}

		output = append(output, []byte(string(value))...)
		index++
	}
	err := ioutil.WriteFile("new_story.txt", output, 0644)
	if err != nil {
		fmt.Print(err)
		return
	}

}
