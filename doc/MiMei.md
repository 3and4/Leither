Leither&弥媒
========  
互联网日益中心化，少数寡头垄断了用户数据，通过规则影响和控制用户行为，并获取巨额利益。打破寡头对互联网的垄断成为民众越来越迫切的需求。  
区块链，IPFS、Solid、dfinity等都是为了这个目标提出来的解决方案，传统中心化互联网发展积累多年，已经高度成熟，目前的去中心化方案和传统方案相比都不太完备。

Leither是一个去中心化的云操作系统，旨在能用去中心的方式肢解和重构互联网上大部分用户相关的核心业务  
Leither实现了构建传统互联网应用所需要的用户安全体系、文件系统、数据库、应用体系、网络体系等模块，在体验和协议上尽可能与传统兼容。

弥媒是系统中的基本对象，用于描述数据和应用。类似于基因和弥母，弥媒是有独立内容或功能的最小单位  
弥媒之间有引用关联，可以构建出复杂的互联网应用和服务  
弥媒可以在节点之间流动，从而实现去中心化的目的。   

### 一、背景——互联网日益中心化
互联网日益中心化，破除垄断，构建去中心化互联网是当前时代的迫切需求

#### 1.1 大数据控制了整个社会  
+ **互联网正在逐步成为社会的核心**  
互联网在当今社会越来越重要  
互联网一步步重构了整个社会  
2019年，中国线上交易额250万亿  
  
+ **互联网通过大数据影响社会**  
大数据的本质是收集数据和控制规则  
进而影响和控制个人行为，并因此创造暴利的商业模式  
  
#### 1.2 寡头控制了大数据  
+ **最核心的数据已属于寡头**  
社交相关的数据归腾讯， 电商相关的数据归阿里  
其他重要的数据也分别落入两者相关的系统内  

+ **通过规则制造了层层壁垒**  
通过系统应用规则进行保护; 通过现行的法律法规保护  
通过生态潜规则进行保护  

+ **利益和机会属于寡头**  
90%的利润被AT系所获取，财富集中速度越来越快  
稍大一些的创业或投资机会，几无例外都落入了寡头手中  
阿里和腾讯市值超过10万亿人民币  

#### 1.3 绝大多数人被困数据牢笼
+ **普通用户，任人宰割**  
没有隐私，是一系列的数据标签  
没有自主，是可以被引导控制的数学模型  

+ **中小企业，进退失据**  
在大数据面前，每个企业都变成了透明体，机会和利润被压到了边缘  
腾讯平台上的游戏厂商分成10%都不可得；大部分电商处在平衡线上  
 
+ **投资环境，机会丧失**   
新投资机会出现的时候，最终都不可避免的落入阿里和腾讯这样的寡头手中  
社会内卷化，我们很久没有听过普通人逆袭的故事了  

#### 1.4 全人类的智慧汗水结晶，被少数公司攫取果实  
+ **美国军方发明了互联网，没有控制获利**  
互联网源自美国军方的阿帕网（国防部高级计划局网络）   
互联网的架构从创建那天就是开放的，去中心化的   

+ **运营商构建了国内网络，没有控制获利**   
国家通过国企业投入巨资，在全国范围通水通电通路通网  
国企运营商不计回报构建了世界上最大的互联网  
投入的4G、5G基站都是占世界一半以上  

+ **少数公司通过几个应用控制了互联网**    
国内互联网被阿里、腾讯系控制  
海外互联网被脸书、谷歌、亚马逊、微软等公司控制  

#### 1.5 结论  
互联网日益中心化，少数寡头垄断了用户数据，通过规则影响和控制用户行为，并获取巨额利益。打破寡头对互联网的垄断成为民众越来越迫切的需求。

### 二、历史需求汇总
我们的目标是构建去中心化互联网
在分析问题提出方案之前，我们需要整理一下计算机和互联网发展各时期的需求，以及当时为解决这些需求而提出的方案模型
这些需求如果不能满足，新的方案就是不完备的，没有办法替代现有的方案  

