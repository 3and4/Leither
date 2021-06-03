MIMEI
======
### Summary
An alarming trend of internet in the past few decades is the fast concentration of traffics on fewer apps or sites owned by tech giants, who used their dominant control of data to influence, even manipulate user behavior for monopolistic profits, at the expense of average users. However, we hold the following truth to be self-evident, that a free and open internet is of the best interest to its users.

Leither, a decentralized Cloud OS, is designed to take the challenge by decoupling and reconstructing most of the core services of regular internet in a decentralized fashion, therefore give the control of data back to user, its real owner. Leither OS has implemented all the subsystems necessary to reconstruct regular internet services, such as authentication and authorization, file system, database, application framework, domain name service and routing.

To reconstruct a regular internet application, for example Youtube or Google, simply creates a Leither Organization and a Consensus Contract to stipulate how to distribute tokens among participants. Token is a proof of contribution that guarantees proportional dividend of the organization's income, or exchanges for services such as traffics, routing, domain name resolution, search, backup, bandwidth, CPU power, contract notary, or mining pledge.

Members of an organization include initiator, investor, developer, content or service provider, device owner, operation and marketing, etc. The distribution of income is guaranteed by Consensus. Device owner does the work of miner in BTC. Devices keep the books as mining machines in BTC, at the same time provide more valuable business services of the organization. Compared with POW-based block-chain Leither block-chain has advantages in every aspect of cost, efficiency, performance and privacy, and last but not least, law abiding.

This paper has two parts:  
Part I: Introduction of MiMei and its application system.  
Part II: <a href="./GongShi-en.md">Organization and Consensus</a>.

### Part I: MiMei and its application
How to decouple and reconstruct regular internet platform services.

Some key concepts are defined as below:
+ **Data**  
Information stored on any media. **Data structure** defines how data is stored and organized. It is a set of data items that collaborate with each other following certain rules.

+ **Application**  
Data must be consumed or processed by program. There are two types of programs in general, system software and application program. The latter is used to process user data, aka **APP**.  

+ **MiMei**  
The key concept of Leither, coined referring to Gene or MEME. MEME is the gene of civilization, an element of a culture or system of behavior passed from one individual to another by imitation or other non-genetic means.

    MiMei is the carrier of information on internet. It defines the structure of data storage, and all kinds of operations by associated applications, such as create, open, save, read, search, render and user-interact. 

    Everything in Leither-based ecosystem is defined in term of MiMei. The referential relationships among MiMei makes it possible to develop comprehensive internet applications and services.

+ **Leither**  
is a container of MiMei, an ultra-light decentralized cloud OS.

    Leither is designed to support frequently used internet applications with extremely limited resources, for example to run office applications for a small business on a Raspberry Pi Zero.

    Leither OS has implemented all the subsystems necessary to reconstruct regular internet services, including authentication and authorization, file system, database, application framework, independent domain name service and routing. Leither API supports about 40 kinds of development languages.

    PS: Leither written in Kanji character is 弥媒. 弥 means extremely small and ubiquitous, usually something imperceptible. 媒 means media. 弥媒 put together almost as good as Ether, a substance once believed to be the building block of universe.

+ **Node**  
A smart device running Leither OS is a **Node**. As small as 6MB in size and runs on all regular operation systems, Leither supports almost any smart devices, like PC, NAS, router and Raspberry Pi.  
1. One independent node can serve as an ultralight server;
2. Two nodes can backup each other in services like data backup, error tolerance, load balance, domain name resolution and routing;
3. Multiple nodes can network to support decentralized internet applications.

    MiMei flows among nodes to facilitate the decentralization of centralized internet.

### I. Background -- Centralized Internet
#### 1.1 Big Data has conquered the world
Internet has changed our daily life. Big Data is running internet, therefore our lives. Online retail up to US$35 trillion in 2019 in China alone. The essence of Big Data is the accumulation of data and formation of rules, for the purpose of manipulating user behavior and generating monopolistic profit.

