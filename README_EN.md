Leither
========
Leither OS is an ultra-lightweight, decentralized cloud operating system designed to transform traditional internet services by decentralizing them.

Users can either create or join a decentralized network, where they can access or publish application services and information. In Leither network, the processes of generating, storing, publishing, displaying, and disseminating data (both data and applications) are fully decentralized. The rules governing each step are transparent, allowing users to establish their own, with no single component being irreplaceable. This network is characterized by being user-built, user-controlled, and user-owned.

Inspired by the concepts of Gene and Meme, a Mimei represents the smallest unit of information in Leither OS, capable of storing both data and application code. Mimei support file systems and databases, enabling information to flow between different nodes in the form of mimei. Mimei can suppport complex business logics through associations similar to hyperlinks.

Using Mimei and its application system, developers can reconstruct internet applications in a decentralized manner. The Mimei system can support internet applications with minimal resources, such as websites, gateways, mini-programs, and apps. It is only 17 megabytes in size and can run on various devices such as PCs, servers, phones, NAS, routers, and Raspberry Pi. A single Leither node can work as a lightweight server in home office at nearly zero cost. Two nodes can collaborate for data backup, fault tolerance, load balance, domain name resolution, and routing. Multiple nodes can form decentralized network, implementing common internet applications in a decentralized manner .

The Leither system is compatible with most operating systems including Linux, MacOS, Windows and Android. It incorporates authentication systems, application systems, file and database systems, as well as domain name and network services. Developers can create applications with Leither API that support over 40 programming languages. Developer can also interact with the system via command line or browser.

### **Installation and Operation**
Application changes and programs are available at the following URL:
<a href="http://vzhan.cn/mm/Fc1BRTFafOGzq5P8KmkVJqwS2v2/" target="_blank"> Download applications of various versions</a>

Download the program locally and then extract it. For details on each file, refer to <a href="./doc/Directory-en.md"> System Directory Structure</a>.

Rename the specified version of the program to Leither.

Set the execution permissions for Leither:
```shell
chmod +x ./Leither
```

Initialization is required for the first run. You can specify the port with -p, defaulting to 4800 if not specified. If using port 80, root permissions are required, and commands must be prefixed with sudo. You can specify the network entry with -b, defaulting to: mimei.org leither.cn vzhan.cn. The generated information is saved in Systemvar.json and can be manually modified.
```shell
Leither init -p 4800 -b mimei.org
```

Start as a background service:
```shell
./Leither run -d
```

Stop the service:
```shell
./Leither stop
```

### **Feature Experience**
You can experience features through various methods such as command line, browser, and API development. Below, the system features are directly demonstrated via command line. You can also run corresponding scripts in the TestCase directory through the API to experience specific APIs.

### **Network**
After a node starts, it enters the network through a bootstrap node. Display the ID of this node:
```shell
./Leither swarm id
```
Returns:
```shell
9I6JPEqsxWHN7dr2WG9C0CJ-VnN
12D3KooWBiuFhtpQL2fs3CasdDZ2yZsHHGWdbGTEuVM3BHaj4Aj
```
There are two IDs, short and long. All resources in Leither have a 27-length ID, generated similarly to Bitcoin wallet addresses. The long ID is compatible with the IPFS network, generated in the same way as IPFS network rules (multihash).

Display the local machine address:
```shell
./Leither swarm local
```

Returns:
```shell
/ip4/192.168.3.7/tcp/4800
/ip4/60.186.9.237/tcp/4800
/ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/4800
```
Leither network addresses are in multiaddr format.

Display nearby network nodes:
```shell
./Leither swarm addrs
```
Returns:
```shell
Ngacq50-IRX_DfcndFwZ0c9S0Nh (3)
        /ip4/192.168.3.7/tcp/4800
        /ip4/60.186.9.237/tcp/4800
        /ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/4800
l86HuY4FuRDezLEPHOHBjnaQczp (2)
        /ip4/172.31.47.58/tcp/80
        /ip4/18.222.243.60/tcp/80
...
```

Find a node:
```shell
./Leither dht findpeer tNP93yuZhNXd-om4izWQkYHfS50
```
Returns:
```shell
/ip4/99.79.46.219/tcp/80
/ip6/2600:1f11:ec1:3001:4d57:55c2:aec6:e279/tcp/80
/ip4/10.0.17.253/tcp/80
```

### **IPFS Files**
IPFS is a well-known decentralized file storage system. The Leither network supports the IPFS protocol and is compatible with the IPFS network.

**Add an IPFS file to the network:**
```shell
./Leither ipfs add Leither.txt
```
Output:
```shell
ipfs add ok  /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

You can view the file on any node in the network via URL:
```shell
curl 127.0.0.1:4800/ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

### **Mimei**
A mimei is an information container that can store files or databases. Mimeis facilitate data storage and presentation and can flow between network nodes. They can describe complex content through resource references and specify associated applications for data operations.

Traditional internet services tightly couple applications and data. The main purpose of mimeis is to deconstruct and reconstruct these contents, enabling most traditional internet functions.

