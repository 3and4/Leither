网络模块(草)
========  

## 一、域名模块  
### 1.1 设置域名  
SetDomain(sid, domain string, info map[string]string) error
### 1.2 删除域名  
DelDomain(sid, domain string) error

### 1.3 显示域名  
ShowDomain(sid string) ([]string, error)
### 1.4 更换域名地址  
UpdateAddrs func(sid string, addrs ...string) error
### 1.5 其它  
//通过getvar 可以获取到节点的ip,域名等信息

## 二、节点管理模块  
### 2.1 邀请节点
Invite(sid string, validity uint, credit uint, friendCount int) (string, error)
### 2.2 接受邀请
Accept(sid, inviteInfo string) error
### 2.3 添加节点
AddNode(id, addrs, invCode string) error
### 2.4 节点登录
type NodeLoginReply struct {  
	Services Services  
	Sid      string  
}  
NLogin(ppt string) (*NodeLoginReply, error)  
### 2.5 获取互信
NGetGift func(ppt string) (string, error)  

## 三、隧道    
节点可以为其它节点开设一条隧道，这样第三方节点可以访问到没有独立ip的节点  
### 3.1 打开隧道
OpenTunnel(sid, sUrl, uuid, tagurl string) (string, error)
### 3.2 关闭隧道
CloseTunnel(sid, sUrl, turl string) error
### 3.3 设置外网地址
SetExtAddr  func(sid, addr string) error

## 四、其它    
### 4.1 利用节点发送Get调用  
Proxyget(sid, url string) (string, error)
### 4.2 利用节点发送Post调用  
Proxypost(sid, url string, header map[string]string, postContent interface{}) (string, error)
### 4.3 利用节点发送邮件  
SendMail(sid, to, subject, body, mailtype string) error