#### 1.2 Big Techs own Big Data
Big Techs own core data. For example, FB owns social media and Amazon online retail. Data protocols are manipulated to serve as business barriers. Google’s spider cannot access FB pages, vice versa. Monopoly is protected by corporation regulations, laws and ecosystem rules. Big Techs also take the lion’s share of profit, so their fortune concentrates accelerating.

#### 1.3 Users are caged by Big Data
Individual user is helpless against the invasion of Big Data into its privacy. All user behaviors are data tagged to feed data models, which is then used to predict and influence user behavior. Small business struggles under the scrutinizing peek of Big Data, to which the ledger of small business is an open book. Small game developers can barely share 10% of income on Tencent Game. Most online retailers are hard pressed to break even.

#### 1.4 Intellectual products of all mankind taken by Big Tech
The internet originated from ARPA net of US Army as non-profit project. It was borne an open system without a center. Telecom firms has invested billions of dollars to build telecom infrastructure, without controlling the data flowing through it. However, Tech Giants, such as Alibaba, Tencent, FB, Google, etc., has taken control of internet with a few APPs.

#### 1.5 Summary
The internet of Status quo has become a centralized playground of big players. By controlling user data and influencing user behaviors, monopolistic profit is being made at the cost of individuals, for whom a weapon to fight back is in dire need.

### II. System Requirement
Before presenting a plan to build decentralized internet, it is necessary to briefly review the different requests that computer and internet must address at each stage of their development, and the corresponding solutions presented. If a new solution cannot solve old problems satisfactorily, it will not beat the incumbent.
#### 2.1 File -- Carrier of Data
In the beginning of computer, program was developed to solve problems. A program is a sequence of data. With its development, program generated large amount of data in return. 

In Unix, File was designed as carrier of data. File System was designed to manage files and storage medias. Everything is file. Any information related device is of File type.
#### 2.2 MiMei -- Type of Data
With the evolution of applications, it became necessary to describe the relationship between file and application. DOS introduced the concept of File Extension. In Windows, file extension is bound with a corresponding application. In E-mail, MIME maintains a strict definition of file type as attachment. All existing web browsers still follow the MIME protocol.
#### 2.3 Database -- Relationship of Data
Database was designed to manage complicated data structure. A database is a closed set of datum and their relationships.  
Shortcomings: Data stored in closed space with little liquidity, value not maximized.
#### 2.4 HTML -- connections of data over network
WWW regulates relationships of data over a network. With the birth of internet, information explosion ensued. The inter-connection of data cross network breaks the physical limits of data stored in a single machine. Gateway sites flourished benefiting from this innovation.  
Shortcomings: connections based on linkage is unstable.
#### 2.5 Google -- Search Engine
Information explosion creates the request for search engine. Google provides content search service based on tags. Extra dimensions of properties are added to tagged data from then on.  
Shortcomings: Data is isolated and monopolized. Data protocol began to be tainted by private interest.
#### 2.6 Facebook -- Social Model
The early social media, ICQ, MSN, QQ solved the request for instant messaging. Twitter provides a platform for broadcasting short messages quickly.  
From MySpace to Facebook, a more social-media oriented model appeared. Social model usually consists of 4 parts:
|--|--|--|
|1|My settings|Personal information, settings and preferences|
|2|Friends|User contacts|
|3|App & Services|Application and services for the user|
|4|Message|Interactions with friends, applications and services|

This model is remarkably successful. All the internet social Apps adopted a similar framework, for example Wechat, Facebook, Twitter and some OA system for enterprise users.
Shortcomings: Internet becomes more monopolized, more isolated, and data protocols manipulated.
#### 2.7 Summary of Design
Design principles for breaking data monopoly and creating decentralized internet:
+ **User oriented in every aspect**
A new birth of freedom, and that internet of the people, by the people, for the people.
+ **Free and open system**
Open data protocols  
Open operation systems protocols
Only in an open system, user has the freedom of choice and power of ownership.
+ **Free Data Migration**
With user’s authorization and reasonable security discretion, user data shall be allowed to move from one platform to another and transform from one format to another. Only then, a third party can provide valuable applications and services.

