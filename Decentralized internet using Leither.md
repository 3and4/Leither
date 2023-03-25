# Decentralized Internet using Leither OS

### Introduction to Leither

In the past decades, a worrying trend has emerged on the internet,
characterized by the increasing concentration of data traffic on a
limited number of APPs controlled by tech giants. These corporations
exploit their dominant hold on data to influence or even manipulate user
behavior to secure massive profits, often at the expense of ordinary
users.

This monopolistic environment creates significant wealth disparities
among Internet companies, hindering competition and innovation while
causing the Internet\'s growth potential to stagnate. However, it is
essential to recognize that a free and open internet serves the
interests of its users the best.

Leither, an ultra-light cloud-based OS, is designed to take the
challenge by dissembling and reconstructing most of the core services of
conventional internet in a decentralized manner, thereby returning data
control to the user, its rightful owner.

The decentralized internet powered by Leither OS has the following
advantages.

-   Straightforward app development

-   Minimal system resources demand

-   Low bandwidth expenses

-   Effortless deployment of elastic cloud services

#### Key concepts of Leither:

-   **MiMei**\
    This concept of MiMei is inspired by Gene and Meme. Meme is the gene
    of civilization, an element of a culture or system of behavior
    passed from one individual to another by imitation or other
    non-genetic means.\
    \
    MiMei is the carrier of information on internet. It defines the
    structure of data storage and operations of associated applications,
    such as create, open, save, read, search, render and user-interact.

    Everything in Leither-based ecosystem is a MiMei, just like everything
    is an object in Java. The referential relationships among MiMei
    objects make it possible to define complex data structures.

-   **Leither**\
    is a container of MiMei, an ultra-light decentralized cloud OS. It
    is designed to serve a limited number of users from each node in a
    decentralized network, aka, a distributed application platform.

    Leither applications are designed to function using significantly
    fewer resources compared to traditional apps, for example to run
    office applications for a small business on a Raspberry Pi Zero.
    
    Leither OS has implemented all the subsystems necessary to support
    internet services, including authentication and authorization, file
    system, database, application framework, independent domain name
    service and routing. Leither API supports about 40 kinds of
    development languages.
    
    PS: Leither, if written in Kanji character, is 弥媒. 弥 means
    extremely small and ubiquitous, usually something imperceptible. 媒
    means media. 弥媒 put together is almost as good as Ether, a substance
    once believed to be the building block of the universe.

-   **Node**\
    Any smart device running Leither OS is a **Node**. As small as 10MB
    in size and compatible with Windows and all flavors of Linux,
    Leither can run on almost any smart devices, like PC, NAS, router
    and Raspberry Pi.

1.  A standalone node can serve as an ultralight server.

2.  Two nodes can backup each other in services like data backup, error
    tolerance, load balance, domain name resolution and routing.

3.  Multiple nodes can form a cloud to support decentralized internet
    applications.

# Part I: LeitherOS and MiMei

### I. System analysis of decentralized internet

Before proposing a strategy for constructing a decentralized internet,
it is essential to briefly examine the varying demands that computers
and the internet have faced during each phase of their evolution, as
well as the corresponding solutions offered. A new solution must
satisfactorily address previous challenges in order to replace the
existing system.

#### 1.1 File \-- Carrier of data

At the dawn of the computer era, files were created as data carriers.
File systems were developed to manage files and storage media. In Unix,
everything is treated as a file, and I/O device is considered a type of
File.

#### 1.2 MIME \-- Type of data

As applications evolved, it became essential to establish the
relationship between files and applications. DOS introduced the concept
of File Extensions. In Windows, a file extension is associated with a
specific application. MIME defined a strict set of rules for file types
when used as email attachments. To this day, all existing web browsers
still adhere to the MIME protocol.

#### 1.3 Database \-- Relationship of data

Database was designed to manage complicated data structure. A database
is a closed space of datasets and their relationships.\
Shortcomings: Data stored in closed space has little liquidity, value
not maximized.

#### 1.4 HTML \-- connections of data over network

The World Wide Web (WWW) facilitates the linking of data across machines
over a network. This interconnection of data across networks transcends
the physical constraints of data storage on mainframes. With the advent
of the internet, an information explosion followed, leading to the
proliferation of gateway sites.\
Shortcomings: connections based on linkage is vulnerable.

#### 1.5 Google \-- search engine

The information explosion led to the demand for search engines. By
incorporating tags into data, Google offers a content search service
based on keyword queries.\
Shortcomings: As data becomes increasingly isolated and monopolized,
data protocols have started to be influenced and manipulated by various
entities.

