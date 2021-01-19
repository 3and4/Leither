## 认证Api  
认证体系的主要功能是去定义一个用户和节点，定义节点服务和权限机制，定义节点之间的访问机制。  
详细文档参见<a href="./doc/Pki.md"> Pki体系和相关命令行</a>  
### 1、Login  
Login  
```golang
Login(pa ...string) (*LoginReply, error)
type LoginReply struct {
	Sid string  //Session ID，用于绝大多数Api
	Uid string  //当前用户的ID,用户的身份标识
}
```

用户可以通过多种方式登录访问节点：  
+ Guest方式登录
+ 通行证方式登录
+ 用户名密码方式
+ 其他兼容拓展方式  
  
登录函数  
登录函数会返回一个会话id，会话id代表用户的身份和权限，通过会话id可以访问绝大多数API，会话id也代表用户操作的上下文。


#### 1.1 Guest登录模式  
对于Guest用户，可以不登录以sid为空的方式访问节点。但是没有sid, 就没有上下文，节点没办法把用户当成一个完整的用户来对待，相关操作会受很大限制。  
调用示例：
```golang
reply, err := Login()
if err != nil{
    fmt.Println("Login err=%s", err.Error())
    return
}
fmt.Println("sid=%s, uid=%s", reply.Sid, reply.Uid)
```

#### 1.2 通行证登录模式  
通行证有多种获取方式：
+ 密钥和证书生成ppt，包含用户信息
+ 节点签发的Mac通行证,包含内网用户的设备信息
+ 微信节点签发的微信通行证，包含代表微信用户的唯一身份
+ 其它代表用户身份的唯一信息
没有包含用户信息的通行证，需要Register或SetUserInfo绑定一下用户方可登录。  
系统中有用于测试的mac通行证，后续测试案例中会使用相关信息进行测试。  

```golang
//strPPT为代表用户身份的通行证
reply, err := Login(strPPT) 
if err != nil{
    fmt.Println("Login err=%s", err.Error())
    return
}
fmt.Println("sid=%s, uid=%s", reply.Sid, reply.Uid)
```

#### 1.3 用户名密码登录模式  
这种方式用户的密钥保存在当前节点中，通过用户名密码进行保护。 通常这个节点是用户自己的节点或者高度可信的节点。  
```golang
//strPPT为代表用户身份的通行证
reply, err := Login("lsb", "123456") 
if err != nil{
    fmt.Println("Login err=%s", err.Error())
    return
}
fmt.Println("sid=%s, uid=%s", reply.Sid, reply.Uid)
```

#### 1.4 兼容拓展模式  
略  
  
  
### 2.Register  
注册用户，返回值是用户id和错误
```golang
Register(par ...string) (string, error)
```

#### 2.1 用户名密码注册模式  
```golang
name := "lsb"
pass := "123456"
email:= "lsb@gmail.com" //这个参数当前可以忽略
userid, err := Register(name, pass, email)
if err != nil{
    fmt.Println("Register err=%s", err.Error())
    return
}
fmt.Println("userid=%s", userid)
```

#### 2.2 map注册模式  
注册相关的参数信息比较多，后期也需要拓展，可以使用map方式存放参数
参数定义可以参考附件中的用户信息项
```golang
param := make(map[string]string)
param["newname"] = "lsb"
param["newpass"] = "123456"
ay, _ := json.Marshal(param)

userid, err := Register(string(ay))
if err != nil{
    fmt.Println("Register err=%s", err.Error())
    return
}
fmt.Println("userid=%s", userid)
```

### 3.Logout
退出登录.  
调用之后，sid会失效，并向所有的消息队列发送下线通知

```golang
Logout(sid, info string) error  
```
调用示例：  
```golang
info := "这个信息会记录在日志文件中"
err := Logout(sid, info) 
if err != nil{
    fmt.Println("Logout err=%s", err.Error())
    return
}
```  

### 4.SetUserInfo  
设置用户信息  
sid为会话id
param为参数，定义参考附件中的用户信息项

```golang
SetUserInfo(sid string, param map[string]string) error
```

### 5.RequestService  
用户向节点申请服务   
ppt是申请者自己签的ppt,附加申请的服务内容和申请者身份   
返回值中第一层map的key是服务名，第二重是服务的属性参数  
```golang  
RequestService(ppt string) (map[string]map[string]string, error) //服务清求
```  

<!--
### 6.签发通行证  
```golang
SignPPT(sid string, info string, period int) (string, error)
```
-->
### 附件
#### 用户信息项
用于Api参数（Register,SetUserInfo）  
  
```golang
const (
	AuthKeyEmail   = "email"   //email,联系方式
	AuthKeyFather  = "father"  //父对象id,有权修改子用户信息
	AuthKeyName    = "name"    //用户名，鉴权用
	AuthKeyPass    = "pass"    //密码，鉴权用
	AuthKeyPPT     = "ppt"     //通行证，鉴权用
	AuthKeyUserID  = "userid"  //用户id，当前操作用户的id，针对没有鉴权信息的用户,修改用户信息用到
	AuthKeyNewName = "newname" //新的用户名，设置用
	AuthKeyNewPass = "newpass" //新的密码，设置用
	AuthKeyNewPPT  = "newppt"  //新的通行证，设置用
	AuthKeyDelPPT  = "delppt"  //通行证Bindinfo，删除用
)
```  