The following requests shall be supported in a decentralized system:
+ File, file system, database
+ Association of data type and application
+ Connection between contents over network
+ Search of content over network
+ Data migration over network
+ Decentralized social media
+ Light-weight container system
+ Easy to use development environment
+ User experience compatible with existing internet
+ Law-abiding
### III. Requirement Analysis
First and foremost, an understanding of how data is monopolized.
#### 3.1 The origin of data monopoly
In nascent internet, data access follows open protocol across the whole network. Major platforms like Google and Microsoft, all provided open API for accessing its data.

The earliest abuse of open data policy was committed by Facebook, which used other platforms’ data to construct its own social media, then shut its own from outside. User data get stuck in Facebook’s platform and the first monopoly of data was borne.

Afterward Apple and Amazon created their own kingdom of user data, respectively. When Tencent published QQ, it used a proprietary communication protocol. When Baidu block Ad from Alibaba, the latter responded in kind.

Thus, the internet has become segregated. User data is fortified by impenetrable barriers created by every developer who can get his hands on it.

There are two forms of data access.  
**Data->App**
User has access to data and can choose the App to process the data. For example, in Windows user can select different applications to open certain type of file, or use different browsers to open a web page.  
In this case, User has the freedom of choice and is not restricted to single App.  

**App->Data**
User's right to access data is restricted to certain Apps, usually only one. For example, when user login an e-commerce site for shopping, or get on social media to chat with friends.

In this case, the algorithm decides the path for user to access data, so user behavior is controlled by the App.

The two approaches above corresponds to **Object Oriented Programming** and **Procedure Oriented Programming**, respectively.  
#### 3.2 Principles for breaking data monopoly
**Open Source Platform**  
Microsoft dominated the PC ecosystem with DOS and Windows, and strangled competition in both office software and browser. Google uses search engine to control data and influences user behavior. FB does the same with social media.

On the other hand, open source community turned operation system into a public platform with Linux. It is the first attempt to open source a platform. Similar cases are Android for smartphone and RISC-V for IC instruction set.

There shall be similar open source platform for collecting user data, especially for **Social Media** and **Search Engine**.  

**Free Data Migration**  
In early operation system, data file can be copied and transferred among different mainframes of different systems. Free movement of data shall be a consensus and primary design norm of operation system.  
Three measures can be adopted to achieve the above goal:
+ Promote data-freedom as a mainstream ideology
+ Legislate that any platform that collects user data must allow the owner to migrate its personal data.
+ Develop opensource program, such as browser plugins, to facilitate user to migrate data to personal storage devices. Once the ownership of user data taken back by its creator, the data can be revived and reconnected with each other.

**Separation of Data and Rule**  
If data can be migrated freely, user will have the choice of rules that can be applied to it. In another word, try to be **Object Oriented**, avoid Procedure Oriented development whenever possible. Data structure is relatively simple to be reverse engineered. It is easier to reconstruct program structure with the knowledge of data structure.  

**Granulation of User Data**  
There is no way to decouple convoluted and complicated dataset. It is best to granulate data in design, following the pattern of gene, keep the information relatively comprehensive and independent. A piece of granulated information is easy to be processed, transferred, and its connection with other objects clear to be understood.

**Connections of Data**  
Isolated data have little value compared with inter-connected dataset, whose connections contain much more information. HTML is perfect protocol to describe relationships, if not link could be unstable.  

#### 3.3 Data Container  
Many data centers are required to process data of billions of users in regular internet. A large data center has capacity for tens of thousands of servers. Every task is processed by groups of servers using parallel algorithm. A army of maintenance personals are required for daily operation.

No data center is required when storing personal data by individual. However, some kind of container, a light-weight clouds OS, will be necessary, like Pod promoted by Solid platform. The container shall be low cost, lightweight, and easy to operate by layman. Compared with the Homomorphic Encryption that certain block-chain alogrithm is using, Leither method is simpler and more efficient. One who owns the storage device owns the data and defines the rules to apply.