#### 1.6 Facebook \-- social model

The early social medias, ICQ, MSN, QQ, solved problem of instant
messaging. Twitter provides a platform for broadcasting short messages
quickly. From MySpace to Facebook, a more systematic social media model
is created.

  ---------------------------------------------------------------------------
  **\#**   **Title**         **Detail**
  -------- ----------------- ------------------------------------------------
  1        My settings       Personal information, settings and preferences

  2        Friends           User contacts

  3        App & Services    Application and services for the user

  4        Message           Interactions with friends, applications and
                             services
  ---------------------------------------------------------------------------

This model is remarkably successful. All the afterward internet social
Apps adopted a similar framework, for example WeChat, Facebook, Alipay
and some OA system for enterprise users.\
Shortcomings: Internet becomes more monopolized, more isolated, and data
protocols manipulated.

#### 1.7 Design principles of decentralized internet

Design principles for breaking data monopoly and creating decentralized
internet:

-   **User oriented in every aspect**\
    > A new birth of freedom, and that internet of the people, by the
    > people, for the people.

-   **Free and open system**

1.  Open data protocols.

2.  Open operation system protocols.

Only in an open system, user has the freedom of choice and power of
control.

-   **Unrestricted data migration**\
    With user's authorization and reasonable security discretion, user
    data shall be allowed to move from one platform to another and
    transform from one format to another. Only then, a third party can
    provide valuable applications and services.

The following requests shall be supported in a decentralized system:

-   File, file system, database

-   Association of data type and application

-   Association of contents over network

-   Search of content over network

-   Data migration over network

-   Decentralized social model

-   Light-weight container system

-   User friendly development framework

-   User experience compatible with existing internet

-   Law-abiding

### II. How is data monopolized?

In nascent internet, data access protocols were open. Major platforms,
such as Google and Microsoft, provided open APIs for accessing their
data. The first notable misuse of open data policy was by Facebook,
which leveraged data from other platforms to build its social media
presence, and then locked out its competitors.

As a result, user data became confined within Facebook\'s platform,
giving rise to the initial data monopoly. Later, Apple and Amazon
established their own separate domains of user data. Consequently, the
internet became fragmented. User data is now protected by impenetrable
barriers erected by each monopolist who gains control over it.

There are two routes of data access.

**Data-\>App**

> Users have the freedom to access data and then select the desired
> application for processing that data. For instance, in Windows, users
> can choose different applications to open specific file types or use
> various browsers to open a web page.
>
> In this scenario, users enjoy the liberty of choice and are not
> confined to a specific App.

**App-\>Data**

> User\'s right to access data is restricted to certain Apps, usually
> only one. For example, when user login an e-commerce site for
> shopping, or get on social media to chat with friends.
>
> In this scenario, the algorithm of App determines the way users access
> data, resulting in user behavior controlled by the App\'s algorithm or
> rules.

These two approaches correspond to **Object Oriented
Programming** and **Procedure Oriented Programming**, respectively.

#### 2.1 Principles for breaking data monopoly

-   **Opensource platform**\
    Microsoft dominated the PC ecosystem with DOS and Windows, and
    stifled competition in both office software and browser. Google
    exercises control over data and influences user behavior using its
    search engine. FB does the same with social media.\
    \
    On the other hand, the open-source community transformed the
    operating system into a public platform with Linux. This marked the
    first attempt to open-source a platform. Similar cases are Android
    for smartphone and RISC-V for IC instruction set.\
    \
    There shall be similar opensource platform for collecting user data,
    especially for ***Social Media*** and ***Search Engine***.

```{=html}
<!-- -->
```
-   **Unrestricted data migration**\
    Ensuring the free migration of data should be a widely accepted
    principle and a primary design consideration for operating systems.
    Three measures can be implemented to achieve these objectives:

1.  Promote data-freedom as a mainstream ideology.

2.  Legislation to allow user to migrate personal data from any platform
    > that collects it. (European Commission proposed a similar
    > legislation recently)

3.  Develop open-source tools, such as browser extensions, to enable
    > users to transfer data to their personal storage devices easily.
    > By returning data ownership to its creators, the data can be
    > revitalized and reconnected with each other.

-   **Rule subordinates to data**\
    When data can be migrated freely, users will have the choice to
    apply their preferred rules to it. In other words, try to
    be ***Object-Oriented***, avoid ***Procedure-Oriented*** development
    whenever possible. Data structures are relatively simple to reverse
    engineer. It is easier to reconstruct program structure with the
    knowledge of data structures.

