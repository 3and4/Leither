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

    Leither OS has implemented all the subsystems necessary to reconstruct regular internet services, including authentication and authorization, file system, database, application framework, domain name service and routing. Leither API supports about 40 kinds of development languages.

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