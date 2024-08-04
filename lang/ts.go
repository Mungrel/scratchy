package lang

type TypescriptGenerator struct{}

var _ ScratchEnvGenerator = TypescriptGenerator{}

func (g TypescriptGenerator) Generate(dir string) error {
	if err := create(dir, "main.ts",
		`// tsx was installed globally, run me with 'tsx main.ts'.
const main = async () => {
    console.log('hey');
}

main();
`); err != nil {
		return err
	}

	if err := command(dir, "npm", "install", "@types/node"); err != nil {
		return err
	}

	if err := command(dir, "npm", "install", "-g", "tsx"); err != nil {
		return err
	}

	return nil
}

func (g TypescriptGenerator) MainFile() string {
	return "main.ts"
}
