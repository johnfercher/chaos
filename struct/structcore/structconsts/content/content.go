package content

type Type string

const (
	Go         Type = "go"
	GoMod      Type = "mod"
	Dockerfile Type = "dockerfile"
	Makefile   Type = "makefile"
	Unknown    Type = "unknown"
)
