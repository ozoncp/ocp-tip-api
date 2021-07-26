package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("This is Ozon Tip API")

	openFile := func(path string) error {
		file, openingErr := os.Open(path)
		if openingErr != nil {
			return openingErr
		}
		defer func() {
			closingErr := file.Close()
			if closingErr != nil {
				log.Fatal(closingErr)
			}
		}()
		return nil
	}
	for i := 0; i < 10; i++ {
		if err := openFile("config.txt"); err != nil {
			log.Fatal(err)
		}
	}
}
