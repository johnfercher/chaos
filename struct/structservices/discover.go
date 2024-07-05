package structservices

import (
	"github.com/johnfercher/chaos/struct/regex"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/content"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/file"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"github.com/johnfercher/chaos/struct/structcore/structservices"
	"log"
	"os"
)

const goMod = "/go.mod"

type Discover struct {
	loader               structservices.File
	fileClassifier       structservices.FileClassifier
	files                []structmodels.File
	interfaceInterpreter structservices.InterfaceInterpreter
	functionInterpreter  structservices.FunctionInterpreter
}

func NewDiscover(loader structservices.File, fileClassifier structservices.FileClassifier,
	interfaceInterpreter structservices.InterfaceInterpreter, functionInterpreter structservices.FunctionInterpreter) Discover {
	return Discover{
		loader:               loader,
		fileClassifier:       fileClassifier,
		interfaceInterpreter: interfaceInterpreter,
		functionInterpreter:  functionInterpreter,
	}
}

func (d *Discover) Project(path string) ([]structmodels.File, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		innerFileDirType := file.File
		if e.IsDir() {
			innerFileDirType = file.Dir
		}

		if innerFileDirType == file.Dir {
			err := d.findDir(path+"/"+e.Name(), e.Name(), innerFileDirType)
			if err != nil {
				return nil, err
			}
		} else {
			filePath := path + "/" + e.Name()
			fileContent, err := d.loader.Read(filePath)
			if err != nil {
				return nil, err
			}

			file := structmodels.File{
				Name:        e.Name(),
				Path:        path,
				Type:        file.File,
				ContentType: d.fileClassifier.Classify(fileContent),
				Content:     fileContent,
				Imports:     regex.GetImports(fileContent),
				Package:     regex.GetPackageName(fileContent),
				Interfaces:  d.interfaceInterpreter.Interpret(fileContent),
				Functions:   d.functionInterpreter.Interpret(fileContent),
			}
			d.files = append(d.files, file)
		}
	}

	var files []structmodels.File
	for _, file := range d.files {
		if file.ContentType == content.Go || file.Name == "go.mod" {
			files = append(files, file)
		}
	}

	return files, nil
}

func (d *Discover) findDir(path string, name string, fileDirType file.Type) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		innerFileDirType := file.File
		if e.IsDir() {
			innerFileDirType = file.Dir
		}

		if innerFileDirType == file.Dir {
			err := d.findDir(path+"/"+e.Name(), e.Name(), innerFileDirType)
			if err != nil {
				return err
			}
		} else {
			filePath := path + "/" + e.Name()
			fileContent, err := d.loader.Read(filePath)
			if err != nil {
				return err
			}
			file := structmodels.File{
				Name:        e.Name(),
				Path:        path,
				Type:        file.File,
				ContentType: d.fileClassifier.Classify(fileContent),
				Imports:     regex.GetImports(fileContent),
				Content:     fileContent,
				Package:     regex.GetPackageName(fileContent),
				Interfaces:  d.interfaceInterpreter.Interpret(fileContent),
				Functions:   d.functionInterpreter.Interpret(fileContent),
			}
			d.files = append(d.files, file)
		}

	}

	return nil
}
