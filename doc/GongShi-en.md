Part II Consensus Mechanism
========
### Summary
### I. Organization
Chapter Nine of _MiMei and Application System_ mentioned individual centered model of relationship chain, and credit model based on the relationship-chain. In social life, people always belong to one social group or another, which is called **Organization** in this paper.  

Organization is a group centered structure. User chooses to join a group after its foundation. There are three types of group identities: supervisor (founder), member and guest. Supervisor and member maintain the group together. In traditional internet, both structures are maintained by platform operator, who also set the corresponding regulations.

The maintenance of a group is a public service, which need a block-chain like solution to coordinate the interests of its members. When more potent public service attached to a group, the group becomes an organization. Any organizational entity on internet can be created in the form of a Leither group.
### II. Organization Services
Organization service is public service attached to an organization. Most of the business currently on internet can be reconstructed as organization service.
#### 2.1 Categories of Internet Business Services
Internet business can be divided into categories of search, content, social media and tool.  
+ Search  
Google, Amazon, Alibaba and Baidu all profit from **Information Search**.  
In Leither network，any user can create its own search preference and strategy, or filter and summarize search results from 3rd parties. Therefore search service will not be controlled by a few providers, instead user can take control of and decide its own search information.  
+ Social service  
FB, Twitter, Wechat make huge profits by controlling the **Social Relationship** data of users.  
In Leither network, data of social connections are saved by user owned devices, therefore all of the attached messages, Apps and services are also selected by user.  
+ Content and Service  
On-line video, music and news media provide **Content and Service**.
In Leither network, content nodes implement content services in the form of MiMei. MiMei can flow among nodes that are not under the control of oligarchies any more. User can collect its own data from internet using plug-ins and save it on its own node.  
+ Tool  
Google Doc, online office, XMind make profits providing **Tool Services**.  
In Leither, Tool App is also a MiMei. 

#### 2.2 Split of Business Services  
The above business service can be split and reconstructed through the coordination of the following members of Leither ecosystem.  
+ Application Development  
Programmer and developer who develop applications for Leither ecosystem, similar to R&D in regular business.
+ Service or Content Production  
Service provider and content producer, similar to editor of a website or operation department.
+ Service Support  
People who provide physical devices for the whole ecosystem, different from regular internet, devices do not have to belong to a company. Similar to the role of miner in BTC.  
+ Service Promoter  
One way is to promote web service of Leither, such as website, link, app and applet, in regular internet. Another is to create tags for content, and provide them to Search Node as public services. Similar to what operation department of a website does.  
+ Value Appreciation  
Every node can provide services, which can be monetized and used to pay for other services. Information search and online traffics can generate great financial income, which is the foundation of the ecosystem's value.  
Usually a block-chain like system can coordinate the investment and return of its members. More details in later chapters.  

#### 2.3 Service Support
Similar to the work of miner in BTC, including node management, MiMei migration, information search and group consensus.  
+ MiMei Storage  
Leither decouples internet content and service into MiMei form. MiMei is saved on node. MiMei flow among nodes on demand. Bill of services between nodes is settled with credit. Node participates in the settlement of group services by group consensus.  
+ Information Index  
Information of search index is stored on the Distribute Hash Table (DHT) of every online node. Node participates in the settlement of group service by group consensus.  
+ Docking Service  
Applications such as APP, DNS and applet, that connect MiMei with regular internet. Those services run servers of regular internet and participate in the settlement of group services by group consensus.
+ Consensus  
Public service of a group. Consensus runs on all the nodes.

### III. Organization Functions
#### 3.1 Create Organization  
Any user can create an organization. In the beginning, Leither creates a pair of encryption keys for the organization. Default public key ID is the organization ID. Creator is a organization member and supervisor by default. Organization has the same authorization scheme as MiMei. User identity in a group includes supervisor, member and guest.  

Organization runs on the node where it is created. Distributed organization also starts to work on its creation node. The default network structure is DHT network which stores public information of the organization.

Public information of the organization can be inquired through DHT network. Information of member's node can be searched through routing service. Members can send P2P message. Group chatting data belongs to public and saved on DHT network.
#### 3.2 Join Organization  
User can apply to join a organization directly, or send application through a member. Once supervisor confirms the application, the user becomes a member and DHT is updated with corresponding information. Member can check public information of the organization.

