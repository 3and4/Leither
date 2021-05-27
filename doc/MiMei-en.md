MIMEI
======
### Summary
An alarming trend of internet in the past few decades is the fast concentration of traffics on fewer apps or sites owned by tech giants, who used their dominant control of data to influence, even manipulate user behavior for monopolistic profits, at the expense of average users. However, we hold the following truth to be self-evident, that a free and open internet is of the best interest to its users.

Leither, a decentralized Cloud OS, is designed to take the challenge by decoupling and reconstructing most of the core services of regular internet in a decentralized fashion, therefore give the control of data back to user, its real owner. Leither OS has implemented all the subsystems necessary to reconstruct regular internet services, such as authentication and authorization, file system, database, application framework, domain name service and routing.

To reconstruct a regular internet application, for example Youtube or Google, simply creates a Leither Organization and a Consensus Contract to stipulate how to distribute tokens among participants. Token is a proof of contribution that garantees proportional dividend of the organizaiton's income, or exchanges for services such as traffics, routing, domian name resolution, search, backup, bandwidth, CPU power, contract notary, or mining pledge.

Members of an organization include initiater, investor, developer, content or service provider, device owner, operation and marketing, etc. The distribution of income is garanteed by Consensus. Device owner does the work of miner in BTC. Devices keep the books as mining machines in BTC, at the same time provide more valuable business services of the organization. Compared with POW-based block-chain Leither block-chain has advantages in every aspect of cost, efficiency, performance and privacy, and last but not least, law abiding.

This paper is divided into 2 parts:  
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

    MEMEI is the carrier of information on internet. It defines the structure of data storage, and all kinds of operations by associated applications, such as create, open, save, read, search, render and user-interact. 

    Everything in Leither-based ecosystem is defined in term of MEMEI. The referential relationships among MEMEI makes it possible to develop comprehensive internet applications and services.

+ **Leither**  
is a container of MEMEI, an ultra-light decentralized cloud OS.

    Leither is designed to support frequently used internet applications with extremely limited resources, for example to run office applications for a small business on a Raspberry Pi Zero.

    Leither OS has implemented all the subsystems necessary to reconstruct regular internet services, including authentication and authorization, file system, database, application framework, independent domain name service and routing. Leither API supports about 40 kinds of development languages.

    PS: Leither writtten in Kanji character is 弥媒. 弥 means extremely small and ubiquitous, usually something imperceptable. 媒 means media. 弥媒 put togather almost as good as Ether, a substance once believed to be the building block of universe.

+ **Node**  
A smart device running Leither OS is a Node. As small as 6MB in size and runs on all regular operation systems, Leither supports almost any smart devices, like PC, NAS, router and Raspberry Pi.  
1. One independent node can serve as an ultralight server;
2. Two nodes can backup each other in services like data backup, error tolerance, load balance, domain name resolution and routing;
3. Multiple nodes can network to support decentralized internet applications.

    MiMei flows among nodes to facilitate the decentralization of centralized internet.

### I. Background -- Centralized Internet
#### 1.1 Big Data has conquered the world
Internet has changed our daily life. Big Data is running internet, therefore our lives. Online retail up to US$35 trillion in 2019 in China alone. The essence of Big Data is the accumulation of data and formation of rules, for the purpose of manipulating user behavior and generating monopolistic profit.

#### 1.2 Big Techs own Big Data
Big Techs own core data. For example, FB owns social media and Amazon online retail. Data protocols are manipulated to serve as business barriers. Google’s spider cannot access FB pages, vice versa. Monopoly is protected by corporation regulations, laws and ecosystem rules. Big Techs also take the lion’s share of profit, so their fortune concentrates acceleratingly.

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
|1|My settings|Peronal information, settings and preferences|
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

In this case, User has the freedom of choice and is not restricted by a single App.  
**App->Data**
The right of user accessing data is restricted to certain Apps, usually only one. For example, when user login an e-commerce site do shopping, or get on social media to chat with friends.

In this case, the algorithm decides the path of user accessing data, and user behavior is controlled by the App.