#### 2.1 文件——数据载体  
计算机诞生初期，最早的数据需求是应用  
后续随着应用功能复杂，应用也产生了大量的相关数据    
在unix中，设计了文件这样的数据载体，通过文件系统管理数据。  
一切皆文件，所有信息相关的设备都通过文件表达  
文件接口包括打开，读，写，关闭这样通过的数据接口。  
同时为了管理文件和存储介质，设计了文件系统

#### 2.2 MIME——数据类型  
随着应用的增加，需要表达文件和应用之间的关联。  
在dos中有了扩展名的概念，这时候还是比较粗的文件类型表达  
到了windows中，扩展名已经可以和应用进行绑定  
在电子邮件的MIME中，对文件（附件）类型有了严格的定义    
目前的浏览器都对这个定义进行了支持

#### 2.3 数据库——数据间关联  
为了处理复杂的数据关系，数据库应需求而生。  
数据库还是封闭的数据集合，是一个体系内的关联关系   
缺点：数据存放在封闭空间，不能自由流动，限制了价值最大化  

#### 2.4 HTML——数据跨点关联  
www的出现，从规则上定义了数据和数据之间跨文件跨主机的关联  
互联网诞生，信息极大爆炸  
数据可以跨节点关联，突破了主机空间对数据的限制  
门户网站充分利用了这一特征，获得的蓬勃的发展  
缺点：基于链接的关联不够精确稳定。

#### 2.5 google——数据检索  
信息大爆炸，产生了检索的需求  
Goolge提供了一种根据关键字进行内容检索的服务  
给数据附加了标签的属性  
缺点：日益垄断，日益封闭，数据获取规则被干扰控制

#### 2.6 facebook——社交模型  
最早的社交通讯工具是icp、msn、qq  
解决了人和人之间的实时消息通讯问题  
twitter提供了一个更快速简短的信息分享平台
myspace到facebook一个更系统的社交模型诞生了  
这个模型分四部分：  
|序号|名称|说明|
|--|--|--|
|1|我|主要是管理用户个人相关的的信息|
|2|朋友|用户的关系链|
|3|应用和服务|用户使用的应用和服务 |
|4|消息流|用户和外界的互动区，互动对象是朋友、应用、服务等|


这个系统模型非常成功，后续出的互联网平台产品都在向这个架构靠拢  
典型代表是微信，微博，阿里来往，支付宝，一些企业的OA系统，都或多或少有这个框架的影子。

缺点：互联网日益垄断，日益封闭，规则被干扰控制

#### 2.7 需求汇总  
基于破除垄断、构建去中心互联网目标，设计原则如下：  
+ **一切以用户为中心**  
用户构建、用户控制、用户享有  
a new birth of freedom and that  internet of the people, by the people, for the people 

+ **规则自由开放**  
网络底层协议是开放的；  
底层操作系统规则是开放的  
规则开放，用户才有选择权，用户才有控制权

+ **数据自由流动**  
在用户许可、安全可靠的情况下，用户数据可以从一个平台流动到另一个平台，可以从一种形式转换成另一种形式。  
数据能自由流动，第三方才可以为用户提供丰富多采的应用和服务  


结合计算机和互联网的历史，我们列出更细的设计需求如下：
+ 支持文件，文件系统和数据库  
+ 支持数据和应用的关联  
+ 支持跨节点的内容关联  
+ 支持跨节点的检索  
+ 支持跨节点数据流动  
+ 支持去中心化的社交模型  
+ 支持轻量级的容器系统
+ 支持简洁的应用开发方式
+ 支持体验要和当前互联网兼容
+ 合法合规兼容当下的社会制度  

### 三、问题分析  
首先，我们要知道问题是怎么产生的，也就是数据是如何被垄断的。 

#### 3.1 相关概念  
这个问题涉及到两组概念：数据和规则;耦合和解构

