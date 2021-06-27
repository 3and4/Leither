Part II Consensus Mechanism
========
### Summary
### I. Organization
Chapter Nine of _MiMei and Application System_ mentioned individual oriented relationship model, and credit model based on relationship network. In social life, people always belong to all kinds of social groups, which is called **Organization** in this paper.  

Organization is a group oriented structure. A group is created first, then people join it to become a member. There are three types of roles within an organization: supervisor (founder), member and guest. Supervisor and member work with each other to maintain and manage the group. In traditional internet, both structures are maintained by platform operator, who also set the rules.

The maintenance of a group is a public service best operated with a block-chain like solution to coordinate the interests of all of its members. A group must become an organization when more potent public service is required of it. Any organizational entity on internet can be reconstructed in the form of a Leither group.
### II. Organization Services
Organization service is a public or private service that organization provides. Most of the businesses currently on internet can be reconstructed as organization service.
#### 2.1 Categories of Internet Business Services
Internet business can be divided into categories of search engine, content, social media and tool.  
+ Search engine  
Google, Amazon, Alibaba and Baidu all profit from **Information Search**.  
In Leither network, user can select its preferred search engines, set its own search strategy. Therefore search service will not be controlled by a few providers, instead user can take control of and decide its own search information.  
+ Social service  
FB, Twitter, Wechat make huge profits by controlling the **Social Relationship** of users.  
In Leither network, data of social connections are saved by user owned devices, therefore all of the attached messages, Apps and services are also selected and customized by user.  
+ Content and Service  
On-line video, music and news media provide **Content and Service**.
In Leither network, content and service can be split into finer parcels, and packaged into MiMei format.  
Content producer can publish directly to its own node, or collect content from regular internet and put on its node.  
MiMei can flow freely among nodes, users can also customize content search method and resources, not to be subjected to criminal abuse by Gig Techs anymore.    
+ Tool  
Google Doc, online office, XMind make profits  by providing **Tool Services**.  
In Leither, Tool application is also a MiMei. 

#### 2.2 Division of Services  
The internet services mentioned in previous section can all be divided and reconstructed through the coordination of the following entities, in decentralized fashion.  
+ Application Development  
Develop application and platform for Leither ecosystem, undertaken by programmers. A similar role of the R&D in an enterprise.
+ Service or Content Production  
Provide service or produce content for the ecosystem. For content service, similar to editor of a website or operation department of a company.
+ Service Support  
Provide physical devices for the whole ecosystem, run by operators. Different from regular internet, devices are not owned by one company. Similar to the role of maintenance department in a regular company.  
+ Service Promotion  
One method is to promote web service of Leither, such as website, link, app and applet, in regular internet. Another is to generate tags of content and provide public service from Search Node. Similar to operation department of a company.  
+ Value Appreciation  
Every node can provide services, which can be monetized and pay for other services. Information search and online traffics have great financial value, which is the **foundation of the ecosystem's value**.  
Usually a block-chain like system can coordinate the investment and return of its members. More details in later chapters.  

#### 2.3 Service Support
Node management, MiMei migration, information search and group consensus, etc.  
+ MiMei Storage  
Leither decouples internet content and service into MiMei form. MiMei is stored on nodes according to business request. MiMei flow among nodes on demand. Service exchanged between nodes is settled with credit. Node participates in the settlement of group services by group consensus.  
+ Index Information  
Index is used for search service. The information is stored on Distributed Hash Table (DHT) of all the online nodes. Node participates in the settlement of group service by group consensus.  
+ Docking Service  
Services that connect MiMei with regular internet, such as APP, DNS and applet. Those services run on servers of regular internet, and participate in the settlement of group services by group consensus.
+ Consensus  
Public business logic of a group. Consensus runs on all the nodes.