-   **Granulation of user data**\
    Disentangling complex and intertwined datasets is challenging. It is
    preferable to design data in a granular manner, following the
    gene-like pattern, ensuring the information is relatively complete
    and independent. A piece of granular information is more manageable
    to process, transfer, and comprehend its relationship with other
    objects.

-   **inter-connections of data**\
    Cold data has limited value. While inter-connected, hot data
    possesses significantly more information. HTML is a perfect protocol
    to describe relationships, if not link could be unstable.

#### 2.2 Data Container

There is no need for a data center when individuals store their personal
data. Nevertheless, a user-friendly, affordable, and low-maintenance
container, such as the Pod advocated by the Solid Platform, is
essential. The owner of the storage device possesses the data and
determines the rules governing its usage. Compared to the Homomorphic
Encryption employed by certain blockchain technologies for data
security, the Leither approach is simpler and considerably more
efficient.

### III. Introduction to MiMei

Entropy is a crucial measure in both the universe and human society. The
second law of thermodynamics indicates the direction in which the
universe evolves, which is why entropy is often referred to as the Arrow
of Time. Erwin Schrödinger, in his book *What is Life*, asserts that an
organism essentially feeds on negative entropy.

Life stores information about its environment in genes. Heredity
preserves adaptations, while variation addresses environmental changes.
Life simply executes the instructions locked in its gene. Once human
developed self-conscious, the human brain started simulating adjustment
to its environment and processing the information. Richard Dawkins
coined the term **Meme** to describe the basic unit of this type of
information.

With the emergence of the internet, methods for processing information
have transformed, prompting the definition of a novel information unit
to accommodate these changes, which is called **MiMei** in this article.
Like Gene and Meme, MiMei contains a piece of information and the rules
applicable to it.

***MiMei flows among nodes to facilitate the decentralization of
centralized internet***.

#### 3.1 Implementation of MiMei

Based on **Sec 1.7 Summary of Design**, the following functions have
been implemented.

-   MiMei operation: create, manage, save, render

-   MiMei created with content\'s topic, each with a unique ID

-   User rights to set permissions on MiMei

-   Referential relationship among MiMei objects

-   MiMei can run or be saved on multiple nodes

-   MiMei data versioning, similar to Git, to solve problem of data
    > consistency

-   Synchronization among nodes in real-time or after backup

-   Support for database and file, stream media

-   File system is a special system type, based on database or file.

-   Docker-like data container in database and file system for every
    > MiMei

-   Heredity and evolution of MiMei by continuously backing up changes

-   User can also facilitate variation of MiMei by forking

-   Leither is implemented as MiMei container and a decentralized cloud
    > OS

#### 3.2 MiMei Features

  -----------------------------------------------------------------------
  **Feature**        **Detail**
  ------------------ ----------------------------------------------------
  Unique label       MiMei ID with version number can locate a piece of
                     information consistently and precisely

  Complex            Support database, file system, app and referential
  information        relationship to express complicated information.
  rendering          

  Unrestricted data  Support constructing relationship chain or group.
  migration          MiMei can flow freely among different nodes
                     according to user behaviors.

  Multi-dimension    User can add tags for MiMei and customize
  search             multi-dimension search result through different
                     service providers.
  -----------------------------------------------------------------------

### IV. MiMei Label

#### 4.1 MiMei ID

In the decentralized network of Leither, all resources, such as user,
node, content and application, are identified by a unique ID. User ID
and node ID are generated by encryption key pairs. Label ID of file or
data block is based on synopsis of its content.

There are two sets of labels for external and internal reference.

-   External label is MiMei ID\
    ***MiMei ID is the unique ID of a MiMei object and never changes
    after its creation**.* MiMei ID is generated based on its creator,
    associated application, MiMei type and label. MiMei data is stored
    in a file or database. Every time its data is changed, a new version
    of backup is created with an ID based on the synopsis of its
    content. The version number increments from 0. For convenience the
    latest version of backup data is called *last*.

-   Internal label is ID of synopsis of content\
    Every backup version of MiMei points to an ID of a file or database,
    which is generated by hashing synopsis of the content.

#### 4.2 MiMei Information

MiMei can be defined as:

type MiMeiInfo struct {

Author string //creator

AppType string //associated application type

Ext string //extension

Mark string //MiMei ID

Create time //timestamp of creation

Right uint64 //authorized rights

}

**MiMei data is accessed with MiMei ID and version number.** New version
is created while MiMei being edited or backed up. New version ID is
based on synopsis of its content, which is read-only after backup.
Historical MiMei data can be retrieved using MiMei ID combined with
version number.

### V. MiMei storage and association