+ **数据和规则**  
数据结构和算法，是计算机科学中重要的两个概念;  
数据结构描述了数据存放的方式，通常描述内存，数据库或文件中的对象
算法描述了怎么处理这些数据对象，打开，生成，存取，保存，检索，展现等  
下文中的**数据**是遵守同一数据结构的数据对象， **规则**指的是使用这些数据的功能模块或应用。

+ **耦合和解耦**  
耦合是各部分组合在一起，互相都有复杂的关联关系，使各部分无法拆分，形成一个整体  
解耦是理顺各模块之间的关系，对之进行归纳整理，使各部分之间的关系简洁明了，降低相互的依赖  
如果使各部分形成一个整体，就强化各模块之间的关联关系  
如果使问题变的更简单高效，就使关联关系变的简单明了便于优化  
如果使各部分脱离一个整体，就弱化简化各模块之间的关联关系

#### 3.2 数据访问的两种形式  
我们在访问数据进行操作的时候，通常有两种形式：  
+ **数据->规则**  
用户可以自由接触到数据，然后调用数据相应的规则
windows中，用户可以选择不同的程序打开同一类型文件
电子邮件中，用户可以选用不同的程序打开附件
同一个网址，可以选用不同的浏览器打开。
在这种情况下，用户可以可以自由选择规则。不被单一规则所限制。

+ **规则->数据**  
用户访问数据时，数据被特定的规则保护，限定了用户只能通过一套规则进行访问。  
用户在电商平台购物，登录电商平台，查找商品，进行购物  
用户在社交媒体上访问社交信息  

#### 3.3 数据垄断的形成  
在传统的操作系统上，文件可以在各终端和平台之间自由存取复制  
互联网初期，用户数据访问是通过统一规则是互相开放的。  
谷歌，微软等大平台都提供了访问用户数据的api.  
最早的数据垄断来自于facebook,facebook利用各大公司提供的关系链数据构建了自己的社交媒体，然后断绝了第三方对用户数据的访问，用户数据被限制在了社交媒体内部。  
随后亚马逊苹果等公司纷纷建立了自己的数据王国。  
国内互联网巨头，qq起家的时候没有采用通用的通讯协议。  
百度通过广告封杀淘宝的时候，阿里系对百度也进行了内容封锁。  
腾讯阿里大战的时候，也互相进行了封杀。  
互联网日益孤岛化，堡垒化，用户的数据被割裂保存在不同的寡头平台。  


#### 3.4 垄断的破除  
**平台公器化**  
	计算机刚诞生的时候，为了便于管理设备和应用，操作系统诞生了。  
	随着pc的发展，微软借助操作系统控制了整个pc的应用生态。在办公软件和浏览器等领域都利用操作系统扼杀了竞争对手。  
	另一边，开源社区利用linux把操作系统变成了公共平台。是一次平台公器化的尝试。  
	类似的例子还有手机操作系统android，芯片开源指令集架构RISC-V。  
	对于存放用户数据的平台，应该有开源的平台架构。  
	当前互联网最应该公器化的是**社交体系**和**内容检索**两个部分  
	阿里、百度、谷歌、亚马逊是通过内容检索控制用户数据和行为  
	腾讯、脸书是通过社交体系控制用户数据和行为  

**数据能自由流动**  
	在最早的操作系统里，用户的文件是可以在不同主机不同系统之间自由复制。
	数据能自由流动应该成为大家的共识，也应该成为基本的设计规范。  
	我们可以做三个工作：  
+ 广为宣传使之成为主流的价值观;  
+ 促进各国制定法律法规成为行业规范，对于保存在用户平台上的数据，必须支持用户导入和导出的接口;  
+ 是发动类似浏览器插件这样的开源项目，组织研发人员把用户数据从寡头的平台上导入到用户的空间，数据到了用户空间，便可以对用户的数据复活及再关联。  

**数据和规则分离**  
	数据能自由流动的情况下，用户可以选用不同规则去操作数据

**数据和数据关联**  
	孤立的数据价值是有限的，不同数据之间应该能够互联关联，能够表达更复杂的内容

