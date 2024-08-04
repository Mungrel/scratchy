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
	"go": lang.GoGenerator{},
	"ts": lang.TypescriptGenerator{},
}

func main() {
	var noCode bool

	flag.BoolVar(&noCode, "no-code", false, "don't open VS code")
	flag.Parse()

	args := flag.Args()

	if err := run(noCode, args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(noCode bool, args []string) error {
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
	if !noCode {
		if err := exec.Command("code", "-g", mainPath, tmpDir).Run(); err != nil {
			return errors.New("couldn't open VS Code")
		}
	}

	return nil
}
