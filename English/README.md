Leither
========
Leither is a highly efficient and decentralized cloud operating system, designed to deconstruct and rebuild conventional internet services in a decentralized way.  

Nowadays, the internet has become essential to our daily lives, but it is dominated by a small number of powerful corporations. These companies monopolize control over data and rules, capturing a significant portion of the profits across various industries. This has resulted in a loss of innovation and motivation within the internet, with users unknowingly becoming exploited and manipulated.  
In light of this situation, it is crucial to create a decentralized internet that is built, controlled, and enjoyed by its users.  

Leither allows users to create or join decentralized networks where they can access or provide application services and information.  
Throughout the network, information (including applications) is generated, stored, published, displayed, and shared in a fully decentralized manner.  
All of these operations are governed by transparent rules that user can determine, with no single point being indispensable.  

By utilizing MiMei and the application system, developers can break down and reconstruct internet platform services in a decentralized fashion.  
Organizers can employ organizational and consensus mechanisms to balance the interests of all parties involved, encouraging broader participation.  
Together, these approaches can help reclaim the internet from the grip of oligopolies and return it to the users.        

The system is capable of running internet applications with minimal resource requirements, including but not limited to websites, public accounts, mini-programs, and apps.  
With a size of just 12 MB, it can operate on a variety of devices, such as PCs, servers, smartphones, NAS devices, routers, and Raspberry Pis.  
A single node can serve as a lightweight host, delivering data center-grade performance on home networks.  
Two nodes can work together, providing services like data backup, fault tolerance, load balancing, domain name resolution, and routing.  
Multiple nodes can form a decentralized distributed network, allowing the implementation of standard internet applications in a decentralized way.  

Leither can run on popular base operating systems, including Windows, Linux, FreeBSD, Darwin, and Android.  
Its functionalities include authentication, application system, file system and database, and domain and network system.  
Developers can create applications using APIs that support over forty common programming languages and can also interact with the system via the command line or a web browser. 

### **Install**  
Versions of application program at the following address:  
<a href="./bin/README.md">Download</a>  

Download and extract it.  
For details of each file, refer to <a href="./doc/Directory.md"> System Directory Structure</a>  

Setup permissions for Leither app  
```bash
chmod +x ./Leither
```

Run  
```bash
./Leither&
```

### **Feature Experience**  
You can experience the features in multiple ways, such as through the command line, browser, and API application development.  
Below, the system features are demonstrated directly through the command line.  
You can also experience specific APIs by running the corresponding script programs in the TestCase directory.  

### **Network**
After the node starts, it enters the network through the bootstrap node.  
Display local node id  
```bash
./Leither lpki id
```
Return value:
```bash
Ngacq50-IRX_DfcndFwZ0c9S0Nh
```

Display local node address
```bash
./Leither swarm local
```

Return value:  
```bash
/ip4/192.168.3.7/tcp/8000
/ip4/60.186.9.237/tcp/8000
/ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/8000
```

Display neighbour nodes
```bash
./Leither swarm addrs
```
Return value:
```bash
Ngacq50-IRX_DfcndFwZ0c9S0Nh (3)
        /ip4/192.168.3.7/tcp/8000
        /ip4/60.186.9.237/tcp/8000
        /ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/8000
l86HuY4FuRDezLEPHOHBjnaQczp (2)
        /ip4/172.31.47.58/tcp/80
        /ip4/18.222.243.60/tcp/80
..
...  
....
```

Search a node
```bash
./Leither dht findpeer tNP93yuZhNXd-om4izWQkYHfS50
```
Return value:

```bash
/ip4/99.79.46.219/tcp/80
/ip6/2600:1f11:ec1:3001:4d57:55c2:aec6:e279/tcp/80
/ip4/10.0.17.253/tcp/80
```

### **IPFS**
IPFS is a well-known decentralized file storage system.   
Leither supports IPFS protocol and is compatible with IPFS network.

**Add IPFS file to network**
```bash
./Leither ipfs add Leither.txt
```
Result:
```
IpfsAdd  /home/pi/sdb/Leither/Leither.txt
wa.sid= f6503de3a81ee92c2e0c8f04c1c1be71ce6af912
add /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2 50773
ipfs add ok  /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

Within Leither network, file can be checked on any node by URL:
```
curl 127.0.0.1:8000/ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
......
file content...
......
```
### **MiMei**
MiMei is an information container that can store files or databases.    
MiMei enables data storage, retrieval, and presentation.  
It can flow between various network nodes, complex content can be described through resource references,
associated applications can be associated to operate data.      

In conventional internet, applications and data are highly coupled.    
The primary purpose of MiMei is to deconstruct and reconstruct these contents.  
Using MiMei, most traditional internet functions can be implemented.      

**Create MiMei object**   
```bash
./Leither mimei create
```

Result:  
```
Create MiMei  ok 
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
```
A MiMei object is created.  
id=RXN74QNeiY08LRSaoeQhx3nOLTC  
Its default type is File.  

After its creation, data can be filled into the MiMei.  

**Fill a IPFS file into MiMei container**  
MiMei supports IPFS file and file system.

Given a MiMei, fill a IPFS file in it.  
```bash
./Leither mimei setcid RXN74QNeiY08LRSaoeQhx3nOLTC QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