### III. Organizational Functions
#### 3.1 Create Organization  
Any user can create an organization. First, system creates a pair of encryption keys for the organization. The public key ID is the default organization ID. Founder is automatically a member and supervisor by default. Organization has the same authorization scheme as MiMei. User identity includes supervisor, member and guest.  

Organization runs on the node of its creation. Distributed organization also starts from the first node where it is created. The default network structure is DHT network which stores public information of the organization.

Public information of the organization can be inquired through DHT network. Information of member's node can be searched through routing service. Members can send P2P message. Group chatting data belongs to public and saved on DHT network.
#### 3.2 Join Organization  
User can apply for membership of an organization directly, or ask a member to apply on its behavior. Once supervisor confirms the application, the user becomes a member and DHT is updated with corresponding information. Member can check public information of the organization. User can also perform certain task for the organization and get rewarded a membership.  

#### 3.3 Business Development  
Organization service is actually a series of Leither applications, which can be published in the public information area of the organization network.  
For example a video website can be separated into the following sub-services:
+ Content collection and production  
+ Content tagging  
+ Application support with storage, bandwidth, CPU, etc.  
+ Content access channel, such as APP, applet, DNS, routing, search, etc.  
+ Promotion to send content to people who need it.  
+ Value appreciation  
    There two types of value appreciation.  
    Type one: content contains prepaid information of a third party.  
    Type two: content is valuable to average users, who would pay for it or watch the imbedded commercials. This business model is tried and true, a highly matured business practice in regular internet.

### IV. Valuation-based System
**Value is created through meaningful work**. It is the basic principle for designing value system.  

Every quantum of token shall be based on meaningful work. The quantity of its value is based on its contribution to the ecosystem, or how useful the work is.  

With consensus, an organizational ledger, which is actually a sparse Merkle tree with snapshot function, can be created. Any incentives issued by the organization will be announced with reason and proof for organization wide publicity.  
This mechanism is called **Proof of Performance (PoP)**, different from POW of BTC.  

System rewards include:  
+ Foundation reward:  
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
Same as traditional mining, all the online nodes can vote in bookkeeping and smart contract. Reward for new block will be used to support the network, manage public information and routing service.  

    The reward makes sure that organization stays active.

+ Operation reward  
In the beginning when the system cannot sustain itself with insufficient content and service. Extra incentives can be rewarded to accelerate the growth of an infant system.

    The reward helps the organization to mature faster.  

### V. Basic Concept
**Distributed Hash Table (DHT) network**  
In DHT network, node can be quickly searched by its ID to access its information. Public information of an organization can also be saved over the whole network distributively using DHT.  

**Sparse Merkle Tree (SMT)**  
Jump-table is a key data structure used in Redis and LevelDB. It introduced the concept of **Equilibrium Probability**, which makes it possible to manipulate a tree without changing its structure substantially. Compared with other algorithms, jump-table has a small impact on performance, but still keeps the computation complexity at the same order of magnitude. Using the same algorithm, Merkle Tree can be improved to create so called Sparse Merkle Tree, which is used to hold core information of an organization, similar to the public ledger of a block-chain.

Each leaf node of a SMT stores account information of an organization member. Account number is the member's ID that is the hash of the member's public key. Each ID is 160 bit long. All of the possible IDs can fill up the leaf nodes of a binary tree of height 160. Because the actual number of IDs is far less than the number of possible leaves, the tree is a sparse tree. The root node has two children nodes: 0x0m 0x1, and four grandchildren nodes: 0x00, 0x01, 0x10, 0x11, and such. Leaf node with account information is originally on level 160.  

If a tree-node has only one child, the branch can be shorten by moving the child node up one level to replace its parent. This simple optimization can reduce the height of the tree tremendously. For N IDs that are randomly and independently distributed, most of the IDs will be on leaf-node at level log2(n)+1. The closer to the root-node, denser the tree. The addition or removal of a tree-node, or change of account information, only effects the branch where the node is on.  

