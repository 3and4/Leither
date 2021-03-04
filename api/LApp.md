## lapp命令集  

**1、上传应用**  
上传应用到当前（cur）版本
<a id="upload"></a>  
```bash  
Leither lapp uploadapp -i 应用目录 -p 通行证 -n 节点地址 -s    
```  
-i 应用目录。应用目录名就是应用名  
-p 应用开发者的通告证。在系统目录中缺省使用节点身份运行。  
-n 节点地址。在系统目录中缺省使用当前目录下的节点地址。  
-s 加这个选项表示要实时同步最新变动（有文件改动就自动上传到当前文件）,

输出结果如下:
```bash  
UploadApp node=[http://192.168.1.7:8000/] pptfile=[newuserforlogin.ppt] infile=[dist/dav] syncFile=false
create tar file dist/dav.tar    
（打包过程日志略）
Upload File [dist/dav.tar]
Upload File  dist/dav.tar
mac =  siNKzYDh_WWhR1_NOuw3Myp5SCS
UploadFile fszie=41939
uploadfile ok
mac=[siNKzYDh_WWhR1_NOuw3Myp5SCS], fid=[siNKzYDh_WWhR1_NOuw3Myp5SCS]Upload File ok fileid= siNKzYDh_WWhR1_NOuw3Myp5SCS        
UploadApp  siNKzYDh_WWhR1_NOuw3Myp5SCS type= .tar
upload app ok

Author  : h7Nkwe4rfb1d15En8yQ65Lhhq9b
AppName : dav
AppID    htpoEXiE6IlCAqVbCvjvkY_XNfu
Last    : 2
Release :
url     : http://192.168.1.7:8000/entry?author=h7Nkwe4rfb1d15En8yQ65Lhhq9b&app=dav&ver=cur
Done in 0.84s.
```  
最后显示的是当前版本的防问链接

**2、更新文件**  
<a id="uploadfile"></a>  
更新应用当前版本中的某个文件 
```bash  
Leither lapp uploadfile -n 应用名 -i 要更新的文件 -n 节点地址 -p 通行证  
```  
-a 应用名。同一个开发者，应用不能重名。  
-i 要更新的应用名   
-p 应用开发者的通告证。在系统目录中缺省使用节点身份运行。  
-n 节点地址。在系统目录中缺省使用当前目录下的节点地址。  

**3、删除应用版本**  
<a id="delvers"></a>  
移除应用的某个或所有版本  
```bash  
Leither lapp delvers -a appName -v vers  -p pptfile -n address`
```  
-a 应用名。同一个开发者，应用不能重名。  
-v 要卸载的版本，如果为空删除整个应用  
-p 应用开发者的通告证。在系统目录中缺省使用节点身份运行。  
-n 节点地址。在系统目录中缺省使用当前目录下的节点地址。  

**4、显示应用**  
<a id="showapp"></a>  
显示应用信息  
```bash  
Leither lapp showapp -a appName -v version -p pptfile -n address`
```  

-a 应用名。同一个开发者，应用不能重名。 如果为空，显示所有应用    
-p 应用开发者的通告证。在系统目录中缺省使用节点身份运行。  
-n 节点地址。在系统目录中缺省使用当前目录下的节点地址。  
-v 要显示的应用版本信息  

**5、域名设置**   
<a id="setdomain"></a>  
显示应用信息  
```bash  
Leither lapp setdomain -a appName -d domain -m message -p pptfile -n address`  
Leither lapp setdomain -d homepi.leither.cn -n http://192.168.1.7:8000/ -a dav -p newuserforlogin.ppt -m gwaddr=leither.cn
```  
-a 应用名。同一个开发者，应用不能重名。 如果为空，显示所有应用    
-p 应用开发者的通告证。在系统目录中缺省使用节点身份运行。  
-n 节点地址。在系统目录中缺省使用当前目录下的节点地址。  
-d 要设置的域名  
-m 用逗号分开的键值的序列，包含网关地址版本信息等  
  
**6、版本备份**  
把当前版本生成一个版本号进行备份，同时把last版本指向这个版本
<a id="backup"></a>  
```bash  
Leither lapp backup -a 应用名 -v 应用版本 -p 通行证 -n 节点地址  
```  
-a 应用名。同一个开发者，应用不能重名。  
-v 要发布的应用版本。可以是特殊版本：cur,last。  
-p 应用开发者的通告证。在系统目录中缺省使用节点身份运行。
-n 节点地址。在系统目录中缺省使用当前目录下的节点地址。
  
输出结果如下:
```bash  
Backup
appid= htpoEXiE6IlCAqVbCvjvkY_XNfu
version= last
Current Release :[9]
ShowApp [dav] begin
Author  : h7Nkwe4rfb1d15En8yQ65Lhhq9b
AppName : dav
AppID    htpoEXiE6IlCAqVbCvjvkY_XNfu
Last    : 9
Release : 9
url     : http://192.168.1.7:8000/entry?author=h7Nkwe4rfb1d15En8yQ65Lhhq9b&app=dav&ver=last       
Version:[1,2,3,4,5,6,7,8,9,cur,last,release,]
```
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
url     : http://192.168.1.7:8000/entry?author=h7Nkwe4rfb1d15En8yQ65Lhhq9b&app=dav&ver=release       
Version:[1,2,3,4,5,6,7,8,9,cur,last,release,]
```
  
**8、版本恢复**   
把应用的指定版本恢复到当前版本
<a id="restore"></a>、  
```bash  
Leither lapp restore -a 应用名 -v 要恢复的版本 -p 通行证 -n 节点地址  
```  
-a 应用名。同一个开发者，应用不能重名。  
-v 要恢复的版本。可以是特殊版本：cur,last。  
-p 应用开发者的通行证。在系统目录中缺省使用节点身份运行。
-n 节点地址。在系统目录中缺省使用当前目录下的节点地址。
  
输出结果如下:
```bash  
Restore node=[http://192.168.1.7:8000/] pptfile=[newuserforlogin.ppt] appname=[dav]
appid= htpoEXiE6IlCAqVbCvjvkY_XNfu
Restore version[last] to cur ok
ShowApp [dav] begin
Author  : h7Nkwe4rfb1d15En8yQ65Lhhq9b
AppName : dav
AppID    htpoEXiE6IlCAqVbCvjvkY_XNfu
Last    : 2
Release :
url     : http://192.168.1.7:8000/entry?author=h7Nkwe4rfb1d15En8yQ65Lhhq9b&app=dav&ver=cur
Version:[1,2,cur,last,]
```
