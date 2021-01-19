认证体系  
======== 
## 简介  
在整个分布式网络中，两个节点（或对象）之间是最初是没有关联的，需要解决以下问题：  
标识一个用户或对象  
证明一个用户或对象身份  
描述用户权限  
描述用户之间的关系  
描述用户信用  
描述用户互相服务能力  

## 基础知识
**密钥**  
这里讲的密钥是非对称加密公钥  
使用的密钥算法是ed25519

**对象标识**  
对于确定的文件或内容，我们对之取摘要，再转化成可视的28字节字符串，可以用来标识一个确定内容的对象。  
  
**身份标识**  
一对密钥代表一个身份，密钥为二进制，也过长，不适合去标识  
公钥可以取摘要再转化成可视的28字节字符串  
这个ID唯一标识了密钥对。可以用来表示用户或节点的身份。  
以上是参考了比特币钱包的方式。  
  
**证书**  
证书通常由以下内容构成：   
内容信息，通常是一系列的key value串。  
签名信息  
签名者信息  
  
**自签名证书**  
证书中的内容如果是证明用户自己的身份，这个证书就是自签名证书
  
**CA证书**  
用于证明他人身份的证书，是CA证书  
  
**通行证**  
通行证是一个登录某个节点，具备某些权限的证书。  
通行证的内容通常包括访问者身份，时效，服务内容  
  
**Mac通行证**  
在内网中，手机电脑等设备访问节点时，可以直接获取节点签发的基于网卡mac能行证。可以和用户绑定，可以实现自动登录。  

**微信节点**  
在Leither架构中和公众号绑定的服务节点叫微信节点  
微信节点并不保存用户信息和数据，主要功能是提供用户设备和微信的对接入口和身份鉴别  

**微信通行证**  
微信用户在访问公众号时，腾讯服务器会向微信节点发送代表用户的唯一信息。  
这个唯一信息如果绑定在用户的信息中，可以做到微信用户自动登录用户节点。
微信节点签发的代表微信用户身份的通行证叫微信通行证  

  
## 生成key  
节点的密钥放在根目录，名称为hostkey.cfg  
Leither第一次启动时，系统检查根目录，如果没有这个文件会自动生成一个。   
在根目录下执行lpki命令，参数中需要key,如果不特别指定，缺省使用节点密钥。  

**生成密钥对**  
<a href="../api/LPki.md#genkey"> 命令行生成密钥对</a>  
```bash  
Leither lpki genkey -o my.key   
```  
输出结果如下：  
```bash  
GenKey private key file: new.key

id= xTTk7gZTywmHiR_y54TIpxQWRaa
Key= pk=kVIfR1dAIypas6nPhDPwiKKY
54Xco7NK80f9pUfePJY=
sk=fbyI_vwOQ2tTtSFyTrnwc9VeVFbC-
KgkRRS0WGAaROWRUh9HV0AjKlqzqc-EM
_CIopjnhdyjs0rzR_2lR948lg==

create key to file [new.key] ok

```  

**导出公钥**  
```bash  
Leither lpki genpk -i new.key -o new.pub
```  
输出结果如下：  
```bash  
genpk private key file: new.key public Key file  new.pub
Read KeyFile [new.key]ok
Key:
pk=2T1U4-IUMfYpppIcPi8ITJazt8I2PpOI8W1jctmm1Hs=
sk=eoddzYXYh2XTDp_aIGUxUey0XSDotYyCRl4oIQ9KodDZPVTj4hQx9immkhw-LwhMlrO3wjY-k4jxb
WNy2abUew==

Public Key = pk=2T1U4-IUMfYpppIcPi8ITJazt8I2PpOI8W1jctmm1Hs=

genpk public key [new.pub] ok!
```  

## 生成证书  

**生成包含私钥的证书**  
```bash  
Leither lpki gencert -k new.key -m "name=weixin" -o o
utput.cert
```
-k 证书的私钥
-m 证书信息
-c ca证书，如果为空是自签名证书  
生成的证书文件中包含私钥，对外签名用，不可以公开。  


输出结果如下:
```bash  
GenCert
Read message [name=weixin]
name=weixin
name weixin
Read cert key
Key:
pk=2T1U4-IUMfYpppIcPi8ITJazt8I2PpOI8W1jctmm1Hs=
sk=eoddzYXYh2XTDp_aIGUxUey0XSDotYyCRl4oIQ9KodDZPVTj4hQx9immkhw-Lwh
MlrO3wjY-k4jxbWNy2abUew==

caFile is empty, CreateSelfCertByKey
Sign Cert to file [output.cert] ok!
Read Cert File output.cert
id= h3PPmr6HVHrmaV_WAbnEP6t3x87
{"CertPK":"CQwABnNvZGl1bSn_gQMBAQ1Tb2RpdW1LZXlQYWlyAf-CAAECAQJTawE
KAAECUGsBCgAAACX_ggIg2T1U4-IUMfYpppIcPi8ITJazt8I2PpOI8W1jctmm1HsA"
,"CertPKID":"h3PPmr6HVHrmaV_WAbnEP6t3x87","name":"weixin"}
```    