On SMT, account information on leaf node can be quickly located, which is essential to the implementation of time-space snapshot, network pulse, node grouping, fund transfer and smart contract. Information of each account is on the leaf node of the tree. Each branch uses hashes of its root node's children and account sum to generates its own hash.  

**Node Group**  
With the growth of network and traffic, storage, data processing and communication will eventually overload network nodes. If not optimized, Leither network will fall in to the same conundrum of 7 throughput like BTC.

Neighboring nodes on a sub-tree or branch of SMT can be partitioned into a **node group** to solve the above problem. The max number of nodes in one group is 256, and the max height is 64.  

Node group enables two unique advantages. One is elastic concurrency support, the other is an opaque network with variable shades of gray.

**Node Election**  
Every node group has bookkeepers, which is responsible to record, check and verify transactions of the group, and also broadcast and communicate related information. Account within a group can vote for bookkeepers. The accumulated vote cannot exceed its account balance.

The rules of election is part of the group consensus, which is designed by the organizer, who will take factors, such as the amount of asset pledged and network speed, into consideration. A bookkeeper and backup bookkeeper will be elected. The bookkeeper is in charge of updating the general ledger and the backup bookkeeper verifies it.

Node group can vote as a member in its parent group, to elect bookkeeper on higher hierarchy. The bookkeepers at the lowest level deal with detailed business transactions. In the middle and upper level they combine changes of information from groups below, generate synopsis and update the overall account balance.

Top level information that is close to the root node is synchronized network wide, which describe the status quo of the whole network.  

**Network Pulse**  
In order to coordinate all nodes among the network, bookkeepers at the root generate an incremental sequence number and broadcast it top down in every **network pulse cycle**. A pulse cycle is one second by default.  

The sequence number serves as sync timer of each node, and version number of its data. Within each pulse cycle, every node consolidates and saves the newly added data from last cycle. Every node group is responsible to process data of several levels, depending on the height of the SMT. The work done is called **Proof of Performance**, similar to PoW in BTC, and will be rewarded by the system.  

**Time-space Snapshot**  
LevelDB is an excellent database optimized for writing operation. The writing of new data using writeStream can reach the speed limit of storage media, even faster than database reading. On the other hand, sequence number of key-value provides revised version number, with which historical record can be quickly inquired. 

Network pulse is equivalent of the sequence version number in LevelDB, with which changed data within each pulse cycle can be quickly recorded and labelled. The same method can be used to take snapshot of time-space SMT.

In every pulse cycle, the structure of the current SMT is identical to its previous version, except the newly added changed data. When taking a snapshot, only the changed information need to be saved. With the version sequence number of key-value, data of any node in any pulse cycle can be quickly retrieved. Regular node only has to record its own account information and information of its node group. 

The underlying MiMei database of Leither has built in support to time-space snapshot.

### VI. Related Procedures
#### 6.1 Network Construction  
After a user joins an organization, it joins the DHT network and becomes a member. The newcomer writes the routing information of its node into DHT, so that other nodes can find it and use its services.  
#### 6.2 Ledger Construction  
Ledger is a SMT with snapshot function. In order to keep network wide messages in sync, the whole network must share a time sequence variable, aka **Time Sequence**. Its initial value is 0, system generates one pulse periodically (1s by default). If system state changes, time sequence increments. Other wise, it stays put.  

The system periodically elects a few **bookkeepers** that backups each other. The first one is chief bookkeeper. The term of a bookkeeper is one **Election Cycle** (30min by default).    
#### 6.3 Fund Transfer Procedure  
**Distribution by organization**  
The very first token distribution will be announced network wide, with reasons for scrutiny by the members. However, afterward the transfer of tokens between users is private, or publicized only to relevant nodes.  
**Transfer between users**  
The transaction between two users concerns only themselves. After the transaction is confirmed, information signed by both users will be sent to their parent nodes on the SMT respectively. Each parent node will record the changes in its own branch and broadcast the information among backup bookkeepers at the same level. After time sequence increments (in less than 1s), transaction data becomes read only. Bookkeepers begin to check the branch bottom up and summarize branch information to generate synopsis. The chief bookkeeper summarizes overall information, generates synopsis of the SMT, and broadcasts top down to everyone below.

