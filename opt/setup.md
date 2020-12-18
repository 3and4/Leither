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