**用户数据颗粒化**  
	对于耦合在一起的大规模数据，我们是没办法单独剥离的。在设计上应该尽可能的让数据颗粒化。需要比文件更小更灵活的数据对象。描述内容的基本单位

#### 3.5 数据容器  
传统互联网的数据存放着十亿级用户规模的数据，需要构建大量的数据中心。  
大的数据中心超过十万台服务器，每个业务通过各种技术方式进行集群处理。  
这些数据中心需要大量的运维人员进行日常管理。   
对于个人数据，规模和复杂度都要低的多，没必要再使用这样的数据中心模式。  
数据回归个人，需要有新的机制存放用户数据，需要有轻量级的云操作系统。
类似Solid中推荐的Pod，数据要回归用户，首先需要一个容器来存放用户数据   
这个容器能够最好是存放在用户的个人空间，成本要可控，轻量级免维护的。  
普通人不具备很强的专业知识


#### 3.6 参考模型  
**文件**  	
	文件是记录在在存储介质上的程序和数据的集合。  
	文件有一定的格式，我们通常用数据结构来描述文件结构  
	操作系统中通常有文件系统管理文件
	操作系统通过树状目录来存放文件
	
**Web（World Wide Web）**  
  	在目前的Web标准中 html5表达页面内容，CSS定义页面的风格排版布局，JS执行页面的逻辑。可以相对完整存放并表达数据。页面间通过链接进行关联，可以组织出大规模的高复杂度的信息内容。  
	缺点：  
	只定义到了前端，缺失了后端功能。寡头通过后端控制了整个用户数据  
	页面间的关联是基于链接的，边接中的主机和位置是不稳定的  
	没有对主机，用户等元素进行定义

**基因（GENE）**  
  	基因是生命信息的最小片段，DNA或RNA保存一段完整的遗传物质  
  	遗传物质记录了生物对环境的适应信息  
  	这些信息的执行规则保存在生物体内。
  	生物通过遗传的方式复制之前的环境应对，通过变异加筛选的方式适应新的环境变化  
	缺点：  
	就是每迭一次需要一代生命周期。  
	信息量相对有限，不能表达复杂的变化，改造环境能力有限  

**弥母（MEME）**  
	弥母是道金斯在自私的基因一书提出，和基因类似，弥母是文明信息的最小因子。  
	弥母记录了人类对这个世界的所有知识
	这些知识保存在人类的大脑里，在人和人之间传播，也可以保存在各种媒介上  
	人类可以在大脑中对外界进行建模，在和外界的互动过程中，每一次反省就是一次迭代进化。“吾日三省吾身”相当于每天进化了三次。相对基因，弥母的进化速度有了数量级的提高。

 ### 四、弥媒（MIMEI）的由来
信息熵是宇宙和人类发展过程中重要的一个物理量   
熵增定律为宇宙演变指明了一个方向。熵也被称为时间之箭    
薛定谔在《生命是什么》中提到：人活着就是在对抗熵增定律，生命以负熵为生    
生命把对环境的信息记录在了基因（gene）里
通过遗传适应之前的环境，通过变异筛选适应环境的变化  
人类有了意识之后，可以在大脑中对环境进行模拟调整，在大脑中处理相应的信息    
这些信息的基础单位被查德-道金斯称为弥母（meme）  
计算机诞生以后，信息的处理方式发性变化，需要一种新的基本单位来描述信息  
这种新的信息单位，我们称之为弥媒（mimei）  

弥媒的相关功能如下：    
弥媒通过应用生成、管理、存取、展现、传播  
弥媒基于内容主题创建，每一个弥媒都有一个唯一ID  
用户可以设置弥媒的各种权限  
弥媒和弥媒之间通过引用进行关联  
弥媒可以运行保存在一个或多个节点上  
象Git一样，弥媒有版本的概念  
版本主要是解决数据的一致性问题  
节点之间的数据库可以实时同步变化  
也可以在生成确定版本的时候再进行同步   
弥媒目前描述的数据类型有数据库和文件，后续会支持流  
文件系统是特殊的类型，基于数据库和文件完成  
每个弥媒都可以有独立的数据库和文件系统空间  
弥媒通过不断的修改备份完成内容的继承进化  
用户也可以用分支的方式完成弥媒内容的变异  
Leither是我们实现的一个弥媒容器系统，同时也是一个云中心化的云操作系统