**导出对外的公钥证书**  
公钥证书是发给其他人，用于验证签发的信息
```bash  
Leither lpki gencert -k new.key -m "name=weixin" -o o
utput.cert
Leither lpki genpkcert -i in.cert -o out.pkcert
```

输出结果如下:
```bash  
genpkcert cert file: in.cert

Read Cert file  in.cert

CertPKID:h3PPmr6HVHrmaV_WAbnEP6t3x87
name    :weixin
CertPK  :CQwABnNvZGl1bSn_gQMBAQ1Tb2RpdW1LZXlQYWlyAf-CAAECAQJTawEKAAECUGsBCgAA
ACX_ggIg2T1U4-IUMfYpppIcPi8ITJazt8I2PpOI8W1jctmm1HsA
Signature: [119 104 8 88 233 100 228 116 248 126 66 252 167 18 170 147 124 7
75 41 21 162 49 96 35 244 75 214 253 203 150 243 236 211 160 111 145 206 230
58 150 248 168 235 89 159 192 115 184 26 15 53 68 161 85 25 183 23 217 24 237
 206 224 5]
SignerID: h3PPmr6HVHrmaV_WAbnEP6t3x87
KeyType : sodium
Version : 1
pk=2T1U4-IUMfYpppIcPi8ITJazt8I2PpOI8W1jctmm1Hs=
sk=eoddzYXYh2XTDp_aIGUxUey0XSDotYyCRl4oIQ9KodDZPVTj4hQx9immkhw-LwhMlrO3wjY-k4
jxbWNy2abUew==

==>PkCert
CertPK  :CQwABnNvZGl1bSn_gQMBAQ1Tb2RpdW1LZXlQYWlyAf-CAAECAQJTawEKAAECUGsBCgAA
ACX_ggIg2T1U4-IUMfYpppIcPi8ITJazt8I2PpOI8W1jctmm1HsA
CertPKID:h3PPmr6HVHrmaV_WAbnEP6t3x87
name    :weixin
Signature: [119 104 8 88 233 100 228 116 248 126 66 252 167 18 170 147 124 7
75 41 21 162 49 96 35 244 75 214 253 203 150 243 236 211 160 111 145 206 230
58 150 248 168 235 89 159 192 115 184 26 15 53 68 161 85 25 183 23 217 24 237
 206 224 5]
SignerID: h3PPmr6HVHrmaV_WAbnEP6t3x87
KeyType : sodium
Version : 1

genpkcert public key Cert[out.pkcert] ok!
```  

## 生成通行证  
**命令行生成**  
<a href="../api/LPki.md#signppt"> 命令行生成通行证</a>  

```bash  
Leither lpki signppt -c ca.cert -p 720 -m "CertFor=Self,Userid=h3PPmr6HVHrmaV_WAbnEP6t3x87," -o test.ppt
```
-p 通行证的有效期，单位小时
-m 通行证信息。信息中的内容定义基于业务
-c 为通行证信息背书的证书  


输出结果如下:
略   

**获取Mac通行证**  
可以通过Rpc或http两种Api方式获取Mac通行证


## 申请服务  
向节点申请权限  
用户登录之后，在使用节点资源前，先要申请服务访问权限  
  
命令行示例：
```bash  
Leither lpki reqservice -c my.cert -m RequestService=mimei -n http://192.168.3.29:4800/  
```
-c 后面是申请者的通行证,证明申请者的身份  
-n 后面是申请的节点地址，包括本地调试程序的开发节点或实际运行的应用节点  
-m 服务的内容  
格式：RequestService=服务1(属性1:数值&属性2),服务2(属性1&属性2)；

节点处理服务权限的角本都放在“根目录/service/RequestService”,参考结构如下  
根目录  
｜  
｜--service---------------服务目录,存放服务端角本  
｜   ｜----RequestService--处理服务请求的角本  
  
服务定义  
```golang  
Service_DNS            = "dns"            //查找节点，域名解析
Service_Tunnel         = "tunnel"         //隧道
Service_Mimei          = "mimei"          //弥媒
```  
