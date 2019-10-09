package commands

type ResolvePath struct {
	BaseCommand
	Path string
}

func NewResolvePath(path string) *ResolvePath {
	return &ResolvePath{
		BaseCommand: *NewBaseCommand("ResolvePath"),
		Path:        path,
	}
}
