Leither
========
Leither OS is *an ultra-lightweight, decentralized cloud operating system* designed to revolutionize traditional internet services by decentralizing them.

Users can either create or join a decentralized network to publish application services and information. In the Leither network, the processes of generating, storing, publishing, displaying, and disseminating data and applications are fully decentralized. The governance rules for each step are transparent, ensuring no single component is irreplaceable, allowing users to establish their own rules. This network is characterized by being user-built, user-controlled, and user-owned.

Inspired by the concepts of Gene and Meme, a Mimei represents the smallest unit of information in the Leither network, capable of storing both data and application code. Mimei supports file systems and databases, enabling information to be rendered by associations among different nodes, similar to hyperlinks. Using Mimei and its application system, developers can reconstruct internet applications in a decentralized manner. Mimei supports internet applications with minimal resources, such as websites, gateways, mini-programs, and apps.

The Leither executable is only 17MB in size and can run on various devices, including PCs, servers, phones, NAS, routers, and Raspberry Pi. A single Leither node can function as a lightweight server in a home office at nearly zero cost. Two nodes can collaborate for data backup, fault tolerance, load balancing, domain name resolution, and routing. Multiple nodes can form a decentralized network.

Leither OS is compatible with most operating systems, including Linux, MacOS, Windows, and Android. It incorporates authentication systems, application systems, file and database systems, as well as domain name and network services. Developers can create applications using the Leither API, which supports over 40 programming languages. They can also interact with Leither via command line or browser.

### **Install**
Obtain the Leither executable from 
<a href="http://vzhan.cn/mm/Fc1BRTFafOGzq5P8KmkVJqwS2v2/" target="_blank">Leither executables</a>

Rename the downloaded file to ***Leither*** and place it in an empty directory. This directory will serve as the home directory for this Leither instance. Leither will automatically generate all necessary supporting files in this directory. For more information on these files, refer to <a href="./doc/Directory-en.md"> System Directory Structure</a>.

Before running Leither, ensure it has the correct execution permissions.
```shell
chmod +x ./Leither
```

When starting Leither, user can specify a port number with -p, defaulting to *4800*, and network entry with -b, defaulting to *mimei.org, leither.cn, vzhan.cn*. The settings are saved in SystemVars.json within Leither directory, which can be manually modified or by the following command.
```shell
./Leither init -p 4800 -b mimei.org
```
This command initializes a Leither instance on Port 4800 and register it with a bootstap node at mimei.org. Note that any Leither node can function as a bootstrap node.

Running as a background service:
```shell
./Leither run -d
```

Stopping the service:
```shell
./Leither stop
```
To ensure the Leither service starts automatically after a system reboot, configure it using system services like _systemctl_.
### **Features**
User can explore various features vie command line, browser, and development API. In the following sections, the system features are demonstrated via command line.

### **Network**
After a Leither node starts, it joints the Leither network by registering at a bootstrap node. The following command shows current node's ID:
```shell
./Leither swarm id
```
Returns:
```shell
9I6JPEqsxWHN7dr2WG9C0CJ-VnN
12D3KooWBiuFhtpQL2fs3CasdDZ2yZsHHGWdbGTEuVM3BHaj4Aj
```
There are two types of IDs: the shorter one is a 27-character Leither ID, generated with algorithm similar to Bitcoin wallet. The longer one is compatible with the <a href="https://github.com/ipfs/ipfs">IPFS</a> protocol and is created using the same multi-hash method.

***Local machine's IP addresses:***
```shell
./Leither swarm local
```

Returns:
```shell
/ip4/192.168.3.7/tcp/4800
/ip4/60.186.9.237/tcp/4800
/ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/4800
```
Leither network addresses use multi-addr format.

***Nearby Leither nodes:***
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

***Find a node:***
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
 Leither network supports <a href="https://github.com/ipfs/ipfs">IPFS</a> protocol and is compatible with <a href="https://github.com/ipfs/ipfs">IPFS</a> network.

