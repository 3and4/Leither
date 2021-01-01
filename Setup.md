节点配置
========  
## 简介  
本文演示内容如下：  
不同节点的配置  
演示服务授权  
演示应用的开发和发布  
演示如何將家用节点和应用映射到指定的域名上  
演示命令行角本方式调用api
演示js通过rpc方式调用Api
演示用户认证鉴权API  
演示不同方式的文件操作  

## 基础知识  

**目录结构**  
  
根目录  
｜Leither-----------------系统程序   
｜SystemVars.json---------配置文件  
｜temp.html---------------域名模板  
｜--service---------------服务目录,存放服务端角本  
｜   ｜----RequestService--处理服务请求的角本  
｜--webdav----------------缺省的webdav目录     
  
<a href="./Directory.md"> 完整目录结构介绍</a>  
  
**获取Leither执行程序**  
Leither可执行程序大约6MB，可以从Leither/bin下载，根据所用硬件选择正确版本。Leither即是云操作系统，也是功能强大的命令行工具。下述内容默认的系统构成是Leither云服务节点一台，开发机一台，域名节点一个。开发机上执行命令行前也需要启动Leither，并在后台运行。本文描述开发机的设置和使用。  
<a href="./bin/"> 点击下载</a>  
  
**设置配置文件**  
文件格式为json，参考如下：  
```json
{
    "ServicePort":8000,
    "LogAdapter":"file",
    "LogConfig":"{"filename":"Leither.log","level":1}",
}
```

SystemVars.json文件跟Leither可执行文件放同一目录下。设置运行参数，一般只需要设置ServicePort, Gateway, FixedAddr, LAN  

1. ServicePort  
服务端口  
不设置的情况，系统缺省值为9000，对于家用节点，不可以设置80，8080端口。通常是被运营商封掉了  
配置示例  
    ```json
    {
        ......
        "ServicePort":80,
        ......
    }
    ```

2. FixedAddr  
域名或固定的ip  
用于有独立ip的云主机
不设置的情况，程序会通过网关或好友节点自动探查自己的外网地址  
配置示例  
    ```json
    {
        ......
        "FixedAddr":"121.44.244.138",
        ......
    }
    ```

3. LogConfig  
日志配置  
配置示例  
   + 文件日志:  
    ```json
    {
        ......
        "LogAdapter":"file",
        "LogConfig":"{\"filename\":\"Leither.log\",\"level\":1}",  
        ......
    }
    ```
   + 控制台日志:  

    ```json
    {
        ......
        "LogAdapter":"console",
        "LogConfig":"{\"level\":1}",  
        ......
    }
    ```

4. LAN  
多网卡的情况下，绑定服务的网卡  
配置示例   
   + 绑定Mac  
    ```json
    {
        ......
        "LAN":"98-5F-D3-46-17-A8",
        ......
    }
    ```  
    + 绑定IP  
    ```json  
    
    {
        ......
        "LAN":"192.168.1.24",
        ......
    }
    ```

5. Gateway  
对于有动态ip的节点，启动时会通过网关节点获取自己当前的外网ip  
节点定期（当前是15分钟）会向网关节点上报自己的地址  
可以通过网关节点，找到指定的节点或内容  
通过网关节点帮助，可以把节点、应用、内容绑定在设定的域名上

    ```json
    {
        ......
        "Gateway":["18.222.243.60","121.40.244.135","112.126.60.40"],  
        ......
    }
    ```
  
**运行Leither**  
```
./Leither &  
```

Leither运行成功后，会在本地创建多个目录，包括service，WebDav等。  

## 配置节点  
本文涉及三个节点：开发节点、应用节点、域名节点  
**开发节点**  
开发程序时使用的节点   
通常在开发机上，这个节点用于应用的开发和调试  
需要配置的文件  
Leither(.exe), SystemVar.json  
初次使用，只需要这两个文件便可运行,根目录会生成密钥文件 hostkey.cfg。这个文件代表节点身份。  
  
**域名节点**
域名节点是帮应用节点解析域名的节点  
有固定ip的主机，节点会把提前设定的域名请求转到家用节点，整个过程中除了开始的2.7k的引导包，其他流量是直接访问家用节点

需要的文件比开发节点多一个模板文件temp.html  

**应用节点**
这个节点通常布置在用户家里的设备上，使用家用网络,运行应用，保存用户数据。

## 用户和权限  
本文案例涉及多个用户：
1. 三个节点用户  
   对应管理三个节点，密钥存放在节点程序的根目录,文件名为hostkey.cfg。
2. 开发者用户  
   开发者用户开发维护并发布程序,密钥需要通过Leither lssl 命令单独生成，单独保存  
   
3. 数据所有者用户  
   应用数据的所有者  
   为了简单，在下面的示例中缺省使用了应用节点的用户存放测试数据。  
   这个用户便是下文中讲的测试用户"lsb"  
   为了演示简单，省去注册用户环节，这用户写在了代码里。

4. 数据查看者用户
   查看使用数据的用户  
   在下面的示例中，是通过角本执行register产生的，这种方式密钥保存在当前节点中，需要对当前节点有依赖关系。

