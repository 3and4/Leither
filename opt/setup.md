## 设置节点的基础运行环境
1. **获取Leither执行程序**  
Leither可执行程序大约6MB，可以从Leither/bin下载，根据所用硬件选择正确版本。Leither即是云操作系统，也是功能强大的命令行工具。下述内容默认的系统构成是Leither云服务节点一台。开发机一台，执行命令行前也需要启动Leither，并在后天运行。本文描述开发机的设置和使用。
  
2. **设置Leither参数**  
在SystemVars.json文件跟Leither可执行文件放同目录下。设置运行参数，一般只需要设置ServicePort, Gateway, FixedAddr  
{ "ServicePort":8000,  
  "LogAdapter":"file",  
  "LogConfig":"{\"filename\":\"Leither.log\",\"level\":1}",  
  "Gateway":["18.222.243.60","121.40.244.135","112.126.60.40"],&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#Leither域名节点  
  "FixedAddr":["121.40.213.15"],&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#有固定公网IP的情况，没有公网地址不需要设置，会通过其他线上的节点自动识别本机节点  
  "NOStart":["checkmac"]}  
  
3. **运行Leither**  
./Leither &&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#多个Leither可同时运行  
Leither运行成功后，会在本地创建多个目录，包括service，WebDav等。  
 
4. **生成测试用户，测试程序**  
继续用Leither命令行运行脚本，生成一个测试用户  
./Leither lssl runscript -s "local auth=require('auth'); return auth.Register('lsb', '123456');"  
授权新用户访问mimei。在service目录下建立RequestService，把mimei.lua放进去。允许在节点上操作弥媒。    
./Leither lssl runscript -s "local node=require('mimei'); return node.MMSetRight(request.sid, 'mmroot', '', 0x07276707);"  
  
生成测试用户并授权后，可以在./WebDav目录下放一些图片视频，下面的测试应用会显示这些文件列表。  
在Leither同目录下创建一个应用程序目录，目录名即默认为测试应用名，本示例测试应用名为lapp。把应用代码拷入./lapp  

5. **生成通行证**  
  用以准备发布代码数据到服务节点上
  a. 生成key  
  ./Leither lssl genkey -o my.key  
  b. 生成ca  
  ./Leither lssl genca -k my.key -m "name=my" -o my.ca&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;#-m message 可以省略  
  c. 生成自签名cert  
  ./Leither lssl gencert -k my.key -c my.ca -m "name=forapp" -o my.cert  
  d. 生成登录用passport(ppt)  
  ./Leither lssl signppt -c my.cert -m "CertFor=Self" -o mylogin.ppt  
  
6. **发布到Leither服务节点**  
  a. 在目标节点上申请弥媒权限  
  ./Leither lssl reqservice -c my.cert -m RequestService=mimei -n http://192.168.3.29:4800/  
  b. 上传代码  
  ./Leither deploy uploadapp -p mylogin.ppt -i ./lapp -n http://192.168.3.29:4800
  c. 发布
  ./Leither.exe deploy backup -a lapp -p mylogin.ppt -n http://192.168.3.29:4800/

7. **绑定Leither外网域名**  
  a. 内网URL绑定Leither公网URL。需要在路由器设置NAT端口转发  
  ./Leither.exe deploy setdomain -d fangpi.leither.cn -n http://192.168.3.29:4800/ -a lapp -p mylogin.ppt -m gwaddr=leither.cn  
  fangpi是开发者创建的Leither二级域名，lapp是测试应用名称
  b. 固化：在多个版本的应用中选定当前版本，默认为最新版  
  ./Leither.exe deploy backup -a lapp -p mylogin.ppt -n http://192.168.3.29:4800/  

至此，应用发布完成，可以在外网访问 http://fangpi.leither.cn, 其结果与内网访问 http://192.168.3.29:4800 相同。  