Both parties of the transaction record time sequence, synopsis of each level and its own account information, and finally transaction is committed for good. One transaction waits at most 2 pulse cycles (2s) to be confirmed.

**Dispute Handling**  
Every transaction must have sufficient security deposit and enough time for more nodes to verify it. Transaction is processed by multiple nodes simultaneously, and bookkeeper and backup bookkeeper are randomly selected to avoid collusion. If any node disputes the transaction, disputation resolution procedure kicks in. During the procedure, all relevant funds are frozen.

All of the nodes check the disputed transaction and vote. Security deposit of the erroneous node will be confiscated. Bookkeeper can only process the transaction amount that the frozen deposit can cover.

**Legitimacy Check**  
    1. The trading parties check the legitimacy of the transaction.   
    2. The bookkeeper checks if the general ledger is in balance.   
    3. The branch bookkeeper verifies the legitimacy of the transaction on its branch.   
    4. During synchronization, all leaf nodes check the balance and sanity of their neighbors' accounts.  

**Redundancy Backup**  
All the account information on a branch are stored by a few bookkeepers. In order to keep the network robust, all nodes are encouraged to redundantly backup account information of its neighboring branches.

There are two ways for reward. Regualr nodes can audit the transaction information of nearby branches and earn reward. For illegal transaction, account information of neighboring nodes are frozen until the state of their branches recuperate. Number of neighboring nodes could be 1, 3, 7, 15, 31, 63, 255.  

All of the nodes will strive to maintain the health of the network, in order to earn reward and avoid collateral damage.  
Because of the redundant backup, a small number of nodes can recover the whole network after a crash.  
Redundant backup happens when bookkeepers broadcast information after finishing their tasks.  
Network recovery happens at information verification when a node getting online.  

#### 6.4 Dispute Handling**  
**Publicity**  
Every transaction must have a period of publicity, 24hrs by default. If other nodes find any problem, the transaction will be reported and dispute handling procedure kicks in.  
**Fundtime**  
Every transaction involves different amount of fund. Each bookkeeper or auditor pledges different amount of credit. Therefore a new concept _Amount-time_ is coined for the convenience of calculation.  
amount-time = Amount of fund * Time during which the fund is occupied  
For transaction,   
Amount = transaction amount, time = period of publicity  
For auditing and bookkeeping,  
Amount = amount of credit pledged, time = period the fund frozen  
**Accumulated Period**  
= accumulated (for auditing or bookkeeping) amount-time / transaction amount  
**Transaction Security Threshold**  
+ Minimum number of users who audit a transaction  
+ Minimum accumulated audit period  
    Security threshold sets the minimum auditing cost, so the cost of mining and service fee can be estimated. Accumulated audit period of a normal transaction will be 1.  
    Different audit period has different value, therefore different system reward.  
**Dispute Procedure and Penalty**  
All of the online nodes will be responsible to audit transactions within a certain scope.  

#### 6.5 Smart Contract
Smart contract is similar to transfer procedure. A contract includes smart code, signature of consignor, upper limit of contract. The contract will be executed by the bookkeeper in upper level by default. User can also appoint a consignee node. For complicated contract, if unforeseeable risk is too high, the security deposit of nodes executing and checking the contract shall be raised to cover the risk. It is like a third party that endorses the contract with its own credit and earns corresponding reward.

When dispute happens, the whole network enters dispute handling procedure. The nodes that vote must freeze some of its own asset. Afterward the losing side will lost its frozen asset.  

