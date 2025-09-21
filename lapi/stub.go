package lapi

import (
	"bytes"
	"encoding/gob"
	"os"
	"time"
)

// Package lapi provides RPC functionality for the Leither project
type LApiStub struct {
	AuthStub
	VarActStub
	MiMeiStub
	NetStub
}

// 后续应该是单独文件，先放这里
type MiMeiRefs = map[string]map[string]int
type Refs = map[string]int

type MiMeiStub struct {
	MDbStub
	MFileStub

	MMCreate   func(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)
	MMOpen     func(sid, mid, ver string, opt ...string) (string, error)
	MMOpenUrl  func(sid, ps string) (string, error)
	MMClose    func(mmsid string) error
	MMSetRight func(sid, mid, member string, right uint64) error
	MMGetRight func(sid, mmid, uid string) (right uint64, err error)
	MMBackup   func(sid, mid, memo string, opt ...string) (string, error) //snapshot
	MMRestore  func(sid, mid, ver string) error
	MMDelVers  func(sid, mid string, vers ...string) (int64, error)
	MMRelease  func(sid, mid, ver string) (string, error)
	MMAddRef   func(sid, mid string, fileids ...string) (int, error)
	MMDelRef   func(sid, mid string, fileids ...string) (int, error)
	MMGetRef   func(sid, mid, ver string) (ret Refs, err error)
	MMGetRefs  func(sid, mid string, vers ...string) (MiMeiRefs, error)
	MMSum      func(sid, mid, ver, tp string) (string, error) //

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
	Act             func(sid, name string, args ...string) (any, error)
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
	FName    string      // base name of the file
	FSize    int64       // length in bytes for regular files; system-dependent for others
	FModTime time.Time   // modification time
	FIsDir   bool        // abbreviation for Mode().IsDir()
	FMode    os.FileMode // mode
	FID      string      //文件id
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

type NetStub struct {
	FilesStub
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

// IAuthStub 接口定义了认证相关的操作方法
type IAuthStub interface {
	LoginWithPPT(strPPT string) (*LoginReply, error)
	Logout(sid, info string) error
}

type IVarActStub interface {
	GetVar(sid, name string, args ...string) (any, error)
	Act(sid, name string, args ...string) (any, error)
	GetGobVar(sid, name string, args ...string) ([]byte, error)
	// GetGobVar(sid, name string, args ...string) ([]byte, error)
}

// IMiMeiStub 接口定义了弥媒相关的操作方法
type IMiMeiStub interface {
	IMDbStub
	IMFileStub
	MMCreate(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)
	MMOpen(sid, mid, ver string, opt ...string) (string, error)
	MMOpenUrl(sid, ps string) (string, error)
	MMClose(mmsid string) error
	MMSetRight(sid, mid, member string, right uint64) error
	MMGetRight(sid, mmid, uid string) (right uint64, err error)
	MMBackup(sid, mid, memo string, opt ...string) (string, error)
	MMRestore(sid, mid, ver string) error
	MMDelVers(sid, mid string, vers ...string) (int64, error)
	MMRelease(sid, mid, ver string) (string, error)
	MMAddRef(sid, mid string, fileids ...string) (int, error)
	MMDelRef(sid, mid string, fileids ...string) (int, error)
	MMGetRef(sid, mid, ver string) (ret Refs, err error)
	MMGetRefs(sid, mid string, vers ...string) (MiMeiRefs, error)
	MMSum(sid, mid, ver, tp string) (string, error)
}

// INetStub 接口定义了网络相关的操作方法
type INetStub interface {
	IFilesStub
}

type IFilesStub interface {
	FilesCopy(sid, src, dst string, flush bool) error
	FilesLs(sid, ps string) ([]LsLink, error)
	FilesMkdir(sid, ps string, flush bool) error
	FilesRm(sid, ps string, recursive, flush bool) error
	FilesMv(sid, src, dst string, flush bool) error
	FilesFlush(sid, ps string) (string, error)
	FilesStat(sid, ps string) (*StatInfo, error)
}

// 在LApiStub上实现IAuthStub接口方法
func (s *LApiStub) LoginWithPPT(strPPT string) (*LoginReply, error) {
	if s.AuthStub.LoginWithPPT != nil {
		return s.AuthStub.LoginWithPPT(strPPT)
	}
	return nil, nil
}

func (s *LApiStub) Logout(sid, info string) error {
	if s.AuthStub.Logout != nil {
		return s.AuthStub.Logout(sid, info)
	}
	return nil
}

// 在LApiStub上实现IVarActStub接口方法
func (s *LApiStub) GetVar(sid, name string, args ...string) (any, error) {
	if s.VarActStub.GetVar != nil {
		return s.VarActStub.GetVar(sid, name, args...)
	}
	return nil, nil
}

func (s *LApiStub) GetVarByContext(sid, name string, mapOpt map[string]string) (any, error) {
	if s.VarActStub.GetVarByContext != nil {
		return s.VarActStub.GetVarByContext(sid, name, mapOpt)
	}
	return nil, nil
}

func (s *LApiStub) Act(sid, name string, args ...string) (any, error) {
	if s.VarActStub.Act != nil {
		return s.VarActStub.Act(sid, name, args...)
	}
	return nil, nil
}

func (s *LApiStub) GetGobVar(sid, name string, args ...string) ([]byte, error) {
	if s.VarActStub.GetGobVar != nil {
		return s.VarActStub.GetGobVar(sid, name, args...)
	}
	return nil, nil
}

func (s *LApiStub) GetVarObj(obj any, sid string, name string, args ...string) error {
	return s.VarActStub.GetVarObj(obj, sid, name, args...)
}

// 在LApiStub上实现IMiMeiStub接口方法（部分示例）
func (s *LApiStub) MMCreate(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error) {
	if s.MiMeiStub.MMCreate != nil {
		return s.MiMeiStub.MMCreate(sid, appid, ext, mark, tp, right)
	}
	return "", nil
}

func (s *LApiStub) MMOpen(sid, mid, ver string, opt ...string) (string, error) {
	if s.MiMeiStub.MMOpen != nil {
		return s.MiMeiStub.MMOpen(sid, mid, ver, opt...)
	}
	return "", nil
}

func (s *LApiStub) MMOpenUrl(sid, ps string) (string, error) {
	if s.MiMeiStub.MMOpenUrl != nil {
		return s.MiMeiStub.MMOpenUrl(sid, ps)
	}
	return "", nil
}

func (s *LApiStub) MMClose(mmsid string) error {
	if s.MiMeiStub.MMClose != nil {
		return s.MiMeiStub.MMClose(mmsid)
	}
	return nil
}

func (s *LApiStub) MMSetRight(sid, mid, member string, right uint64) error {
	if s.MiMeiStub.MMSetRight != nil {
		return s.MiMeiStub.MMSetRight(sid, mid, member, right)
	}
	return nil
}

func (s *LApiStub) MMGetRight(sid, mmid, uid string) (right uint64, err error) {
	if s.MiMeiStub.MMGetRight != nil {
		return s.MiMeiStub.MMGetRight(sid, mmid, uid)
	}
	return 0, nil
}

func (s *LApiStub) MMBackup(sid, mid, memo string, opt ...string) (string, error) {
	if s.MiMeiStub.MMBackup != nil {
		return s.MiMeiStub.MMBackup(sid, mid, memo, opt...)
	}
	return "", nil
}

func (s *LApiStub) MMRestore(sid, mid, ver string) error {
	if s.MiMeiStub.MMRestore != nil {
		return s.MiMeiStub.MMRestore(sid, mid, ver)
	}
	return nil
}

func (s *LApiStub) MMDelVers(sid, mid string, vers ...string) (int64, error) {
	if s.MiMeiStub.MMDelVers != nil {
		return s.MiMeiStub.MMDelVers(sid, mid, vers...)
	}
	return 0, nil
}

func (s *LApiStub) MMRelease(sid, mid, ver string) (string, error) {
	if s.MiMeiStub.MMRelease != nil {
		return s.MiMeiStub.MMRelease(sid, mid, ver)
	}
	return "", nil
}

func (s *LApiStub) MMAddRef(sid, mid string, fileids ...string) (int, error) {
	if s.MiMeiStub.MMAddRef != nil {
		return s.MiMeiStub.MMAddRef(sid, mid, fileids...)
	}
	return 0, nil
}

func (s *LApiStub) MMDelRef(sid, mid string, fileids ...string) (int, error) {
	if s.MiMeiStub.MMDelRef != nil {
		return s.MiMeiStub.MMDelRef(sid, mid, fileids...)
	}
	return 0, nil
}

func (s *LApiStub) MMGetRef(sid, mid, ver string) (ret Refs, err error) {
	if s.MiMeiStub.MMGetRef != nil {
		return s.MiMeiStub.MMGetRef(sid, mid, ver)
	}
	return nil, nil
}

func (s *LApiStub) MMGetRefs(sid, mid string, vers ...string) (MiMeiRefs, error) {
	if s.MiMeiStub.MMGetRefs != nil {
		return s.MiMeiStub.MMGetRefs(sid, mid, vers...)
	}
	return nil, nil
}

func (s *LApiStub) MMSum(sid, mid, ver, tp string) (string, error) {
	if s.MiMeiStub.MMSum != nil {
		return s.MiMeiStub.MMSum(sid, mid, ver, tp)
	}
	return "", nil
}

// 在LApiStub上实现INetStub接口方法
func (s *LApiStub) FilesCopy(sid, src, dst string, flush bool) error {
	if s.NetStub.FilesCopy != nil {
		return s.NetStub.FilesCopy(sid, src, dst, flush)
	}
	return nil
}

func (s *LApiStub) FilesLs(sid, ps string) ([]LsLink, error) {
	if s.NetStub.FilesLs != nil {
		return s.NetStub.FilesLs(sid, ps)
	}
	return nil, nil
}

func (s *LApiStub) FilesMkdir(sid, ps string, flush bool) error {
	if s.NetStub.FilesMkdir != nil {
		return s.NetStub.FilesMkdir(sid, ps, flush)
	}
	return nil
}

func (s *LApiStub) FilesRm(sid, ps string, recursive, flush bool) error {
	if s.NetStub.FilesRm != nil {
		return s.NetStub.FilesRm(sid, ps, recursive, flush)
	}
	return nil
}

func (s *LApiStub) FilesMv(sid, src, dst string, flush bool) error {
	if s.NetStub.FilesMv != nil {
		return s.NetStub.FilesMv(sid, src, dst, flush)
	}
	return nil
}

func (s *LApiStub) FilesFlush(sid, ps string) (string, error) {
	if s.NetStub.FilesFlush != nil {
		return s.NetStub.FilesFlush(sid, ps)
	}
	return "", nil
}

func (s *LApiStub) FilesStat(sid, ps string) (*StatInfo, error) {
	if s.NetStub.FilesStat != nil {
		return s.NetStub.FilesStat(sid, ps)
	}
	return nil, nil
}

// 接口验证
var _ IAuthStub = (*LApiStub)(nil)
var _ IVarActStub = (*LApiStub)(nil)
var _ IMiMeiStub = (*LApiStub)(nil)
var _ INetStub = (*LApiStub)(nil)

// var _ IMDbStub = (*LApiStub)(nil)

// MDbStub接口定义
type IMDbStub interface {
	Begin(mmsid string, timeout int) error
	Commit(mmsid string) error
	Rollback(mmsid string) error
	GetLastSeq(mmsid string) (uint64, error)
	Set(mmsid, key string, value any) error
	Get(mmsid, key string) (any, error)
	Del(mmsid string, key ...string) (int64, error)
	Incr(mmsid, key string) (int64, error)
	IncrBy(mmsid, key string, delta int64) (int64, error)
	Strlen(mmsid, key string) (int64, error)
	Hmclear(mmsid string, key ...string) (int64, error)
	Hclear(mmsid string, key string) (int64, error)
	Hdel(mmsid, key string, field ...string) (int64, error)
	Hlen(mmsid, key string) (int64, error)
	Hset(mmsid, key, field string, value any) (int64, error)
	Hget(mmsid, key, field string) (any, error)
	Hmget(mmsid, key string, fields ...string) ([]any, error)
	Hmset(mmsid, key string, values ...FVPair) error
	Hgetall(mmsid, key string) ([]FVPair, error)
	Hkeys(mmsid, key string) ([]string, error)
	Hscan(mmsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)
	Hrevscan(mmsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)
	HincrBy(mmsid, key, field string, delta int64) (ret int64, err error)
	Lpush(mmsid, key string, value ...any) (int64, error)
	Lpop(mmsid, key string) (any, error)
	Rpush(mmsid, key string, value ...any) (int64, error)
	Rpop(mmsid, key string) (any, error)
	Lrange(mmsid, key string, start, stop int32) ([]any, error)
	Lclear(mmsid, key string) (int64, error)
	Lmclear(mmsid string, keys ...string) (int64, error)
	Lindex(mmsid, key string, index int32) (any, error)
	Llen(mmsid, key string) (int64, error)
	Lset(mmsid, key string, index int32, value any) error
	Zadd(mmsid, key string, args ...ScorePair) (int64, error)
	Zaddwithseq(mmsid, key string, members ...string) (ret int64, err error)
	Zcard(mmsid, key string) (int64, error)
	Zcount(mmsid, key string, mins, maxs int64) (int64, error)
	Zrem(mmsid, key string, members ...string) (int64, error)
	Zscore(mmsid, key, member string) (int64, error)
	Zrank(mmsid, key, member string) (int64, error)
	Zrange(mmsid, key string, start, stop int) ([]ScorePair, error)
	Zrangebyscore(mmsid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error)
	Zremrangebyscore(mmsid, key string, mins, maxs int64) (int64, error)
	Zrevrange(mmsid, key string, start, stop int) (ret []ScorePair, err error)
	Zrevrangebyscore(mmsid, key string, mins, maxs int64, offset int, count int) (ret []ScorePair, err error)
	Zmclear(mmsid string, key ...string) (int64, error)
	Zclear(mmsid, key string) (int64, error)
	ZincrBy(mmsid, key string, delta int64, member string) (int64, error)
	Sadd(mmsid, key string, args ...string) (int64, error)
	Scard(mmsid, key string) (int64, error)
	Sclear(mmsid, key string) (int64, error)
	Sdiff(mmsid string, keys ...string) ([]string, error)
	Sinter(mmsid string, keys ...string) ([]string, error)
	Smclear(mmsid string, key ...string) (int64, error)
	Smembers(mmsid, key string) ([]string, error)
	Srem(mmsid, key string, m string) (int64, error)
	Sunion(mmsid string, keys ...string) ([]string, error)
	Scan(mmsid string, begin, match string, count int, inclusive bool, tp byte) (keys []string, err error)
	MMSyncData(sid, mid, tp string, mark []byte, count uint64, verseq ...uint64) ([]byte, error)
	CheckIntegrity(mmsid string, dataType uint8, key string, bRepair bool) error
}

// MFileStub接口定义
type IMFileStub interface {
	MFOpenByPath(sid, mid, path string, flag int) (string, error)
	MFOpenMacFile(sid, mid, fileid string) (string, error)
	MFOpenTempFile(sid string) (string, error)
	MFPath2TempID(sid, localpath string) (string, error)
	MFLocal2MacFile(sid, mid, ps string) (string, error)
	MFTemp2MacFile(fsid, mid string) (string, error)
	MFTemp2Ipfs(fsid, mid string) (string, error)
	MFTemp2Files(fsid, ps string) (string, error)
	MFTruncate(fsid string, size int64) error
	MFSetObject(fsid string, obj any) error
	MFGetObject(fsid string) (any, error)
	MFSetData(fsid string, data []byte, start int64) (count int, err error)
	MFGetData(fsid string, start int64, count int) ([]byte, error)
	MFGetSize(fsid string) (int64, error)
	MFStat(fsid string) (*FileInfo, error)
	MFIsExist(fsid, fileid string) (bool, error)
	MFReaddir(fsid string, count int) ([]*FileInfo, error)
	MFGetMimeType(fsid string) (string, error)
	MFSetCid(sid, mid, cid string) (string, error)
}

// 在LApiStub中实现IMDbStub接口方法
func (stub *LApiStub) Begin(mmsid string, timeout int) error {
	if stub.MiMeiStub.MDbStub.Begin == nil {
		return nil
	}
	return stub.MiMeiStub.MDbStub.Begin(mmsid, timeout)
}

func (stub *LApiStub) Commit(mmsid string) error {
	if stub.MiMeiStub.MDbStub.Commit == nil {
		return nil
	}
	return stub.MiMeiStub.MDbStub.Commit(mmsid)
}

func (stub *LApiStub) Rollback(mmsid string) error {
	if stub.MiMeiStub.MDbStub.Rollback == nil {
		return nil
	}
	return stub.MiMeiStub.MDbStub.Rollback(mmsid)
}

func (stub *LApiStub) GetLastSeq(mmsid string) (uint64, error) {
	if stub.MiMeiStub.MDbStub.GetLastSeq == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.GetLastSeq(mmsid)
}

func (stub *LApiStub) Set(mmsid, key string, value any) error {
	if stub.MiMeiStub.MDbStub.Set == nil {
		return nil
	}
	return stub.MiMeiStub.MDbStub.Set(mmsid, key, value)
}

func (stub *LApiStub) Get(mmsid, key string) (any, error) {
	if stub.MiMeiStub.MDbStub.Get == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Get(mmsid, key)
}

func (stub *LApiStub) Del(mmsid string, key ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Del == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Del(mmsid, key...)
}

func (stub *LApiStub) Incr(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Incr == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Incr(mmsid, key)
}

func (stub *LApiStub) IncrBy(mmsid, key string, delta int64) (int64, error) {
	if stub.MiMeiStub.MDbStub.IncrBy == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.IncrBy(mmsid, key, delta)
}

func (stub *LApiStub) Strlen(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Strlen == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Strlen(mmsid, key)
}

func (stub *LApiStub) Hmclear(mmsid string, key ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Hmclear == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Hmclear(mmsid, key...)
}

func (stub *LApiStub) Hclear(mmsid string, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Hclear == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Hclear(mmsid, key)
}

func (stub *LApiStub) Hdel(mmsid, key string, field ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Hdel == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Hdel(mmsid, key, field...)
}

func (stub *LApiStub) Hlen(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Hlen == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Hlen(mmsid, key)
}

func (stub *LApiStub) Hset(mmsid, key, field string, value any) (int64, error) {
	if stub.MiMeiStub.MDbStub.Hset == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Hset(mmsid, key, field, value)
}

func (stub *LApiStub) Hget(mmsid, key, field string) (any, error) {
	if stub.MiMeiStub.MDbStub.Hget == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Hget(mmsid, key, field)
}

func (stub *LApiStub) Hmget(mmsid, key string, fields ...string) ([]any, error) {
	if stub.MiMeiStub.MDbStub.Hmget == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Hmget(mmsid, key, fields...)
}

func (stub *LApiStub) Hmset(mmsid, key string, values ...FVPair) error {
	if stub.MiMeiStub.MDbStub.Hmset == nil {
		return nil
	}
	return stub.MiMeiStub.MDbStub.Hmset(mmsid, key, values...)
}

func (stub *LApiStub) Hgetall(mmsid, key string) ([]FVPair, error) {
	if stub.MiMeiStub.MDbStub.Hgetall == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Hgetall(mmsid, key)
}

func (stub *LApiStub) Hkeys(mmsid, key string) ([]string, error) {
	if stub.MiMeiStub.MDbStub.Hkeys == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Hkeys(mmsid, key)
}

func (stub *LApiStub) Hscan(mmsid, key, beginfield, match string, count int, inclusive bool) ([]FVPair, error) {
	if stub.MiMeiStub.MDbStub.Hscan == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Hscan(mmsid, key, beginfield, match, count, inclusive)
}

func (stub *LApiStub) Hrevscan(mmsid, key, beginfield, match string, count int, inclusive bool) ([]FVPair, error) {
	if stub.MiMeiStub.MDbStub.Hrevscan == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Hrevscan(mmsid, key, beginfield, match, count, inclusive)
}

func (stub *LApiStub) HincrBy(mmsid, key, field string, delta int64) (int64, error) {
	if stub.MiMeiStub.MDbStub.HincrBy == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.HincrBy(mmsid, key, field, delta)
}

func (stub *LApiStub) Lpush(mmsid, key string, value ...any) (int64, error) {
	if stub.MiMeiStub.MDbStub.Lpush == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Lpush(mmsid, key, value...)
}

func (stub *LApiStub) Lpop(mmsid, key string) (any, error) {
	if stub.MiMeiStub.MDbStub.Lpop == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Lpop(mmsid, key)
}

func (stub *LApiStub) Rpush(mmsid, key string, value ...any) (int64, error) {
	if stub.MiMeiStub.MDbStub.Rpush == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Rpush(mmsid, key, value...)
}

func (stub *LApiStub) Rpop(mmsid, key string) (any, error) {
	if stub.MiMeiStub.MDbStub.Rpop == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Rpop(mmsid, key)
}

func (stub *LApiStub) Lrange(mmsid, key string, start, stop int32) ([]any, error) {
	if stub.MiMeiStub.MDbStub.Lrange == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Lrange(mmsid, key, start, stop)
}

func (stub *LApiStub) Lclear(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Lclear == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Lclear(mmsid, key)
}

func (stub *LApiStub) Lmclear(mmsid string, keys ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Lmclear == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Lmclear(mmsid, keys...)
}

func (stub *LApiStub) Lindex(mmsid, key string, index int32) (any, error) {
	if stub.MiMeiStub.MDbStub.Lindex == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Lindex(mmsid, key, index)
}

func (stub *LApiStub) Llen(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Llen == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Llen(mmsid, key)
}

func (stub *LApiStub) Lset(mmsid, key string, index int32, value any) error {
	if stub.MiMeiStub.MDbStub.Lset == nil {
		return nil
	}
	return stub.MiMeiStub.MDbStub.Lset(mmsid, key, index, value)
}

func (stub *LApiStub) Zadd(mmsid, key string, args ...ScorePair) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zadd == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zadd(mmsid, key, args...)
}

func (stub *LApiStub) Zaddwithseq(mmsid, key string, members ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zaddwithseq == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zaddwithseq(mmsid, key, members...)
}

func (stub *LApiStub) Zcard(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zcard == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zcard(mmsid, key)
}

func (stub *LApiStub) Zcount(mmsid, key string, mins, maxs int64) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zcount == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zcount(mmsid, key, mins, maxs)
}

