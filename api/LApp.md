## lapp命令集  

**1、上传应用**  
**2、更新文件**  
**3、卸载应用**  
**4、显示应用**  
**5、域名设置**   
**6、版本备份**  
**7、版本发布**   
指定一个版本为发布版（release）
<a id="release"></a>  
```bash  
Leither lapp release -a 应用名 -v 应用版本 -p 通行证 -n 节点地址

```  
-a 应用名。同一个开发者，应用不能重名。  
-v 要发布的应用版本。可以是特殊版本：cur,last。  
-p 应用开发者的通告证。在系统目录中缺省使用节点身份运行。
-n 节点地址。在系统目录中缺省使用当前目录下的节点地址。
  
输出结果如下:
```bash  
Release
appid= htpoEXiE6IlCAqVbCvjvkY_XNfu
version= last
Current Release :[9]
ShowApp [dav] begin
Author  : h7Nkwe4rfb1d15En8yQ65Lhhq9b
AppName : dav
AppID    htpoEXiE6IlCAqVbCvjvkY_XNfu
Last    : 9
Release : 9
url     : http://192.168.1.7:8000/entry?author=h7Nkwe4rfb1d15En8yQ65Lhhq9b&app=dav&ver=cur       
Version:[1,2,3,4,5,6,7,8,9,cur,last,release,]
```
**8、版本恢复**   