#### 6.6 Account Inquiry  
Information of system account is public, but regular user account is private and cannot be inquired by 3rd party. However authenticity of account information that user chooses to demonstrate can be verified by checking the synopsis recorded in public ledger. The truthfulness of synopsis can be confirmed by checking synopsis of child branches.  
### VII. Privacy and Transaction Security
#### 7.1 Privacy and Consensus  
**Conflict**  
In traditional block-chain, consensus is based on transparency. If account information is hidden, there is no way to confirm the legitimacy of a transaction. The buyer might overspend.  
**Solution**  
For parties not involved in a transaction, the purpose of knowledge of the transaction is more likely to avoid potential harm of ignorance, rather than voyeurism. If a proper balance of translucency can be reached, there is no need for the others to know the detail of a deal.  
**Premise and Principles**  
The premise and principles for anonymity is that the parties involved in a deal and those not do not harm each other.  
#### 7.2 Transaction Security  
Based on not-be-a-sucker principle, it is necessary to make every member of an organization believes a deal is legal.  
#### 7.2.1 Safety Conditions  
Legal transaction must meet the following conditions.  
+ Truthful intention  
The deal is confirm by both parties.
+ True deal  
The seller has sufficient goods to trade.  
Fund transferred cannot be negative.  
No negative account after the transaction.  
+ Balanced Account  
Without system reward, the overall account balance does not change after a transaction.  
System reward can be counted as a transfer from organization to users, therefore total account balance still does not change.  
#### 7.2.2 Safety Analysis  
+ Traders' responsibility   
The seller will strictly check the safety conditions above, otherwise it will not receive the fund.  
The buyer will also double check the parameters and procedure of the transaction, for illegal deal will be punished by the organization.  
+ Other members  
The other parties need to check if there is any fake transaction or account balance issue, in order to protect themselves from collateral damage.  
+ Other Child Organization  
Child Organization: Every account ID is a 160 bit number on a Merkle tree. If subnet mask of Ethernet can be applied to mask account ID one bit a time, one account can belong to 160 child organizations (sub-tree) simultaneously.  

The question of collateral damage can be restated. Any child organization that does not include trading user only need to verify there is no fake deal or account balance issue within the child organization that includes a trading user.  

#### 7.3 Anonymity Strategy  
**Child Organization Opacity**  
If an child organization is a unit of settlement, transaction information can be hidden within the child-organization. Any other user only need to worry about the general account balance of the child organization and legitimacy.  
**Third Party Opacity**  
If there is a third party endorsement to ensure no collateral damage to the others irrelevant to the deal, transaction information can be revealed in delayed time.  
**Translucent Network**  
During redundant backup, information of an account is backed up only by a few accounts nearby. Accounts faraway can only acquire summarized information through a few top level nodes, without knowing any details of the lower level nodes. This strategy not only guarantees the transaction information safely confirmable, but also the anonymity of information.  
**Capital flow unimportant**  
In reality, the anonymity of capital flow is more important than account balance. However in Leither network 3rd party will be more interested in account balance and legitimacy, instead of capital flow. If system purges transaction information periodically, after a while capital flow will not be traceable.
#### 7.4 Anonymity Method  
The following methods implement information anonymity.
**Credit Endorsement**  
A third party, bookkeeper by default, can be appointed to endorse a transaction. For the other users, the transaction runs on branch and is pledged by the endorsor, so they are collateral damage free. 
**Neighbor Check**  
After a transaction is committed and Merkle tree changed, bookkeeper must broadcast the information. The neighbors of the trading nodes will get details because they provide redundant backup, so must verify the transaction and check the Merkle tree. If problem detected, dispute procedure kicks in. Except a few bookkeepers and neighbors, most of the users will not receive transaction details.  
**Multiple transfers**  
If someone wants to know the details of transaction, it needs to access the trading nodes during the transaction. It will cost time and money. User can choose to transfer fund multiple hops through different branches, which make the tracing of a deal virtually impossible.  
**Transfer across organizations**  
Every organization has its own consensus system. Different organizations can exchange their tokens. Transfer across organizations will also increase the cost of tracing exponentially.

