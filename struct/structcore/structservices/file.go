package structservices

type File interface {
	Read(pathName string) (string, error)
	Write(pathName string, content string) error
}
