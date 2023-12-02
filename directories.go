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

const (
	nDays    = 25
	basePath = "input/days"
)

type dayInput struct {
	exampleOne string
	exampleTwo string
	mainInput  string
}

func newDayInput(exampleOne string, exampleTwo string, mainInput string) *dayInput {
	return &dayInput{
		exampleOne: exampleOne,
		exampleTwo: exampleTwo,
		mainInput:  mainInput,
	}
}

func generateInputPaths() [nDays]*dayInput {
	var inputs [nDays]*dayInput
	var exampleOne, exampleTwo, mainInput string
	for i := 0; i < nDays; i++ {
		dayDir := fmt.Sprintf("%s/%d", basePath, i+1)
		exampleOne = fmt.Sprintf("%s/example1.txt", dayDir)
		exampleTwo = fmt.Sprintf("%s/example2.txt", dayDir)
		mainInput = fmt.Sprintf("%s/input.txt", dayDir)
		inputs[i] = newDayInput(exampleOne, exampleTwo, mainInput)
	}
	return inputs
}

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

	inputPaths := generateInputPaths()
	for i, inputPath := range inputPaths {
		err := os.Mkdir(fmt.Sprintf("%s/%d", basePath, i+1), 0755)
		check(err)
		createEmptyFile(inputPath.exampleOne)
		createEmptyFile(inputPath.exampleTwo)
		createEmptyFile(inputPath.mainInput)
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
	check(err)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