<a href="../api/MiMei.md"> 点击查看弥媒相关的API</a> 


#### 4.2、弥媒的数据结构
**弥媒ID**
在Leither定义的去中心化网络生态中，所有的资源都是用ID来描述的
包括不限于：用户、节点、内容、应用等  
用户和节点的ID是基于公私钥产生的  
一个文件和数据块的的标识ID是基于内容取摘要产生的  
弥媒的ID是<a href="../api/MiMei.md#MMCreate"> 弥媒创建</a>时产生  

**弥媒信息**  
创建时基于以下参数:
|名称|说明|
|--|--|
|Author|创建者|
|AppType|关联应用|
|Ext|扩展类型|
|Mark|标识|
|Create|创建时间|
|Right|权限|

弥媒的信息用以下结构表达：
```golang
type MiMeiInfo struct {
	Author  string    //作者
	AppType string    //应用类型
	Ext     string    //扩展信息
	Mark    string    //标识
	Create  time.Time //创建时间
	Right   uint64    //权限
}
```

### 五、弥媒的标识
我们分析一下，各架构下资源的描述方式：  
在WWW中，是通过链接描述资源的，链接格式如下:  
http://:服务器地址[:端口号]/路径/文件名[参数=值]   

在常见的操作系统中是通过路径来描述资源的,路径格式如下:  
盘符:/路径/文件名  
/路径/文件  
在上面的描述方式中，所有的环节都是不确定的，服务器地址，端口号，路径，文件名，参数都是可变的。
现行的方式都无法准确描述一个资源

在ipfs协议中的基于对文件取摘要的方式生成唯一ID。      
这个唯一ID是基于内容的。内容改变就会生成新的ID  
这种不确定性不利于正常的操作和关联  
在弥媒体系中我们提出了新的方案：  

**弥媒ID是稳定的**
弥媒ID是基于创建者、关联应用、弥媒类型、弥媒标识产生的。
创建之后，不管怎么操作，弥媒ID都是不变的

**版本机制**
弥媒在编辑的过程中，备份的时候生成版本，这个版本对应的数据是基于内容摘要的。版本备份之后内容确定不再修改。  
版本相关的操作有：备份、恢复、发布、清除、查询状态等  

**特殊版本**
cur代表当前修改中未备份的版本，代表在实时修改  
last代表最后一备份的版本，代表最新的固化内容
release代表审核过，确定可以对外发布的版本。

有了版本机制，我们便可以通过唯一ID获取不同版本的弥媒数据。
可以获取最新的数据  
可以获取最后的备份  
也可以获取最稳定的数据  
也可以获取任一指定的版本

### 六、弥媒文件
本章节描述系统中的文件对象 

#### 6.1、文件对象  
**弥媒当前文件**  
指的是弥媒文件的当前版本，当前文件是可读可写的  

**弥媒版本文件**  
指的是弥媒的备份文件，备份文件是只读的    
由当前版本文件备份生成， 同系统动态分配一个版本号指向这个备份文件。  
  
**Mac文件**  
以文件内容的mac id为标识的文件，Mac文件是只读的  
Mac文件可以被弥媒对象引用  
弥媒的版本文件就是一个版本号指向一个mac文件
Mac文件也可以由临时文件转换生成

**临时文件**  
由MFOpenTempFile创建生成，进行数据写操作之后。
由MFTemp2MacFile转换为mac文件  