5. 域名节点用户  
   域名节点的所有者，帮助其它节点解析域名，在本案例中是使用域名节点用户

**生成测试用户**  
测试用户指的是上文中的数据查看者用户。
在应用节点根目录下，用Leither命令行运行脚本，生成一个测试用户"lsb"  
```lua  
./Leither lssl runscript -s "local auth=require('auth'); return auth.Register('lsb', '123456');"  
```  
在这个目录下执行命令，缺省会使用当前节点信息，使用当前节点身份，连接当前节点进行操作
本例中，lssl是安全相关的命行集，Register是认证API中的用户注册

**授权测试用户访问权限**  
设置数据目录权限  
后面应用中操作的数据目录是webdav目录，对应的名称是“mmroot”,缺省是不对外开放的。 需要修改对外权限  
```bash  
./Leither lssl runscript -s "local node=require('mimei'); return node.MMSetRight(request.sid, 'mmroot', '', 0x07276707);"  
```  
生成测试用户并授权后，lsb用户便可以查看这个目录下的数据。  
可以在WebDav目录下放一些图片视频，下面的测试应用会显示这些文件列表。  


**生成开发者用户**  
生成密钥和通行证， 用以发布代码数据到服务节点上  
密钥相关信息应该单独保存，切勿和节点放在一起  
+  生成key  
包含公钥和私钥，公钥取摘要生成代表用户身份的唯一id  
    ```bash  
    Leither lssl genkey -o my.key  
    ```  
+ 生成ca  
    ```bash  
    Leither lssl genca -k my.key -m "name=my" -o my.ca
    ```  
    -m message 可以省略  
+ 生成自签名cert  
    ```bash  
    Leither lssl gencert -k my.key -c my.ca -m "name=forapp" -o my.cert  
    ```  
+ 生成登录用passport(ppt)  
    通行证中包含了一个时效信息,缺省72小时
    ```bash  
    Leither lssl signppt -c my.cert -m "CertFor=Self" -o mylogin.ppt  
    ```  
  
**开发者获取权限**  
开发者获取发布应用权限  
+ 节点设置授权角本  
节点处理服务权限的角本都放在“根目录/service/RequestService”,参考结构如下  
根目录  
｜  
｜--service---------------服务目录,存放服务端角本  
｜   ｜----RequestService--处理服务请求的角本  
  
    在Leither中，应用和数据都是弥媒对象，处理弥媒权限的角本是mimei.lua  
    下载角本<a href="./opt/mimei.lua"> 点击下载</a>  
    角本中代码如下:  
    ```lua  
    local log = require("log");
    log.Debug("request service mimei request=%v", request)
    return true
    ```  
    在service目录下建立目录RequestService，把mimei.lua放进去，节点重启后生效。  

+ 申请应用发布权限（弥媒权限）
    -c 后面是开发者的通行证  
    -n 后面是申请的节点地址，包括本地调试程序的开发节点或实际运行的应用节点  
    ```bash  
    ./Leither lssl reqservice -c my.cert -m RequestService=mimei -n http://192.168.3.29:4800/  
    ```

**应用发布**  
下载要发布的测试应用  <a href="./opt/dav.zip"> 点击下载</a>  
在开发的工作目录中展开压缩包，目录名就是测试应用名，当前的应用名就是“dav”,可以手工修改  

```bash  
Leither deploy uploadapp -p mylogin.ppt -i 应用路径 -n http://192.168.3.29:4800  
```
-i 后面是要发布的应用路径  
-n 后面是要发布的节点链接，可以是应用节点，也可以是开发节点
这时候，应用发布的是一个当前版本，可以通过显示信息中的链接检查应用是否正常。在确定应用无误之后，可以把当前版本入库。    
固化发布  
把版本号的cur内容入库，并更新版本号，第一个版本为1，第二个版本为2，last通常为最后版，也可以是指定版本。  
```bash  
./Leither.exe deploy backup -a dav -p mylogin.ppt -n http://192.168.3.29:4800/  
```

**绑定Leither外网域名**  
1. 确保节点能够被外网访问  
有些用户没有分配外网ip，这时候需要向运营商索取,或者向第三方借用ip(这个话题暂不展开)  
通常路由器都屏蔽了局域网内的设备，这时候可以通过配置端口转发使外网可以访问到节点

2. 域名设置  
```bash  
./Leither.exe deploy setdomain -d fangpi.leither.cn -n http://192.168.3.29:4800/ -a dav -p mylogin.ppt -m gwaddr=leither.cn  
```  
fangpi是开发者创建的Leither二级域名，dav是测试应用名称  
```bash  

也可设置自己的域名，需要把域名解析指向域名节点的ip便可

至此，应用发布完成，可以在外网访问 http://fangpi.leither.cn, 其结果与内网访问 http://192.168.3.29:4800 相同。  
如果从内网访问发布后的公网URL，同时又开启VPN的话，可能会出现网页无法打开。此时关闭vpn即可。

