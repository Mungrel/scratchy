package lang

import (
	"os"
	"os/exec"
	"path"
)

type ScratchEnvGenerator interface {
	Generate(dir string) error
	MainFile() string
}

func create(dir, name, content string) error {
	f, err := os.Create(path.Join(dir, name))
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		return err
	}

	return nil
}

func command(dir, name string, args ...string) error {
	c := exec.Command(name, args...)
	c.Dir = dir
	return c.Run()
}