**弥媒文件系统**    
系统内置的一个弥媒类型,描述一个弥媒的文件系统  
创建时系统代码如下：  
```golang
	fs.MTreeID, _, err = mc.CreateMiMei(
		dir.OwnerID,
		"",                 //需要一个系统app的id,目前用空表示
		api.MM_EXT_TREE,    //		"mmtree",
		api.MM_Auto,        //自动生成唯一id
		api.MM_File,        //内部保存为一个文件
		api.DataRight_Default_Group,
		false,
	)
```  
在webdav目录下，可以通过设置扩展名为.mmfs的配置文件配置或自动生成弥媒文件系统。  
可以通过MFOpenByPath打开文件系统中的对象  
注：目前是通过json文件的方式实现的，在文件比较多的时候性能较差，正在规划变成数据库方式实现  

**操作系统文件和目录**  
可以把操作系统中的文件目录link到webdav目录下
这时候可以通过MFOpenByPath函数打开操作系统中的文件和目录

**弥媒根目录**  
webdav目录是节点对外展示的弥媒总入口。  
可以把操作系统中的文件或目录link到这个目录
也可以生成一个配置文件指向节点内的一个弥媒对象  
  
#### 6.2、打开文件
**弥媒文件**
创建弥媒文件  
接口函数MMCreate，参数类型为api.MM_File
返回值为弥媒id  
打开弥媒文件  
接口函数为MMOpen(sid, mid, ver string) (string, error)
返回值为会话id,通过会话id可以操作文件内容
ver如果是"cur",可以进行写操作。  
其它版本文件为备份过的文件，只能进行读操作。  

**通过路径打开**    
这种方式打开的是弥媒文件系统或操作系统文件系统中的文件  
相关API：  
MFOpenByPath(sid, fsid, path string, flag int) (string, error)

**打开Mac文件**   
打开挂在某个弥媒下的mac文件  
相关API：  
MFOpenMacFile(sid, mid, fileid string) (string, error)

**打开临时文件**   
打开一个临时文件，用于读写操作
MFOpenTempFile(sid string) (string, error)  
读写操作完成之后转换为Mac文件    
MFTemp2MacFile func(sid, mid string) (string, error)  

**关闭文件**  
除临时文件外，其它文件操作完成后都需要关闭文件  
考虑到远程操作文件经常会有异常掉线情况。   
操作对象的会话id在超时情况下都会自动关闭文件  


#### 6.3、操作文件  
**对象方式读写**  
MFSetObject(fsid string, obj interface{}) error  
MFGetObject(fsid string) (interface{}, error)

**字节数组读写**  
MFSetData(fsid string, data []byte, start int64) (count int, err error)  
MFGetData(fsid string, start int64, count int) ([]byte, error)  

**查询状态**   
MFGetSize(fsid string) (int64, error)
MFStat(fsid string) (*FileInfo, error)
MFIsExist(fsid, fileid string) (bool, error)  
MFGetMimeType(fsid string) (string, error)

**目录操作**  
MFReaddir(fsid string, count int) ([]*FileInfo, error)
  
**文件系统**  
FSFind(sid, mmfsid, path string) (*FindResult, error)  
FSMkDir(sid, mmfsid, path string) error  
FSRemoveAll(sid, mmfsid, path string) error  
FSStat(sid, mmfsid, path string) (*FileInfo, error)  
FSRename(sid, mmfsid, oldpath, newFullName string) error

**其它**  
MFTruncate(fsid string, size int64) error  
MFCopy(fsid, dst, src, srcVer string) error  


### 七、数据库
数据库的底层有两种，一种是基于LevelDB,一种是基于BoltDB。  
两种数据库都进行过底层改造。  
LevelDb用于当前版本，可写可读，一致性是基于时序。  
写事务时会检查改动的部分是否有第三方同时修改过相应的内容。  
有改动则提示错误，通知调用端重新执行相关操作，避免写冲突。
BoltDb用于历史版本，只用于读。

Api参考Redis  
可以操作字符串，哈希表，列表，集合，有序集五组数据类型  
支持事务    

#### 7.1、事务 
Begin(dbsid string, timeout int) error  
Commit(dbsid string) error  
Rollback(dbsid string) error 

