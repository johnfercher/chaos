package services

import (
	"os"
)

type File struct{}

func NewFile() *File {
	return &File{}
}

func (f *File) Read(pathName string) (string, error) {
	file, err := os.ReadFile(pathName)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func (f *File) Write(pathName string, content string) error {
	return os.WriteFile(pathName, []byte(content), os.ModePerm)
}