In programming, the above two approaches corresponds to **Subjest Oriented Programming** and **Procedure Oriented Programming**, respectively.
#### 3.2 How to break data monopoly
**Opensource Platform**  
Microsoft took control of the PC ecosystem with DOS and Windows, and strangled competition in both office software and browser. Google uses search engine to control data and influences user behavior. FB does the same with social media.

On the other hand, opensource community turned operation system into a public platform with Linux. It is the first attempt to opensource a platform. Similar cases are Android for smartphone and RISC-V for IC instruction set.

There shall be similar opensource platform for collecting user data, especially for **Social Media** and **Search Engine**.  
**Free Data Migration**  
In early operation system, data file can be copied and transferred among different mainframes of different systems. Free movement of data shall be a consensus and primary design norm of operation system.  
Three measures can be adopted to achieve the above goal:
+ Promote data-freedom as a mainstream ideology
+ Legislate that any platform that collects user data must allow the owner to migrate its personal data.
+ Develop opensource program, such as browser plugins, to facilitate user to migrate data to personal storage devices. Once the ownership of user data taken back by its creator, the data can be revived and reconnected with each other.

**Seperation of Data and Rule**  
If data can be migrated freely, user will have the choice of rules that can be applied to it. In another word, try to be **Object Oriented**, avoid **Procedure Oriented** development whenever possible. Data structure is relatively simple to be reverse engineered. It is easier to reconstruct program structrue with the knowledge of data structure.  
**Granulation of User Data**  
There is no way to decouple convoluted and complicated dataset. It is best to granulate data in design, following the principle of structure of gene, keep the information relatively comprehensive and independent. A piece of granulated information is easier to be applied to certain algorithm independently, and its connection with other objects clearly defined.  
**Connections of Data**  
Isolated data have little value compared with inter-connected dataset, whose connections contain much more information. HTML is perfect protocol to describe relationships, if not link could be unstable.  
#### 3.3 Data Container  
Large number of data centers are indispensable for regular internet to process data of billions of users. A large data center can have tens of thousands of servers. Every task will be processed by groups of servers using certain technique. A army of maintenance personals are required for daily operation.

No data center is required when storing personal data by individual. However, some kind of container, a light-weight clouds OS, will be handy, like Pod promoted by Solid platform. The container shall be low cost, lightweight, and easy to operate by layman.Compared with the Homomorphic Encryption that certain block-chain alogrithm is using, Leither method is simpler and more efficient. One who owns the storage device owns the data and defines the rules to apply.
#### 3.4 Data Model for Reference  
**File**  
File is a set of program and data stored on media. File is constructed according to certain format, which is usually referred to as data structure. Operation system manages files with File System in a tree-like data structure. File capsulates the access of data on media and optimizes the procedure.  
**WWW (World Wide Web)**  
In a standard webpage, HTML5 is used to render content of webpage, CSS style and JS tasks. A page can relatively store and render its data comprehensively. Through interconnections among pages, large scale and complicated information can be represented.

Shortcomings:
Webpage only contains information about frontend, without backend data. Big Techs take monopolistic control of user data by owning the backend servers. The connections among webpages are link based which is unstable for lacking definition about server or user.
**Gene**  
The basic unit that contains information of life. DNA or RNA stores a complete segment of hereditary material, which records the adaptation of life to its environment. Life copies hereditary information from generation to generation, while adapting to new environment by variation.

Limitation: Every iteration takes a generation, with restrictions on the amount of information inherited and weak capability to deal with complicated changes.  
**MEME**  
Richard Dawkins introduced the concept of MEME in his original book, the Selfish Gene. Similar to gene, MEME is defined as the basic unit contains information of a civilization. MEME can be used to record all the knowledge of human society. It can be remembered in human brain, broadcasted among people, or stroed in any kind of media.

Human brain uses pattern recognition to model the world. During its interaction with the surroundings, every reflection by the mind is an iteration of evolution. Compared with gene, MEME can evolute faster than gene by orders of magnitude.  
### IV. Introduction to MiMei
**Origin of MiMei**  
Entropy is one of the most important measurements in the universe and human society. The 2nd law of thermodynamics points the direction along which the universe evolves. That is why entropy is also called Arrow of Time. Erwin Schrodinger in his book, What is Life, declares that what an organism feeds upon is negative entropy.
