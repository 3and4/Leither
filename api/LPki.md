## lpki命令集  
lpki是Leither关于安全认证的命令集  
<a href="../doc/Pki.md"> 详细的安全认证文档</a>  

用户可以通过lpki命令生成,所有的身份信息都保存在密钥中。  

**1、生成密钥对**  
<a id="genkey"></a> 
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
    
**2、导出公钥**    
<a id="genpk"></a> 
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

**3、生成私钥证书**  
<a id="gencert"></a> 
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

**4、导出公钥证书**  
<a id="genpkcert"></a> 
公钥证书是发给其他人，用于验证签发的信息
```bash  
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

**4、生成通行证**  
<a id="signppt"></a>
```bash  
Leither lpki signppt -c ca.cert -p 720 -m "CertFor=Self,Userid=h3PPmr6HVHrmaV_WAbnEP6t3x87," -o test.ppt
```
-p 通行证的有效期，单位小时
-m 通行证信息。信息中的内容定义基于业务
-c 为通行证信息背书的证书  


输出结果如下:
略   

**5、验证通行证**  
<a id="verifyppt"></a>
```bash  
$ ./Leither8000.exe lpki verifyppt -c ca.cert  -i test.ppt
```
-c 为通行证信息背书的证书  


输出结果如下:
略   

**6、显示对象**  
<a id="show"></a>  
```bash  
Leither lpki show -i in.cert
```
-i 输入文件

输出结果如下:
```bash  
show
file type is Key
id= h7Nkwe4rfb1d15En8yQ65Lhhq9b
pk=yPrZX6l9AWNwNVwpOh10a1Mdas56i6nEmV3LUNOjFZg=
sk=G3p1HaeaR0kuvNVrB-YxDV0s3VpwjCuXweFIQySLa5LI-tlfq
X0BY3A1XCk6HXRrUx1qznqLqcSZXctQ06MVmA==
```
