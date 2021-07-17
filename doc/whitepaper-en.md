Decentered network and its application in smart contract
========
### Introduction
An alarming trend of internet in the past few decades is the fast concentration of traffics on fewer apps or sites owned by tech giants, who used their dominant control of data to influence, even manipulate user behavior for monopolistic profits, at the expense of average users. However, we hold the following truth to be self-evident, that a free and open internet is of the best interest to its users.

Leither, a decentralized Cloud OS, is designed to take the challenge by decoupling and reconstructing most of the core services of regular internet in a decentralized fashion, and give the control of data back to user, its real owner. Leither OS has implemented all the subsystems necessary to reconstruct regular internet services, such as authentication and authorization, file system, database, application framework, domain name service and routing.

The user experience in Leither system is compatible with regular internet, but Leither has the following advantages.
+ Easy application development
+ Extremely low requirement of system resources
+ Low cost of bandwidth
+ Simple implement of elastic cloud services

To reconstruct a regular internet application, for example Youtube or Google, simply creates a Leither Organization and a Smart Contract to regulate how to distribute tokens among participants. Token is a proof of contribution that guarantees proportional dividends of the organization's income, or can exchange for organization services such as traffics, routing, domain name resolution, search, backup, bandwidth, CPU power, contract notary, or mining pledge.

Members of an organization include founder, investor, developer, content or service provider, device operator, operation and marketing, etc. The distribution of income is guaranteed by consensus. Device owner does the work of miner in BTC. Devices keep the books as mining machines in BTC, at the same time provide more valuable business services of the organization. Compared with POW-based block-chain, Leither has advantages in all aspects of cost, efficiency, performance and privacy, and last but not least, law abiding.

+ **MiMei**  
The key concept of Leither, inspired by Gene and MEME. MEME is the gene of civilization, an element of a culture or system of behavior passed from one individual to another by imitation or other non-genetic means.

    MiMei is the carrier of information on internet. It defines the structure of data storage and operations by associated applications, such as create, open, save, read, search, render and user-interact. 

    Everything in Leither-based ecosystem is defined in term of MiMei. The referential relationships among MiMei makes it possible to develop comprehensive internet applications and services.

+ **Leither**  
is a container of MiMei, an ultra-light decentralized cloud OS.

    Leither is designed to support frequently used internet applications with extremely limited resources, for example to run office applications for a small business on a Raspberry Pi Zero.

    Leither OS has implemented all the subsystems necessary to reconstruct regular internet services, including authentication and authorization, file system, database, application framework, independent domain name service and routing. Leither API supports about 40 kinds of development languages.

    PS: Leither written in Kanji character is 弥媒. 弥 means extremely small and ubiquitous, usually something imperceptible. 媒 means media. 弥媒 put together is almost as good as Ether, a substance once believed to be the building block of universe.

+ **Node**  
A smart device running Leither OS is a **Node**. As small as 6MB in size and compatible with all regular operation systems, Leither can run on almost any smart devices, like PC, NAS, router and Raspberry Pi.  
    1. One independent node can serve as an ultralight server;
    2. Two nodes can backup each other in services like data backup, error tolerance, load balance, domain name resolution and routing;
    3. Multiple nodes can network to support decentralized internet applications.
Part I: LeitherOS and MiMei
========
### I. System analysis of decentralized internet
Before presenting a plan to build decentralized internet, it is necessary to briefly review the different requests that computer and internet must address at each stage of their developments, and the corresponding solutions presented. If the new solution cannot solve old problems satisfactorily, it cannot beat the incumbent.  
#### 1.1 File -- Carrier of data
In the beginning of computer age, file was designed as carrier of data. File system was designed to manage files and storage medias. In Unix everything is file and any information related device is of File type.  
#### 1.2 MiMei -- Type of data
With the evolution of applications, it became necessary to describe the relationship between file and application. DOS introduced the concept of File Extension. In Windows, file extension is bound with a corresponding application. MIME defines a strict rule for file type as E-mail attachment. All existing web browsers still follow the MIME protocol.
#### 1.3 Database -- Relationship of data
Database was designed to manage complicated data structure. A database is a closed space of datasets and their relationships.  
Shortcomings: Data stored in closed space with little liquidity, value not maximized.
#### 1.4 HTML -- connections of data over network
WWW regulates connections of data cross machines over a network. The inter-connection of data cross network breaks the physical limits of data storage on mainframe. With the birth of internet, information explosion ensued and gateway sites flourished.  
Shortcomings: connections based on linkage is vulnerable.
#### 1.5 Google -- Search engine
Information explosion creates the request for search engine. By adding tags to data, Google provides content search service by keywords.  
Shortcomings: Data is more and more isolated and monopolized. Data protocol began to be influenced and controlled.
#### 1.6 Facebook -- Social model
The early social medias, ICQ, MSN, QQ, solved problem of instant messaging among people. Twitter provides a platform for broadcasting short messages quickly.  
From MySpace to Facebook, a more systematic social media model is created. The consists of 4 parts:  
|#|Title|Detail|
|--|--|--|
|1|My settings|Personal information, settings and preferences|
|2|Friends|User contacts|
|3|App & Services|Application and services for the user|
|4|Message|Interactions with friends, applications and services|

This model is remarkably successful. All the afterward internet social Apps adopted a similar framework, for example Wechat, Facebook, Alipay and some OA system for enterprise users.  
Shortcomings: Internet becomes more monopolized, more isolated, and data protocols manipulated.
#### 1.7 Design principles of decentralized internet
Design principles for breaking data monopoly and creating decentralized internet:
+ **User oriented in every aspect**  
A new birth of freedom, and that internet of the people, by the people, for the people.
+ **Free and open system**  
    1. Open data protocols.  
    2. Open operation system protocols.  
Only in an open system, user has the freedom of choice and power of control.
+ **Unrestricted data migration**  
With user’s authorization and reasonable security discretion, user data shall be allowed to move from one platform to another and transform from one format to another. Only then, a third party can provide valuable applications and services.