**Create a mimei:**
```shell
./Leither mimei create
```

Returns:
```shell
Create MiMei  ok 
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
```
A mimei is created with ID RXN74QNeiY08LRSaoeQhx3nOLTC, default type is file.

After creation, data can be filled into the mimei.

**Fill an IPFS file into a mimei:**
Mimeis support IPFS files and file systems.

Place an IPFS file into a mimei:
```shell
./Leither mimei setcid RXN74QNeiY08LRSaoeQhx3nOLTC QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

Returns:
```shell
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
cid= QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiSetCid ver= 1
```

The mimei's latest version is now 1, pointing to the above IPFS file.

**Directly fill a file into a mimei:**
```shell
./Leither mimei add RXN74QNeiY08LRSaoeQhx3nOLTC Leither.txt
```

Returns:
```shell
add /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2 50773
MiMeiAdd cid= /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiAdd ver= 2
```
The mimei's latest version is now 2, pointing to the above IPFS file.

**Publish a mimei:**
Publish mimei information to the network:
```shell
./Leither mimei publish RXN74QNeiY08LRSaoeQhx3nOLTC
```

Returns:
```shell
MiMeiPublish mids= [RXN74QNeiY08LRSaoeQhx3nOLTC] EOL = 168h
MiMeiPublish ok
```
The mimei's information is now published on the network, and all supporting nodes will update the corresponding content in real-time.

**View mimei information:**
Query local and network mimei information:
```shell
./Leither mimei show RXN74QNeiY08LRSaoeQhx3nOLTC
```

Returns:
```shell
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

**View mimei content:**
Query mimei content:
```shell
./Leither mimei get RXN74QNeiY08LRSaoeQhx3nOLTC
```
Returns:
```shell
.......
File content omitted
.......
```

**Mimei synchronization:**
Synchronize mimeis from the network or specified nodes. Enter the following command on another node:
```shell
./Leither mimei sync RXN74QNeiY08LRSaoeQhx3nOLTC
```
Returns:
```shell
MiMeiSync mid = RXN74QNeiY08LRSaoeQhx3nOLTC
mimei sync ok
```
Mimei information and data are now synchronized to the local node.

**Mimei support:**
Nodes providing mimei data are called support nodes or data providers. The provide command broadcasts to the network that this node provides all data for this mimei.

```shell
./Leither mimei provide RXN74QNeiY08LRSaoeQhx3nOLTC
```
Returns:
```shell
MiMeiProvide cids= [RXN74QNeiY08LRSaoeQhx3nOLTC]
MiMeiProvide ok
```

View support nodes:
```shell
./Leither mimei findprovs RXN74QNeiY08LRSaoeQhx3nOLTC
```
Returns:
```shell
MiMeiFindProvide mid= RXN74QNeiY08LRSaoeQhx3nOLTC
mimei findprovs  ver=2
mimei findprovs  addrs=[{Ngacq50-IRX_DfcndFwZ0c9S0Nh [/ip6/240e:390:86b:f630:2900
:1e4:50f8:cd6a/tcp/4800 /ip4/192.168.3.7/tcp/4800 /ip4/60.186.9.237/tcp/4800]}]
mimei findprovs ok
```

When users publish new mimei data, all online support nodes will synchronize the data in real-time.

**Mimei usage:**
In addition to reading mimei data directly via API and commands, mimeis can specify applications for display. For more details, see the sections on "Applications" and "Domain Name Display of Mimeis."

### **Applications**
Users can develop applications based on APIs. The API protocol is based on the hprose protocol, supporting most common development languages. Users can directly use system functions in their applications. The system has special optimizations for HTML5 applications.

**Application backup:**
```shell
./Leither lapp backup -a dav -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```

**Upload to node:**
```shell
./Leither lapp uploadapp -i dist/dav -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```

**View application information:**
```shell
./Leither lapp showapp -a dav -v cur -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```

**Publish application to the network:**
```shell
./Leither mimei publish htpoEXiE6IlCAqVbCvjvkY_XNfu -p newuserforlogin.ppt
```
After publishing the application to the network, it can be accessed and run on any node via a browser.

Applications are stored in the dist/dav directory, with the application name being the directory name dav. The developer's issued pass is newuserforlogin.ppt. The node address is http://127.0.0.1:4800/. The above commands can be scripted and written in the configuration file of the development tool (e.g., package.json).

**Synchronize application to node:**
Synchronize applications from the network or specified nodes. Enter the following command on another node:
```shell
./Leither mimei sync htpoEXiE6IlCAqVbCvjvkY_XNfu
```
Returns:
```shell
MiMeiSync mid = htpoEXiE6IlCAqVbCvjvkY_XNfu
mimei sync ok
```
Application information and data are now synchronized to the local node.

**Specify default application for mimei:**
When creating a mimei, you can specify a default application:
```shell
./Leither mimei create -a application_id
```
When opening a mimei, the specified application will be used by default. You can also specify an application when opening a mimei.

