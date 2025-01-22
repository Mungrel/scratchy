package lang

type JSONGenerator struct{}

var _ ScratchEnvGenerator = JSONGenerator{}

func (j JSONGenerator) Generate(dir string) error {
	if err := create(dir, "data.json",
		`{
	"hello": "world"
}
`); err != nil {
		return err
	}

	return nil
}

func (j JSONGenerator) MainFile() string {
	return "data.json"
}