#### 3.4 Data Model for Reference  
**File**  
File is a set of program and data stored on media. File is constructed according to certain format, usually referred to as data structure. Operation system manages files with File System in a tree-like data structure. File wraps the access of data on media and optimizes the procedure.  

**WWW (World Wide Web)**  
In a standard webpage, HTML5 is used to render content of webpage, CSS style and JS tasks. A page can relatively store and render its data comprehensively. Through interconnections among pages, large scale and complicated information can be represented.

Shortcomings:
Webpage only contains information about frontend, without backend data. Big Techs take monopolistic control of user data by owning the backend servers. The connections among webpages are link based which is unstable for lacking definition about server or user.

**Gene**  
The basic unit that contains information of life. DNA or RNA stores a complete segment of hereditary material, which records the adaptation of life to its environment. Life copies hereditary information from generation to generation, while adapting to new environment by variation.

Limitation: Every iteration takes a generation, with restrictions on the amount of information inherited and weak capability to deal with complicated changes.  

**MEME**  
Richard Dawkins introduced the concept of MEME in his original book, the Selfish Gene. Similar to gene, MEME is defined as the base unit that contains information of a civilization. MEME can be used to record all the knowledge of human society. It can be remembered in human brain, propagated among people, or stored in any kind of media.

Human brain uses pattern recognition to model the world. During its interaction with environment, every reflection by the mind is an iteration of evolution. Compared with gene, MEME can evolute faster than gene by orders of magnitude.  
### IV. Introduction to MiMei
**Origin of MiMei**  
Entropy is one of the most important measurements in the universe and human society. The 2nd law of thermodynamics points out the direction along which the universe evolves. That is why entropy is also called Arrow of Time. Erwin Schrodinger in his book, What is Life, declares that what an organism feeds upon is negative entropy.

Life stores information of its environment in gene. Heredity keeps the adaptation and variant deals with the changes of environment. Life simply executes the instructions locked in its gene. Once human developed self-conscious, human brain begins to simulate adjustment made to its environment and process the information. The basic unit of such information is called **Meme** by Richard Dawkins.  

After the birth of internet, the method for processing information has changed and a new unit of information is required. We call it **MiMei**. Similar to Gene and Meme, MiMei contains a piece of information and the rules applicable to it.

**Implementation of MiMei**  
Based on **Sec 2.7 Summary of Design**, the following functionalities have be implemented.
+ MiMei operation: create, manage, save, render and send
+ MiMei creation by content's topic, each MiMei wiht a unique ID
+ Right to set permissions on MiMei by user
+ Referential relationship among MiMei
+ MiMei can run or be saved on multiple nodes
+ Version of MiMei, similar to Git, to solve problem of data consistency
+ Synchronization among nodes in real-time or after backup
+ Type of database and file, stream type coming soon
+ File system as a special system type, based on database and file
+ Independent space in database and file system for every MiMei
+ Heredity and evolution of MiMei by continuously backuping changes
+ User can also facilitate variation by fork
Leither is implemented as MiMei container system, and also a decentralized cloud OS.
<a href="../api/MiMei.md">MiMei API</a>  

**MiMei Features**
|Feature|Detail|
|--|--|
|Identification consistency|MiMei ID combined with its version can describe a piece of information stably and precisely|
|Complex information render|Support database, file system, referential relationship, App to express complicated information|
|Free Data Migration|MiMei can flow among different nodes in a relationship-chain or group, according to user behaviors|  
|Multi-dimension search|User can add tag to MiMei and search it through multiple nodes over multiple tags|

The above features will be explained in detail later.

### V. MiMei Identification
**Traditional method**  
Path of a resource in regular protocols,  
WWW: _http://server_ip[:port#]/path/file_name[parameter=value]_  
OS: _partition/path/file_name_  