MiMei support file system and database, meeting the data storage request
of most internet applications with its file system and database.

#### 5.1 Granulation of Information

MiMei granulates information. It is recommended to redefine the
following types of information with MiMei data type, aka Mimeimization:

-   Content can be searched.

-   Content for sharing among users.

-   Content that migrates among nodes.

Mimeimization can be executed beforehand, or on demand. The latter is in
fact a split. A new MiMei object detached from the original one. A
referential relationship keeps the tie.

User can use traditional development method if only its own data is
concerned. In this case, a standalone MiMei object is created for the
user or whole application.

#### 5.2 MiMei version

Both MiMei file system and database support versioning. Different data
version can be retrieved by unique MiMei ID.

-   *cur*: the current working copy, not yet backed up.

-   *last*: the latest backup copy, represents the newest confirmed
    > content

-   *release*: tested and ready for release

#### 5.3 Association of MiMei

It is difficult to represent complicated information with isolated data.
With unique label of MiMei, it is possible to establish relatively
stable relationship structure. Information can be saved in MiMei file or
database, with format defined by its associated application.

The relationship between MiMei objects is determined by references,
which include MiMei ID and reference count. Ordinarily the semantic
associations of data are interpreted by its associated application.
Leither system cannot access App data, so it is the App\'s job to
maintain the reference integrity by calling corresponding API.

Usually, MiMei object contains only granulated piece of information. The
complete and more complex content can be expressed through the
association of MiMei objects.

#### 5.4 File system

Originally file system was designed for mainframe where the number of
applications and volume of data is limited. File system increased
efficiency by saving applications from the task of storage medias. With
the development computer and internet, especially mobile internet, the
number of users, Apps and data volume all increased explosively. The
following shortcomings of file system began to appear.

1.  File name cannot precisely label a file

2.  Insufficient index information, only path and file information.

3.  One file might have multiple duplicated copies

In Leither, traditional file system is converted into MiMei framework
with the following improvements,

-   Unique label: Every data file has a unique ID generated from its
    > synopsis, as its label.

-   *cur* version: Current version as target of frequent data access.

-   Reference count: Do not copy data, increase count of reference
    > instead. Use unique label to access its data.

The directory structure of file system is also **granulated** (Ch6.1).
Directory that is indexed or shared with high hit count shall be
Mimeimized. Currently Leither manages file directory with JSON file. In
future database will replace JSON. Leither has comprehensive instruction
set and API for file and database operations.

### VI. Flow of MiMei Information

#### 6.1 Why information must flow

-   **The meaning of life is the spread of its information**\
    Gene and Meme are the carriers of information of life, whose purpose
    is to occupy time and space as much as possible. Similarly, the
    purpose of MiMei is the propagation of meaningful information.

-   **Hot data is more valuable**\
    File Coin conveyed a misleading concept that data is asset. Cold
    (static) data is in fact negative asset, for device, storage and
    bandwidth all cost money. The value of data appreciates only when
    data turned into traffic (get hot). Services to sustain data is
    mainly the operations of save and retrieval. However, services to
    support data traffic involve the creation or collection of content,
    save, propagation, render, value realization and business
    development. Every time some content is displayed, value is created.
    The number of renders and the value created each time determine the
    value created by the content.

-   **The flow of information is a business requirement**\
    On internet, data backup, error tolerance, load balance and elastic
    cloud services fundamentally all require the flow of information.

#### 6.2 The procedure of information flow

Information flows under two different situations:

1.  **Voluntary Save**\
    Usually when the MiMei data is highly valuable, it can be saved for
    future benefit.

2.  **On demand Save\
    **When a node is in danger of overloading, it may ask other nodes to
    help in backup, error tolerance, load balance, etc.

Procedure of save:

-   **Mimeimize information\
    **After information is MiMeimized, information can be backed up as
    file blocks.

-   **Duplicate data over node network\
    **With proper authorization, copy MiMei block from one node to
    another

-   **Update routing information of MiMei\
    **Terminal content consumers first obtain the routing information of
    MiMei and then access its data. The routing information includes
    node information and the latest changes of MiMei. Once the routing
    information is updated, user can get the most recent MiMei data.

### VII. Leither Application System

Leither API supports 40+ development languages, especially HTML5 because
majority of the frequently used applications are developed in HTML5,
including website, App, Applet, etc. HTML5 lacks standards for backend
data processing and business logic. Leither supplements the missing
functionalities when building cloud applications with Html5, including
user authentication, application system, cloud file system and database,
decentralized domain name resolution, as well as fault tolerance and
load balance.

#### 7.1 Leither Solution

