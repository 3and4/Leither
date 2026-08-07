// imimei.go
package lapi

import (
	"os"
	"time"
)

// IMiMei 接口定义了弥媒相关的操作方法
// 弥媒是Leither平台中的核心数据对象，支持版本控制、权限管理、备份恢复等功能
// 该接口提供了对弥媒对象的完整生命周期管理
type IMiMei interface {
	IMDb
	IMFile

	// MMCreate 创建新的弥媒对象
	// 参数: sid-会话ID, appid-应用ID, ext-扩展标识, mark-标记字符串, tp-弥媒类型, right-权限位掩码
	// 返回值: mid-新创建的弥媒ID
	MMCreate(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)

	// MMOpen 打开指定版本的弥媒
	// 参数: sid-会话ID, mid-弥媒ID, ver-版本号(空字符串表示最新版本), opt-可选参数
	// 返回值: 会话句柄，用于后续操作
	MMOpen(sid, mid, ver string, opt ...string) (string, error)

	// MMOpenUrl 通过URL路径打开弥媒
	// 参数: sid-会话ID, ps-URL路径
	// 返回值: 会话句柄
	MMOpenUrl(sid, ps string) (string, error)

	// MMClose 关闭弥媒会话
	// 参数: mmsid-弥媒会话ID
	MMClose(mmsid string) error

	// MMSetRight 设置弥媒成员权限
	// 参数: sid-会话ID, mid-弥媒ID, member-成员标识, right-权限位掩码
	MMSetRight(sid, mid, member string, right uint64) error

	// MMGetRight 获取成员对弥媒的权限
	// 参数: sid-会话ID, mmid-弥媒ID, uid-用户ID
	// 返回值: right-权限位掩码
	MMGetRight(sid, mmid, uid string) (right uint64, err error)

	// MMBackup 创建弥媒快照备份
	// 参数: sid-会话ID, mid-弥媒ID, memo-备份备注信息, opt-可选参数
	// 返回值: 备份版本号
	MMBackup(sid, mid, memo string, opt ...string) (string, error)

	// MMRestore 从指定版本恢复弥媒
	// 参数: sid-会话ID, mid-弥媒ID, ver-要恢复的版本号
	MMRestore(sid, mid, ver string) error

	// MMDelVers 删除指定版本
	// 参数: sid-会话ID, mid-弥媒ID, vers-要删除的版本号列表
	// 返回值: 删除的版本数量
	MMDelVers(sid, mid string, vers ...string) (int64, error)

	// MMRelease 发布指定版本
	// 参数: sid-会话ID, mid-弥媒ID, ver-要发布的版本号
	// 返回值: 发布后的版本标识
	MMRelease(sid, mid, ver string) (string, error)

	// MMAddRef 为弥媒添加文件引用
	// 参数: sid-会话ID, mid-弥媒ID, fileids-要引用的文件ID列表
	// 返回值: 成功添加的引用数量
	MMAddRef(sid, mid string, fileids ...string) (int, error)

	// MMDelRef 删除弥媒的文件引用
	// 参数: sid-会话ID, mid-弥媒ID, fileids-要删除引用的文件ID列表
	// 返回值: 成功删除的引用数量
	MMDelRef(sid, mid string, fileids ...string) (int, error)

	// MMGetRef 获取指定版本的引用关系
	// 参数: sid-会话ID, mid-弥媒ID, ver-版本号
	// 返回值: ret-文件引用关系映射
	MMGetRef(sid, mid, ver string) (ret Refs, err error)

	// MMGetRefs 获取多个版本的引用关系
	// 参数: sid-会话ID, mid-弥媒ID, vers-版本号列表
	// 返回值: 各版本的引用关系映射
	MMGetRefs(sid, mid string, vers ...string) (MiMeiRefs, error)

	// MMSum 计算弥媒的摘要信息
	// 参数: sid-会话ID, mid-弥媒ID, ver-版本号, tp-摘要类型
	// 返回值: 摘要字符串
	MMSum(sid, mid, ver, tp string) (string, error)

	// InlineAppCss 将 html 中引用应用内文件的 CSS 内联为 <style>，返回转换后 html。
	// appid/root: 应用 mmroot 定位（mm://<appid>/<root>/<name>），用于读应用内 CSS。
	// mid/ver: 构造 url() 改写基准 /mm/<mid>:<ver>/。
	InlineAppCss(sid, appid, mid, ver, root string, html []byte) ([]byte, error)
}

type FVPair struct {
	Field string
	Value any
}

type ScorePair struct {
	Score  int64
	Member string
}