func (stub *LApiStub) Zrem(mmsid, key string, members ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zrem == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zrem(mmsid, key, members...)
}

func (stub *LApiStub) Zscore(mmsid, key, member string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zscore == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zscore(mmsid, key, member)
}

func (stub *LApiStub) Zrank(mmsid, key, member string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zrank == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zrank(mmsid, key, member)
}

func (stub *LApiStub) Zrange(mmsid, key string, start, stop int) ([]ScorePair, error) {
	if stub.MiMeiStub.MDbStub.Zrange == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Zrange(mmsid, key, start, stop)
}

func (stub *LApiStub) Zrangebyscore(mmsid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error) {
	if stub.MiMeiStub.MDbStub.Zrangebyscore == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Zrangebyscore(mmsid, key, mins, maxs, offset, count)
}

func (stub *LApiStub) Zremrangebyscore(mmsid, key string, mins, maxs int64) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zremrangebyscore == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zremrangebyscore(mmsid, key, mins, maxs)
}

func (stub *LApiStub) Zrevrange(mmsid, key string, start, stop int) ([]ScorePair, error) {
	if stub.MiMeiStub.MDbStub.Zrevrange == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Zrevrange(mmsid, key, start, stop)
}

func (stub *LApiStub) Zrevrangebyscore(mmsid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error) {
	if stub.MiMeiStub.MDbStub.Zrevrangebyscore == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Zrevrangebyscore(mmsid, key, mins, maxs, offset, count)
}