### **Domain Name Display of Mimeis**
Through domain name nodes, users can display their home nodes to others via domain names. The experience is similar to a host in an IDC data center.

The following link opens a mimei with the default application, mimei ID is dJM6X7OTmJXbGqPQaFdAZ3kGpBl:
http://www.vzhan.cn/tpt/dJM6X7OTmJXbGqPQaFdAZ3kGpBl/

The following link directly runs an application, application ID is htpoEXiE6IlCAqVbCvjvkY_XNfu:
http://www.leither.cn/tpt/,htpoEXiE6IlCAqVbCvjvkY_XNfu/

Note: Ensure that the home node is accessible from the external network. All operators support IPv6, and telecom operators support dynamic IPv4 addresses. Both need to set up local routers or optical modems to ensure external access. If there is no accessible address locally, the tunnel function can be used to provide services externally through others' networks.

### **Mimei Database**
When creating a mimei, you can specify a database type. The database is compatible with most Redis functions. HTML5 database functions are correspondingly implemented. You can operate this database using command line and API. Using the above backup and synchronization, you can publish the database to the network. Support nodes will synchronize database changes in real-time.

### **Load Balancing and Fault Tolerance**
**Fault Tolerance:**
Node addresses in a distributed network are complex, with some being IPv4, some IPv6, and some through NAT. When users access a domain name, a template page is returned, detecting the optimal network path to access nodes.

**Load Balancing:**
Through the provide mechanism, multiple nodes can support user data and applications. When users access applications via domain names or links, domain and routing nodes will fill all support node information in the template page. Through the browser, the optimal access node is selected locally, achieving load balancing. Users can set or modify the template page to specify different load balancing strategies.

### **More Functional Usage**
You can interact with the system through the following methods: system application interface, command line, system API, and browser. For convenience, the document introduces and demonstrates the functions of each module via command line. Developers can use these functions in their applications through the API.

For details, refer to <a href="./api/Api.md"> API Documentation</a>.

### **Security**
When a node starts, a pair of keys is automatically created, and a user ID is generated based on the keys, which is also the user's identity. The lpki command set includes key management functions and some identity authentication-related operations.

### **Network**
When the service starts, it enters the network through the bootstrap node in the configuration. The swarm command set includes node address connection filtering and other functions.

The entire network is a DHT network. The dht command set includes network read and write functions and node search network routing functions.

### **Mimei**
All applications and data in the system are mimei types. Mimeis support file and database types. Mimeis can be published on the network, and other users in the network can directly access mimeis through mimei IDs. Data can also be synchronized to local nodes. Users can also provide data support for others' mimeis. Each change in a mimei generates a version, and these changes are updated in real-time on all support nodes (data providers).

Combined with domain names and load balancing mechanisms, fault tolerance and load balancing functions can be achieved.

Mimei command set related operations.

The system is compatible with the IPFS network. The IPFS command set includes common IPFS commands, and the files command set includes how to operate the local IPFS file system.

### **Application System**
When creating a mimei, an application type is specified. When users open a mimei, the application is called to display the mimei. The application itself is also a mimei. The lapp command set manages application backup and release, including version and domain management.

### **Domain Name and Load Balancing**
The system can run on home nodes. The advantage of home networks is speed and low cost. The disadvantage is the lack of stable IP and port 80, and stability is not as good as IDC data centers. The system solves this problem through domain name nodes. Users can bind resources (nodes, applications, content) to a domain name and point the domain name resolution to a domain name node. Third-party users access the domain name node through a browser, and the domain name node returns a very small (3k) data packet. The browser parses the data and node information in the packet, selects the best-performing node to provide services to third-party users. The process of the browser selecting nodes is a fault tolerance and load balancing process, completed entirely in the browser without additional traffic from the domain name node.

For the viewer, there is no page jump, equivalent to directly accessing the user's node. The experience is comparable to IDC data center services.

### **Introduction to Related Technical Principles**
<a href="./doc/MiMei.md"> Mimei and Application System</a>
Introduces how to dismantle internet platform businesses through decentralized applications.
<a href="./doc/GongShi.md"> Organization and Consensus Mechanism</a>
How to coordinate the interests of all parties through consensus mechanisms to encourage more participation.

### **Related Documents**
Links to project-related information PPT:
<https://zhuanlan.zhihu.com/p/330261216>

<a href="./doc/Pki.md"> Security System</a>

<a href="./doc/Applition.md"> Application System</a>

<a href="./doc/Setup.md"> Node and Application Installation</a>

<a href="./opt/dav/"> Example Application Code</a>

<a href="./api/Api.md"> API Documentation</a>

<a href="./doc/MiMei.md"> Mimei and Application System</a>

<a href="./doc/GongShi.md"> Organization and Consensus Mechanism</a>

<a href="./doc/PaaS.md"> Leither vs Traditional Cloud</a>   
The system can be used for cloud services, offering a significant reduction in costs related to CPU, memory, and bandwidth compared to virtual hosting and Docker.