User can also help with certain task of the organization and become a member as reward.  

#### 3.3 Business Development  
Organization business is a pack of Leither applications, which can be published as public information of the organization network.  
For example a video website can be separated into the following sub-business:
+ Content collection and production  
+ Add tags to content  
+ Application support with resource, including storage, bandwidth, CPU, etc.  
+ Content access channel, such as APP, applet, DNS, routing, search, etc.  
+ Promotion to send content to people who need it.  
+ Value appreciation  
    There two types of value appreciation. In the first type, content contains information of a 3rd party who will pay for the promotion in advance. In the second type, the content is valuable to average users, who would pay for it or watch the attached commercials. Commercial business model is tried and true, a highly matured business practice in regular internet.

### IV. Valuation-based System
**Value is created meaningful work**. It is the basic principle for designing valuation-based system.  

Every quantum of token shall be based on a meaningful work and the quantity of its value is based on the contribution of the work to the ecosystem, or how useful the work is.  

An organization ledger in the form of a sparse Merkle Tree with versioning mechanism will be kept by consensus. Any incentives issued by the organization will have to be announced with reason and proof for publicity within the organization. Different from POW, it is called **Proof of Performance (PoP)**.  

System rewards include:  
+ Structural reward  
Rewards to member who has made fundamental contribution to establish the organization, such as founder, investor, system and App developer, who did all the absolutely necessary work to make it happen. The reward is recommended by supervisor and issued after vote of consensus.

The reward is for creating the organization.  

+ Reward to service and content  
Organization provides valuable service or content to users. At every step of value creation, contribution is recorded. At the point of appreciation of the value, smart contract is triggered and incentives rewarded to each member who did valuable work.

The reward can be distributed by fixed proportion or dynamically balanced following the following mimic process.  
1. Developer writes a content collection plugin.  
2. Collector uses the plugin to gather content from internet, turn it into MiMei, and publish the MiMei onto Leither network. MiMei will indicate collector of its content.  
3. Operator push the content to users through channels such as App, applet, website, or group message.  
4. User browse the content, commercials are displayed, maybe even clicked.  
5. Checking mechanism is built in the commercials. Once displayed, it will trigger the smart contract, which then transfers tokens from advertiser's account to each participators of the whole process, respectively.  

The reward mechanism will feedback the healthy development of the organization and the growth of usefulness.  

+ Mining Reward  
Same as any other mining, all the online nodes can join vote and contract service. Reward for new block will be used to support the network, manage public information and routing service.  

The reward makes sure that organization stays active.

+ Operation Reward  
In the beginning when the system cannot sustain itself with insufficient content and service. Extra incentives can be rewarded to accelerate the growth of an infant system.

The reward helps the organization to mature faster.  

### V. Basic Concept
**Distributed Hash Table (DHT) network**  
In DHT network, node can be quickly located by its ID, to access its information. Public information of an organization can also be saved over the whole network distributively.  

**Sparse Merkle Tree**  
Jump-table is a key data structure used in Redis and LevelDB. It introduced the concept of **equilibrium probability**, which makes it possible to manipulate a tree without changing its structure substantially. Compared with other algorithms, jump-table has a little bit of impact on performance, but still keeps the computation complexity at the same order of magnitude. Using the same algorithm, Merkle Tree is improved to create so called **Sparse Merkle Tree (SMT)**, which is used to hold core information of an organization, similar to the public ledger of a block-chain.

Each leaf node of a SMT stores account information of an organization member. Account number is the member's ID that is the hash of the member's public key. Each ID is 160 bit long. All of the possible IDs can fill up the leaf nodes of a binary tree of height 160. Because the actual number of IDs is far less than the number of possible leaves, the tree is a sparse tree.  

If a tree-node has only one child, the branch can be shorten by moving the child node up one level to replace its parent. This simple optimization can reduce the height of the tree tremendously. If the number of randomly distributed IDs is n, most of the IDs will be on leaf-node at level log2(n)+1. The closer to the root-node, denser the tree. The addition or removal of a tree-node, or change of account information, only effects the branch where the node is on.  

On SMT, account information on leaf node can be quickly located, which is essential to the implementation of time-space snapshot, network pulse, network partition, fund transfer and smart contract. Information of each account is on the leaf node of the tree. Each branch uses hashes of its root node's children and account sum to generates its own hash.  

