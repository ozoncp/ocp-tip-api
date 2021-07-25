package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is Ozon Tip API")

	openFile := func(path string) error {
		file, err := os.Open(path)
		defer file.Close()
		return err
	}
	for i := 0; i < 10; i++ {
		openFile("config.txt")
	}
}
