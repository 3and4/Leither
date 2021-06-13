Part II Consensus Mechanism
========
### Summary
### I. Organization
Chapter Nine of _MiMei and Application System_ mentioned individual centered model of relationship chain, and credit model based on the relationship-chain. In social life, people always belong to one social group or another, which is called **Organization** in this paper.  

Organization is a group centered structure. User chooses to join a group after its foundation. There are three types of group identities: supervisor (founder), member and guest. Supervisor and member maintain the group together. In traditional internet, both structures are maintained by platform operator, who also set the corresponding regulations.

The maintenance of a group is a public service, which need a block-chain like solution to coordinate the interests of its members. When more potent public service attached to a group, the group becomes an organization. Any organizational entity on internet can be created in the form of a Leither group.
### I. Organization Services
Organization service is public service attached to an organization. Most of the business currently on internet can be reconstructed as organization service.
#### 2.1 Categories of Internet Business Services
Internet business can be divided into categories of search, content, social media and tool.  
+ Search  
Google, Amazon, Alibaba and Baidu all profit from **Information Search**  
In Leither network search data is stored by special searching node in distributed mechanism. User can search for multi-dimension results by selecting different searching node.  
+ Social  
FB,Twitter, Wechat make huge profits by controlling the **Social Relationship** data of users.  
In Leither network, data of social connections are saved by user itself. Corresponding messages, App and service are also selected by user, the owner.  
+ Content and Service  
On-line video, music and news media provide **Content and Service**.
In Leither network, content nodes implement content services in the form of MiMei. MiMei can flow among nodes and not under the influence of monopoly any more. User can collect its own data from internet using plug-ins and save it on its own node.  
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
**The foundation of value is meaningful work**. It is the basic principle for designing valuation-based system.  

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
**DHT network**  
Distributed Hash Table (DHT) is used to build basics of network topology where nodes are organized by their IDs, so that DHT can provide routing function for the organization.

Node ID can be used to locate a node quickly and access information of the node. Information of the organization can be saved over the whole network in distributed method.  

**Sparse Merkle Tree**  
is constructed for consensus data, to save the core information of an organization, similar to the public ledger of block-chain.

Jump-table is a key data structure used in Redis and LevelDB. It introduced the concept of equilibrium probability, which makes it possible to manipulate the tree without changing its structure extensively. Compared with other method, the merit of jump-table costs a little bit of performance, but still on the same order of magnitude. Leither introduced a similar algorithm, Sparse Merkle Tree, to improve Merkle tree.

Leither ID is a hash key of 160 bit. All of the possible IDs can occupy a binary tree of height 160. Every ID has a fixed location on the tree. Because the real number of IDs is far less than the number of possible leaves, the tree is a sparse tree.  

If a tree-node has only one child, the branch can be shorten by moving the child node up one level to replace its parent. This simple optimization can reduce the height of the tree tremendously. If ID is randomly generated, the number of IDs is N, then most of the IDs will be on leaf-node at level log2(n)+1. The closer to the root-node, denser the tree. The addition or removal of a tree-node only effects its own branch.  

Fast search of a tree-node is essential to implement functions such as time-space versioning, pulse, partition, transfer and contract, which will be described later in this chapter. Information of each account  

**Network Pulse**  
is an incremental sequence number generated by some trusted Leither nodes that are chosen periodically. Each pulse is 1 second and Each **Pulse Cycle** is 30 minutes by default. 

The sequence number serves as the synchronization timer of each node, or the version number of its data. Within each cycle, every node backups the new data from last cycle. Each node is responsible to process data for several layers of Merkle tree, according to the height of the tree. The work done is similar to the proof of work in BTC, and will be rewarded by the system.  

**Time-space Versioning**  
In LevelDB, write operation of new data using writeStream can be optimized up to the speed limit of storage media, even faster than database reading. On the other hand, sequence number concatenated at the end of the key serves as the revised version number, so that historical record can be inquired with it quickly. Network Pulse is equivalent of the sequence number in levelDB. With version number, bookkeeping data can be quickly recorded in each cycle.

The same operation can be applied to multi-dimension Merkle tree. Because branches of other nodes do not change, the current version of Merkle tree is almost identical to the previous one, so only the limited variations of the branch where the current node resides need to be processed.

The underlying MiMei database has already built in with similar functions to support time-space versioning.

**Network Partition**  
Partition enables two unique advantages. The first one is elastic concurrency support. The second is an opaque network with variable shades of gray.

In each pulse cycle, trustworthy Leither nodes will be chosen to process several layers of nodes on the tree. Each layer is appointed with a backup candidate to safeguard the network. The Leither nodes that are assigned to the lowest leaf layers process specific business instructions. Upper layers only merge changed child-branches into the tree, and generate synopsis.

Top layer information is synchronized over the whole Leither network.

### VI. Related Procedures
#### 6.1 Network Construction  
After user joins an organization, it becomes a member of DHT network. User writes routing information of its node into DHT, so that other nodes can find it for communication or services.  
#### 6.2 Ledger Construction  
Ledger is a sparse Merkle tree with time-space versioning. In order to keep network wide messages in sync, the whole network will share a time sequence variable, aka Time Sequence. Initial value is 0, system generates one pulse periodically (1s by default). If system state changes, time sequence increments. Other wise, stay put.  
#### 6.3 Transfer Procedure  
**Organization Distribute**  
The very first token distribution will be announced network wide, with reasons for scrutiny by the members. However the transfer of tokens between users is private, or public only to the relevant nodes.  
**Transfer between users**  
The transaction between two users is relevant only to themselves. After the transaction is confirmed, information attached with signatures of both users will be sent to the parent nodes of the users on Merkle tree respectively. Each parent node will record the changes in its own branch and broadcast the information among backup nodes in the same level. After new time sequence is created (in less than 1s), transaction data becomes read only. Trustee nodes begin to check the branch bottom up and assemble branch information, generate synopsis.  Level 1 trustee node assemble information and generate synopsis for the topmost node, and broadcast to nodes on each level. Both parties of a transaction record time sequence, synopsis of each level and its own account information. Transaction is finalized. One transaction waits at most 2 pulses (2s).
