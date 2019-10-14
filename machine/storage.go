package machine

type Storage struct {
	Mounted   bool
	Speed     uint64
	Capacity  uint64
	Free      uint64
	OpenFiles uint64
	Path      string
}
