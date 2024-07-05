package services

type GenerationOrchestrator interface {
	Generate(filePath string, name string) error
}
