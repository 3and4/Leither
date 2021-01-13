认证体系  
======== 
## 简介  
在整个分布式网络中，两个节点（或对象）之间是最初是没有关联的，需要解决以下问题：  
标识一个用户或对象；  
证明一个用户或对象身份；  
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


## 生成key  
**生成密钥对**  
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

**生成公钥**  
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
Leither lpki gencert -k certkey.key -c ca.cert -m "name=weixin" -o output.cert
```    

**生成对外的公钥证书**  
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
