package lapi

import "time"

// 这里定义所有的api接口
type Service map[string]string
type Services map[string]map[string]string

type Auth interface {
	RequestService(ppt string) (Services, error) //服务清求
	Register(par ...string) (string, error)
	Login(pa ...string) (*LoginReply, error)
	LoginWithPPT(strPPT string) (*LoginReply, error)
	Logout(sid, info string) error
	SetUserInfo(sid string, param map[string]string) error
	SignPPT(sid string, info map[string]string, period int) (string, error)
	Sign(sid string, message []byte) (sig []byte, err error)
	PPTStr2Map(strPPT string) (map[string]string, error)
	SignInfo2Map(strInfo string) (map[string]string, error)
}

type LNet interface {
	Tunnel
	DNS
	Node
	Ipfs
	// FilesCmdStub
	Proxyget(sid, url string) (string, error)
	Proxypost(sid, url string, header map[string]string, postContent any) (string, error)
	SendMail(sid, to, subject, body, mailtype string) error
}

type Tunnel interface {
	OpenTunnel(sid, sUrl, uuid, tagurl string) (string, error)
	CloseTunnel(sid, sUrl, turl string) error
	RequestNodeTunnel(sid, nodeid, uuid string) (addr string, err error)
}

type DNS interface {
	SetDomain(sid, domain string, info map[string]string) error
	DelDomain(sid, domain string) error
	ShowDomain(sid string) ([]string, error)
	//UpdateAddrs (sid string, addrs ...string) error
	//UpdateIp    (sid, ips string) error
	//AddGateway (sid, url string) error
	//通过getvar 可以获取到节点的ip,域名等信息
}

// 节点api
type NodeLoginReply struct {
	Services Services
	Sid      string
}

type Node interface {
	//	AddNodeStub
	NLogin(ppt string) (*NodeLoginReply, error)
	NGetGift(ppt string) (string, error) //TODO:暂时保留
}

type AddrInfo struct {
	ID    string //ipfs中可以识别的id
	Addrs []string
}

type Ipfs interface {
	DHT
	Swarm
	MiMeiCmd //这个不应该在这里应该移出云
	IPFSPinCmd
	IpfsNode
	Repo
	Dag
	IpfsAdd(sid, ps string) (string, error)
	IpfsProvide(sid string, rec bool, keys ...string) error
	IpfsFindProvs(sid, cid string, n int) ([]AddrInfo, error)
	IpfsPublish(sid string, mids ...string) error
}

type DHT interface {
	DhtFindPeer(sid, dhts, pid string) (*AddrInfo, error)
	DhtGetAllKeys(sid string) ([]string, error)
	DhtGet(sid, key string) ([]string, error)
	DhtGets(sid string, keys ...string) (map[string][]string, error)
	DhtPutValue(sid, dhts, key, value string) error
	DhtGetValue(sid, dhts, key string) (string, error)
	DhtProvide(sid, dhts, strCid string, bAnnounce bool) error
	DhtFindProvs(sid, dhts, strCid string, count int) ([]AddrInfo, error)
	DhtNetworkSize(sid, strdhts string) (nSize int, err error)
}

type Swarm interface {
	SwarmAddrs(sid string, pids ...string) (map[string][]string, error)
	SwarmLocal(sid string) ([]string, error)
	SwarmListen(sid string) ([]string, error)
	SwarmPeers(sid string) ([]string, error)
	SwarmConnect(sid, addr string) error
	SwarmDisconnect(sid, addr string) error
	FiltersAdd(sid string, cidrs []string) error
	FiltersRm(sid string, cidrs []string) error
	//SwarmFilters   (sid string) ([]string, error)
}

// TODO:必要的话，在这里增加一个最大序号
type MiMeiVersions struct {
	//StructVer  uint8            //数据结构的版本，当前为0，后期兼容用
	Versions    []*MiMeiVersion   //所有的版本
	SpecialVers []*SpecialVersion //特殊版本
	MinDifSeq   uint64            //最小的差值序号，当前数据版本大于这个序号可以差值获取数据。20230308
}

// type MiMeiRefs = map[string]map[string]int
// type Refs = map[string]int

type MiMeiVersion struct {
	Version string //记录的版本
	MacRes  string //同时写上两个变量，这是过渡期兼容。FileID后续废弃
	MacRef  string
}

type SpecialVersion struct {
	VerName string //Release, Cur
	Version string //版本
}

type MiMeiShowReply struct {
	Versions *MiMeiVersions
	Rights   map[string]uint64
	//Refs     MiMeiRefs
}

type DhtReply struct {
	DhtName string
	Info    string
}

type Ver2ProvsInfo = map[uint64][]AddrInfo

