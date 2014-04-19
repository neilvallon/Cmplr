package project

type Job struct {
	Type string

	Outputfile string
	InputFiles []string

	Options map[string]string
}