The following requests shall be supported in a decentralized system:
+ File, file system, database
+ Association of data type and application
+ Association of contents over network
+ Search of content over network
+ Data migration over network
+ Decentralized social model
+ Light-weight container system
+ User friendly development framework
+ User experience compatible with existing internet
+ Law-abiding
### II. How is data monopolized?
In nascent internet, data access protocol was open. Major platforms like Google and Microsoft, all provided open API for accessing their data. The earliest abuse of open data policy was committed by Facebook, which used other platforms’ data to construct its own social media, then shut the others out.

User data get stuck in Facebook’s platform and the first monopoly of data was borne. Afterward Apple and Amazon created their own kingdoms of user data, respectively. Thus internet became segregated. User data is fortified by impenetrable barriers created by every monopolist who can get its hands on it.

There are two types of data access.  
**Data->App**  
User can freely access data, then choose whatever application to process the data. For example, in Windows user can select different applications to open certain type of file, or use different browsers to open a web page.  
In this case, User has the freedom of choice and is not restricted by certain App.  

**App->Data**  
User's right to access data is restricted to certain Apps, usually only one. For example, when user login an e-commerce site for shopping, or get on social media to chat with friends.
In this case, the algorithm of App decides the path for user to access data, so user behavior is controlled by the App's algorithm or rules.

The two approaches above correspond to **Object Oriented Programming** and **Procedure Oriented Programming**, respectively.  
#### 2.1 Principles for breaking data monopoly
**Open source platform**  
Microsoft dominated the PC ecosystem with DOS and Windows, and strangled competition in both office software and browser. Google uses search engine to control data and influences user behavior. FB does the same with social media.

On the other hand, open source community turned operation system into a public platform with Linux. It is the first attempt to open source a platform. Similar cases are Android for smartphone and RISC-V for IC instruction set.

There shall be similar open source platform for collecting user data, especially for **Social Media** and **Search Engine**.  

**Unrestricted data migration**  
Free migration of data shall be a consensus and primary design norm of operation system. Three measures can be adopted to achieve the above goals:
+ Promote data-freedom as a mainstream ideology
+ Mandate that any platform that collects user data must allow the owner to migrate its personal data. (European Commission proposed a similar legislation recently)
+ Develop open source programs, such as browser plugins, to facilitate user to migrate data to personal storage devices. Once the ownership of user data taken back by its creator, the data can be revived and reconnected with each other.

**Rule subordinates to data**  
If data can be migrated freely, user will have the choice of rules to apply to it. In another word, try to be **Object Oriented**, avoid **Procedure Oriented** development whenever possible. Data structure is relatively simple to be reverse engineered. It is easier to reconstruct program structure with the knowledge of data structure.  

**Granulation of User Data**  
There is no way to decouple convoluted and complicated dataset. It is best to granulate data in design, following the pattern of gene, keep the information relatively comprehensive and independent. A piece of granulated information is easy to be processed, transferred, and its connection with other objects clear to be understood.

**inter-connections of data**  
Isolated data have limited value. Inter-connected data can express much more complicated information. HTML is a perfect protocol to describe relationships, if not link could be unstable. The connections of information within a database is still of limited scope.  

#### 2.2 Data Container  
No data center is required when storing personal data by individual. However, some kind of easy-to-use, cheap and maintenance free data-container will be necessary, like Pod promoted by Solid Platform. One who owns the storage device owns the data and defines the rules of application. Compared with the Homomorphic Encryption that certain block-chain is using to secure data, Leither method is simpler and more efficient.

### III. Introduction to MiMei
Entropy is one of the most important measurements in the universe and human society. The 2nd law of thermodynamics points out the direction along which the universe evolves. That is why entropy is also called Arrow of Time. Erwin Schrodinger in his book, What is Life, declares that what an organism feeds upon is negative entropy.

Life stores information of its environment in gene. The heredity keeps the adaptation and the variation deals with the changes of environment. Life simply executes the instructions locked in its gene. Once human developed self-conscious, human brain begins to simulate adjustment made to its environment and process the information. The basic unit of such information is called **Meme** by Richard Dawkins.  

After the birth of internet, the method for processing information has changed and a new unit of information is required. We call it **MiMei**. Similar to Gene and Meme, MiMei contains a piece of information and the rules applicable to it.

_**MiMei flows among nodes to facilitate the decentralization of centralized internet**_.

#### 3.1 Implementation of MiMei
Based on **Sec 1.7 Summary of Design**, the following functionalities have be implemented.
+ MiMei operation: create, manage, save, render and send
+ MiMei creation by content's topic, each MiMei with a unique ID
+ Right to set permissions on MiMei by user
+ Referential relationship among MiMei
+ MiMei can run or be saved on multiple nodes
+ Version of MiMei, similar to Git, to solve problem of data consistency
+ Synchronization among nodes in real-time or after backup
+ Type of database and file, stream type coming soon
+ File system as a special system type, based on database and file
+ Independent space in database and file system for every MiMei
+ Heredity and evolution of MiMei by continuously backing up changes
+ User can also facilitate variation by fork
Leither is implemented as MiMei container system, and also a decentralized cloud OS.
<a href="../api/MiMei.md">MiMei API</a>  

#### 3.2 MiMei Features
|Feature|Detail|
|--|--|
|Unique label|MiMei ID with version number can describe a piece of information consistently and precisely|
|Complex information rendering|Support database, file system, referential relationship and App. Describe complicated information|
|Free Data Migration|Construct relationship chain or group. MiMei can flow among different nodes according to user behaviors|  
|Multi-dimension search|User can add tags to MiMei and customize multi-dimension search result through multiple nodes|

### IV. MiMei Label
#### 4.1 MiMei ID  
In the decentralized network of Leither, all resources are described by ID, such as user, node, content, application, etc. ID of user and node is generated by encryption key pairs. Label ID of file or data block is based on synopsis of its content.

There are two sets of labels for external and internal reference. 
+ External label is MiMei ID  
 **MiMei ID is the unique ID of a MiMei object and never changes after its creation**.  
 MiMei ID can be used in index, operation and reference of the object.  
 MiMei data is stored in file or database. Each time the data is changed, a new version of backup is created with an ID based on the synopsis of its content. The version number increments from 0.  
 For convenience the latest version of backup data is called _last_. 
 + Internal label is content synopsis ID  