**Node Grouping**  
With the growth of network scale and traffic, storage, data processing and communication will eventually overload a network node. If not optimized, Leither network will fall in to the same conundrum of 7 throughput like BTC.

Nodes within an organization can be grouped to solve the above problem. Neighboring nodes on SMT are grouped togather as a block.

**Network Pulse**  
is an incremental sequence number broadcasted top down every **Pulse Cycle**. A pulse cycle is 1 second.  

The sequence number serves as the synchronization timer of each node, and version number of its data. Within a pulse cycle, every node backups the new data received in last one. According to the hight of Merkle tree, each node is responsible to process data for several layers. The work done is called **Proof of Performance**, similar to PoW in BTC, and will be rewarded by the system.  

**Time-space Snapshot**  
In LevelDB, writing of new data using writeStream can be optimized up to the speed limit of storage media, even faster than database reading. On the other hand, sequence number concatenated at the end of the key serves as the revised version number, with which historical record can be inquired. Network Pulse is equivalent of generating sequence version number in levelDB, with which ledger data can be quickly recorded in each cycle.

The same operation can be applied to multi-dimension time-space Merklenodes do not change, the current version of Merkle tree is almost identical to the previous one, so only the limited variations on the current branch need to be processed.

The underlying MiMei database of Leither has already built in with similar functions to support time-space versioning.

**Network Partition**  
Partition enables two unique advantages. The first one is elastic concurrency support. The second is an opaque network with variable shades of gray.

A **Election Cycle** is 30 minutes (by default), during which trustworthy nodes are elected as **Bookkeeper**. Each branch is also appointed a backup candidate to safeguard the network. The Leither nodes that are assigned to the lowest leaf layers process specific business instructions. Upper layers only merge changed child-branches into the tree, and generate synopsis.

Top layer information is synchronized over the whole Leither network.

### VI. Related Procedures
#### 6.1 Network Construction  
After user joins an organization, it becomes a member of DHT network. User writes routing information of its node into DHT, so that other nodes can find it for communication or services.  
#### 6.2 Ledger Construction  
Ledger is a sparse Merkle tree with time-space versioning. In order to keep network wide messages in sync, the whole network will share a time sequence variable, aka Time Sequence. Its initial value is 0, system generates one pulse periodically (1s by default). If system state changes, time sequence increments. Other wise, it stays put.  

System periodically elects **Bookkeepers**, who backups each other. The first one is chief bookkeeper. The term of bookkeeper is one **Election Cycle** (30min by default). Each branch is responsible to a few layers.  
#### 6.3 Transfer Procedure  
**Organization Distribute**  
The very first token distribution will be announced network wide, with reasons for scrutiny by the members. However the transfer of tokens between users is private, or public only to the relevant nodes.  
**Transfer between users**  
The transaction between two users is relevant only to themselves. After the transaction is confirmed, information attached with signatures of both users will be sent to the parent nodes of the users on Merkle tree respectively. Each parent node will record the changes in its own branch and broadcast the information among backup nodes in the same level. After new time sequence is created (in less than 1s), transaction data becomes read only. Bookkeepers begin to check the branch bottom up and summarize branch information to generate synopsis. The chief bookkeeper summarizes summarize general information and generate synopsis for the tree, and broadcast top down to each level.

Both parties of a transaction record time sequence, synopsis of each level and its own account information, and finally transaction is committed. One transaction waits at most 2 pulse cycles(2s).

**Dispute Handling**  
Every transaction must have sufficient security deposit and enough time for other nodes to verify it. Transaction will be processed by multiple nodes simultaneously. Bookkeeper and backup bookkeeper are randomly assigned to avoid collusion. If any node disputes the transaction, dispute resolution procedure kicks in. During the procedure, all relevant funds are frozen.

All of the nodes check the disputed transaction and vote according to the result. Deposit of the erroneous node will be confiscated. Bookkeeper can only handle the amount of transaction that the frozen fund can cover.

**Legitimacy Check**  
The relevant parties check the legitimacy of the transaction.   
The general bookkeeper checks the overall account balance.   
The branch bookkeeper verifies the legitimacy of transactions on its branch.   
Each leaf node checks the balance of its neighbors and the sanity of their account, each time when it synchronizes with them.  

