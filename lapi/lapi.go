package lapi

import "time"

// 这里定义所有的api接口
type Service map[string]string
type Services map[string]map[string]string

// 目前对外的只有之两个函数
type IAuth interface {
	IAuthStub
}

// type IDNS interface {
// 	SetDomain(sid, domain string, info map[string]string) error
// 	DelDomain(sid, domain string) error
// 	ShowDomain(sid string) ([]string, error)
// 	//UpdateAddrs (sid string, addrs ...string) error
// 	//UpdateIp    (sid, ips string) error
// 	//AddGateway (sid, url string) error
// 	//通过getvar 可以获取到节点的ip,域名等信息
// }

// 节点api
type NodeLoginReply struct {
	Services Services
	Sid      string
}

// type INode interface {
// 	//	AddNodeStub
// 	NLogin(ppt string) (*NodeLoginReply, error)
// 	NGetGift(ppt string) (string, error) //TODO:暂时保留
// }

type AddrInfo struct {
	ID    string //ipfs中可以识别的id
	Addrs []string
}

// type Ipfs interface {
// 	DHT
// 	Swarm
// 	MiMeiCmd //这个不应该在这里应该移出云
// 	IPFSPinCmd
// 	IpfsNode
// 	Repo
// 	Dag
// 	IpfsAdd(sid, ps string) (string, error)
// 	IpfsProvide(sid string, rec bool, keys ...string) error
// 	IpfsFindProvs(sid, cid string, n int) ([]AddrInfo, error)
// 	IpfsPublish(sid string, mids ...string) error
// }

// type IDHT interface {
// 	DhtFindPeer(sid, dhts, pid string) (*AddrInfo, error)
// 	DhtGetAllKeys(sid string) ([]string, error)
// 	DhtGet(sid, key string) ([]string, error)
// 	DhtGets(sid string, keys ...string) (map[string][]string, error)
// 	DhtPutValue(sid, dhts, key, value string) error
// 	DhtGetValue(sid, dhts, key string) (string, error)
// 	DhtProvide(sid, dhts, strCid string, bAnnounce bool) error
// 	DhtFindProvs(sid, dhts, strCid string, count int) ([]AddrInfo, error)
// 	DhtNetworkSize(sid, strdhts string) (nSize int, err error)
// }

// type ISwarm interface {
// 	SwarmAddrs(sid string, pids ...string) (map[string][]string, error)
// 	SwarmLocal(sid string) ([]string, error)
// 	SwarmListen(sid string) ([]string, error)
// 	SwarmPeers(sid string) ([]string, error)
// 	SwarmConnect(sid, addr string) error
// 	SwarmDisconnect(sid, addr string) error
// 	FiltersAdd(sid string, cidrs []string) error
// 	FiltersRm(sid string, cidrs []string) error
// 	//SwarmFilters   (sid string) ([]string, error)
// }

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

// type IMiMeiCmd interface {
// 	//发布弥媒到网上
// 	MiMeiPublish(sid, dhts string, mid string) ([]DhtReply, error)
// 	MiMeiUnpublish(sid, dhts string, mid string) ([]DhtReply, error)

// 	//告知网络，当前节点提供这个弥媒的数据支撑
// 	MiMeiProvide(sid, dhts, mid string) ([]DhtReply, error)
// 	MiMeiUnprovide(sid, dhts, mid string) ([]DhtReply, error)

// 	//查找弥媒的提供者
// 	MiMeiFindProvs(sid, dhts, mid string, count int) (Ver2ProvsInfo, error)

// 	//网络查询这个弥媒的信息。目前是版本信息，后续加入提供者
// 	MiMeiShow(sid, dhts, mid string) (*MiMeiShowReply, error)

// 	//同步弥媒信息
// 	MiMeiSync(sid, dhts, mid string, param map[string]string) error

// 	//查询弥媒信息，针对ipfs文件系统
// 	MFLs(sid, ps string) ([]LsLink, error)
// 	MFCopy(sid, src, dst string, bBackup, bForce bool) error
// 	MFMkdir(sid, ps string, flush bool) error

// 	//读出内容，需要是文件类型
// 	//命令中已弃用
// 	//MiMeiGet (sid, ps string) ([]byte, error)
// }

type PinLsObject struct {
	Cid  string
	Type string
}

