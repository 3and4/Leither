index.html和styples.css为已生成好的首页内容
当前目标是生成好“安装页面”
文件名为install.html
页面框架风格参考首页
页面内容包括三个部分：角本安装、手工安装、初始化、管理服务
在index.html中点击安装进入本网页
```shells标注的内容是在控制台操作的命令，需要有相应的css风格


二、角本安装
支持bash的终端下可以直接执行以下角本安装在当前目录:
```shell
curl -fsSL http://vzhan.cn/install.sh | bash
```

三、手工安装
到以下网址下载程序到本地
<a href="./dl.html"> 各版本的应用</a>  

把指定版本的程序改为Leither  

设置好Leither的运行权限  
```shell
chmod +x ./Leither
```

了解各文件细节参考<a href="./doc/Directory.md"> 系统目录结构</a>  

三、初始化
第一次运行需要初始化  
可以通过-p指定端口，不指定缺省为4800.如果是80端口，需要root权限，命令都要加sudo 
可以通过-b指定网络入口，不指定为：mimei.org vzhan.cn  
生成的信息保存在Systemvar.json中，可手工修改。
```shell
Leither init -p 4800 -b mimei.org
```

四、管理服务
后台服务方式启动  
```shell
./Leither run -d
```

关闭服务
```shell
./Leither stop
```
