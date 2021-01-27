## lapp命令集  

**1、上传应用**  
指定一个版本为发布版（release）
<a id="upload"></a>  
```bash  
Leither lapp upload -i 应用目录 -p 通行证 -n 节点地址 -s    
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