***Add a file to IPFS network:***
```shell
./Leither ipfs add filename.txt
```
Output:
```shell
ipfs add ok /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

User can view the file on any node in the network via URL:
```shell
curl http://60.186.9.237:4800/ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
or open the url in a web browser.
```

### **Mimei**
Mimei is the core concept and innovation in Leither system. It is an information container that can store files or databases, and facilitate data presentation. Mimei can be synchronized among network nodes, describe complicated data structure through resource references, and specify associated applications for data operations.

Data are tightly coupled with applicaiton in regular internet. The main purpose of Mimei is to decouple and reconstruct these relationships.

***Create a mimei:***
```shell
./Leither mimei create
```

Returns:
```shell
Create MiMei  ok 
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
```
The command creates a Mimei ID RXN74QNeiY08LRSaoeQhx3nOLTC. The default Mimei type is file. Another type is Redis database.

After its creation, data can be stored in the Mimei.

***Store an IPFS file into a Mimei:***\
Mimei support IPFS files and file systems.

Set an IPFS file into a Mimei container:
```shell
./Leither mimei setcid RXN74QNeiY08LRSaoeQhx3nOLTC QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

Returns:
```shell
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
cid= QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiSetCid ver= 1
```

Now the Mimei's last version is 1, pointing to the above IPFS file.

***Store a file into a Mimei:***
```shell
./Leither mimei add RXN74QNeiY08LRSaoeQhx3nOLTC filename.txt
```

Returns:
```shell
add /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2 50773
MiMeiAdd cid= /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiAdd ver= 2
```
The mimei's last version is 2 now, pointing to the above IPFS file.

***Publish a Mimei:***\
Publish Mimei information to the network:
```shell
./Leither mimei publish RXN74QNeiY08LRSaoeQhx3nOLTC
```

Returns:
```shell
MiMeiPublish mids= [RXN74QNeiY08LRSaoeQhx3nOLTC] EOL = 168h
MiMeiPublish ok
```
The Mimei is now published on Leither network. All of its providers, which are nodes that support the Mimei, will update their copies in real-time.

***View Mimei information:***\
Query local and network Mimei information:
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
Right   : 7276704

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

***View Mimei content:***\
Query Mimei of File type:
```shell
./Leither mimei get RXN74QNeiY08LRSaoeQhx3nOLTC
```
Returns:
```shell
.......
File content shown here
.......
```
Query Mimei of Database type:
```shell
./Leither mimei grep Wg5Bgcs3B3j_kWtz1rJF3XvP24N
```
Returns:
```shell
mid: Wg5Bgcs3B3j_kWtz1rJF3XvP24N
ver: last
MMOpen ok mmsid= 89825eac2a87ab03d716c227a15fda354bcb084f
Data Type KV
data_of_author	map[avatar:QmdwTr6Rf9toV4uTYqtx2grMP7ofGaib35xWK9KYs5J6o5...
.......
```

***Mimei synchronization:***\
Synchronize Mimei from the network or specified nodes. Run the following command on a different Leither node:
```shell
./Leither mimei sync RXN74QNeiY08LRSaoeQhx3nOLTC
```
Returns:
```shell
MiMeiSync mid = RXN74QNeiY08LRSaoeQhx3nOLTC
mimei sync ok
```
Mimei information and data are now synchronized to current node.

***Mimei provider:***\
Node providing data of a Mimei is called its **Provider**. The *provide* command broadcasts to the network that this node provides all data for this Mimei.

```shell
./Leither mimei provide RXN74QNeiY08LRSaoeQhx3nOLTC
```
Returns:
```shell
MiMeiProvide cids= [RXN74QNeiY08LRSaoeQhx3nOLTC]
MiMeiProvide ok
```