// MDbStub接口定义
type IMDb interface {
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

// IMFile 接口定义了弥媒文件操作的核心功能
// 该接口提供了对弥媒中文件的完整生命周期管理，包括打开、读写、转换、查询等操作
// 所有文件操作都通过会话ID进行权限控制和状态管理
//
// 弥媒文件存储机制：
// - 备份版本：备份后会生成数字版本号，对应两种数据类型：
//   - mac文件：保存在mac目录中，文件名为文件hash转化成的字符串 (PATH_MacFILE="mac")
//   - ipfs文件：保存在系统的ipfs文件仓库中
//
// - 未备份版本：
//   - mac文件：保存在mid目录中 (PATH_MIDFILE="mid")
//   - ipfs文件：保存在files文件系统的.mm目录中 (Path_Files_MM="/.mm/")
//
// - 版本信息存储在Redis数据库中，通过GetVerInfo等函数管理
type IMFile interface {
	// MFOpenByPath 通过路径打开文件
	// 参数: sid-会话ID, mid-弥媒ID, path-文件路径, flag-打开标志
	// 返回值: 文件会话ID，用于后续文件操作
	MFOpenByPath(sid, mid, path string, flag int) (string, error)

	// MFOpenMacFile 通过文件ID打开Mac文件
	// 参数: sid-会话ID, mid-弥媒ID, fileid-文件ID
	// 返回值: 文件会话ID
	MFOpenMacFile(sid, mid, fileid string) (string, error)

	// MFOpenTempFile 创建并打开临时文件
	// 参数: sid-会话ID
	// 返回值: 临时文件会话ID
	MFOpenTempFile(sid string) (string, error)

	// MFPath2TempID 将本地路径转换为临时文件ID
	// 参数: sid-会话ID, localpath-本地文件路径
	// 返回值: 临时文件ID
	MFPath2TempID(sid, localpath string) (string, error)

	// MFLocal2MacFile 将本地文件转换为Mac文件格式
	// 参数: sid-会话ID, mid-弥媒ID, ps-本地文件路径
	// 返回值: 转换后的Mac文件ID
	// 说明: 转换后的Mac文件存储在mac目录中，文件名为hash值
	MFLocal2MacFile(sid, mid, ps string) (string, error)

	// MFTemp2MacFile 将临时文件转换为Mac文件
	// 参数: fsid-文件会话ID, mid-弥媒ID
	// 返回值: 转换后的Mac文件ID
	// 说明: 转换后的Mac文件存储在mac目录中，文件名为hash值
	MFTemp2MacFile(fsid, mid string) (string, error)

	// MFTemp2Ipfs 将临时文件转换为IPFS文件
	// 参数: fsid-文件会话ID, mid-弥媒ID
	// 返回值: 转换后的IPFS CID
	// 说明: 转换后的IPFS文件存储在系统的IPFS仓库中
	MFTemp2Ipfs(fsid, mid string) (string, error)

	// MFTemp2Files 将临时文件转换为Files系统文件
	// 参数: fsid-文件会话ID, psfiles-目标路径
	// 返回值: 转换后的CID
	// 说明: 转换后的文件存储在files文件系统的.mm目录中
	MFTemp2Files(fsid, ps string) (string, error)

	// MFTruncate 截断文件到指定大小
	// 参数: fsid-文件会话ID, size-目标大小
	MFTruncate(fsid string, size int64) error

	// MFSetObject 将对象序列化写入文件
	// 参数: fsid-文件会话ID, obj-要写入的对象
	MFSetObject(fsid string, obj any) error

	// MFGetObject 从文件中读取并反序列化对象
	// 参数: fsid-文件会话ID
	// 返回值: 反序列化后的对象
	MFGetObject(fsid string) (any, error)

	// MFSetData 在文件指定位置写入数据
	// 参数: fsid-文件会话ID, data-要写入的数据, start-起始位置
	// 返回值: 实际写入的字节数
	// 说明: start为正表示从头部偏移，start为负表示从尾部偏移
	MFSetData(fsid string, data []byte, start int64) (count int, err error)

	// MFGetData 从文件指定位置读取数据
	// 参数: fsid-文件会话ID, start-起始位置, count-读取字节数
	// 返回值: 读取的数据
	// 说明: start为正表示从头部偏移，start为负表示从尾部偏移
	MFGetData(fsid string, start int64, count int) ([]byte, error)

	// MFGetSize 获取文件大小
	// 参数: fsid-文件会话ID
	// 返回值: 文件大小
	MFGetSize(fsid string) (int64, error)

	// MFStat 获取文件信息
	// 参数: fsid-文件会话ID
	// 返回值: 文件信息结构
	MFStat(fsid string) (*FileInfo, error)

	// MFIsExist 检查文件是否存在
	// 参数: fsid-文件会话ID, fileid-文件ID
	// 返回值: 文件是否存在
	MFIsExist(fsid, fileid string) (bool, error)

	// MFReaddir 读取目录内容
	// 参数: fsid-文件会话ID, count-读取条目数
	// 返回值: 目录条目列表
	MFReaddir(fsid string, count int) ([]*FileInfo, error)

	// MFGetMimeType 获取文件MIME类型
	// 参数: fsid-文件会话ID
	// 返回值: MIME类型字符串
	MFGetMimeType(fsid string) (string, error)

	// MFSetCid 设置文件的IPFS CID
	// 参数: sid-会话ID, mid-弥媒ID, cid-IPFS CID
	// 返回值: 设置后的文件ID
	MFSetCid(sid, mid, cid string) (string, error)
}
