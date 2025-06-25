package rpc

import (
	"bytes"
	"encoding/gob"
	"os"
	"time"
)

// Package rpc provides RPC functionality for the Leither project
type LApiStub struct {
	AuthStub
	VarActStub
	MiMeiStub
	FilesStub
}
type MiMeiStub struct {
	MFileStub
	MMOpenUrl func(sid, ps string) (string, error)
	MMClose   func(mmsid string) error
}

// NOTE:这个sid是有有效期的
type LoginReply struct {
	Sid string
	Uid string
}

type AuthStub struct {
	LoginWithPPT func(strPPT string) (*LoginReply, error)
	Logout       func(sid, info string) error
}

type VarActStub struct {
	GetVar          func(sid, name string, args ...string) (any, error)
	GetVarByContext func(sid, name string, mapOpt map[string]string) (any, error)
	Act             func(sid, name string, args ...string) error
	GetGobVar       func(sid, name string, args ...string) ([]byte, error)
}

func (wa *VarActStub) GetVarObj(obj any, sid string, name string, args ...string) error {
	data, err := wa.GetGobVar(sid, name, args...)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(obj)
}

// 弥媒文件相关的操作
// 1、MFile ID的生成
// mid, err := MMCreate(sid, appid, ext, api.MM_File, right, opt)
//
// 2、读写操作
// 管理所有的文件，替换原来的lfile
// MFSetObject(sid, mid string, obj any) error
// MFGetObject(sid, mid string) (any, error)
// MFSetData(sid, mid string, start int64, data any) (count int, err error)
// MFGetData(sid, mid string, start int64, count int, ver int64) ([]byte, error)
// MFTruncate(sid, mid string, size int64) error
// FSCopy(sid, mid string, mid2 string, ver int64) error

// 4、版本
// MFGetVers(sid, mid string, start int64, count uint) ([]int64, error)
// MFDelVers(sid, mid string, start int64, count uint) error
// MFSetRelease(sid, mid string, ver int64) error
// MFDelRelease(sid, mid string, vers ...int64) error
// type FileInfo = pub.FileInfo

var _ os.FileInfo = FileInfo{}
var _ os.DirEntry = FileInfo{}

type FileInfo struct {
	FName    string    // base name of the file
	FSize    int64     // length in bytes for regular files; system-dependent for others
	FModTime time.Time // modification time
	FIsDir   bool      // abbreviation for Mode().IsDir()
	FMode    os.FileMode
}

func (fi FileInfo) Sys() interface{} {
	return &fi
}

func (fi FileInfo) Size() int64 {
	return fi.FSize
}

func (fi FileInfo) Name() string {
	return fi.FName
}

func (fi FileInfo) Mode() os.FileMode {
	return fi.FMode
}

func (fi FileInfo) ModTime() time.Time {
	return fi.FModTime
}

func (fi FileInfo) IsDir() bool {
	return fi.FIsDir
}

// 兼容目录DirEntry
func (fi FileInfo) Type() os.FileMode {
	return fi.FMode
}

func (fi FileInfo) Info() (os.FileInfo, error) {
	return fi, nil
}

type MFileStub struct {
	MFOpenByPath    func(sid, mid, path string, flag int) (string, error)
	MFOpenMacFile   func(sid, mid, fileid string) (string, error)
	MFOpenTempFile  func(sid string) (string, error)
	MFPath2TempID   func(sid, localpath string) (string, error) //这个函数有争议
	MFLocal2MacFile func(sid, mid, ps string) (string, error)
	MFTemp2MacFile  func(fsid, mid string) (string, error)
	MFTemp2Ipfs     func(fsid, mid string) (string, error)
	MFTemp2Files    func(fsid, ps string) (string, error)
	MFTruncate      func(fsid string, size int64) error
	MFSetObject     func(fsid string, obj any) error
	MFGetObject     func(fsid string) (any, error)
	MFSetData       func(fsid string, data []byte, start int64) (count int, err error)
	MFGetData       func(fsid string, start int64, count int) ([]byte, error)
	MFGetSize       func(fsid string) (int64, error)
	MFStat          func(fsid string) (*FileInfo, error)
	MFIsExist       func(fsid, fileid string) (bool, error)
	MFReaddir       func(fsid string, count int) ([]*FileInfo, error)
	MFGetMimeType   func(fsid string) (string, error)
	MFSetCid        func(sid, mid, cid string) (string, error)
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
	FilesRm    func(sid, ps string, recursive, flush bool) error
	FilesMv    func(sid, src, dst string, flush bool) error
	FilesFlush func(sid, ps string) (string, error)
	FilesStat  func(sid, ps string) (*StatInfo, error)
}