Every backup version of MiMei is pointed to by a file or database ID, which is generated by hashing synopsis of content.

#### 4.2 MiMei Information
MiMei can be defined as:  
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
**Access MiMei data with MiMei ID and version number**  
New version is created while MiMei being edited or backed up. New version ID is based on synopsis of its content, which is read-only after backup. Historical MiMei data can be retrieved by MiMei ID combined with version number.
### V. MiMei storage and association
MiMei can support data storage of most internet applications with its support of file system and database.
#### 5.1 Granulation of Information
MiMei granulates information. It is recommended to redefine the following types of information with MiMei data type, aka Mimeimization:  
+ Content needed to be indexed
+ Content for sharing among users
+ Content that migrate among nodes  
Mimeimization can be executed beforehand, or on demand. The latter is actually a split. A new MiMei object detached from the original one. A referential relationship keeps the tie.  

Leither supports database and file system. User can use traditional development method if only its own data is concerned. In this case, one stand alone MiMei object is constructed for the user or application.

#### 5.2 MiMei version
Both MiMei file system and database support version mechanism. Data version can be retrieved by MiMei ID.
+ _cur_: the current working copy, not yet backed up.
+ _last_: the latest backup copy, represents the newest confirmed content
+ _release_: tested and ready for release  

#### 5.3 Association of MiMei
Isolated data cannot describe complicated information. With unique label of MiMei, it is possible to establish relatively stable relationship structure. Information can be saved in MiMei file or database, with format defined by its associated application.

Association between MiMei is established by reference. Referential information includes MiMei ID and number of references. Ordinarily the semantic association of data is interpreted by its associated application. Leither system cannot access App data, so it is the App's job to maintain the reference of MiMei by calling corresponding API.

Most MiMei contains only granulated piece of information, the whole picture can be described through association of MiMei.  
#### 5.4 File system  
Originally file system was designed for mainframe where the number of applications and volume of data is limited. File system increased efficiency by saving applications from the task of managing and accessing storage media. With the development computer and internet, especially mobile internet, the number of users, Apps and data volume all increased explosively. The following shortcomings of file system began to appear.  
1. File name cannot precisely label a file
2. Insufficient index information, only path and file information.
3. One file might have multiple duplicated copies  

In Leither, traditional file system is converted into MiMei framework with the following improvements,  
+ Unique label: Every data file has a unique ID generated from its synopsis, as its label.  
+ _cur_ version: Current version as target of frequent data access.  
+ Reference count: Do not copy data, increase count of reference instead. Use unique label to access its data.  

The directory structure of file system is also **granulated** (Ch6.1). Directory that is indexed or shared with high hit count shall be Mimeimized. Currently Leither manages file directory with JSON file. In future database will replace JSON. Leither has comprehensive instruction set and API for file and database operations.

### VI. Flow of MiMei Information
#### 6.1 Why information must flow  
+ **The meaning of life is the spread of its information**  
Gene and Meme are the carriers of information of life, whose purpose is to occupy time and space as much as possible. Similarly the purpose of MiMei is the propagation of meaningful information.  

+ **Hot data is more valuable**  
FileCoin conveyed a misunderstanding that data is asset. Cold (static) data is actually a debt, for device, storage and network all cost money. The value of data appreciates only when data turn into traffic (get hot). Services to support data is mainly the operations of save and retrieval. However, services to support data traffic involve the creation or collection of content, save, propagation, render, value realization and business development. Each time a content is displayed, more value is created. The number of renders and the value of each display determine the value created by the content.  

+ **The flow of information is a business requirement**  
In traditional network, data backup, error tolerance, load balance and elastic cloud services are all migration of data in essence.
#### 6.2 The procedure of information flow
Information can flow in two different situations:  
1. Save voluntarily.   
Usually when the MiMei data is highly valuable, it can be saved for future benefit.
2. Save on demand.   
When a node is in danger of overloading, it may ask other nodes to help in backup, error tolerance, load balance, etc.

Procedure of save:  
+ Mimeimize information  
After information is MiMeimized, shared information can be backed up as file block.  
+ Duplicate data over node network  
With proper authorization, copy MiMei block from one node to another  
+ Update routing information of MiMei  
Terminal content consumer will get routing information of MiMei first, then accessing its data. Routing information includes node information and the newest changes of MiMei. After routing information is updated, user can get updated MiMei information.
### VII. Leither Application System
Leither API supports 40+ development languages, particular HTML5 because large number of frequently used applications are developed in HTML5, including website, App, Applet, etc. HTML5 lacks standards for backend data processing and business logic. Leither makes up the missing functionalities of HTML5 in building cloud applications, including user authentication, application system, cloud file system and database, decentralized domain name resolution, data redundancy and load balance.

#### 7.1 Leither Solution
**Business Logic in Front-end**  
By default business logics shall be executed in the front-end. The benefit is that Leither application development will be similar to single page application development, simple and easy.

**Thin Node**  
PC and mobile phone are not suitable for long term stable service. Server is maintenance heavy and expensive. Router and TV box are usually proprietary system with restricted accessability. NAS and other customizable hardware are preferred devices. Less powerful CPU is ideal to support fine granulated Leither applications.

**Special Optimization**  
Because Leither node device is intended to be maintained by layman, node setup and application installation must be dummy proof. Node is optimized for home network, with the unique decentralized domain name resolution of Leither system, Leither can provide service similar to commercial cloud server.  

Leither node can also be optimized to support batch execution and back-end tasks.

**Mimeimization of Application**  
Leither provides comprehensive application system, within which an application is actually a type of MiMei. 

#### 7.2 Advantages of Leither Cloud
**Leither is ideal to build cloud service**  
Besides capable of building complete cloud applications, all of the applications and data within Leither system is of MiMei type, which can freely migrate among nodes. MiMei also enables load balance mechanism of Leither. All those powerful features make Leither ideal for building cloud service platform. The best practice of Leither App development is to execute business logic in the frontend, therefore greatly reduce the load on the backend. The same server running Leither can outperform traditional platform 100 times or more.

