<!--
    AuthStub      //认证部分
	NetStub       //网络部分	TunnelStub，	DNSStub,NodeStub
	AppStub       //应用部分
	TaskStub      //任务管理
	MsgStub       //消息
	MiMeiStub     //弥媒
	VarActStub    //封装了大批Api
	UserGroupStub //用户组
	DataRightStub //数据权限
	SystemStub    //废弃的都放这里了
-->

以下Api使用Golang格式，其它语言的参数和调用格式类似  

## 一、认证部分  
### 1.用户生成  
用户可以通过
### 1.1 通过命令行生成
用户可以通过lpki命令生成,所有的身份信息都保存在密钥中。  
<a href="./Pki.md"> 更完整的命令行信息在这里</a>  
+  生成key  
包含公钥和私钥，公钥取摘要生成代表用户身份的唯一id  
    ```bash  
    Leither lpki genkey -o my.key  
    ```  
+ 生成自签名cert  
    ```bash  
    Leither lpki gencert -k my.key -c my.cert -m "name=lsb" -o my.cert  
    ```  
+ 生成登录用passport(ppt)  
    通行证中包含了一个时效信息,缺省72小时
    ```bash  
    Leither lpki signppt -c ca.cert -p 720 -m "CertFor=Self,Userid=h3PPmr6HVHrmaV_WAbnEP6t3x87," -o test.ppt
    ```  

### 1.2 通过Register生成
也可以通过Register函数生成，这时生成的密钥保存在用户节点中，需要对这个节点绝对信任。
Register  
注册用户，返回值是用户id和错误

```golang
Register(par ...string) (string, error)
```

#### 1.2.1 用户名密码注册模式  
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

#### 1.2.2 map注册模式  
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


### 2.用户登录  
Login  
```golang
Login(pa ...string) (*LoginReply, error)
type LoginReply struct {
	Sid string  //Session ID，用于绝大多数Api
	Uid string  //当前用户的ID,用户的身份标识
}
```

#### 1.1 Guest登录模式  
```golang
reply, err := Login()
if err != nil{
    fmt.Println("Login err=%s", err.Error())
    return
}
fmt.Println("sid=%s, uid=%s", reply.Sid, reply.Uid)
```

#### 1.2 通行证登录模式  
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


### 3.退出登录  
```golang
Logout(sid, info string) error  
```
调用之后，sid会失效，并向所有的消息队列发送下线通知

```golang
info := "这个信息会记录在日志文件中"
err := Logout(sid, info) 
if err != nil{
    fmt.Println("Logout err=%s", err.Error())
    return
}
```  

### 4.设置用户信息  
sid为会话id
param为参数，定义参考附件中的用户信息项

```golang
SetUserInfo(sid string, param map[string]string) error
```


### 5.申请服务  
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