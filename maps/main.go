package main

import (
	"fmt"
)

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#F00",
		"blue":  "#00F",
		"green": "#0F0",
	}

	colors["white"] = "#FFF"

	printMap(colors)
}

func printMap(c map[string]string) {
	for c, v := range c {
		fmt.Println(c, v)
	}
}