#### 7.2、字符串 
Set(dbsid, key string, value interface{}) error  
Get(dbsid, key string) (interface{}, error)  
Del(dbsid string, key ...string) (int64, error)  
Incr(dbsid, key string) (int64, error)  
IncrBy(dbsid, key string, delta int64) (int64, error)  
Strlen(dbsid, key string) (int64, error)  
  
#### 7.3、哈希表   
Hmclear(dbsid string, key ...string) (int64, error)  
Hdel(dbsid, key string, field ...string) (int64, error)  
Hlen(dbsid, key string) (int64, error)  
Hset(dbsid, key, field string, value interface{}) (int64, error)  
Hget(dbsid, key, field string) (interface{}, error)  
Hmget(dbsid, key string, fields ...string) ([]interface{}, error)  
Hmset(dbsid, key string, values ...FVPair) error  
Hgetall(dbsid, key string) ([]FVPair, error)  
Hkeys(dbsid, key string) ([]string, error)  
Hscan(dbsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)  
Hrevscan(dbsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)  
HincrBy(sid, key, field string, delta int64) (ret int64, err error)  

#### 7.4 列表
Lpush(dbsid, key string, value ...interface{}) (int64, error)  
Lpop(dbsid, key string) (interface{}, error)  
Rpush(dbsid, key string, value ...interface{}) (int64, error)  
Rpop(dbsid, key string) (interface{}, error)  
Lrange(dbsid, key string, start, stop int32) ([]interface{}, error)  
Lclear(dbsid, key string) (int64, error)  
Lmclear(dbsid string, keys ...string) (int64, error)  
Lindex(dbsid, k string, index int32) (interface{}, error)  
Llen(dbsid, k string) (int64, error)  
Lset(dbsid, k string, index int32, value interface{}) error  

#### 7.5 集合
Sadd(dbsid, key string, args ...string) (int64, error)  
Scard(dbsid, key string) (int64, error)  
Sclear(dbsid, key string) (int64, error)  
Sdiff(dbsid string, keys ...string) ([]string, error)  
Sinter(dbsid string, keys ...string) ([]string, error)  
Smclear(dbsid string, key ...string) (int64, error)  
Smembers(dbsid, key string) ([]string, error)  
Srem(dbsid, key string, m string) (int64, error)  
Sunion(dbsid string, keys ...string) ([]string, error)  
Scan(dbsid string, begin, match string, count int, inclusive bool, tp byte) (keys []string, err error)

#### 7.6 有序集
Zadd(dbsid, key string, args ...ScorePair) (int64, error)  
Zcard(dbsid, key string) (int64, error)  
Zcount(dbsid, key string, mins, maxs int64) (int64, error)  
Zrem(dbsid, key string, members ...string) (int64, error)  
Zscore(dbsid, key, member string) (int64, error)  
Zrank(dbsid, key, member string) (int64, error)  
Zrange(dbsid, key string, mins, maxs int) ([]ScorePair, error)  
Zrangebyscore(dbsid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error)  
Zremrangebyscore (dbsid, key string, mins, maxs int64) (int64, error)
Zrevrange(dbsid, key string, start, stop int) (ret []ScorePair, err error)  
Zrevrangebyscore (dbsid, key string, mins, maxs int64, offset int, count int) (ret []ScorePair, err error)  
Zmclear(dbsid string, key ...string) (int64, error)  
Zclear(dbsid, key string) (int64, error)  
ZincrBy(dbsid, key string, delta int64, member string) (ret int64, err error)  


### 七、应用和数据


### 八、弥媒的关联性
孤立的数据不能表示复杂的信息  
弥媒有了唯一标识，可以构建相对稳定的关系结构  
Html中
MMAddRef(sid, mid string, fileids ...string) (int, error)  
MMDelRef(sid, mid string, fileids ...string) (int, error)  
MMGetRef(sid, mid string) (ret map[string]int, err error)  




### 九、弥媒信息流动

### 十、冗错、均衡和弹性伸缩

### 十一、群和共识机制

### 十二、信息检索

### 十三、社交模型

### 十四、共识机制

### 后记

  

### 2.4、权限  
MMSetRight(sid, mid, member string, right uint64) error  