Result:
```bash
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
cid= QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiSetCid ver= 1
```

In the code above, the *last* version of MiMei is 1.  
Version 1 points to the IPFS file.


**Fill MiMei wih a regular file**  
```bash
./Leither mimei add RXN74QNeiY08LRSaoeQhx3nOLTC Leither.txt
```

Result:
```bash
add /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2 50773
MiMeiAdd cid= /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiAdd ver= 2
```
The *last* version is 2  
Version 2 points to the IPFS file.


**Publish MiMei**  
Publish MiMei information on network
```bash
./Leither mimei publish RXN74QNeiY08LRSaoeQhx3nOLTC
```

Result:   
```bash
MiMeiPublish mids= [RXN74QNeiY08LRSaoeQhx3nOLTC] EOL = 168h
MiMeiPublish ok
```
MiMei is published on network. All nodes that support this MiMei will update in real time.  
  
  
**Query MiMei**  
Search MiMei information on local node and network  
```bash
./Leither mimei show RXN74QNeiY08LRSaoeQhx3nOLTC
```

Result:

```bash
MiMeiShow mid= RXN74QNeiY08LRSaoeQhx3nOLTC
Author  : Ngacq50-IRX_DfcndFwZ0c9S0Nh
AppType :
Ext     :
Mark    : 16650373773415379251557603663
Create  : 2022-10-06T14:22:57+08:00
Right   :7276704

from local
-------------------------------------------------------------       
2        QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
last     2

from net
-------------------------------------------------------------       
2        QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
last     2

conflict
-------------------------------------------------------------       
no conflict

compare version
-------------------------------------------------------------       
local version is same with net

mimei show ok
```

**Review MiMei content**    
```bash
./Leither mimei get RXN74QNeiY08LRSaoeQhx3nOLTC
```
Result:  
```bash
.......
content
.......
```


**Synchronize MiMei**  
After MiMei published, on another node run the following command    
```bash
./Leither mimei sync RXN74QNeiY08LRSaoeQhx3nOLTC
```  
Result:  
```bash
MiMeiSync mid = RXN74QNeiY08LRSaoeQhx3nOLTC
mimei sync ok
```  
MiMei data and information will be synchronized at current node.  

**Support MiMei**  
Node that provide MiMei data is called Provider. 
*provide* is to broadcast a message on the network: The MiMei data of an ID is availabe here.  
```bash
./Leither mimei provide RXN74QNeiY08LRSaoeQhx3nOLTC
```  
result:
```bash
MiMeiProvide cids= [RXN74QNeiY08LRSaoeQhx3nOLTC]
MiMeiProvide ok
```  

Look for provider for a MiMei 
```bash
./Leither mimei findprovs RXN74QNeiY08LRSaoeQhx3nOLTC
```
Result:
```bash
MiMeiFindProvide mid= RXN74QNeiY08LRSaoeQhx3nOLTC
mimei findprovs  ver=2
mimei findprovs  addrs=[{Ngacq50-IRX_DfcndFwZ0c9S0Nh [/ip6/240e:390:86b:f630:2900
:1e4:50f8:cd6a/tcp/8000 /ip4/192.168.3.7/tcp/8000 /ip4/60.186.9.237/tcp/8000]}]
mimei findprovs ok
```  

When a user publishes new MiMei data, all online providers will update their data in realtime.

**Use MiMei**  
Besides reading MiMei data with API or command, MiMei can be rendered by associated App.  

### **App**
Users can develop applications based on APIs. The API protocol is built upon the Hprose protocol and supports most common programming languages.  
Users can directly use system features within their applications.  
For HTML5 applications, the system has been specially optimized.  

**Backup App**  
```bash
./Leither lapp backup -a dav -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```  
  
**Upload App to a node**
```bash
./Leither lapp uploadapp -i dist/dav -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```  

**Lookup App information**  
```bash
./Leither lapp showapp -a dav -v cur -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```  

**Publish App on network**  
```bash
./Leither mimei publish htpoEXiE6IlCAqVbCvjvkY_XNfu -p newuserforlogin.ppt
```  
After deploying an App to the network, you can run the App on any node through a browser.  

The App is stored in the dist/dav directory, and the App name is the directory name "dav"  
The passport signed by the developer is in newuserforlogin.ppt  
The node address is http://127.0.0.1:4800/  
The above commands can saved as script and written in the configuration file of the development tool, such as package.json  

**Assign default App for MiMei**  
When MiMei is created, a default App can be associated with it.
```bash
./Leither mimei create -a 应用id
```  
When opening MiMei, the default App will be used. A different App can also be specified.  

### **Domain MiMei**
Using service of domain name node, a household node can be accessed through domain name, like a commercial server in IDC.  

