package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"path"
	"time"
)

//go:embed logs
var f embed.FS

func main() {
	files, err := f.ReadDir("logs")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

		if err := printFile(path.Join("logs", file.Name())); err != nil {
			panic(err)
		}
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func printFile(path string) error {
	file, err := f.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(50 * time.Millisecond)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