**Simple Application Development**  
Knowledge of HTML5 is sufficient for most of the work for developing cloud services or website, App and applet independently. The workflow is almost identical to a HTML single page application.

**Extremely low demand on system resources**  
Leither is developed with weak CPU and small memory devices in mind. Leither applications consume very little system resources of the node, because data and business logic are executed in the frontend, usually user's browser or App.

**Low bandwidth cost**  
For traffic heavy applications, traffic load can be efficiently balanced by nodes in users' home network, which is cheap. The cost of traffic is only 1/10th of regular web service.

**Elastic service**  
Both Leither system applications and data support docker-like sandbox. Both file and database support versioned backup. One physical machine can support multiple users, multiple applications and their datum. Application and data can flow among different physical machines freely without changing domain name or link address.   

Elastic service of Leither platform is enabled by decentralized domain name and MiMei mechanism.
### VIII. Leither credit system
During the exchange of information and services between user(node) and user(node), there must be a mechanism of quantitative settlement. Currency is the tool to settle business in daily life. In Leither network it is the credit that can be exchanged between users.

Every node can grant other nodes credit, which authorize the others to use its services. The difference between services provided and received is the equivalent of a IOU that a borrower gives the lender, or an cryptocurrency issued by the borrower. Cryptocurrency issued by a group is actually a certificate of contribution made by its members. The contribution could be anything valuable to the group, such as resource, money or work done for the group. This certificate can be used to exchange for group services or other value from its members.  

**Object of credit**  
In reality fiat currency is the object of the highest credit rating. In Leither network, services exchanged between nodes, such as bandwidth, storage and CPU, can serve as the object. Their order of importance is bandwidth > storage > CPU. Because of the characteristics of node network, bandwidth is the preferred base unit of settlement between nodes. All the other forms of services can be converted into bandwidth with a certain coefficient. Fiat currency can also be used as object of credit.  

**Individual credit model**  
is a type of User-to-User (P2P) L/C or IOU.  
+ Letter of Credit (L/C)
L/C can be understood as the amount of services allowed to use before payment, whereas IOU is the amount of services that have been used without payment. A L/C includes information of the authorizer, the authorized, unit of credit, maximum credit, signature of authorizer.  
+ IOU and Currency  
An IOU includes information about lender, borrower, unit of credit, amount, signature of borrower. For borrower, IOU is equivalent to issuing **Cryptocurrency**.

**How to buildup credit**  
If there were a relationship between nodes, there should have been certain basis for mutual trust, so that certain amount of credit could be granted according to the magnitude of that relationship. After a relationship between nodes established, credit can be increased periodically. 

Each node is capable of providing useful services. Credit can be acquired by proactively providing valuable service to the others, including data backup, load balance, routing, search, or mediator service. User can also exchange for more credit through 3rd party.

Part II: Organization and consensus
========
Part I describes the key functions, features and important data structures of Leither platform. Part II will explain how to establish highly efficient, translucent and privacy safe block-chain smart contract services.

### I. Organization
In social life, people always belong to some kinds of social groups. **Organization** is a group with a purpose and enhanced functionalities. An organization is first created as a group, then people join the membership. There are three types of roles within an organization: supervisor (founder), member and guest.

The maintenance of a group is a public service, best operated with a consensus mechanism to coordinate the interests of all of its members. A group must become an organization when more potent public service is required of it. Any organizational entity on internet can be reconstructed in the form of a Leither group.
### II. Organization Services
Organization service is a public or private service that organization provides. Most of the businesses currently on internet can be reconstructed as organization service.
#### 2.1 Categories of Internet Business Services
Internet business can be divided into categories of search engine, content, social media and tool.  
+ Search engine  
Google, Amazon, Alibaba and Baidu all profit from **Information Search**.  
In Leither network, user can select its preferred search engine and set its own search strategy. Therefore search service will not be controlled by a few providers, instead user can take control of and customize its own search result.  
+ Social service  
FB, Twitter, Wechat make huge profits by controlling the **Social Relationship** of users.  
In Leither network, data of social connections are saved by user owned devices, therefore all of the attached messages, Apps and services are also selected and customized by user.  
+ Content and Service  
On-line video, music and news media provide **Content and Service**.
In Leither network, content and service can be split into finer parcels, and packaged into MiMei format.  
Content producer can publish directly to its own node, or collect content from regular internet and put it on its node.  
MiMei can flow freely among nodes, users can also customize content search method and resources, not to be subjected to criminal abuse by Gig Techs anymore.    
+ Tool  
Google Doc, online office, XMind make profits  by providing **Tool Services**.  
In Leither, Tool application is also a MiMei. 
#### 2.2 Division of Services  
The internet services mentioned in previous section can all be divided and reconstructed through the coordination of the following entities, in decentralized fashion.  
+ Application Development  
Develop applications for Leither ecosystem, undertaken by programmers. Similar to the role of R&D in an enterprise.
+ Service or Content Production  
Provide services or produce contents for the ecosystem. For content service, similar to editor of a website or operation department of a company.
+ Service Support  
Provide physical devices for the whole ecosystem, run by operators. Different from regular internet, devices are not owned by one company. Similar to the role of maintenance department in a regular company.  
+ Service Promotion  
One method is to promote web service of Leither in regular internet, such as website, link, APP and applet. Another is to tag contents and provide public service from Search Node. Similar to operation department of a company.  
+ Value Appreciation  
Every node can provide public services that can be monetized to pay for other services. Information search and online traffics have great financial value, which is the **foundation of the ecosystem's value**. When packed into a block-chain like system, the investment and return of its members can be coordinated through consensus.
#### 2.3 Service Support
Node management, MiMei migration, information search and group consensus, etc.  
+ MiMei Storage  
Leither decouples internet content and service into MiMei form. MiMei is stored on nodes according to business request and flow among nodes on demand. Service exchanged between nodes is settled with credit. Node participates in the settlement of group services by group consensus.  
+ Index Information  
Index is used for search service. The information is stored on Distributed Hash Table (DHT) of all the online nodes. Node participates in the settlement of group service by group consensus.  
+ Docking Service  
Services that connect MiMei with regular internet, such as APP, DNS and applet. Those services run on servers of regular internet, and participate in the settlement of group services by group consensus.
+ Consensus  
Consensus is the public business logic of a group. Consensus runs on all bookkeepers.
### III. Organizational Functions
#### 3.1 Create Organization  
Any user can create an organization. First, system creates a pair of encryption keys for the organization. The public key ID is the default organization ID. Founder is automatically a member and supervisor by default. Organization has the same authorization scheme as MiMei. User identity includes supervisor, member and guest.  