In all of the above methods, every thing is changeable, server IP, port number, path, file name, parameter. Existing method cannot describe a resource precisely. In IPFS, a unique ID is generated by hashing the synopsis of a file. A new ID will need to be created if there is any changes to the file. This kind of uncertainty is troublesome to normal operations and referential integrity.

There is a brand new solution: **MiMei ID**  
In the decentralized network of Leither system, all resources are described by MiMei ID, including user, node, content, application, etc. ID of user and node is generated by encryption key pairs. MiMei ID of file or data block is based on synopsis of its content.

There are two sets of MiMei ID for external and internal reference. MiMei ID is the only unique ID of a MiMei object, created with the creation of the object, and never changes afterward. MiMei ID can be used in index, operation and reference. The data of MiMei is stored in file or database. Each time the data is changed, a new version of backup is created with an ID based on the synopsis of its content.

**MiMei Information**  
Information of MiMei can be defined as:  
```golang
type MiMeiInfo struct {
	Author  string    //creator
	AppType string    //associated application type
	Ext     string    //extension
	Mark    string    //MiMei ID
	Create  time.Time //timestamp of creation
	Right   uint64    //authorized rights
}
```
**MiMei ID unchangeable**  
MiMei ID is based on information of creator, associated application, MiMei type and MiMei mark. Once created, never changes. MiMei ID is invariable.

**Access MiMei data with MiMei ID and version**  
New version is created while MiMei being edited or backed up. Data of the version is based on synopsis of content. Content of the version is read-only. MiMei ID and version combined can identify certian MeMei data.

### VI. Data storage and corelation
MiMei can support data storage of most internet applications with its support of file and database.
#### 6.1 Granulation of Information
MiMei granulates information. It is recommended to redefine the following types of information with MiMei data type, aka Mimeimization:  
+ Content needed to be indexed
+ Content for sharing among users
+ Content might migrate among nodes  
Mimeimization can be executed beforehand, or on demand. The latter is actually a split. A new MiMei object detached from the original one. A referential relationship keeps the tie.  

Leither supprots database and file system. User can use traditional development method if only its own data is concerned. In this case, the whole user or application is one stand alone MiMei object.

#### 6.2 MiMei version
Both MiMei file system and database support version.
+ _cur_: the current working copy, not yet backed up.
+ _last_: the latest backup copy, represents the newest confirmed content
+ _ease_: ready for publish
With versioning mechanism, we can retrieve MiMei data of different version with its ID and version number.

#### 6.3 MiMei 
Isolated data cannot express complex relationship. With unique MiMei ID, it is possible to establish table relational structure. Information can be saved in MiMei file or database, with a format defined by its associated application.

Corelation between MiMeis is formed by their  relationships. Referential information includes MiMei ID and number of references. Ordinarily the semantic relevance of data content is interpreted by its associated application. Leither system cannot access App data, so it is the App's task to maintain the  relationships of MiMei by calling corresponding API.

MiMei application can generate  information according to semantic correlation.  
API: MMAddRef MMDelRef MMGetRef

Most MiMei contains only granulated piece of information, the whole picture can be described through correlation of MiMei.  
#### 6.4 File system  
Originally file system was created for mainframe where number of applications and volume of data is limited. With the development computer and internet, mobile phones, number of users, Apps and data volume all increase explosively. The following shortcomings of file system began to be revealed.  
1. File name cannot precisely identify a file
2. Insufficient index information, only path and file information.
3. One file might have multiple duplicated copies
In Leither, traditional file system is converted into MiMei framework with the following improvements,  
+ Unique MiMei ID: Every file has a unique ID created based on its synopsis as its mark  
+ _cur_ version: Current version as target of data access  
+ Reference count: Do not copy data, increase count of reference instead. Use MiMei ID to access its data.  
The directory structure of file system is also **granulated**(Ch6.1). Directory that is indexed or shared with high hit count shall be Mimeimized. Currently Leither manages file directory with JSON format, which performs poorly when the number of files grow too large. In future database will replace JSON. The algorithm for extracting synopsis is similar to that of Merkle tree.

