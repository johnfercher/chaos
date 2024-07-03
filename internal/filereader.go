package internal

import (
	"os"
)

type File struct {
	workDir string
}

func NewFileReader(workDir string) *File {
	return &File{
		workDir: workDir,
	}
}

func (f *File) Load(pathName string) (string, error) {
	dir := f.workDir + pathName

	file, err := os.ReadFile(dir)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
