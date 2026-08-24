package lapi

import (
	"bytes"
	"encoding/gob"
)

// Package lapi provides RPC functionality for the Leither project
type LApiStub struct {
	AuthStub
	VarActStub
	MiMeiStub
	NetStub
	MsgStub  // Wave B msg：自 api.MsgStub 开放下沉
	StatStub // Wave B stat：自 api.StatStub 开放下沉
}

// MsgStub 消息接口（Wave B msg：自 api.MsgStub 开放下沉）
type MsgStub struct {
	// SendMsg 消息发送
	// 参数：发送者ID，from；接收者ID：to；appID:消息所属应用ID；msg:消息文本内容；
	//		data:消息对象内容；ppt:消息的校验信息
	// 返回值：正常则为空，出错则返回错误信息
	SendMsg func(sid string, msg *Message) error
	// ReadMsg 消息接收
	// 返回值：消息数组；返回不为空就需要重新读数据
	ReadMsg func(sid string) ([]*Message, error)
	// PullMsg 拉取消息（阻塞至 timeout）
	PullMsg func(sid string, timeout int) (*Message, error)
}

// StatStub 状态查询（Wave B stat：自 api.StatStub 开放下沉）
type StatStub struct {
	GetUserMids  func(sid, uid string, start int64, count int) ([]ScorePair, error)
	GetMiMeiStat func(sid, mid string) (stats MiMeiStats, err error)
}

// AppDataStub 应用数据操作（Wave B appdata：自 api.AppDataStub 开放下沉）
type AppDataStub struct {
	MMGetAppDataID    func(sid, aid, tp, owner, mark string, check bool) (appdataid string, err error)
	MMCreateAppData   func(sid, aid, tp, defOwner, mark string, right uint64) (did string, err error)
	MMOpenAppData     func(sid, aid, tp, owner, ver, mark string) (mmsid string, err error)
	MMOpenAppDataApp  func(sid, aid, ver, mark string) (mmsid string, err error)
	MMOpenAppDataNode func(sid, aid, ver, mark string) (mmsid string, err error)
	MMOpenAppDataUser func(sid, aid, owner, ver, mark string) (mmsid string, err error)
}

// MiMeiRefs 表示弥媒引用关系，键为版本号，值为文件ID到引用计数的映射
type MiMeiRefs = map[string]map[string]int

// Refs 表示文件引用关系，键为文件ID，值为引用计数
type Refs = map[string]int

// MiMeiStub 弥媒操作结构体
// 提供对弥媒对象的完整生命周期管理，包括创建、权限控制、版本管理、引用管理等
type MiMeiStub struct {
	MDbStub
	MFileStub

	// MMCreate 创建新的弥媒对象
	// 参数: sid-会话ID, appid-应用ID, ext-扩展标识, mark-标记字符串, tp-弥媒类型, right-权限位掩码
	// 返回值: mid-新创建的弥媒ID
	MMCreate func(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)

	// MMOpen 打开指定版本的弥媒
	// 参数: sid-会话ID, mid-弥媒ID, ver-版本号(空字符串表示最新版本), opt-可选参数
	// 返回值: 会话句柄，用于后续操作
	MMOpen func(sid, mid, ver string, opt ...string) (string, error)

	// MMOpenUrl 通过URL路径打开弥媒
	// 参数: sid-会话ID, ps-URL路径
	// 返回值: 会话句柄
	MMOpenUrl func(sid, ps string) (string, error)

	// MMClose 关闭弥媒会话
	// 参数: mmsid-弥媒会话ID
	MMClose func(mmsid string) error

	// MMSetRight 设置弥媒成员权限
	// 参数: sid-会话ID, mid-弥媒ID, member-成员标识, right-权限位掩码
	MMSetRight func(sid, mid, member string, right uint64) error

	// MMGetRight 获取成员对弥媒的权限
	// 参数: sid-会话ID, mmid-弥媒ID, uid-用户ID
	// 返回值: right-权限位掩码
	MMGetRight func(sid, mmid, uid string) (right uint64, err error)

	// MMBackup 创建弥媒快照备份
	// 参数: sid-会话ID, mid-弥媒ID, memo-备份备注信息, opt-可选参数
	// 返回值: 备份版本号
	MMBackup func(sid, mid, memo string, opt ...string) (string, error) //snapshot

	// MMRestore 从指定版本恢复弥媒
	// 参数: sid-会话ID, mid-弥媒ID, ver-要恢复的版本号
	MMRestore func(sid, mid, ver string) error

	// MMDelVers 删除指定版本
	// 参数: sid-会话ID, mid-弥媒ID, vers-要删除的版本号列表
	// 返回值: 删除的版本数量
	MMDelVers func(sid, mid string, vers ...string) (int64, error)

	// MMRelease 发布指定版本
	// 参数: sid-会话ID, mid-弥媒ID, ver-要发布的版本号
	// 返回值: 发布后的版本标识
	MMRelease func(sid, mid, ver string) (string, error)

	// MMAddRef 为弥媒添加文件引用
	// 参数: sid-会话ID, mid-弥媒ID, fileids-要引用的文件ID列表
	// 返回值: 成功添加的引用数量
	MMAddRef func(sid, mid string, fileids ...string) (int, error)

	// MMDelRef 删除弥媒的文件引用
	// 参数: sid-会话ID, mid-弥媒ID, fileids-要删除引用的文件ID列表
	// 返回值: 成功删除的引用数量
	MMDelRef func(sid, mid string, fileids ...string) (int, error)

	// MMGetRef 获取指定版本的引用关系
	// 参数: sid-会话ID, mid-弥媒ID, ver-版本号
	// 返回值: ret-文件引用关系映射
	MMGetRef func(sid, mid, ver string) (ret Refs, err error)

	// MMGetRefs 获取多个版本的引用关系
	// 参数: sid-会话ID, mid-弥媒ID, vers-版本号列表
	// 返回值: 各版本的引用关系映射
	MMGetRefs func(sid, mid string, vers ...string) (MiMeiRefs, error)

	// MMSum 计算弥媒的摘要信息
	// 参数: sid-会话ID, mid-弥媒ID, ver-版本号, tp-摘要类型
	// 返回值: 摘要字符串
	MMSum func(sid, mid, ver, tp string) (string, error) //

	// Wave B appdata：自 api.AppDataStub 开放下沉（置于末尾保持既有位置字面量兼容）
	AppDataStub

	// Wave B mimei_cmd 本地语义接口下沉（原 api.MiMeiCmdStub 中的本地部分；
	// 网络部分 Publish/Provide/FindProvs/Show/Sync 依赖 DHT，仍不开放）
	MiMeiIsProvider func(sid, mid string) (bool, error)
	MFLs            func(sid, ps string) ([]LsLink, error)
	MFCopy          func(sid, src, dst string, bFlush, bForce bool) error
	MFMkdir         func(sid, ps string, flush bool) error
}