**Redundancy Backup**  
The minority of bookkeepers save all the account information on the branch. In order to keep the network robust, all nodes are encouraged to redundantly backup account information of its neighbors.

There are two methods. Ordinary account can audit the transaction information of nearby branches and earn reward. For illegal transactions, freeze account information of neighboring branches until the state of their accounts recuperate. Number of neighboring nodes could be 1,3,7,15,31,63,255.  

All of the nodes will strive to maintain the health of the network, in order to earn reward and avoid loss.  
Because of the redundant backup, minority of the nodes can recover the whole network after a crash.  
Redundant backup happens when bookkeepers broadcast information after their finished their tasks.  
Network recovery happens at information verification when a node getting online.  

#### 6.4 Dispute Handling**  
**Publicity**  
Every transaction has a period of publicity, 24hrs by default. If other nodes find any problem, the transaction will be reported and dispute handling procedure kicks in.  
**Fundtime**  
Every transaction deals with different amount of fund. Every bookkeeper and auditor pledge different amount of credit. Therefore a new concept _Fundtime_ is coined for the convenience of calculation.  
Fundtime = Amount of Fund x Time during which the fund is occupied  
For transaction,   
Fund = transaction amount, time = time of publicity  
For auditing and bookkeeping,  
Fund = amount of credit pledged, time = time of fund frozen  
**Accumulated Audit Period**  
Accumulated (for auditing or bookkeeping) Fundtime / transaction amount  
**Transaction Security Threshold**  
+ Minimum number of users who audit a transaction  
+ Minimum accumulated audit period  
    Security threshold sets the minimum auditing cost, so that the cost of mining and service fee can be estimated. Accumulated audit period of a normal transaction will be 1.  
    Different audit period has different value, therefore different system reward.  
**Dispute Procedure and Penalty**  
All of the online nodes will be responsible to audit transactions within a certain scope.  

#### 6.5 Smart Contract
Smart contract is similar to transfer procedure. A contract includes smart code, signature of consignor, upper limit of contract. The bookkeeper in upper level by default is responsible to execute the contract. User can also appoint a consignee node. For complicated contract, if unforeseeable risk is too high, the security deposit of nodes executing and checking the contract shall be raised to cover the risk. It is equivalent to a third party endorse the contract with its own credit and earn corresponding reward.

When dispute happens, the whole network joins dispute handling procedure. The nodes that vote must freeze some of its own asset. After the dispute, the wrong side will lost its frozen asset.  

#### 6.6 Account Inquiry  
Information of system account is public, but regular user account is private and cannot be inquired. However authenticity of account information that user chooses to demonstrate can be verified by checking the synopsis recorded in public ledger. The truthfulness of synopsis can be confirmed by checking synopsis of child branches.  
### VII. Privacy and Transaction Security
#### 7.1 Privacy and Consensus  
**Conflict**  
In traditional block-chain, consensus is based on transparency. If account information is hidden, there is not way to confirm the legitimacy of a transaction. The buyer might spend more than it owns.  
**Solution**  
For parties not involved in a transaction, the purpose of information is more likely to avoid potential harm of ignorance, rather than voyeurism. If a proper balance can be reached, there is no need for the others to know the detail of a deal.  
**Premise and Principles**  
The premise and principles for anonymity of personal information is that the parties involved in a deal and the others do no harm to each other.  
#### 7.2 Transaction Security  
To avoid collateral damage, it is necessary to make every member of an organization believes a deal is legal.  
#### 7.2.1 Safety Conditions  
Legal transaction must meet the following conditions.  
+ Truthful intention  
The deal is confirm by both parties.
+ True deal  
The seller has sufficient commodity to trade.  
Fund transferred cannot be negative.  
No negative account after the transaction.  
+ Balanced Account  
Without system reward, the overall account balance do not change after a transaction.  
System reward can be counted as a transfer from organization to users, therefore total account balance still do not change.  
#### 7.2.2 Safety Analysis  
+ Transaction parties' responsibility   
The seller will strictly check the safety conditions above, otherwise it will not receive the fund.  
The buyer will also double check the parameters and procedure of the transaction, for illegal deal will be punished by the organization.  
+ Other members  
The other parties need to check if there is any fake transaction or account balance issue, in order to protect themselves from collateral damage.  
+ Other Child Organization  
Child Organization: Every account ID is a 160 bit number on a Merkle tree. If subnet mask of Ethernet can be applied to mask account ID's bit one at a time, one account can belong to 160 child organization (sub-tree) simultaneously.  

