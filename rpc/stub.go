package rpc

// Package rpc provides RPC functionality for the Leither project
type LApiStub struct {
	VarActStub
	FilesStub
}

type VarActStub struct {
	GetVar          func(sid, name string, args ...string) (any, error)
	GetVarByContext func(sid, name string, mapOpt map[string]string) (any, error)
	Act             func(sid, name string, args ...string) error
	GetGobVar       func(sid, name string, args ...string) ([]byte, error)
}

const (
	TFile = iota
	TDir
)

type LsLink struct {
	Name   string
	Hash   string
	Size   uint64
	Type   int //unixfs_pb.Data_DataType
	Target string
}

func (link *LsLink) IsDir() bool {
	return link.Type == TDir
}

type StatInfo struct {
	Hash           string
	Size           uint64
	CumulativeSize uint64
	Blocks         int
	Type           string
	WithLocality   bool
	Local          bool
	SizeLocal      uint64
}

// ndtype = "directory"
// ndtype = "file"
func (link *StatInfo) IsDir() bool {
	return link.Type == "directory"
}

type FilesStub struct {
	FilesCopy  func(sid, src, dst string, flush bool) error
	FilesLs    func(sid, ps string) ([]LsLink, error)
	FilesMkdir func(sid, ps string, flush bool) error
	FilesRm    func(sid, ps string, recursive /*force,*/, flush bool) error
	FilesMv    func(sid, src, dst string, flush bool) error
	FilesFlush func(sid, ps string) (string, error)
	FilesStat  func(sid, ps string) (*StatInfo, error)
}