func (stub *LApiStub) Zmclear(mmsid string, key ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zmclear == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zmclear(mmsid, key...)
}

func (stub *LApiStub) Zclear(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Zclear == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Zclear(mmsid, key)
}

func (stub *LApiStub) ZincrBy(mmsid, key string, delta int64, member string) (int64, error) {
	if stub.MiMeiStub.MDbStub.ZincrBy == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.ZincrBy(mmsid, key, delta, member)
}

func (stub *LApiStub) Sadd(mmsid, key string, args ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Sadd == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Sadd(mmsid, key, args...)
}

func (stub *LApiStub) Scard(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Scard == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Scard(mmsid, key)
}

func (stub *LApiStub) Sclear(mmsid, key string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Sclear == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Sclear(mmsid, key)
}

func (stub *LApiStub) Sdiff(mmsid string, keys ...string) ([]string, error) {
	if stub.MiMeiStub.MDbStub.Sdiff == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Sdiff(mmsid, keys...)
}

func (stub *LApiStub) Sinter(mmsid string, keys ...string) ([]string, error) {
	if stub.MiMeiStub.MDbStub.Sinter == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Sinter(mmsid, keys...)
}

func (stub *LApiStub) Smclear(mmsid string, key ...string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Smclear == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Smclear(mmsid, key...)
}

