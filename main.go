package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"scratchy/lang"
)

var validGenerators = map[string]lang.ScratchEnvGenerator{
	"go":   lang.GoGenerator{},
	"ts":   lang.TypescriptGenerator{},
	"sh":   lang.ShGenerator{},
	"json": lang.JSONGenerator{},
}

func main() {
	var editor string

	flag.StringVar(&editor, "editor", "vs-code", "editor to open (none, vs-code (default), cursor)")
	flag.Parse()

	args := flag.Args()

	if err := run(editor, args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(editor string, args []string) error {
	langs := make([]string, 0, len(validGenerators))
	for k := range validGenerators {
		langs = append(langs, k)
	}

	if len(args) < 1 {
		return fmt.Errorf("no language provided (valid langs: %s)", strings.Join(langs, ", "))
	}

	arg := strings.TrimSpace(args[0])
	lang := strings.ToLower(arg)
	gen, ok := validGenerators[lang]

	if !ok {
		return fmt.Errorf("invalid lang %q (valid langs: %s)", arg, strings.Join(langs, ", "))
	}

	tmpDir, err := os.MkdirTemp("", "scratch-env-"+lang+"-*")
	if err != nil {
		return err
	}

	if err := gen.Generate(tmpDir); err != nil {
		return err
	}

	mainPath := path.Join(tmpDir, gen.MainFile())

	fmt.Printf("New %s env created at\n\n\t%s\n\n", lang, tmpDir)
	switch editor {
	case "vs-code":
		if err := exec.Command("code", "-g", mainPath, tmpDir).Run(); err != nil {
			return errors.New("couldn't open VS Code")
		}
	case "cursor":
		if err := exec.Command("cursor", "-g", mainPath, tmpDir).Run(); err != nil {
			return errors.New("couldn't open cursor")
		}
	case "none":
		return nil
	}

	return nil
}