// NOTE:这个sid是有有效期的
type LoginReply struct {
	Sid string
	Uid string
}

type AuthStub struct {
	LoginWithPPT func(strPPT string) (*LoginReply, error)
	Logout       func(sid, info string) error
	// —— 以下自 api.AuthStub2 开放下沉（Wave B auth2，评审通过）——
	SetUserInfo  func(sid string, param map[string]string) error
	SignPPT      func(sid string, info map[string]string, period int) (string, error)
	Sign         func(sid string, message []byte) (sig []byte, err error)
	PPTStr2Map   func(strPPT string) (map[string]string, error)
	SignInfo2Map func(strInfo string) (map[string]string, error)
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
	// —— Wave B 批次 3：自 api.NetStub2/IpfsStub 开放下沉（评审通过部分）——
	DNSStub       //域名与路由（net_dns）
	SwarmStub     //节点网络（字段全量下沉；接口仅开放只读 4 方法）
	IpfsNodeStub  //IPFS 节点数据读取（net_ipfs）
	DagStub       //DAG 查询（net_ipfs）
	IpfsAdd func(sid, ps string) (string, error)
}

// DNSStub 域名与路由配置（Wave B net_dns：自 api.DNSStub 开放下沉）
type DNSStub struct {
	SetDomain  func(sid, domain string, info map[string]string) error
	DelDomain  func(sid, domain string) error
	ShowDomain func(sid string) ([]string, error)
	SetRoute   func(sid, path, target string) error
	DelRoute   func(sid, path string) error
	ShowRoute  func(sid string) (map[string]string, error)
}

// SwarmStub 节点网络操作（Wave B net_swarm：整体下沉；lapi.INet 接口仅开放只读 4 方法，
// Connect/Disconnect/Filters* 仍不对 MApp 开放）
type SwarmStub struct {
	SwarmAddrs      func(sid string, pids ...string) (map[string][]string, error)
	SwarmLocal      func(sid string) ([]string, error)
	SwarmListen     func(sid string) ([]string, error)
	SwarmPeers      func(sid string) ([]string, error)
	SwarmConnect    func(sid, addr string) error
	SwarmDisconnect func(sid, addr string) error
	FiltersAdd      func(sid string, cidrs []string) error
	FiltersRm       func(sid string, cidrs []string) error
}

// IpfsNodeStub IPFS 节点数据读取（Wave B net_ipfs：自 api.IpfsNodeStub 开放下沉）
type IpfsNodeStub struct {
	INOpen    func(sid, ps string, level int) ([]byte, error)
	INGetData func(sid, ps string, start int64, count int) ([]byte, error)
}

// DagStub DAG 查询（Wave B net_ipfs：自 api.DagStub 开放下沉）
type DagStub struct {
	DagGet  func(sid, ps string) (DagNodeData, error)
	DagStat func(sid, ps string) (*DagStats, error)
}

type FilesStub struct {
	FilesCopy  func(sid, src, dst string, flush bool) error
	FilesLs    func(sid, ps string) ([]LsLink, error)
	FilesCheck func(sid, ps string, sync bool) error
	FilesMkdir func(sid, ps string, flush bool) error
	FilesRm    func(sid, ps string, recursive, flush bool) error
	FilesMv    func(sid, src, dst string, flush bool) error
	FilesFlush func(sid, ps string) (string, error)
	FilesStat  func(sid, ps string) (*StatInfo, error)
}