Organization runs on the node of its creation. Distributed organization also starts from the first node where it is created. The default network structure is DHT network which stores public information of the organization.

Public information of an organization can be inquired through DHT network, so can the information of a member's node. Members can exchange P2P messages. Group message data belongs to the public and is saved on DHT network.
#### 3.2 Join Organization  
User can apply for membership of an organization directly, or ask a member to apply on its behavior. Once supervisor confirms the application, the user becomes a member and DHT is updated with corresponding information. Member can check public information of the organization. User can also perform certain task for the organization and get rewarded a membership.  

#### 3.3 Business Development  
Organization service is actually a series of Leither applications, which can be listed in the public information area of the organization network.  
For example a video website can be separated into the following sub-services:
+ Content collection and production  
+ Content tagging  
+ Application support with storage, bandwidth, CPU, etc.  
+ Content access channel, such as APP, applet, DNS, routing, search, etc.  
+ Promotion and operation.  
+ Value appreciation  
    There two types of value appreciation.  
    Type one: content contains prepaid information of a third party.  
    Type two: content is valuable to average users, who would pay for it or watch the imbedded commercials. This business model is tried and true, a highly matured business practice in regular internet.

### IV. Value System
**Value is created through meaningful work**, which is the basic principle for designing value system.  

Every quantum of token shall base on meaningful work. The quantity of its value is decided by its contribution to the ecosystem, or how useful the work is.  

With consensus, an organizational ledger, which is actually a sparse Merkle tree with snapshot function, can be created. Any incentives issued by the organization will be announced with reason and proof for organization wide publicity.  
This mechanism is called **Proof of Performance (PoP)**, different from POW of BTC.  

System rewards include:  
+ Construction reward:  
Rewards to member who has made fundamental contribution to the organization, such as founder, investor, system and App developer, who did all the absolutely necessary work to make it happen. The reward is recommended by supervisor and issued after vote of consensus.

    The reward is for creating the organization.  

+ Content service reward:  
Organization provides valuable service or content to users. At every step of value creation, contributor is recorded. At the point of value appreciation, smart contract is triggered and incentives rewarded to each member who did valuable work.
    The reward can be distributed by fixed proportion or dynamically balanced through the following mimic process.  
    1. Developer writes content collection plugins.  
    2. Collector use the plugins to gather content from internet, turn it into MiMei, and publish the MiMei onto Leither network. MiMei will indicate collector of its content.  
    3. Operator push the content to users through channels such as App, applet, website, or group message.  
    4. User browse the content and commercials are displayed.  
    5. Checking mechanism is builtin the commercials' code. Once displayed, it will trigger the smart contract, which then transfers tokens from advertiser's account to each participators of the whole process, respectively.  

    This reward mechanism will feedback positively to the healthy development of the organization and the growth of its usefulness.  

+ Mining reward  
Same as traditional mining, all the online nodes can attend vote, bookkeeping and smart contract service. Reward for new block will be used to support the network, manage public information and routing service.  

    This reward makes sure that organization stays active.

+ Operation reward  
In the beginning when an infant system cannot sustain itself. Extra incentives can be rewarded to accelerate its growth.

    This reward helps the organization to mature faster.  

### V. Basic Concepts
**Distributed Hash Table (DHT) network**  
In DHT network, node can be quickly located by its ID to access its information. Public information of an organization can also be saved over the whole network distributively.  

**Sparse Merkle Tree (SMT)**  
Jump-table is a key data structure used in Redis and LevelDB. It introduced the concept of **Equilibrium Probability**, which makes it possible to manipulate a tree without changing its structure substantially. Compared with other algorithms, jump-table has a small impact on performance, but still keeps the computation complexity at the same order of magnitude. Using the same algorithm, Merkle Tree can be improved to create so called **Sparse Merkle Tree**, which is used to hold core information of an organization, similar to the public ledger of a block-chain.

Each leaf node of a SMT stores account information of an organization member. Account number is the member's ID that is the hash of the member's public key. Each ID is 160 bits long. All of the possible IDs can fill up the leaf nodes of a binary tree of height 160. Because the actual number of IDs is far less than the number of possible leaves, the tree is a sparse tree. The root node has two children nodes: 0x0, 0x1, and four grandchildren nodes: 0x00, 0x01, 0x10, 0x11, and such. Leaf node with account information is originally on level 160.  

If a node has only one child, the branch can be shorten by moving its child up one level to replace the parent. This simple optimization can reduce the height of the tree tremendously. For N IDs that are randomly and independently distributed, most of the IDs will be on leaf-node at level log2(n)+1. The closer to the root-node, denser the tree. The addition or removal of a node, or change of account information, only effects its ancestor nodes.  

On SMT, account information on leaf node can be quickly located, which is essential to the implementation of time-space snapshot, network pulse, node grouping, fund transfer and smart contract. Information of each account is on the leaf node of the tree. Each branch uses hashes of its root node's children and synopses to generates its own hash.  

**Node Group**  
With the growth of network and traffic, storage, data processing and communication will eventually overload network nodes. If not optimized, Leither network will fall in to the same conundrum of 7 throughput like BTC.

Nodes on a sub-tree or branch of SMT can be partitioned into a **node group** to solve the above problem. The max number of nodes in a group is 256, and its max height is 64. When the size of a branch exceeds 256, the longest branch at depth 7 will form a new group. The new group ID will be its root node's ID. Every node within the new group share the same 64-bit branch id.

Node group enables two unique advantages. One is elastic concurrency support, the other is an opaque network with variable shades of gray.

**Node Election**  
Every node group has bookkeepers, which is responsible to record, check and verify transactions of the group, and also broadcast and communicate related information. Account within a group can vote for bookkeepers. The accumulated vote cannot exceed its account balance.