#### 6.4.1 File objects in Leither system
+ MiMei Current File  
The current version of MiMei file, both readable and writable.
+ MiMei Version File  
The read-only backup files of MiMei. Created by backing up the current file. Leither creates a version number for each backup file.
+ Mac file  
File labelled with Mac ID that is generated using content of MiMei file. Read-only. Mac file can be referred by MiMei object. MiMei version file is a Mac file referred to by a version number. Temp file can also be convert into Mac file.  
+ Temp file  
Created by MFOpenTempFile after data is written into it, or Converted by MFTemp2MacFile.
+ MiMei File System  
A built-in special MiMei type of Leither system. It can be created by MMCreate, or placed in _webdav_ directory, or created by Leither automatically using configure file with .mmfs extension, or opened by MFOpenByPath as file system object.  
+ Operation System File and Directory  
Create a link of a directory in _webdav_, the linked directory and file can be opened by MFOpenByPth.
+ MiMei Root Directory  
_webdav_ is the general entrance to access MiMei in a node. File or Directory can be linked into _webdav_, or a configure file can be create that points to a MiMei object in the node.  

#### 6.4.2 Open File
+ MiMei File  
MMCreate: Create MiMei file, take a api.MM_File as input, return MiMei ID.  
MMOpen: Open MiMei file, return handler ID for file operation. If _ver_ is "cur", file is writable, otherwise read-only.
+ Open File by Path  
MFOpenByPath: Open file in MiMei file system or Leither file system.
+ Open Mac File  
MFOpenMacFile: Open a mac file under a MiMei.
+ Open Temp File  
MFOpenTempFile: Open a temporary file for read or write.
MFTemp2MacFile: Convert to Mac file after read or write operation.
+ Close File  
All the files need to be closed after operation, except temporary file. Considering the inevitable off-line scenarios, all handles of file operation will be closed after timeout.

#### 6.4.3 File Operation  
+ Object Method  
MFSetObject  MFGetObject
+ Byte Array Method  
MFSetData MFGetData
+ Query Status
MFGetSize MFStat MfIsExist MFGetMimeType
+ Directory Operation
MFReadDir
+ File System
FSFind FSMkDir FSRemoveAll FSStat FSRename
+ Others
MFTruncate MFCopy
#### 6.5 Database
There are two underlying databases. One is based on LevelDB and the other BoltDB. Both are revised in underlying code. LevelDB is used for current version to support read and write, consensus is based on time sequence. During a write transaction, changed data will be checked to see if it is revised by any other party during the transaction. If it does, transaction fails and initiator will be called to redo the transaction. BoltDB is used for backup version and read-only.

API refers to Redis, support 5 types of data: string operation, hash, list, set, ordered set, and transaction.
+ Transaction
Begin Commit Rollback
+ String
Set Get Incr IncrBy Strlen  
+ Hash
Hmclear Hdel Hlen Hget Hmget Hmset Hgetall Hkeys Hscan Hrevscan HincrBy  
+ List
Lpush Lpop Rpush Rpop Lrange Lclear Lmclear Lindex Llen Lset  
+ Set
Sadd Scard Sclear Sdiff Sinter Smclear Smembers Srem Sunion Scan  
+ Odered Set
Zadd Zcard Zcount Zrem Zscore Zrank Zrange Zrangebyscore Zremrangebyscore Zrevrange Zrevrangebyscore Zmclear Zclear ZincrBy  
### VII. Application System
#### 7.1 History of Application
There a four types of application model:  
1. Local Application  
The earliest form of application. Both data and application are on the same machine. Data is shared through files.  
2. P2P Application
In local network, application is responsible for communication with other terminals. Applications on each node equally takes care of business of itself. When there are too many nodes, one node may be selected for public services. This is the predecessor of server.  
3. Client/Server Model  
The development of large network and internet gave birth to dedicated server. In the beginning, server usually processed core business logic only, most of the specific tasks were handled by client machines.
4. Browser/Server Model
In the time of Internet, browser became the major client to call for remote services. With the enhancement of Javascript, more and more jobs are executed on browsers, aka B/S model. B/S model greatly reduce the complication of end user's task.