type MiMeiCmd interface {
	//发布弥媒到网上
	MiMeiPublish(sid, dhts string, mid string) ([]DhtReply, error)
	MiMeiUnpublish(sid, dhts string, mid string) ([]DhtReply, error)

	//告知网络，当前节点提供这个弥媒的数据支撑
	MiMeiProvide(sid, dhts, mid string) ([]DhtReply, error)
	MiMeiUnprovide(sid, dhts, mid string) ([]DhtReply, error)

	//查找弥媒的提供者
	MiMeiFindProvs(sid, dhts, mid string, count int) (Ver2ProvsInfo, error)

	//网络查询这个弥媒的信息。目前是版本信息，后续加入提供者
	MiMeiShow(sid, dhts, mid string) (*MiMeiShowReply, error)

	//同步弥媒信息
	MiMeiSync(sid, dhts, mid string, param map[string]string) error

	//查询弥媒信息，针对ipfs文件系统
	MFLs(sid, ps string) ([]LsLink, error)
	MFCopy(sid, src, dst string, bBackup, bForce bool) error
	MFMkdir(sid, ps string, flush bool) error

	//读出内容，需要是文件类型
	//命令中已弃用
	//MiMeiGet (sid, ps string) ([]byte, error)
}

type PinLsObject struct {
	Cid  string
	Type string
}

type IPFSPinCmd interface {
	IpfsPinAdd(sid string, r bool, ps string) error
	IpfsPinRm(sid string, r bool, ps string) error
	IpfsPinLs(sid string, opt string) ([]PinLsObject, error)
	IpfsPinType(sid string, ps, tp string) (string, error)
}

type IpfsNode interface {
	INOpen(sid, ps string, level int) ([]byte, error)
	INGetData(sid, ps string, start int64, count int) ([]byte, error)
}
type RepoStat struct {
	RepoSize   uint64 // size in bytes
	StorageMax uint64 // size in bytes
	NumObjects uint64
	RepoPath   string
	Version    string
}

type Repo interface {
	RepoGC(sid, opt string) error
	RepoLs(sid, opt string) error
	RepoStat(sid string) (*RepoStat, error)
}

// 定义成二进制流
type DagNodeData = []byte

// DagStat is a dag stat command response
type DagStat struct {
	Cid       string `json:",omitempty"`
	Size      uint64 `json:",omitempty"`
	NumBlocks uint64 `json:",omitempty"`
}

type DagStats struct {
	RedundantSize uint64     `json:"-"`
	UniqueBlocks  int        `json:",omitempty"`
	TotalSize     uint64     `json:",omitempty"`
	DagStatsArray []*DagStat `json:"DagStats,omitempty"`
}

type Dag interface {
	DagGet(sid, ps string) (DagNodeData, error)
	DagStat(sid, ps string) (*DagStats, error)
}

// 应用
type App struct {
	//这个json标签可以去掉了
	ID      string //`json:"id"` //应用id
	Name    string //应用名称
	Author  string //开发者id
	Release string //release
	Last    string //last ver
}

type LApp interface {
	UploadApp(sid, fileid, tp string) (*App, error)
	UploadAppfile(sid, AppName string, filename, fileid string) error
	UninstallApp(sid, AppName string) error
	//RunScript参数有冲突，暂时屏蔽，后续调整
	//RunScript     (ext, script string, request map[string]string, args []any) (ret any, err error)
	RunMApp(entry string, request map[string]string, args []any, opt ...string) (ret any, err error)
}
type ScriptInfo struct {
	ID   string
	Name string
	Memo string
}

type TaskSetting struct {
	Name     string            //缺省唯一
	ScriptID string            //运行的角本id
	Start    int64             //0，表示立刻执行，正表示unix秒数，负值以后定义，目前考虑事件触发：｛login，消息等｝
	Period   int               //单位秒，-1表示每月，-2表示每年；其它数值可以对应别的特殊信息
	Param    map[string]string //参数
	Status   byte              //0 执行中，1 暂停，/*2，等待任务*/
	//后续添加权限相关
}

type Task interface {
	//RunScript参数有冲突，暂时屏蔽，后续调整
	// RunScript           (ext, script string, glob map[string]string, args []any) (ret any, err error)
	RegisterScript(sid, script string, info *ScriptInfo) (id string, err error)
	RunByScriptID(sid, scriptId string, glob map[string]string) (ret any, err error)
	TaskRunInBackground(sid string, setting TaskSetting) error
}

