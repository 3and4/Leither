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