**The Merging of Models**  
In the powerful HTML5, more and more tasks in B/S model is executed in front end, similar to C/S model. At the same time more and more Apps are developed in HTML5, very much like B/S model.

#### 7.2 Leither Solution
Leither API supports 40+ development languages, particular HTML5 because large number of frequently used applications are developed with HTML5, including website, App, Applet, etc. HTML5 is not perfect for the lack of standards for backend data processing and business logic. Leither supplements the shortcomings in HTML5 for building cloud applications. Leither provides user authentication, application system, cloud file system and database, decentralized domain name resolution, data redundancy and load balance.

**Business Logic in Front-end**  
By default business logics shall be executed in front-end. The benefit is that Leither application development will be similar to single page application development, simple and easy.

**Thin Node**  
PC and mobile phone are not suitable devices to serve as Node. Computer server is maintenance heavy and expensive. Router and TV box are usually proprietary system with restricted accessability. NAS and other customizable hardware are preferred equipments to run as Node. Less powerful CPU is ideal to support fine granulated Leither applications.

**Special Optimization**  
Because Leither node is intended to be maintained by layman, node setup and application installation must be dummy proof. Node is optimized for home network, with the unique decentralized domain name resolution of Leither system, Leither can provide service similar to cloud server.  

Leither node can also be optimized to support batch execution and back-end tasks.

**Mimeimization of Application**  
Leither provides comprehensive application development framework, within which an application is actually a type of MiMei. Refer to <a href="Applition.md"> Application System Development Documentary</a>  

#### 7.3 Design Principle
#### 7.3.1 HTML5 is the core standard and norm of internet
Most frequently used Apps are developed in HTML5 now, including App, website, applet, etc.
#### 7.3.2 HTML5 is not a comprehensive internet application development solution
HTML5 provides only localized data storage (Web Storage and IndexedDB), without setting up the standard for backend data and business logic. The result is internet applications depend heavily on the backend to process business logic and data. Front-end developers became low caste because they do not get in touch with business logic and core data.
#### 7.3.3 Leither complements the missing functions in HTML5
Leither provides functions of authentication and authorization, application development framework, cloud file system and database, decentralized domain name resolution, error tolerance and load balance.
#### 7.3.4 Leither is ideal to build cloud service
Besides constructing
#### 7.4 Advantages of Leither Cloud
#### 7.4.1 App Development is simple
Knowledge of HTML5 is sufficient to build most could services, such as website, App and applet. The workflow is almost identical to a local HTML application.
#### 7.4.2 Extremely low requirement of system resource
Leither is developed with low CPU, low memory devices in mind. Application consumes very little system resources, because business logic is executed in frontend by default, therefore greatly reduces server load. The same server running Leither can withstand up to 100 times of workload of regular system.

#### 7.4.3 Low Bandwidth Cost
For traffic heavy application, traffic load can be balance on home network based network, and cost only 1/10th of regular network.

#### 7.4.4 Elastic Services
System application and data both support docker-like sandbox. Application file and database both support versioned backup. One physical machine can support multiple user, multiple Apps and their data.  
With invariable domain name and link, App and data can flow among different physical machines.   
elastic service is supported by decentralized domain name and MiMei.

### VIII. Flow or MiMei Information
#### 8.1 Why information must flow  
**The meaning of life is the spread of its information**  
Gene and Meme are the carriers of information of life, whose purpose is to occupy more time and space. Similarly the purpose of MiMei is the propagation of meaningful information.  

**The flow of information appreciates its value**  
FileCoin conveyed a misunderstanding that data is asset. Data itself is actually a debt, for device, storage and network all cost money. The higher value of data is appreciated only when data turn into traffic. Business of data is mainly the operations of save and retrieve. However, business of data traffic involves the creation or collection of content, save, propagation, render, value appreciation and business development. Each time a content is displayed, new value is appreciated. The number of renders and the value of each display determine the value created by the content.  

