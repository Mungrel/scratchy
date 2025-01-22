package lang

type ShGenerator struct{}

var _ ScratchEnvGenerator = ShGenerator{}

func (s ShGenerator) Generate(dir string) error {
	if err := create(dir, "run.sh",
		`#!/bin/bash
	
echo "hello"
`); err != nil {
		return err
	}

	if err := command(dir, "chmod", "+x", "run.sh"); err != nil {
		return err
	}

	return nil
}

func (s ShGenerator) MainFile() string {
	return "run.sh"
}
