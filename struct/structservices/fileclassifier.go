package structservices

import (
	"errors"
	"github.com/johnfercher/chaos/struct/regex"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/content"
	"strings"
)

type FileClassifier struct{}

func NewFileClassifier() *FileClassifier {
	return &FileClassifier{}
}

func (c *FileClassifier) Classify(fileContent string) content.Type {
	if c.IsGoFile(fileContent) {
		return content.Go
	}

	return content.Unknown
}

func (c *FileClassifier) IsGoFile(content string) bool {
	_, err := c.getGoPackageName(content)
	if err != nil {
		return false
	}

	return true
}

func (c *FileClassifier) getGoPackageName(content string) (string, error) {
	name := regex.GetPackageName(content)
	if name == "" {
		return "", errors.New("invalid package")
	}

	name = strings.ReplaceAll(name, "package ", "")
	return name, nil
}