func (stub *LApiStub) Smembers(mmsid, key string) ([]string, error) {
	if stub.MiMeiStub.MDbStub.Smembers == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Smembers(mmsid, key)
}

func (stub *LApiStub) Srem(mmsid, key string, m string) (int64, error) {
	if stub.MiMeiStub.MDbStub.Srem == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MDbStub.Srem(mmsid, key, m)
}

func (stub *LApiStub) Sunion(mmsid string, keys ...string) ([]string, error) {
	if stub.MiMeiStub.MDbStub.Sunion == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Sunion(mmsid, keys...)
}

func (stub *LApiStub) Scan(mmsid string, begin, match string, count int, inclusive bool, tp byte) (keys []string, err error) {
	if stub.MiMeiStub.MDbStub.Scan == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.Scan(mmsid, begin, match, count, inclusive, tp)
}

func (stub *LApiStub) MMSyncData(sid, mid, tp string, mark []byte, count uint64, verseq ...uint64) ([]byte, error) {
	if stub.MiMeiStub.MDbStub.MMSyncData == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MDbStub.MMSyncData(sid, mid, tp, mark, count, verseq...)
}

func (stub *LApiStub) CheckIntegrity(mmsid string, dataType uint8, key string, bRepair bool) error {
	if stub.MiMeiStub.MDbStub.CheckIntegrity == nil {
		return nil
	}
	return stub.MiMeiStub.MDbStub.CheckIntegrity(mmsid, dataType, key, bRepair)
}

