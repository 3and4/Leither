## 设置节点的基础运行环境
1. **获取Leither执行程序**  
Leither可执行程序大约6MB，可以从Leither/bin下载，根据所用硬件选择正确版本
  
2. **设置Leither参数**  
在SystemVars.json文件跟Leither可执行文件放同目录下。设置运行参数，一般只需要设置ServicePort, Gateway, FixedAddr  
{ "ServicePort":9000,  
  "LogAdapter":"file",  
  "LogConfig":"{\"filename\":\"Leither.log\",\"level\":1}",  
  "Gateway":["18.222.243.60","121.40.244.135","112.126.60.40"],   #Leither域名节点  
  "FixedAddr":["121.40.213.15"],        #有固定公网IP的情况  
  "NOStart":["checkmac"]}  
  
3. **运行Leither**  
> ./Leither &  
初次安装Leither启动成功后，可以查看一下 http://localhost:service_port/ 会看到无法找到页面的提示。
Leither运行成功后，继续用Leither命令行执行一系列操作，比如下一步的生成信用证。  

生成测试用户，可以在命令行运行脚本  
./Leither lssl runscript -s "local auth=require('auth'); return auth.Register('lsb', '123456');"  
授权新用户访问mimei  
./Leither lssl runscript -s "local node=require('mimei'); return node.MMSetRight(request.sid, 'mmroot', '', 0x07276707);"  
  
4. **生成信用证**  
  a. 生成key  
  ./Leither lssl genkey -o my.key  
  b. 生成ca  
  ./Leither lssl genca -k my.key -m "name=my" -o my.ca   # -m message 可以省略  
  c. 生成自签名cert  
  ./Leither lssl gencert -k my.key -c my.ca -m "name=forapp" -o my.cert  
  d. 生成登录用passport(ppt)  
  ./Leither lssl signppt -c my.cert -m "CertFor=Self" -o mylogin.ppt  
  
5. **发布**  
  a. 在目标节点上申请弥媒权限  
  ./Leither lssl reqservice -c my.cert -m RequestService=mimei -n http://192.168.3.29:4800/  
  b. 发布代码
  ./Leither deploy uploadapp -p mylogin.ppt ./lapp -n http://192.168.3.29:4800

service目录下建立RequestService，把mimei.lua放进去