The following link opens a MiMei with default App. The MiMei Id is dJM6X7OTmJXbGqPQaFdAZ3kGpBl.  
http://www.leither.cn/tpt/dJM6X7OTmJXbGqPQaFdAZ3kGpBl/  

The following link runs an App. The App Id is htpoEXiE6IlCAqVbCvjvkY_XNfu  
http://www.leither.cn/tpt/,htpoEXiE6IlCAqVbCvjvkY_XNfu/  

Note: Make sure the home node is accessible from the internet.    
All carriers support IPv6, and telecom carriers provide dynamic IPv4 addresses.  
Both require setting up router or modem to ensure external access.  
If local address is not accessible, Tunnel service can be used to leverage someone else's network.    

### **MiMei Database**
MiMei can be created as Database, which is compatible with most Redis functions.  
All database functions in HTML5 have corresponding implementations.  
Both command line and APIs can be used to interact with this database.  
Using the same backup and synchronization features as File type, database can be published to the network. All providers will update in real time.  

### **Load balance and fault tolerance**  
**Fault tolerance**  
In a distributed network, node addresses can be complex.  
Some use IPv4, some use IPv6, and some are accessed through NAT.  
When accessesing a domain, a template page is returned. This template page detects the optimal network path to access the node.    

**Load balance**  
Using the *provide* mechanism, multiple nodes can support a user's data and App.  
When accessing an App through a domain or link, the domain and routing nodes fill the template page with information of all providers, and return it to user's browser.  
The browser then selects the optimal node to provide data, thus implementing load balance.  
Users can set or modify the template page to specify different load balancing strategies.  


### **More...**
Users can interact with the system in the following ways:  
System application interface, command line, system API, and browser.  

For convenience, the documentation introduces and demonstrates the functionality of each module through the command line. Developers can utilize these features within their applications via the API.  



For more information, go to <a href="./api/Api.md">API</a>  

### **Security**  
When a node starts, it automatically generates a pair of keys and creates a user ID based on the keys.  
This user ID serves as the user's identity.  
The lpki command set includes key management functions as well as some identity authentication-related operations.  

### **Network**  
When the service starts, it enters the network through the bootstrap nodes specified in the configuration.  
The *swarm* command set includes functions such as node address connection filtering.    

The entire network operates as a Distributed Hash Table (DHT) network. The DHT command set contains read and write functions for the network, as well as network routing functions for node discovery.    


### **MiMei**  
In the system, all applications and data are of the MiMei type.  
MiMei supports file and database types.    
MiMei can be published on the network, and directly accessed by its MiMei ID.  
Users can synchronize MiMei to their local nodes and provide data support for others' MiMei.   
Any change in MiMei generates a new version, and these changes are updated in real-time to all providers.  

By combining domain names and load balancing mechanisms, fault tolerance and load balancing are implemented.  


### MiMei command set
Leither system is compatible with IPFS network.  
*ipfs* command set includes most IPFS commands.  
*files* command set includes commands for operating local IPFS files.  

### **Application System**  
When MiMei is created, it is associated with an App. When user opens the MiMei, the App is invoked to render it.  
App is a MiMei itself.  
*lapp*  command set manages application backup and deployment, as well as version and domain name management.  

### **Domain name and load balance**    
Leither system can run on home-based nodes.  
The advantages of home networks are fast speeds and low costs. The drawbacks are the lack of stable IP addresses and port 80, as well as lower stability compared to IDC data centers.  

Leither system solves this problem through domain name nodes. Users can bind resources (nodes, applications, content) to a domain name and point the domain name resolution to a domain name node.   

Visitors access the domain name node through a browser, which returns a minimal (3k) data package. The browser then parses the data and node information, selecting the best-performing node to provide service to the site visitor.  

The browser's node filtering process is a fault-tolerant and load-balancing process that takes place entirely within the browser without using additional domain name node's traffic.  

For visitors, there is no page redirection, making it seem as if they are directly accessing the user's node. The overall experience is comparable to services provided by IDC data centers.  


### **Introduction to technical principles**  
<a href="./doc/MiMei.md"> MiMei and application system</a>  
Introducing how to dismantle internet platform businesses through decentralized applications.  
<a href="./doc/GongShi.md"> Organization and consensus</a>  
How to coordinate the interests of all participants through consensus mechanisms, encouraging more people to get involved.  

### **Documents**  
Links to related information  
<https://zhuanlan.zhihu.com/p/330261216>


  
<a href="./doc/Pki.md"> 安全体系</a>  

<a href="./doc/Applition.md"> 应用体系</a>  

<a href="./doc/Setup.md"> 节点和应用安装</a>  
  
<a href="./opt/dav/"> 示例应用代码</a>  
  
<a href="./api/Api.md"> Api文档</a>  

<a href="./doc/MiMei.md">  弥媒和应用体系</a>  

<a href="./doc/GongShi.md"> 组织和共识机制</a>  
    
<a href="./doc/PaaS.md"> Leither vs 传统的云</a>   
系统可用于云服务，相较于虚拟主机和docker，CPU内存带宽等方面成本有数量级的降低。

