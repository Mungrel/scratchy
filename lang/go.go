package lang

type GoGenerator struct{}

var _ ScratchEnvGenerator = GoGenerator{}

func (g GoGenerator) Generate(dir string) error {
	if err := create(dir, "main.go",
		`package main

import (
	"fmt"
	"os"
)
	
func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
	
func run(args []string) error {
	return nil
}
`,
	); err != nil {
		return err
	}

	if err := command(dir, "go", "mod", "init", "scratch"); err != nil {
		return err
	}

	return nil
}

func (g GoGenerator) MainFile() string {
	return "main.go"
}
