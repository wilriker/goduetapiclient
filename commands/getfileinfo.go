package commands

type GetFileInfo struct {
	BaseCommand
	FileName string
}

func NewGetFileInfo(fileName string) *GetFileInfo {
	return &GetFileInfo{
		BaseCommand: *NewBaseCommand("GetFileInfo"),
		FileName:    fileName,
	}
}