*View supporting nodes, or providers*:
```shell
./Leither mimei findprovs 9OCLYP-SXzen3e171-Ei_6N3Gwl
```
Returns:
```shell
findprovs mid= 9OCLYP-SXzen3e171-Ei_6N3Gwl
Version:12
PeerID	: 12D3KooWN1J2uCHfNfGetGf7dhVkCtRxuZHrNGF2bVo5941yhBqh
Addrs	:[/ip4/54.151.221.76/tcp/8080 /ip4/172.31.27.218/tcp/8080]

PeerID	: 5TVMyFk-DsUH8n_FeF927OEoZSZ
Addrs	:[/ip6/2001:b011:e606:7bf5:2e0:1dff:feed:3d1/tcp/8080 /ip4/125.229.161.122/tcp/8080 /ip4/192.168.5.4/tcp/8080]
```

When users publish new mimei data, all online support nodes will synchronize the data in real-time.

***Mimei usage:***\
In addition to reading mimei data directly via API and commands, applications can be specified to render Mimei. For details, see the sections on "Applications" and "Domain Name Display of Mimeis."

### **Applications**
Developers can create applications using the Leither API, which is built on the <a href="https://github.com/hprose">Hprose</a> protocol. Hprose supports a wide range of popular programming languages, allowing users to seamlessly integrate Leither system functions into their applications. Special optimizations have been made for HTML5 applications to enhance performance.

**Create Leither App:**\
In Leither directory, create a sub directory, whose name will be the appName used in the folowing sections. Upload *Javascript* or *lua* scripts, which calls the Leither API, into the sub directory and use the following command to manage the new App.

**Application backup:**
```shell
./Leither lapp backup -a appName -k keyfile -n nodeId
```

**Upload to node:**
```shell
./Leither lapp uploadapp -i appName -k keyfile -n nodeId
```

**View App information:**
```shell
./Leither lapp showapp -a appName -v cur -k keyfile -n nodeId
```

**Publish an App to Leither network:**
```shell
./Leither mimei publish htpoEXiE6IlCAqVbCvjvkY_XNfu -k keyfile
```
After publishing the application to the network, it can be accessed and run on any node via a browser.

**Synchronize App to current node:**\
Synchronize applications from Leither network or specified nodes onto local node.
```shell
./Leither mimei sync htpoEXiE6IlCAqVbCvjvkY_XNfu
```
Returns:
```shell
MiMeiSync mid = htpoEXiE6IlCAqVbCvjvkY_XNfu
mimei sync ok
```
**Provide a Mimei:**\
After an App or Mimei synchronized locally, a node can respond to network calls by providing required information.
```shell
./Leither mimei provide htpoEXiE6IlCAqVbCvjvkY_XNfu
```
**Specify default application for a Mimei:**\
When creating a mimei, User can specify a default application to open it:
```shell
./Leither mimei create -a application_id
```
When opening a Mimei, the specified application will be used by default. It can also be specified upon opening a Mimei.

### **Rendering Mimei by Domain Name**
Through domain name nodes, users can expose their home nodes to others, similar to a host in an IDC data center.

The following link opens a Mimei with its default application. Mimei ID is Z5ZbV2tKVzl441nDUbinBITdzCZ. Domain name node is the node where domain name *www.fireshare.us* is resoved to.\
http://fireshare.us/mm/Z5ZbV2tKVzl441nDUbinBITdzCZ

The following link runs an application, application ID is FGPaNfKA-RwvJ-_hGN0JDWMbm9R:\
http://fireshare.us/tpt/FGPaNfKA-RwvJ-_hGN0JDWMbm9R

To ensure the Leither is accessible from the external network, it's important to note that all telecom operators support IPv6 and dynamic IPv4 addresses. Users should configure their local routers or optical modems to enable external access. In cases where a publicly accessible address is unavailable, a tunnel service can be utilized to facilitate external access through other networks.

### **Mimei Database**
Users can create a Mimei database that supports most Redis functionalities. It also implements HTML5 database functions. Users can publish their database to the network, and providers will automatically update their copies of the data in real-time.