### VIII. Data Processing and Quantitative Analysis
**Data Storage**  
In theory Sparse Merkle Tree is an equilibrium probability binary tree. The closer to top, the evener the distribution. Close to the bottom, the length of branches will vary substantially. It is necessary to reduce the disparity. On the other hand, a 160-bit account number is inconvenient to process. Since one byte can differentiate 8 levels of the tree, 8 levels are stratified together. Assume tree node information is 36 bytes (network number 8 bytes + amount 8 bytes + account number 20 bytes ).  
**Node Group**  
One group is a sub-tree that can hold up to 256 nodes, with max height of 64. The max storage space of one group is 10KB.
**Sub-Database**  
Every 256 node groups or less can be stored in one sub-database. One sub-database requires a maximum of 2.5MB of space to store 65536 node accounts.  
**Miner's bookkeeping database**  
In one node (user) database, if there are more than 256 node groups, system will splits some data to other nodes. A bookkeeper can store max 256 sub-databases, 16,777,216 node accounts, which requires 640MB space.  
**Account summary information**  
The topmost node group records general ledger status, including account information of the topmost 256 branches, such as synopses and account balance. The size of all these information is less than 10KB. It must be synced to every online node in the system. A normal user only need to know the general status and information of its own branch.  

**Quantitative Analysis**  
Usee BTC as reference, currently there are about 30 million accounts and the average number of transactions is 6.67 per second.
Assume the following system parameters:
+ Number of accounts 2^32 = 4.3 billions  
+ Transaction concurrency 1 million/s  
+ Pulse cycle 1 second  
+ Election cycle 30 minutes  

**Scale Analysis**  
Most of the accounts are located at Level-32. Bookkeepers are stratified into 3 groups to cover 8, 8 and 16 levels respectively.  
**Miner HD usage**  
If one branch miner process one sub-database (65536 accounts), data size is 2.5MB. If it is one database (16,777,216 accounts), data size will be 640MB.  
**Miner memory usage**  
Only online nodes need to be processed by a miner.  
For Level 1 or 2 miner, most branches need to be processed online. All 10KB data need to be kept in memory.  
For lower level miners, there are only 30 some transactions per second. No need to keep everything in memory, read related data and synopsis of neighboring branches when necessary. The data size is 30*16*2*28 byte = 7.5KB. Times of database access is 90.  
**Miner CPU usage**  
CPU usage is mainly for verifying transactions and generating synopsis.  
Level 1 or 2 miner must calculate for every transaction, which means 16 million synopses. However in reality at most only one calculation for whole network is necessary in each pulse cycle, which is 512 times of synopsis calculation, about 10KB data.    
Branch root node need to verify the whole branch, which requires 30*(16+16)=7KB data, and 96 synopses.  
**Traffic load among miners**  
The search for a node is load-balanced on each node within a DHT network, therefore the amount of traffic is negligible.    
The network structure of ledger is predetermined within an election cycle. Traffic load in fixed network is negligible.  
Transaction between users is handled by the branch node of each trading nodes. Traffic load of average 30 transactions is small.  

Traffic load of root node need to be examined closely below.  
Root node must broadcast to backup miners and lower level miners. The data to backup miner is 10KB/s, so the maximum data rate to 256 miners simultaneously is 2.5MB/s. If the bandwidth is insufficient, miners can be further divided into subgroup of 16 miners per level. Data rate becomes 160KB/s. Number of broadcasts from top to bottom increases from 3 times to 5.  

**Communication time among miners**  
The process of data involves 3 to 4 levels of nodes.  
The submission and verification of a transaction happen in two different pulse cycles.  
Submission is delivered to branch root node, so times of communication is 1.  
Verification request is delivered to root node from branch node, and the result propagates once or twice back to branch node.  
Most of the operations can be done in less than 10 hops of communication. Normally it takes less than 2 pulse cycles.