The rules of election is part of the group consensus, which is designed by the organizer, who will take factors, such as the amount of asset pledged and network speed, into consideration. A bookkeeper and backup bookkeeper will be elected. The bookkeeper is in charge of updating the general ledger and the backup bookkeeper verifies it.

Node group can vote as a member in its parent group, to elect bookkeeper on higher hierarchy. The bookkeepers at the lowest level deal with detailed business transactions. In the middle and upper levels they combine changes of information from groups below, generate synopses and update the overall account balance. If the number of online nodes is small or network load is light, a node group does not have to maintain a bookkeeper. The bookkeeping can be entrusted to upper level bookkeepers.

Top level information that is close to the root node is synchronized network wide, which describe the status quo of the whole network.  

**Network Pulse**  
In order to coordinate all nodes among the network, an incremental sequence number is broadcast top down in every **network pulse cycle** (1s by default). The sequence number is generated by top level bookkeepers through negotiation.  

The sequence number serves as sync timer of each node, and version number of its data. Within each pulse cycle, every node consolidates and saves the newly added data from last cycle. Every node group is responsible to process data of several levels, depending on the height of the SMT. The work done is called **Proof of Performance**, similar to PoW in BTC, and will be rewarded by the system.  

**Time-space Snapshot**  
LevelDB is an excellent database optimized for writing operation. The writing of new data using writeStream can reach the speed limit of storage media, even faster than database reading. On the other hand, sequence number in key-value provides revised version number, with which historical record can be quickly inquired. 

Network pulse is equivalent of the sequence version number in LevelDB, with which changed data within each pulse cycle can be quickly recorded and labelled. The same method can be used to take snapshot of time-space SMT.

In every pulse cycle, the structure of the current SMT is identical to its previous version, except the newly added changed data. When taking a snapshot, only the changed information need to be saved. With the version sequence number of key-value, data of any node in any pulse cycle can be quickly retrieved. Regular account only has to record its own account information and information of its node group. 

The underlying MiMei database of Leither has built in support to time-space snapshot.

### VI. Related Procedures
#### 6.1 Network Construction  
After a user joins an organization, it joins the DHT network and becomes a member. The newcomer writes the routing information of its node into DHT, so that other nodes can find it and use its services.  
#### 6.2 Ledger Construction  
Ledger is a SMT with snapshot function. In order to keep network wide messages in sync, the whole network must share a time sequence variable, aka **Network Pulse**. Its initial value is 0. If system state changes, time sequence increments. Other wise, it stays put.  

The system periodically elects a few **bookkeepers** that backups each other. The first one is chief bookkeeper. The term of a bookkeeper is one **Election Cycle** (30min by default). The bookkeeper is responsible for bookkeeping of its branch.   
#### 6.3 Fund Transfer Procedure  
**Distribution by organization**  
The very first token distribution will be announced network wide, with reasons for scrutiny by the members. However, afterward the transfer of tokens between users is private, or publicized only to relevant nodes.  

**Transfer between users**  
The transaction between two users concerns only themselves. After the transaction is confirmed, information signed by both users will be sent to their bookkeepers on the SMT respectively. Each bookkeeper will record the changes in its own branch and broadcast the information among backup bookkeepers at the same level. After time sequence increments (in less than 1s), transaction data becomes read only. Bookkeepers at each level begin to check their branches bottom up and summarize branch information to generate synopses. The top bookkeeper summarizes overall information, generates synopsis of the SMT, and broadcasts top down to everyone below.

Both parties of the transaction record time sequence, synopsis of each level and its own account information, and finally transaction is committed for good. One transaction spends at most two pulse cycles (2s).

**Dispute Resolution Procedure**  
Every transaction must have sufficient security deposit to endorse enough time for more nodes to verify it. Transaction is processed by multiple nodes simultaneously, and bookkeeper and backup bookkeeper are randomly selected to avoid collusion. If any node disputes the transaction, disputation resolution procedure kicks in. During the procedure, all relevant funds are frozen.

All of the nodes can check the disputed transaction and vote. Security deposit of the erroneous node will be confiscated. Bookkeeper can only process the transaction amount that the frozen deposit can cover.

**Legitimacy Check**  
    1. The trading parties check the legitimacy of the transaction.   
    2. The bookkeeper checks if the general ledger is in balance.   
    3. The branch bookkeeper verifies the legitimacy of the transaction on its branch.   
    4. During synchronization, all leaf nodes check the balance and sanity of nearby branches' accounts.  

**Redundant Backup**  
All the account information on a branch are stored by a few bookkeepers. In order to keep the network robust, all nodes are encouraged to redundantly backup account information of its neighboring branches.

There are two ways for reward. Regualr nodes can audit the transaction information of nearby branches and earn reward. For illegal transaction, account information of neighboring nodes are frozen until the state of their branches recuperate. Number of neighboring nodes could be 1, 3, 7, 15, 31, 63, 255.  

All of the nodes will strive to maintain the health of the network, in order to earn reward and avoid collateral damage.  
Because of the redundant backup, a small number of nodes can recover the whole network after a crash.  
Redundant backup happens when bookkeepers broadcast information after finishing their tasks.  
Network recovery happens at information verification when a node getting online.  

#### 6.4 Dispute Handling**  
**Publicity Period**  
Every transaction must have a period of publicity, 24hrs by default. If other nodes find any problem, the transaction will be reported and disputation resolution procedure kicks in.  

**Amount-time**  
Every transaction involves different amount of fund. Each bookkeeper or auditor pledges different amount of credit. Therefore a new concept _Amount-time_ is coined for the convenience of calculation.  
amount-time = Amount of fund * Time during which the fund is occupied  
For transaction,   
Amount = transaction amount, time = period of publicity  
For auditing and bookkeeping,  
Amount = amount of credit pledged, time = period the fund frozen  

**Accumulated Period**  
= accumulated amount-time (for auditing or bookkeeping) / transaction amount  

