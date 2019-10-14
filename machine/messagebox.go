package machine

type MessageBoxMode uint64

const (
	NoButtons MessageBoxMode = iota
	CloseOnly
	OkOnly
	OkCancel
)

type MessageBox struct {
	Mode         MessageBoxMode
	Title        string
	Message      string
	AxisControls []uint8
	Seq          int64
}