// SendMsg
//
// 消息接口
type Message struct {
	Tm    time.Time //消息发生的时间
	From  string
	To    string
	AppID string //空表示系统消息
	Msg   string //表示命令，是由appid约定的，如果appid为空，则是系统消息?
	Data  any    //应用自定义的数据格式
	Sign  string //发送者签名
	//如果考虑消息的自净，应当加入一个消息的有效期，到期自动删除
}

func (msg *Message) IsEmpty() bool {
	return msg == nil || (len(msg.Msg) == 0 && msg.Data == nil)
}

// 自己发给自己的消息
func (msg *Message) IsSelf() bool {
	return msg.From == msg.To
}

type Msgs []*Message

type Msg interface {
	//SendMsg 消息发送
	//参数：发送者ID，from；接收者ID：to；appID:消息所属应用ID；msg:消息文本内容；
	//		data:消息对象内容；ppt:消息的校验信息
	//返回值：正常则为空，出错则返回错误信息
	SendMsg(sid string, msg *Message) error
	//ReadMsg：消息接收
	//返回值：消息数组；返回不为空就需要重新读数据
	ReadMsg(sid string) ([]*Message, error)
	PullMsg(sid string, timeout int) (*Message, error)
}

type MDb interface {
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
	ZincrBy(mmsid, key string, delta int64, member string) (ret int64, err error)
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
	CheckIntegrity(mmsid string, dataType uint8, key string, bRepair bool) error //校验数据
}

type MFile interface {
	MFOpenByPath(sid, mid, path string, flag int) (string, error)
	MFOpenMacFile(sid, mid, fileid string) (string, error)
	MFOpenTempFile(sid string) (string, error)
	MFPath2TempID(sid, localpath string) (string, error) //这个函数有争议
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

type AppData interface {
	MMGetAppDataID(sid, aid, tp, owner, mark string, check bool) (appdataid string, err error)
	MMCreateAppData(sid, aid, tp, defOwner, mark string, right uint64) (did string, err error)
	MMOpenAppData(sid, aid, tp, owner, ver, mark string) (mmsid string, err error)
	MMOpenAppDataApp(sid, aid, ver, mark string) (mmsid string, err error)
	MMOpenAppDataNode(sid, aid, ver, mark string) (mmsid string, err error)
	MMOpenAppDataUser(sid, aid, owner, ver, mark string) (mmsid string, err error)
}

type MiMei interface {
	MDb   //MDb，数据库部分
	MFile //MFile,这也是弥媒的一部分
	//MFS     //文件系统操作 废弃
	AppData //应用数据
	MMCreate(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)
	MMOpen(sid, mid, ver string, opt ...string) (string, error)
	MMOpenUrl(sid, ps string) (string, error)
	MMClose(mmsid string) error
	MMSetRight(sid, mid, member string, right uint64) error
	MMGetRight(sid, mmid, uid string) (right uint64, err error)
	MMBackup(sid, mid, memo string, opt ...string) (string, error) //snapshot
	MMRestore(sid, mid, ver string) error
	MMDelVers(sid, mid string, vers ...string) (int64, error)
	MMRelease(sid, mid, ver string) (string, error)
	MMAddRef(sid, mid string, fileids ...string) (int, error)
	MMDelRef(sid, mid string, fileids ...string) (int, error)
	MMGetRef(sid, mid, ver string) (ret Refs, err error)
	MMGetRefs(sid, mid string, vers ...string) (MiMeiRefs, error)
	MMSum(sid, mid, ver, tp string) (string, error) //
	//MMIsExist   (sid, id string) (bool, error)

	//MMVersion  (sid, mid, VerName, ver string) (string, error)
}

type VarAct interface {
	GetVar(sid, name string, args ...string) (any, error)
	//GetVarByContext (sid, name string, mapOpt map[string]string) (any, error)
	Act(sid, name string, args ...string) error
	GetGobVar(sid, name string, args ...string) ([]byte, error)
}

type System interface {
	Exit(string, int) error                  //停止本服务
	Restart(string) error                    //重启
	Update(sid, ver, from, opt string) error //自动升级
	Removeuser(sid, uid string) error
	//Test       (sid, bid string) error
}

type MiMeiStat = map[uint8]int64       //tp->StatInfo
type MiMeiStats = map[uint64]MiMeiStat //verseq->StatInfo

type Stat interface {
	GetUserMids(sid, uid string, start int64, count int) ([]ScorePair, error)
	GetMiMeiStat(sid, mid string) (stats MiMeiStats, err error)
}
type LApi interface {
	Auth   //认证部分
	LNet   //网络部分	TunnelStub，	DNSStub,NodeStub
	LApp   //应用部分
	Task   //任务管理
	Msg    //消息
	MiMei  //弥媒
	VarAct //封装了大批Api
	System //废弃的都放这里了
	Stat
}