Based on the above analysis, with equipment and bandwidth of home network, it is completely feasible to establish decentralized network of 4.3 billion users and support 1 million/s concurrency.  

### IV. Structural Summary
The overall plan of this paper can be summarized in the following key points.
#### 9.1 Credential is the result of meaningful work.
Value is created by work that is meaningful to others. Leither created the original credential system with the meaningful of work. Every quantum of credit corresponds to a piece of work result, so credit can be used to exchange for product and service.  

Comparatively, the meaningful work in block-chain is just bookkeeping, which is at lower end in the chain of value.  
#### 9.2 Interaction is based on credit
Credit can endorse any transaction on network. The encryption system is a mechanism to prove user authentication and information. Miner pledges its credit to ensure the others that it will do a good job, otherwise be punished accordingly.  

With security deposit and arbitration mechanism, complicated smart contract can be implemented to support large scale interactions among multiple parties. All of these operations rely on the most basic PKI algorithm, even the cheapest devices are powerful enough to support. Comparatively, traditional block-chain depends on more and more on dedicated devices and squander expensive electricity.  
#### 9.3 Completeness of time-space comes from synopsis and time sequence
Leither system defines a 160-bit account space to accommodate all possible accounts on a sparse Merkle.  
All account information stored on leaf-node, and summarized bottom up to the root-node. The state of the whole system can be represented by one synopsis.  
Time sequence is used to coordinate all the nodes, and record state of time-space at different stages. Time sequence also solved CAP problem in distributed network.  
#### 9.4 Division of work and coordination improve processing power as a whole
Network nodes can be partitioned into layers according to the scale of the network. Transaction data are saved on user nodes. Consensus is confirmed at branch nodes. Data synced network wide is tiny amount of synopses and summaries.  

Network can be more stable and robust by adding proper number of backup nodes. Comparatively, in regular block-chain except large number of miners of PoW, all the bookkeeping is accomplished by all the nodes together, which greatly restricted network scale and throughput.  
#### 9.5 Endorsement and shade of grey solves privacy problem
### X. Advantages of Leither compared with regular block-chain
+ Functionality
With MiMei and Leither application system, most regular internet platforms can be reconstructed as user oriented network.  
With organization and consensus, Leither can implement most of the ethereum functionalities and organize all the participants.  
The above functions combined can give internet back to its users from the grip of oligarchies.  
+ Value  
Traditional block-chain does bookkeeping for others.  
Every organization of Leither has a mission to decentralize an internet application, to give user control of its data, to satisfy user need first and foremost. The action of every participant has real value, in order to take back the outrageous profit of oligarchies and redistribute it user oriented.  
+ Cost  
Mining of regular block-chain is hugely expensive because of all the machines and electricity cost to do the most basic work.  
In Leither, consensus is achieved by doing the very job of internet business. Tv box, router, NAS, smart phone, even Raspberry zero worth $10 can provide valuable services and ROI much higher than mining machine.  
+ Efficiency  
Leither elastic system can scale according to business request. Workload is balanced on to every node through consensus.  
The most business load, the more nodes and powerful the network processing capacity, far exceeds that of regular block-chain. Because of its high efficiency, a few nodes can construct their own business and consensus mechanism.  
+ Privacy  
In regular block-chain, flow of every account is open book. Even though account is anonymous, it is possible to identify the account by analyzing flow information.  
In Leither, except the earliest contribution of a user need to be publicized for auditing. All the transaction afterward only need to open to related parties and a few nodes. The whole network has a shade of grey. Creator can set translucency to protect privacy.  
+ Lawfulness  
The token of Leither has corresponding service as object. It is more like the points, notes, share, Q coin, etc. It has many reference in reality. It is perfectly legal.  
On business form, it is similar to early Taobao business, completely decentralized and conducted in the name of individual.

Smart contract is dealing with real business data, avoid the falling hole of ethereum that it can only process virtual business such as encryption coin.  