// 在LApiStub中实现IMFileStub接口方法
func (stub *LApiStub) MFOpenByPath(sid, mid, path string, flag int) (string, error) {
	if stub.MiMeiStub.MFileStub.MFOpenByPath == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFOpenByPath(sid, mid, path, flag)
}

func (stub *LApiStub) MFOpenMacFile(sid, mid, fileid string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFOpenMacFile == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFOpenMacFile(sid, mid, fileid)
}

func (stub *LApiStub) MFOpenTempFile(sid string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFOpenTempFile == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFOpenTempFile(sid)
}

func (stub *LApiStub) MFPath2TempID(sid, localpath string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFPath2TempID == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFPath2TempID(sid, localpath)
}

func (stub *LApiStub) MFLocal2MacFile(sid, mid, ps string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFLocal2MacFile == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFLocal2MacFile(sid, mid, ps)
}

func (stub *LApiStub) MFTemp2MacFile(fsid, mid string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFTemp2MacFile == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFTemp2MacFile(fsid, mid)
}

func (stub *LApiStub) MFTemp2Ipfs(fsid, mid string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFTemp2Ipfs == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFTemp2Ipfs(fsid, mid)
}

func (stub *LApiStub) MFTemp2Files(fsid, ps string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFTemp2Files == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFTemp2Files(fsid, ps)
}

