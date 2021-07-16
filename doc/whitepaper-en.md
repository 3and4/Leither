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