### **Load Balancing and Fault Tolerance**
**Fault Tolerance:**
In a distributed network, node addresses can be complex, comprising IPv4, IPv6, and NAT-based addresses. When users access a domain name, a template webpage is returned, featuring code that determines the optimal network path to reach the nodes.

**Load Balancing:**
Using the provider mechanism, multiple nodes collaboratively supply user data and applications. When visitors access applications through domain names or URLs, domain and routing nodes update the providers' information on the template webpage. The browser then selects the optimal access node to achieve load balancing. Various load balancing strategies can be customized by modifying the JavaScript code in the template webpage.

### **More Functional Usage**
User can interact with the system through the following methods: system application interface, command line, system API, and browser. For convenience, this document introduces and demonstrates the functions of each module via command line.

For details, refer to <a href="./api/Api.md"> API Documentation</a>.

### **Security**
When a node initiates, it automatically generates a pair of keys, which are used to create a host ID that serves as the host's identity. The ***lpki*** command set encompasses functions for key management and operations related to identity authentication.

### **Network**
Upon node initiation, it joins Leither network through a bootstrap node specified in its configuration file. The ***swarm*** command set encompasses node address connection filtering and other functions.

The entire network is a DHT network. The *dht* command set includes network access and routing functions.

### **Mimei**
The Leither system utilizes Mimei types for all its applications and data, supporting both file and database formats. Mimei can be published on the network, allowing users to access it via its unique ID. Data synchronization across nodes is possible, enabling users to provide data support to others' Mimei. Each modification in a Mimei creates a new version, with real-time updates on all supporting nodes (data providers). 

By integrating domain names and load balancing mechanisms, the system achieves fault tolerance and load balancing capabilities.

***Mimei command set***.

Mimei system is compatible with IPFS network. The ***ipfs*** command set includes common IPFS commands, and the ***file*** command set works on the local IPFS file system.

### **Application System**
When creating Mimei, an application can be specified. When user open the Mimei, the application will be called to render it. The application itself is also a Mimei. The ***lapp*** command set manages application backup and release, including versioning and domain management.

### **Domain Name and Load Balancing**
Leither is compatible with various home devices, offering the benefits of speed and low cost in home networks. However, these networks often face challenges like unstable IPs and lower availability. Leither addresses these issues using domain name nodes (DNN). Users can bind resources (nodes, applications, content, which are all Mimei themselves) to a DNN. Third-party users can access a DNN through web browser, and the DNN returns a very small (3k) data packet. The browser interprets this packet to extract data and node details, selecting the optimal node to deliver services to third-party users. This node selection process, which involves fault tolerance and load balancing, is executed entirely within the browser, eliminating additional traffic from the DNN.

For the webpage viewer, there is no page jump, equivalent to directly accessing the user's node. The user experience is comparable to IDC data center services.

### **Introduction to Key Concepts**
<a href="./doc/MiMei-en.md">**Mimei and Application System**</a>\
How to dismantle internet platform businesses through decentralized applications.\
<a href="./doc/GongShi-en.md">**Organization and Consensus Mechanism**</a>\
How to coordinate the interests of all parties through consensus mechanism to encourage participation and cooperation.

### **Related Documents**
Links to project-related information PPT:
<https://zhuanlan.zhihu.com/p/330261216>

<a href="./doc/Pki-en.md">**Security System**</a>

<a href="./doc/Application-en.md">**Application System**</a>

<a href="./doc/Setup-en.md">**Node and Application Installation**</a>

<a href="./opt/dav/">**Example Application Code**</a>

<a href="./api/Api-en.md">**API Documentation**</a>

<a href="./doc/MiMei-en.md">**Mimei and Application System**</a>

<a href="./doc/GongShi-en.md">**Organization and Consensus Mechanism**</a>

<a href="./doc/PaaS.md">**Leither vs Traditional Cloud**</a>   
Leither system can be used for cloud services, offering a significant reduction in costs related to CPU, memory, and bandwidth compared to virtual hosting and Docker.