func (stub *LApiStub) MFTruncate(fsid string, size int64) error {
	if stub.MiMeiStub.MFileStub.MFTruncate == nil {
		return nil
	}
	return stub.MiMeiStub.MFileStub.MFTruncate(fsid, size)
}

func (stub *LApiStub) MFSetObject(fsid string, obj any) error {
	if stub.MiMeiStub.MFileStub.MFSetObject == nil {
		return nil
	}
	return stub.MiMeiStub.MFileStub.MFSetObject(fsid, obj)
}

func (stub *LApiStub) MFGetObject(fsid string) (any, error) {
	if stub.MiMeiStub.MFileStub.MFGetObject == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MFileStub.MFGetObject(fsid)
}

func (stub *LApiStub) MFSetData(fsid string, data []byte, start int64) (count int, err error) {
	if stub.MiMeiStub.MFileStub.MFSetData == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MFileStub.MFSetData(fsid, data, start)
}

func (stub *LApiStub) MFGetData(fsid string, start int64, count int) ([]byte, error) {
	if stub.MiMeiStub.MFileStub.MFGetData == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MFileStub.MFGetData(fsid, start, count)
}

func (stub *LApiStub) MFGetSize(fsid string) (int64, error) {
	if stub.MiMeiStub.MFileStub.MFGetSize == nil {
		return 0, nil
	}
	return stub.MiMeiStub.MFileStub.MFGetSize(fsid)
}

func (stub *LApiStub) MFStat(fsid string) (*FileInfo, error) {
	if stub.MiMeiStub.MFileStub.MFStat == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MFileStub.MFStat(fsid)
}

func (stub *LApiStub) MFIsExist(fsid, fileid string) (bool, error) {
	if stub.MiMeiStub.MFileStub.MFIsExist == nil {
		return false, nil
	}
	return stub.MiMeiStub.MFileStub.MFIsExist(fsid, fileid)
}

func (stub *LApiStub) MFReaddir(fsid string, count int) ([]*FileInfo, error) {
	if stub.MiMeiStub.MFileStub.MFReaddir == nil {
		return nil, nil
	}
	return stub.MiMeiStub.MFileStub.MFReaddir(fsid, count)
}

func (stub *LApiStub) MFGetMimeType(fsid string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFGetMimeType == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFGetMimeType(fsid)
}

func (stub *LApiStub) MFSetCid(sid, mid, cid string) (string, error) {
	if stub.MiMeiStub.MFileStub.MFSetCid == nil {
		return "", nil
	}
	return stub.MiMeiStub.MFileStub.MFSetCid(sid, mid, cid)
}