**Transaction Security Threshold**  
+ Minimum number of users who audit a transaction  
+ Minimum accumulated audit period  
    Security threshold sets the minimum auditing cost, so the cost of mining and service fee can be estimated. Accumulated audit period of a normal transaction will be 1.  
    Different audit period has different value, therefore different system reward.  

#### 6.5 Smart Contract
Smart contract is similar to transfer procedure. A contract includes smart code, signature of consignor, upper limit of contract. The contract will be executed by the bookkeepers in upper level by default. User can also appoint a consignee node. For complicated contract, if unforeseeable risk is too high, the security deposit of nodes executing and checking the contract shall be raised to cover the risk. It is like a third party that endorses the contract with its own credit and earns corresponding reward.

When dispute happens, the whole network enters dispute handling procedure. The nodes that vote must freeze some of its own asset. Afterward the losing side will lost its frozen asset.  

#### 6.6 Account Inquiry  
Information of system account is public, but regular user account is private and cannot be inquired by 3rd party. However authenticity of account information that user chooses to demonstrate can be verified by checking the synopsis recorded in public ledger. The truthfulness of synopsis can be confirmed by checking synopsis of child branches.  
### VII. Privacy and Transaction Security
#### 7.1 Privacy and Consensus  
**Dilemma of Transparency**  
In traditional block-chain, consensus is based on transparency. If account information is hidden, there is no way to confirm the legitimacy of a transaction. The buyer might overspend.  

For parties not involved in a transaction, the purpose of knowledge of the transaction is more likely to avoid potential harm of ignorance, rather than voyeurism. If a proper balance of translucency can be reached, there is no need for the others to know the detail of a deal.  
**Premise and Principles**  
The premise and principles for anonymity is that trading parties in a deal and the third parties do not harm each other.  
#### 7.2 Transaction Security  
Based on not-be-a-sucker principle, it is necessary to make every member of an organization believes a deal is legal.  
#### 7.2.1 Safety Conditions  
Legal transaction must meet the following conditions.  
+ Truthful intention  
The deal is confirmed by both parties.
+ Real deal  
The seller has sufficient goods to trade.  
Fund transferred cannot be negative.  
No negative account after the transaction.  
+ Balanced Account  
Without system reward, the overall account balance does not change after a transaction.  
If system reward is counted as a transfer from organization to users, total account balance still does not change.  
#### 7.2.2 Safety Analysis  
+ Traders' responsibility   
The seller will strictly check the safety conditions above, otherwise it will not receive the fund.  
The buyer will also double check the parameters and procedure of the transaction, for illegal deal will be punished by the organization.  
+ Other users within the same group  
The other users need to check if there is any fake transaction or account balance issue of the trading accounts, in order to protect themselves from collateral damage.  
+ Other node groups  
Any node group that does not include a trading user only need to verify there is no fake deal or account balance issue within the node group that does include a trading user.  

#### 7.3 Anonymity Strategy  
**Node Group Opacity**  
If a node group is a unit of settlement, transaction information can be hidden within it. Any other user only need to worry about the general ledger of the group and its legitimacy.  
**Third Party Opacity**  
If there is a third party endorsement to ensure no collateral damage to the others irrelevant to the deal, transaction information can be revealed in delayed time.  
**Translucent Network**  
During redundant backup, information of an account is backed up only by a few accounts nearby. Accounts faraway can only acquire summarized information through a few top level nodes, without knowing any details of the lower level nodes. This strategy not only guarantees the transaction information safely confirmable, but also the anonymity of information.  
**Capital flow unimportant**  
In reality, the anonymity of capital flow is more important than account balance. However in Leither network 3rd party will be more interested in account balance and legitimacy, instead of capital flow. If system purges transaction information periodically, after a while capital flow will not be traceable.
#### 7.4 Anonymity Strategy  
The following methods implement information anonymity.
**Credit Endorsement**  
A third party, bookkeeper by default, can be appointed to endorse a transaction. For other users, the transaction runs on branch and is pledged by the endorsor, so they are collateral damage free. 
**Neighbor Check**  
After a transaction is committed and SMT changed, bookkeeper must broadcast the information. The neighbors of the trading nodes will get details because they provide redundant backup, so they must verify the transaction and check the SMT. If problem detected, dispute procedure kicks in. Except a few bookkeepers and neighbors, most of the users will not receive transaction details.  
**Multiple transfers**  
If someone wants to know the details of a transaction, it needs to access the trading nodes during the transaction. It will cost time and money. User can choose to transfer fund multiple hops through different branches, which make the tracking of a deal virtually impossible.  
**Transfer across organizations**  
Every organization has its own consensus system. Different organizations can exchange their tokens. Transfer across organizations will also increase the cost of tracing exponentially.

### VIII. Data Processing and Quantitative Analysis
**Data Storage**  
In theory Sparse Merkle Tree is an equilibrium probability binary tree. The closer to top, the evener the distribution. Close to the bottom, the length of branches will vary substantially. It is necessary to reduce the disparity. On the other hand, a 160-bit account number is inconvenient to process. Since one byte can differentiate 8 levels of the tree, 8 levels are stratified together. Assume tree node information is 36 bytes (network number 8 bytes + amount 8 bytes + account number 20 bytes ).  
**Node Group**  
One group is a sub-tree that can hold up to 256 nodes, with max height of 64. The max storage space of one group is 10KB.
**Sub-Database**  
One sub-database can hold up to 256 node groups and requires a maximum of 2.5MB to store that 65536 node accounts.  
**Miner's bookkeeping database**  
If there are more than 256 node groups in database on one node(user), system will splits some data to other nodes. A bookkeeper can store max 256 sub-databases, 16,777,216 node accounts, which requires 640MB space.  
**Account summary information**  
The topmost node group records general ledger status, including account information of the topmost 256 branches, such as synopses and account balance. The size of all these information is less than 10KB. It must be synced to every online node in the system. A normal user only need to know the general network status and information of its own branch.  

**Quantitative Analysis**  
Usee BTC as reference, currently there are about 30 million accounts and the average number of transactions is 6.67 per second.
Assume the following system parameters:
+ Number of accounts 2^32 = 4.3 billions  
+ Transaction concurrency 1 million/s  
+ Pulse cycle 1 second  
+ Election cycle 30 minutes  

