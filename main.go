package main

import (
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
	if len(os.Args) < 2 {
		fmt.Println("no language provided")
		os.Exit(1)
	}

	arg := strings.TrimSpace(os.Args[1])
	lang := strings.ToLower(arg)
	gen, ok := validGenerators[lang]

	if !ok {
		langs := make([]string, 0, len(validGenerators))
		for k := range validGenerators {
			langs = append(langs, k)
		}
		fmt.Printf("invalid lang %q (valid langs: %s)\n", arg, strings.Join(langs, ", "))
		os.Exit(1)
	}

	tmpDir, err := os.MkdirTemp("", "scratch-env-"+lang+"-*")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := gen.Generate(tmpDir); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mainPath := path.Join(tmpDir, gen.MainFile())

	fmt.Printf("New %s env created at\n\n\t%s\n\n", lang, tmpDir)
	if err := exec.Command("code", "-g", mainPath, tmpDir).Run(); err != nil {
		fmt.Println("Couldn't open VS Code")
	}
}