// IAuth 接口定义了认证相关的操作方法
type IAuth interface {
	LoginWithPPT(strPPT string) (*LoginReply, error)
	Logout(sid, info string) error
	// —— 以下自 api.IAuth 开放下沉（Wave B auth2，评审通过）——
	// SetUserInfo 设置当前会话用户的公开信息
	SetUserInfo(sid string, param map[string]string) error
		// SignPPT 以会话用户身份签名 PPT（period 单位：分钟，与服务端实现一致）
	SignPPT(sid string, info map[string]string, period int) (string, error)
	// Sign 以会话用户身份签名消息
	Sign(sid string, message []byte) (sig []byte, err error)
	// PPTStr2Map 将 PPT 字符串解析为 Map（纯解析，无状态）
	PPTStr2Map(strPPT string) (map[string]string, error)
	// SignInfo2Map 将签名信息字符串解析为 Map（纯解析，无状态）
	SignInfo2Map(strInfo string) (map[string]string, error)
}

type IVarAct interface {
	GetVar(sid, name string, args ...string) (any, error)
	Act(sid, name string, args ...string) (any, error)
	GetGobVar(sid, name string, args ...string) ([]byte, error)
	// GetGobVar(sid, name string, args ...string) ([]byte, error)
}

// INetStub 接口定义了网络相关的操作方法
type INet interface {
	IFilesStub
	IDNS      //域名与路由（Wave B 开放）
	ISwarm    //节点网络只读 4 方法（Wave B 开放；写操作不开放）
	IIpfsNode //IPFS 节点数据读取（Wave B 开放）
	IDag      //DAG 查询（Wave B 开放）
	// IpfsAdd 添加本地文件到 IPFS（Wave B 开放）
	IpfsAdd(sid, ps string) (string, error)
}

// IDNS 域名与路由配置接口（Wave B net_dns：自 api.IDNS 开放下沉）
type IDNS interface {
	SetDomain(sid, domain string, info map[string]string) error
	DelDomain(sid, domain string) error
	ShowDomain(sid string) ([]string, error)
	SetRoute(sid, path, target string) error
	DelRoute(sid, path string) error
	ShowRoute(sid string) (map[string]string, error)
}

// ISwarm 节点网络只读查询接口（Wave B net_swarm：仅开放只读方法，
// Connect/Disconnect/Filters* 不对 MApp 开放）
type ISwarm interface {
	SwarmAddrs(sid string, pids ...string) (map[string][]string, error)
	SwarmLocal(sid string) ([]string, error)
	SwarmListen(sid string) ([]string, error)
	SwarmPeers(sid string) ([]string, error)
}

// IIpfsNode IPFS 节点数据读取接口（Wave B net_ipfs：自 api.IIpfsNode 开放下沉）
type IIpfsNode interface {
	INOpen(sid, ps string, level int) ([]byte, error)
	INGetData(sid, ps string, start int64, count int) ([]byte, error)
}

// IDag DAG 查询接口（Wave B net_ipfs：自 api.IDag 开放下沉）
type IDag interface {
	DagGet(sid, ps string) (DagNodeData, error)
	DagStat(sid, ps string) (*DagStats, error)
}

// IMsg 消息接口（Wave B msg：自 api.IMsg 开放下沉）
type IMsg interface {
	SendMsg(sid string, msg *Message) error
	ReadMsg(sid string) ([]*Message, error)
	PullMsg(sid string, timeout int) (*Message, error)
}

// IStat 状态查询接口（Wave B stat：自 api.IStat 开放下沉）
type IStat interface {
	GetUserMids(sid, uid string, start int64, count int) ([]ScorePair, error)
	GetMiMeiStat(sid, mid string) (stats MiMeiStats, err error)
}

// IAppData 应用数据接口（Wave B appdata：自 api.IAppData 开放下沉）
type IAppData interface {
	MMGetAppDataID(sid, aid, tp, owner, mark string, check bool) (appdataid string, err error)
	MMCreateAppData(sid, aid, tp, defOwner, mark string, right uint64) (did string, err error)
	MMOpenAppData(sid, aid, tp, owner, ver, mark string) (mmsid string, err error)
	MMOpenAppDataApp(sid, aid, ver, mark string) (mmsid string, err error)
	MMOpenAppDataNode(sid, aid, ver, mark string) (mmsid string, err error)
	MMOpenAppDataUser(sid, aid, owner, ver, mark string) (mmsid string, err error)
}

type IFilesStub interface {
	FilesCopy(sid, src, dst string, flush bool) error
	FilesLs(sid, ps string) ([]LsLink, error)
	FilesCheck(sid, ps string, sync bool) error
	FilesMkdir(sid, ps string, flush bool) error
	FilesRm(sid, ps string, recursive, flush bool) error
	FilesMv(sid, src, dst string, flush bool) error
	FilesFlush(sid, ps string) (string, error)
	FilesStat(sid, ps string) (*StatInfo, error)
}
