package lang

type GoGenerator struct{}

var _ ScratchEnvGenerator = GoGenerator{}

func (g GoGenerator) Generate(dir string) error {
	if err := create(dir, "main.go",
		`package main

import "fmt"
	
func main() {
	fmt.Println("hey")
}`,
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