**The flow of information is a business requirement**  
In traditional network, data backup, error tolerance, load balance and elastic cloud services are all the migration of data in nature.
#### 8.2 The procedure of information flow
Information can flow in two different situations:  
1. Save voluntarily.   
Usually when the MiMei data is highly valuable, it can be saved for future benefit.
2. Save on request.   
When current node lacks the capacity to uphold regular services, it may ask other nodes to help doing backup, error tolerance, load balance, etc.

Procedure of save:  
+ Mimeimize Information  
After information is MiMeimized, shared information can be saved in file block.  
+ Duplicate data over node network  
With proper authorization, copy MiMei blob from one node to another  
+ Update routing information of MiMei  
Consumer of MiMei content will get its routing information first to find its data. Routing information includes data of the MiMei node and its newest changes. After routing information is updated, user will get updated MiMei information.

### IX. Social Model and Credential System
Ch2.6 mentioned four parts of social media model: Me, Contact, Application and Message. In Leither the corresponding model looks as below:  
|Category|Detail|
|--|--|
|Me|Authentication and authorization of user|
|App and service|App and service of the user |
|Message|Manage interactions between user and Apps|  
#### 9.2 Credential Model
During the exchange of information and services between user(node) and user(node), there must be a mechanism of quantitative settlement. Currency is the tool to settle business in daily life. In Leither network it is the exchange of credit.

Every node can grant other nodes credit, which authorize the others to use its services. The surplus of services provided over received is the equivalent of a receipt that a borrower gives the lender, or an cryptocurrency issued by the borrower. Within a group, cryptocurrency issued by a member actually is a certificate of services guaranteed receivable by the issuer. This certificate can be used to exchange for services from other members, as long as they recognize it.  

**Object of Credential**  
In reality fiat currency is the object with the highest credit rating. In Leither network, services such as bandwidth, storage and CPU can serve a similar function, aka the object of credit rating in virtuality. Their order of importance is bandwidth > storage > CPU. Because of the characteristics of node network, bandwidth is the preferred base unit to measure settlement between nodes. All the other forms of services can be converted into bandwidth with a certain coefficient.

**Individual Credit Model**  
is a type of User-to-User (P2P) credit model.  
+ Letter of Credit (L/C)
L/C can be understood as the amount of services allowed to use before payment, whereas IOU is the amount of services that have been used without paying. A L/C includes information of the authorizer, the authorized, unit of credit, maximum credit, signature of authorizer.  
+ IOU and Currency  
An IOU includes information about lender, borrower, unit of credit, amount, signature of borrower. For borrower, IOU is equivalent to issuing **Cryptocurrency**.

**How to buildup credential**  
+ Initial Grant of Credit  
If there is a relationship between nodes, there shall be certain basis for mutual trust, so that certain amount of credit can be authorized according to the magnitude of that relationship.  
+ Periodical Credit Increment  
After relationship between nodes established, credit can be increased periodically.  
+ Mutual Help  
Each node is capable of useful services. Credit can be acquired by proactively providing valuable service to the others, including data backup, load balance, routing, search, or mediator service.  
+ Credit Revocation  
User can revoke credit granted to others in case of transgression.  
+  Acquisition of Credit  
User can purchase extra credit through other channels, say, off-line purchase.  
+ Credit Exchange  
User can exchange for more credit through 3rd party. Reference to <a href="./GongShi-en.md"> Organization and Consensus</a>.  

**Trading Procedure**  
With sufficient credential, user can exchange for services with credit and a record will be created in public ledger. When user exceeds its credit cap, service will stop.  

**Agency Service**  
When there is no direct connection between nodes, an agent node can be used to mediate the credit exchange. It is also possible use other trusted media to facilitate the exchange, such as fiat currency, cryptocurrency or **group currency** mentioned in 2nd part.  
### Part 2
<a href="./GongShi-en.md"> Organization and Consensus</a>  