The question of collateral damage can be restated. Any child organization that does not include transaction party only need to verify there is no fake deal or account balance issue within the child organization that includes a transaction party.  

#### 7.3 Anonymity Strategy  
**Child Organization Opacity**  
If an child organization is a unit of settlement, transaction information can be hidden within the child-organization. Any other user only need to worry about the general account balance of the child organization and legal.  
**Third Party Opacity**  
If there is a third party endorsement to ensure no collateral damage to the others irrelevant to the deal, transaction information can be revealed in delayed time.  
**Translucent Network**  
During redundant backup, information of an account is backup only by a few accounts nearby. Accounts faraway can only acquire summarized information by a few top layer accounts, without knowing any details of the lower layers. This strategy both guarantees the transaction information safely confirmable, and also the anonymity of information.  

#### 7.4 Anonymity Method  
The following method to implement information anonymity.
**Credit Endorsement**  
A third party can be appointed to endorse a transaction. Bookkeeper is the default executive. Transaction runs on a branch. Because there is an endorser to pledge for the transaction, third party will not take damage.  
**Neighbor Check**  
After a transaction is committed and Merkle tree changed, bookkeeper must broadcast the information. The neighbors of the transaction parties will get details because they provide redundant backup, so they must verify the transaction and check the Merkle tree.  

If error happens, dispute procedure kicks in. Except a few bookkeepers and neighbors, most of the users will not receive transaction details.

### VIII. Quantitative Analysis
Usee BTC as reference, currently there are about 30 million accounts and the average number of transaction is 6.67 per second.
Analysis Leither network performance with the following assumptions:  
+ Number of accounts 2^32 = 4.3 billions  
+ Transaction concurrency 1 million/s  
+ Pulse cycle 1 second  
+ Election cycle 30 minutes  
Most of the nodes are located at Layer-32. Let's separate the network into 2 16-layers. Each node need 28 bytes (8 bytes for branch, 20 for synopsis). For 1 million/s concurrency, assume each one involves 2 users. On average, each branch processes about 30 transactions per second.  

**Miner HD usage**  
Data storage for one branch is 1.8MB. If a miner need to save 2 layers of data, 3.6MB HD is need.  
**Miner memory usage**  
Only online nodes need memory space from a miner.  
For root node miner, most branches need to be processed online. All 1.6MB need to be kept in memory.  
For sub-branch miners, there are only 30 some transactions per second. No need to keep everything in memory, read related data and synopsis of neighbor branches when necessary. The data amount is 30*16*2*28 byte = 7.5KB. Times of database access is 90.  
**Miner CPU usage**  
CPU usage is mainly for verifying transactions and generating synopsis.  
Root node must calculate every transaction, which means 16 million synopses. However in reality only need to calculate once network wide in each pulse cycle, which is 131,000*20 byte = 2.5MB data.  
Branch root node need to verify the whole branch, which requires 30*(16+16), and calculate 96 synopses.  
**Traffic load among miners**  
The search for a node is load-balanced on each node within a DHT network, therefore the amount of traffic is negligible.    
The network structure of ledger is predetermined within an election cycle. Traffic load in fixed network is negligible.  
Transaction between users is handled by the branch node of each transaction party. Traffic load of average 30 transactions is small.  

Traffic load of root node need to be examined closely below.  
Root node must broadcast to backup miners and branch nodes. The traffic to backup miner is 2.5MB/s. If broadcast is B-tree, total data is 5MB. If data is delivered to node one by one, total data traffic is number 1.5MB * num_of_miners. The total data amount to sub-miners is 16*2*28896 byte = 56 MB/s.

This workload is more than a regular node. A load-balance middle layer can be added to solve the problem. A layer of nodes that handle 8 layers of the B-tree can reduce the communication load to 224 KB/s.

**Communication time among miners**  
The process of data involves 3 to 4 layers of nodes.  
The submission and verification of a transaction happen in two pulse cycles.  
Submission is delivered to branch root node, and done in one hop of communication.  
Verification is delivered to root node, which need to broadcast to child node once or twice after the information is processed.  
Most of the operation can be done with less than 10 hops of communication. Normally it takes less than 2 pulse cycles.

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
