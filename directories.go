package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
)

func directoryExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

const basePath string = "input/days"

func generateDirs() {
	b, err := directoryExists(basePath)
	check(err)
	if b {
		fmt.Println("Input directories already generated")
		return
	}

	err = os.MkdirAll(basePath, 0755)
	check(err)

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	for i := 0; i < 25; i++ {
		dayDir := fmt.Sprintf("%s/%d", basePath, i)
		err := os.Mkdir(dayDir, 0755)
		check(err)
		createEmptyFile(fmt.Sprintf("%s/example.txt", dayDir))
		createEmptyFile(fmt.Sprintf("%s/input.txt", dayDir))
	}

	c, err := os.ReadDir(basePath)
	check(err)

	fmt.Printf("Listing %s\n", basePath)
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	randomDay := rand.Intn(24) + 1
	randomDirectory := fmt.Sprintf("%s/%d", basePath, randomDay)
	err = os.Chdir(randomDirectory)
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Printf("Listing random subdirectory %s\n", randomDirectory)
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Printf("Visiting content of the main sub-dir %s\n", basePath)
	err = filepath.WalkDir(basePath, visit)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
