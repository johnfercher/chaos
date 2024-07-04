package structservices

import "github.com/johnfercher/chaos/struct/structcore/structconsts/content"

type FileClassifier interface {
	Classify(fileContent string) content.Type
}