**Scale Analysis**  
Most of the accounts are distributed in a 32-level network. Bookkeepers are grouped to cover 8, 8 and 16 levels respectively.  
**Miner HD usage**  
If one branch miner process one sub-database (65536 accounts), data size is 2.5MB. If it is one database (16,777,216 accounts), data size will be 640MB.  
**Miner memory usage**  
Only online nodes need to be processed by a miner.  
For Level 1 or 2 miners, most branches need to be processed online. All 10KB data must be in memory.  
For lower level miners, there are only 30 some transactions per second. No need to keep everything in memory, read related data and synopsis of neighboring branches when necessary. The data size is 30* 16*2*28 byte = 7.5KB, and 90 times of database access.  
**Miner CPU usage**  
CPU usage is mainly for verifying transactions and generating synopsis.  
Level 1 or 2 miners must calculate for every transaction, which means 16 million synopses. However in reality at most only one iteration per pulse cycle is necessary for the whole network, which is 512 times of synopsis calculation, about 10KB data.    
Branch root node need to verify the whole branch, which requires 30*(16+16)=7KB data, and 96 synopses.  
**Traffic load among miners**  
The search for a node is load-balanced on each node within a DHT network, therefore the amount of traffic is negligible.    
The network structure of ledger is predetermined within each election cycle. Traffic load in fixed network is negligible.  
Transaction between users is handled by the branch of each user. Traffic load of average 30 transactions is small.  

Traffic load of root node need to be examined closely below.  
Root node must broadcast to backup miners and lower level miners. The data to backup miner is 10KB/s, so the maximum data rate to 256 miners simultaneously is 2.5MB/s. If the bandwidth is insufficient, miners can be further divided into subgroup of 16 miners per level. Data rate becomes 160KB/s. Number of broadcasts from top to bottom increases from 3 times to 5.  

**Communication time among miners**  
The process of data involves 3 to 4 levels of nodes.  
The submission and verification of a transaction happen in two different pulse cycles.  
Submission is delivered to branch root node, so times of communication is one.  
Verification request is delivered to root node from branch node, and the result propagates once or twice back to branch node.  
Most of the operations can be done in less than 10 hops of communication. Normally it takes less than 2 pulse cycles.

Based on the above analysis, with equipment and bandwidth of home network, it is completely feasible to establish decentralized network of 4.3 billion users and support 1 million/s concurrency.  

### IV. Structural Summary
The overall plan of this paper can be summarized in the following key points.
#### 9.1 Credit is the result of meaningful work.
Value is created by work that is meaningful to others. System created the original credit system with meaningful work. Every quantum of credit corresponds to a piece of work result, so credit can be used to exchange for product and service.  

Comparatively, the meaningful work in block-chain is just bookkeeping, which is at lower end in the chain of value.  
#### 9.2 Interaction is based on credit
Any online transaction can be endorsed with credit. The encryption system is a mechanism to prove user authentication and information. Miner pledges its credit to ensure the others that it will do a good job, otherwise be punished accordingly.  

With security deposit and arbitration mechanism, complicated smart contract can be implemented to support large scale interactions among multiple parties. All of these operations rely on the most basic PKI algorithm, which even the cheapest devices are powerful enough to support.  

Comparatively, traditional block-chain depends more and more on dedicated devices and squander expensive electricity.  
#### 9.3 The completeness of time-space comes from synopsis and time sequence
Leither system defines a 160-bit account space to accommodate all possible accounts on a sparse Merkle tree.  
All accounts information are stored on leaf nodes, and synopses are generated bottom up at each level to the root node. The state of the whole system can be represented by one synopsis.  
Time sequence is used to coordinate all of the nodes, and record state of time-space at different stages. Time sequence also solved CAP problem in distributed network.  
#### 9.4 Division of work and coordination improve processing power as a whole
Network nodes can be partitioned into layers according to the scale of the network. Transaction data are saved on user nodes. Consensus is confirmed by branch nodes. Data that must be synced network wide is tiny amount of synopses and summaries.  

Network can be more stable and robust by adding proper number of backup nodes. Comparatively, in regular block-chain all of the bookkeeping is accomplished by all of the nodes together, which greatly restricted network scale and throughput.  
#### 9.5 Endorsement and shade of grey solves privacy problem
### X. Advantages of Leither compared with regular block-chain
+ Functionality
With MiMei and Leither application system, most regular internet platforms can be reconstructed as user oriented network.  
With organization and consensus, Leither can implement most of the ethereum functionalities and organize all the participants.  
The above functions combined can give internet back to its users from the grip of oligarchies.  
+ Value  
Traditional block-chain does bookkeeping for others.  
Every organization of Leither has a mission to decentralize an internet application, to give user control of its data, to satisfy user's need first and foremost. The action of every participant has real value, in order to take back the outrageous profit of oligarchies and redistribute it users who did meaningful work.  
+ Cost  
Mining of regular block-chain is hugely expensive because of all the machines and electricity cost to do the most basic work.  
In Leither, consensus is achieved by doing the very job of internet. Tv box, router, NAS, smart phone, even Raspberry zero worth $10 can provide valuable services and ROI much higher than mining machine.  
+ Efficiency  
Leither elastic system can scale according to business request. Workload is balanced on to every node through consensus.  
The heavier the network load, the more the nodes, and the more powerful the network processing capacity that far exceeds that of regular block-chain. Because of its high efficiency, a few nodes can construct their own business and consensus mechanism.  
+ Privacy  
In regular block-chain, information flow of every account is an open book. Even though account is anonymous, it is possible to identify the owner by analyzing information flow.  
In Leither, except the earliest contribution that need to be publicized for auditing, all the transactions afterward are open only to related parties and a few nodes. The whole network has a shade of grey. Creator can set translucency to protect privacy.  
+ Law Abiding  
A token of Leither has corresponding service as anchor object. It is more like points, notes, share, Q coin, and many other business practices in real life. It is perfectly legal, as a matter of fact very similar to early Taobao business, completely decentralized and conducted in the name of individual.

    Smart contract deals with real business data, therefore avoids the falling hole of ethereum which can only process virtual business such as encryption coin.  
