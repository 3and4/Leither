Part II Consensus Mechanism
========
### Summary
### I. Organization
_Chapter IX: MiMei and Application System_ mentioned user oriented relationship model, and credit model based on relationship network. In social life, people always belong to some kinds of social groups, which is called **Organization** in this paper.  

Organization is a group oriented structure. An organization is first created as a group, then people join it to become a member. There are three types of roles within an organization: supervisor (founder), member and guest. Supervisor and member work with each other to maintain and manage the group.

The maintenance of a group is a public service best operated with a block-chain like structure to coordinate the interests of all of its members. A group must become an organization when more potent public service is required of it. Any organizational entity on internet can be reconstructed in the form of a Leither group.
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
Every node can provide public services that can be monetized to pay for other services. Information search and online traffics have great financial value, which is the **foundation of the ecosystem's value**. When packed into a block-chain like system, the investment and return of its members can be coordinated through consensus. More details in later chapters.

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