-   **Business logic executed in front-end**\
    Business logics shall be executed in the front-end by default,
    unless necessary. The advantage is that Leither application
    development will be similar to single page application development,
    simple and easy.

-   **Thin node**\
    PC and mobile phone are not suitable for long term stable service.
    Server is maintenance heavy and expensive. Router and TV box are
    usually proprietary system with restricted accessibility. NAS and
    other customizable hardware are preferred devices. Less powerful CPU
    is ideal to support fine granulated Leither applications.

-   **Special optimization**\
    Because Leither node device is intended to be maintained by layman,
    node setup and application installation must be dummy proof. Node is
    optimized for home network, with the unique decentralized domain
    name resolution of Leither system, Leither can provide service
    similar to commercial cloud server.\
    \
    Leither node can also be optimized to support batch execution and
    back-end tasks.

-   **Mimeimization of application**\
    Leither offers a comprehensive application ecosystem, where
    applications within the system are essentially a form of MiMei.

#### 7.2 Advantages of Leither Cloud

-   **Leither is ideal to build cloud service**\
    Leither provides authentication and authorization services,
    application framework, cloud-based file system and database,
    decentralized domain name resolution, as well as fault tolerance and
    load balancing. Both application and data within Leither system are
    of MiMei type, which can freely migrate among nodes. Combined with
    the load balancing mechanism, this makes Leither an ideal choice for
    building cloud service platforms.\
    \
    The best practice of Leither App development is to execute business
    logic in the frontend, therefore greatly reduce the load on the
    backend. The same server running Leither can outperform conventional
    platform 100 times or more.

-   **Simple Application Development**\
    Knowledge of HTML5 is sufficient to do most development work of
    cloud services, website, App and applet independently. The workflow
    is almost identical to a HTML single page application.

-   **Extremely low demand on system resources**\
    Leither is developed with weak CPU and small memory devices in mind.
    Leither applications consume very little system resources, because
    data and business logic are executed in the frontend, usually
    user\'s browser or App.

-   **Low bandwidth cost**\
    For traffic heavy applications, traffic load can be efficiently
    balanced by nodes in users\' home network, which is cheap. The cost
    of traffic is less than 1/10th of regular web service.

-   **Elastic service**\
    Both system application and data support container service, like
    that of Docker. Both file and database support versioned backup. One
    physical machine can support multiple users, multiple applications
    and their datum. Application and data can flow among different
    physical machines freely without changing domain name or link
    address.\
    \
    Elastic service of Leither is enabled by decentralized domain name
    resolution and MiMei mechanism.

### VIII. Leither credential system

During the exchanges of information and services between user(node) and
user(node), there must be a mechanism for quantified settlement. In real
world fiat currency is the tool for settlement. In Leither network it is
credit record that serves the purpose.

Every node can grant other nodes a line of, allowing them to use its
services. The difference in service volume between nodes is essentially
an IOU issued by the borrower to the lender, which can also be
understood as cryptocurrency issued by the borrower.\
\
Cryptocurrency issued within a group is in fact a certificate of
contribution made by its members. The contribution could be anything
valuable to the group, such as resource, money or work done for the
group. This certificate can be used to exchange for services or other
valuables from its members.

-   **Object of credit**\
    The highest credit standard in real life is fiat currency. Within
    the Leither network, the mutual assistance services between nodes
    can serve as a reference standard, for example bandwidth, storage or
    CPU. Their order of importance is bandwidth \> storage \> CPU.
    Because of the characteristics of node network, bandwidth is the
    preferred standard of settlement between nodes. All the other forms
    of services can be converted into bandwidth with a given
    coefficient. Fiat currency can also be used as standard of credit.

-   **Individual credit model**\
    is a type of User-to-User (P2P) L/C or IOU.


-   Letter of Credit (L/C) L/C can be understood as the amount of
    services allowed to use before payment, whereas IOU is the amount of
    services that have been used without payment. A L/C includes
    information of the authorizer, the authorized, unit of credit,
    maximum credit, signature of authorizer.

-   An IOU includes information about lender, borrower, unit of credit,
    amount, signature of borrower. For borrower, IOU is equivalent to
    issuing **Cryptocurrency**.


-   **How to buildup credit**\
    If there were connection between nodes, both parties have a
    foundation of mutual trust and can authorize varying amounts of
    credit based on the proximity of their relationship, credit can be
    adjusted periodically based on their interaction.\
    \
    Each node can provide useful services. Credit can be earned by
    proactively providing valuable service to the others, including data
    backup, load balance, routing, search, or mediator service. Users
    can also acquire more credit through 3rd party.