// type IIPFSPinCmd interface {
// 	IpfsPinAdd(sid string, r bool, ps string) error
// 	IpfsPinRm(sid string, r bool, ps string) error
// 	IpfsPinLs(sid string, opt string) ([]PinLsObject, error)
// 	IpfsPinType(sid string, ps, tp string) (string, error)
// }

//	type IIpfsNode interface {
//		INOpen(sid, ps string, level int) ([]byte, error)
//		INGetData(sid, ps string, start int64, count int) ([]byte, error)
//	}
type RepoStat struct {
	RepoSize   uint64 // size in bytes
	StorageMax uint64 // size in bytes
	NumObjects uint64
	RepoPath   string
	Version    string
}

// type Repo interface {
// 	RepoGC(sid, opt string) error
// 	RepoLs(sid, opt string) error
// 	RepoStat(sid string) (*RepoStat, error)
// }

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

// type IDag interface {
// 	DagGet(sid, ps string) (DagNodeData, error)
// 	DagStat(sid, ps string) (*DagStats, error)
// }

// 应用
type App struct {
	//这个json标签可以去掉了
	ID      string //`json:"id"` //应用id
	Name    string //应用名称
	Author  string //开发者id
	Release string //release
	Last    string //last ver
}

type ILApp interface {
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

// type ITask interface {
// 	//RunScript参数有冲突，暂时屏蔽，后续调整
// 	// RunScript           (ext, script string, glob map[string]string, args []any) (ret any, err error)
// 	RegisterScript(sid, script string, info *ScriptInfo) (id string, err error)
// 	RunByScriptID(sid, scriptId string, glob map[string]string) (ret any, err error)
// 	TaskRunInBackground(sid string, setting TaskSetting) error
// }

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

// type IMsg interface {
// 	//SendMsg 消息发送
// 	//参数：发送者ID，from；接收者ID：to；appID:消息所属应用ID；msg:消息文本内容；
// 	//		data:消息对象内容；ppt:消息的校验信息
// 	//返回值：正常则为空，出错则返回错误信息
// 	SendMsg(sid string, msg *Message) error
// 	//ReadMsg：消息接收
// 	//返回值：消息数组；返回不为空就需要重新读数据
// 	ReadMsg(sid string) ([]*Message, error)
// 	PullMsg(sid string, timeout int) (*Message, error)
// }
// IAppData //应用数据

// type IAppData interface {
// MMGetAppDataID(sid, aid, tp, owner, mark string, check bool) (appdataid string, err error)
// MMCreateAppData(sid, aid, tp, defOwner, mark string, right uint64) (did string, err error)
// MMOpenAppData(sid, aid, tp, owner, ver, mark string) (mmsid string, err error)
// MMOpenAppDataApp(sid, aid, ver, mark string) (mmsid string, err error)
// MMOpenAppDataNode(sid, aid, ver, mark string) (mmsid string, err error)
// MMOpenAppDataUser(sid, aid, owner, ver, mark string) (mmsid string, err error)
// }

type IMiMei interface {
	IMiMeiStub
}

type IVarAct interface {
	IVarActStub
}

// type ISystem interface {
// 	Exit(string, int) error                  //停止本服务
// 	Restart(string) error                    //重启
// 	Update(sid, ver, from, opt string) error //自动升级
// 	Removeuser(sid, uid string) error
// 	//Test       (sid, bid string) error
// }

type MiMeiStat = map[uint8]int64       //tp->StatInfo
type MiMeiStats = map[uint64]MiMeiStat //verseq->StatInfo

// type IStat interface {
// 	GetUserMids(sid, uid string, start int64, count int) ([]ScorePair, error)
// 	GetMiMeiStat(sid, mid string) (stats MiMeiStats, err error)
// }

// 节点内使用的api
// 写示例时逐个添加
type LApi interface {
	// IBackEnd

	IAuth   //认证部分
	IVarAct //封装了大批Api
	IMiMei  //弥媒
	INet
}
type INet = INetStub

// 跨节点获取的api,通过rpc方式获取生成
type IRPC interface {
	IAuthStub   //认证部分
	IVarActStub //封装了大批Api
	IMiMeiStub  //弥媒
	INetStub    //网络部分	TunnelStub，	DNSStub,NodeStub

	// LApp   //应用部分
	// Task   //任务管理
	// Msg    //消息
	// System //废弃的都放这里了
	// Stat
}

// 下面是后端专用
type IBackEnd interface {
	ISession
	IBEAppData
}
