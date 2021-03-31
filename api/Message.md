消息模块
========  

## 一、发送消息 
```golang 
type Message struct {
	Tm    time.Time //消息发生的时间
	From  string
	To    string
	AppID string      //空表示系统消息
	Msg   string      //表示命令，是由appid约定的，如果appid为空，则是系统消息?
	Data  interface{} //应用自定义的数据格式
	Sign  string      //发送者签名
	//如果考虑消息的自净，应当加入一个消息的有效期，到期自动删除
}  

type Msgs []*Message
//SendMsg 消息发送
//参数：发送者ID，from；接收者ID：to；appID:消息所属应用ID；msg:消息文本内容；
//		data:消息对象内容；ppt:消息的校验信息
//返回值：正常则为空，出错则返回错误信息
SendMsg(sid string, msg *Message) error
```  

## 二、读取消息 
```golang
//ReadMsg：消息接收
//返回值：消息数组；返回不为空就需要重新读数据
ReadMsg(sid string) ([]*Message, error)
```    

## 三、拉取消息 
```golang
PullMsg(sid string, timeout int) (*Message, error)
```  
