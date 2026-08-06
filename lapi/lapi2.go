// lapi2.go
package lapi

import "time"

type Service map[string]string
type Services map[string]map[string]string

// 节点api
type NodeLoginReply struct {
	Services Services
	Sid      string
}

type AddrInfo struct {
	ID    string //ipfs中可以识别的id
	Addrs []string
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

type PinLsObject struct {
	Cid  string
	Type string
}

type RepoStat struct {
	RepoSize   uint64 // size in bytes
	StorageMax uint64 // size in bytes
	NumObjects uint64
	RepoPath   string
	Version    string
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

type MiMeiStat = map[uint8]int64       //tp->StatInfo
type MiMeiStats = map[uint64]MiMeiStat //verseq->StatInfo
