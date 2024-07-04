package services

type Orchestrator interface {
	Generate(filePath string, name string) error
}
