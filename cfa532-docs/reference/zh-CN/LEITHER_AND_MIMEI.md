# Leither 与弥媒技术基础

**维护位置：** `cfa532/Leither/cfa532-docs/`；2026-10-16 从 LifeAlbum 迁入。
**翻译状态：** 保留原有中文版本；后续技术更新以[英文参考](../LEITHER_AND_MIMEI.md)为准。

**状态：** 与具体应用无关的技术参考
**核对版本：** Leither CLI V0.24.02，API 合约版本 1
**源码参考：** `github.com/3and4/Leither`，2026 年 8 月
**目的：** 说明 Leither、弥媒、存储类型、版本、分发、身份和应用与节点之间的边界。

## 1. 核心结论

Leither 不是单纯的 IPFS 节点，而是一个轻量级的数据和应用容器系统。它把稳定的数据身份、版本化存储、应用、身份权限和兼容 IPFS 的分布式网络组合在一起。

系统的核心存储单元是**弥媒（MiMei）**。弥媒对外有固定不变的 `mid`，每个已提交版本再指向不可变的内容。这使应用既能长期引用一个稳定对象，又能利用内容寻址实现完整性校验、去重和分布式获取。

弥媒主要有两种数据类型：

- **文件型：** 保存一个带版本的文件或目录树，底层内容寻址。
- **数据库型：** 保存带版本的键值数据库，提供更复杂的数据操作。

对于强调耐久、便携和文件化存储的应用，文件型弥媒是已经可以用于生产的选择。数据库型弥媒提供更丰富的数据操作，但实现和恢复复杂度明显更高。不同应用应根据自身需求以及对应 Leither 子系统的成熟度选择存储模型。

## 2. Leither 的系统角色

一个 Leither 实例可以同时承担：

- **容器：** 保存弥媒数据和 Leither 应用。
- **存储节点：** 保存不可变数据块、文件型弥媒目录树、版本和副本。
- **服务节点：** 向获准的客户端提供弥媒数据。
- **支撑节点（Provider）：** 向网络声明本节点可以提供某个弥媒或 IPFS 对象。
- **网络节点：** 通过 DHT 和 IPFS 兼容协议发现节点和支撑者。
- **身份端点：** 使用节点/用户密钥、证书、通行证、会话和权限。
- **应用端点：** 保存并运行前端资源以及可选的应用逻辑。

IPFS 负责不可变内容块和分布式获取；弥媒在其上增加稳定身份、可修改当前态、版本、引用、权限、应用和发布语义。

Leither 有两个相关但不同的文件系统界面：

- 节点级 **Files 文件系统**是一个以普通根路径寻址、通过 `Files*` 操作的 IPFS 可变文件系统。
- **文件型弥媒文件系统**属于一个稳定的 `mid`，先打开指定版本，再通过返回的弥媒/文件句柄和 `MF*` 操作。

节点 Files 文件系统可以用来组装目录 DAG、执行 flush，再把根 CID 赋给文件型弥媒；它是构建工作区，不是弥媒本身的身份或版本边界。

## 3. 弥媒身份

### 3.1 `mid`、版本与 CID

| 标识 | 含义 | 内容变化时是否变化 |
|---|---|---:|
| 弥媒 ID（`mid`） | 一个逻辑数据对象的稳定身份 | 否 |
| 弥媒版本 | 该对象某个不可变的已提交状态 | 产生新版本 |
| CID 或 Mac ID | 不可变内容或内容树的地址 | 是 |

应用对外应使用 `mid` 和应用对象 ID；CID/Mac ID 是其下层的内容身份。

### 3.2 唯一创建与确定性创建

`mimei create` 缺省使用 `--mark "{{auto}}"`，每次生成新的弥媒。使用非自动 mark 时，相同创建参数可以再次定位同一个弥媒身份。

应用可以有意识地使用两种方式：

- 需要重复定位的单例资源使用带命名空间的确定性 mark。
- 身份必须唯一的独立对象使用自动 mark。
- mark 只作内部身份用途，不作为用户可见名称。

## 4. 弥媒版本

- **`cur`：** 可读写的工作版本。
- **`last`：** 最近一次备份提交的版本，普通读取使用它。
- **`release`：** 主动选定的稳定发布版本。
- **数字版本：** 不可变的历史备份。

标准写入流程为：

```text
打开 cur -> 完成同一业务操作的全部修改 -> backup -> last 前移
```

`backup` 是提交边界。文件写入成功不代表普通读者或其它节点已经能看到完整的新状态。应用应把一次逻辑操作的内容和元数据放在同一次备份中，不应因自动 flush 产生半完成的中间历史。

`restore` 把历史版本复制回 `cur`，不会自动删除历史；`release` 选择稳定版本；`delver` 按保留策略删除指定历史版本。

## 5. 文件型弥媒

### 5.1 文件型弥媒的特点

文件型弥媒适合生产级文件存储，因为它提供：

- 文件与目录恢复方式直观、可检查。
- 底层内容不可变并可寻址。
- Merkle DAG 可以复用未变化的数据块。
- 天然适合照片、视频、扫描件等大文件。
- JSON 是便携、透明、长期可迁移的对象格式。
- 不依赖某个数据库实现和数据库修复工具。
- 可以使用通用文件工具完成导出和灾难恢复。

选择文件型弥媒的应用仍可使用本地数据库或搜索索引加速，但这些缓存应能从权威文件完整重建。

### 5.2 两个文件系统界面

| 界面 | 身份与范围 | 代表操作 | 典型用途 |
|---|---|---|---|
| 文件型弥媒文件系统 | 一个 `mid`，打开 `cur`、`last`、`release` 或数字版本 | `MMOpen`、`MFOpenByPath`、`MFOpenMacFile`、`MFReaddir`、`MFSetData`、`MFGetData`、`MFStat` | 权威的版本化应用数据 |
| 节点 Files 文件系统 | 节点本地、以 `/` 为根的 IPFS 可变目录树 | `FilesMkdir`、`FilesLs`、`FilesCopy`、`FilesMv`、`FilesRm`、`FilesStat`、`FilesFlush` | 暂存、组装、移动并 flush IPFS 目录树 |
| 不可变 IPFS/Mac 内容 | 使用内容 ID，不使用可变路径 | `MFTemp2MacFile`、`MFTemp2Ipfs`、IPFS add/get | 保存和复用不可变内容 |

`MMOpen(sid, mid, ver)` 返回操作会话 `mmsid`，不是新的弥媒 ID。对于文件型弥媒，该会话可以代表所选版本的文件系统根。仓库中的文件型弥媒示例打开版本后直接调用 `MFReaddir(mmsid, -1)` 枚举根目录；也可以使用 `MFOpenByPath(sid, mmsid, path, flag)` 打开根下的文件，再用返回的资源会话读写字节、对象和状态。

直接的弥媒文件系统流程是：

```text
登录 -> MMCreate 或定位 mid -> MMOpen(mid, "cur")
     -> 通过弥媒文件系统句柄打开并修改文件/目录
     -> 关闭资源句柄 -> MMBackup -> last 前移
```

仓库的旧版 API 文档还描述了 `FSFind`、`FSMkDir`、`FSStat`、`FSCopy` 和 `FSRename`。这些接口能说明架构意图，但没有出现在仓库当前的 Go `LApiStub` 中。应用采用之前，必须在实际安装版本中核对函数是否存在、参数、flag 和 JavaScript 名称。

另一条节点 Files 组装流程是：

```text
临时文件 -> MFTemp2Files(节点路径)
         -> Files* 修改 -> FilesFlush(path) 返回根 CID
         -> MFSetCid(mid, 根 CID) 把完整目录树赋给文件型弥媒
```

这条流程适合大文件上传和整棵目录树组装，但节点路径不是弥媒身份；稳定身份仍然是 `mid`，正式读取应打开弥媒版本。

#### `/.mm/<mid>`：Files 命名空间中文件型弥媒的可写目录树（实验性）

**状态：实验性。** 2026-10-14 在 gen8（Leither V0.24.62）上以节点系统用户运行 CLI 观察到。
先以节点系统用户运行 CLI，再用非系统会话测试。本文已查阅的 Leither 仓库文档没有记载它。
通过它的读取不受访问控制（见下文），不要让应用代码依赖它。

节点 Files 根目录下有一个 `.mm` 目录，节点上每个文件型弥媒在其中有一个以 `mid` 命名的子目录。
`/.mm/<mid>` 就是该弥媒的 `cur` 目录树，可以用普通 `Files*` 调用访问：

- 下面测试的每一步里，`files stat` 返回的哈希都等于 `mm://<mid>:cur`；只要 `cur` 有未提交的改动，
  它就不同于 `mm://<mid>:last`。`cur` 与最新提交相同时，三者一致。
- 目录在第一次写入前不存在：新建的弥媒对 `/.mm/<mid>`、`mm://<mid>:cur` 和 `mm://<mid>:last`
  都返回 `file does not exist`；第一次 `files mkdir /.mm/<mid>/photos` 才创建它。
- 写入只进入 `cur`。提交之前 `last` 不出现；出现之后也不会移动。

在一个临时弥媒上的测试（`mimei create -t file -e probe.test`）：

| 步骤 | `/.mm/<mid>` | `mm://<mid>:cur` | `mm://<mid>:last` |
|---|---|---|---|
| 创建后 | 不存在 | `file does not exist` | `ver[last] not find!` |
| `files mkdir /.mm/<mid>/photos` | 新根 | 相同 | `ver[last] not find!` |
| `files cp ipfs/<cid> /.mm/<mid>/photos/a.txt` | 新根 | 相同 | `ver[last] not find!` |
| `files flush /.mm/<mid>`（返回根 CID） | 相同 | 相同 | `ver[last] not find!` |
| `mimei setcid --mid <mid> <根 CID>` | 相同 | 相同 | 相同（版本 1） |
| `files mkdir /.mm/<mid>/docs` 并 flush | 新根 | 相同 | 不变 |

第二次写入后，未改动的 `photos` 子树保持原 CID，只有根变化，符合 Merkle 树的预期。单独的 `FilesFlush`
不提交；提交的是 `mimei setcid`（`MFSetCid`），与 5.6 节一致。

2026-10-14 在同一节点上用非系统会话做了测试。身份是新生成的密钥、自签证书和 `CertFor=Self` 通行证，
通过 `LoginWithPPT` 登录，方式与 LifeDrive 身份服务登录家庭用户相同。除非另有说明，它没有任何授权。
对照组：普通 Files 目录加上 `files setright -u <uid> -d <目录> -r write` 后，该身份的 `mkdir`、`cp`、`rm`
都成功，所以探针本身是可靠的。

| 情形 | 结果 |
|---|---|
| 无授权，对 `/.mm/<mid>` 做 `FilesLs`/`FilesStat` | **允许**。权限为 `72f6f04` 的弥媒和权限为 `0` 的私有弥媒都可以；对第二个私有弥媒还列出了 `secret/a.txt` 及其 CID 和大小。 |
| 无授权，`FilesLs("/.mm")` | **允许**，返回节点上全部 12 个 `mid`。 |
| 非所有者会话对 `mm://<私有 mid>:last/…` 和 `:cur/…` 做 `FilesLs`/`FilesStat` | **允许**，与 `/.mm/<mid>` 相同。（本行早先版本记录为 `invalid mimei id`，那是探针的缺陷：它在第一个冒号处截断了 URL。） |
| 完全不登录（guest）的同样调用，包括从互联网访问 gen8 的公网端口 | **允许** |
| 非所有者或 guest 对 `mm://<私有 mid>:last/…` 或 `/.mm/<mid>/…` 做 `MMOpenUrl`/`MFGetData` | `1101: Permission denied`（这条路由保护内容） |
| 非所有者会话 `MFOpenTempFile`（MApp 式 `MFTemp2Files` 写入的第一步） | `1101: Permission denied`；维护者所说「`MFTemp2Files` 不判权」无法验证 |
| 无授权，在 `/.mm/<mid>` 下 `FilesMkdir` | `5006: Insufficient permissions Run as administrator` |
| 执行 `files setright -u <uid> -d /.mm/<mid> -r write`（对 `/.mm/<mid>/photos` 也试过）之后 | 仍为 `5006`；`listrights` 显示有授权，但写路径不理会 |
| 执行 `mimei setright --mid <mid> <uid> 0x07276704` 之后 | 仍为 `5006`；`getright` 把该 uid 读回为仅 `Read` |
| 执行 `mimei setright --mid <mid> <uid> 0x07`（读+写+执行）之后 | **允许**：`mkdir`、`cp ipfs/<cid>`、`mv`、`rm`、`flush` 全部成功 |
| 执行 `mimei setright --mid <mid> <uid> 0` 之后 | 再次 `5006` |

结论：

- `/.mm/<mid>` 下的写入由弥媒按 uid 授予的权限决定，需要写位。目录级 `files setright` 在这里不起作用。
  这与「按弥媒作者定的权限走」一致，但不是 LifeDrive 对 `/lifedrive-work/<用户>` 使用的机制。
- 会话的写入进入 `cur`。非系统会话 `flush` 之后 `last` 没有变化（仍是 `QmSTU…`），只有提交才会移动它。
  测试没有从该会话调用 `MFSetCid`，所以谁能通过它提交尚未验证。
- **目录和元数据的读取不受弥媒权限限制，任何路径形式都一样。** `FilesLs` 和 `FilesStat` 对 `/.mm/<mid>`、
  `mm://<mid>:cur/` 和 `mm://<mid>:last/` 都会回答非所有者，也回答完全不登录的调用，包括权限为 `0` 的弥媒。
  gen8 的公网端口回答了来自另一台机器的匿名调用；匿名 `FilesLs("/.mm")` 返回了节点上全部 13 个文件型弥媒 id。
  所以任何能访问节点的人，无论是否知道 `mid`（知道 `mid` 就不需要 `/.mm`），都能看到私有目录树的文件名、大小和 CID。
  字节读取路由 `MMOpenUrl` 会判权，但从列表得到的 CID 可以不带会话直接 `GET /ipfs/<cid>`：
  gen8 公网端口对未登录请求返回 `200`；minipc 的环回端口对一份家庭身份记录也是如此（只记录了状态码和字节数，没有读取内容）。
- **这对 minipc 上的 LifeDrive 意味着什么。** 身份弥媒是文件型弥媒；对 `mm://<身份 mid>:last/` 做匿名 `FilesStat`
  返回了它的哈希和大小（5,271 字节），`/ipfs/<哈希>` 返回同样大小。因此任何持有身份 `mid` 的人都能不经会话读取该用户的身份记录。
  `mid` 在当前设计中并不保密（它在配对载荷和资料回复里）。目录迁移之后，那里的旧网盘和备份弥媒是空目录，没有暴露内容。
  数据库型的目录弥媒和 user-root 弥媒返回 `type is MiMei` 和 `not ipfs file id`，**不能**通过 Files RPC 访问。
- 这是 Leither 的行为，应向其维护者报告，应用无法修复。在修复之前，应把任何文件型弥媒的名称、大小、CID 以及这些 CID 背后的字节
  都视为人人可读。

Leither 维护者的备忘录（2026-10-14，「files-mm-read-permission」，在临时 V0.24.62 节点上实测）
证实并解释了上述结果：

- `/.mm/<mid>` 下的读取不判权：`FilesLs`、`FilesStat`、`MFLs` 都不做权限判断；未授权用户或 guest 通过 MApp
  读取 `/.mm/<mid>/secret.txt`，得到了一个「Guest 无权限」弥媒中的内容。所以可读的不只是名称和 CID，还有内容。
- 数据面（经 `MMOpenUrl`/`MFGetData` 读取 `mm://<mid>:<版本>`）会判 `DataRight_Read`：对 Guest 无权限的弥媒返回
  `1101: Permission denied`。泄露出在 Files 命名空间的路径，而不在弥媒本身。
- `/.mm/<mid>` 下的写入判弥媒的 `DataRight_Write`，不看 `files setright` 授权，与上表一致。
- 按用户 `setright` 的值只取低字节（`0x07276704` 变成 `0x04`，即 Read）再加上组槽默认位，所以读回只有 Read。
  `Right :072f6f04` 这种完整位图用于 `mimei create -r`：低字节是 Guest 槽，所以 `0x072f6f04`（`mimei create`
  的默认值）Guest 可读，`0x072f6f00` 不可读。
- 他们把它列为缺口 A，建议加一个读闸，有 `off`、`observe`、`enforce` 三档（与 `FilesGateMode` 同族），
  覆盖 `FilesLs`、`FilesStat`、`MFLs` 和 `MMOpenUrl("files://…")`，判 `DataRight_Read`，主体未知时拒绝。
  V0.24.62 没有。
- 他们还说明：`FilesGateMode=off`（缺省）时，`MFTemp2Files` 写入 `/.mm/<mid>` 路径同样不判权。本文没有测试这一点。

LifeDrive 对这类读闸的要求：

- 分享校验 `apps/lifedrive-identity/shares.go` 中的 `FilesStat("", mm://<share>:last/<name>)` 用空会话访问以
  `0x07276704`（Guest 可读）创建的分享弥媒。「主体未知时拒绝」必须理解为检查 Guest 槽，而不是直接拒绝。
- 身份服务用用户的 authority 会话经 `mm://<mid>:last` 读取网盘和目录；该会话是弥媒作者，恒可通过。
- 备忘录针对的是 `/.mm/<mid>` 路径，而 LifeDrive 处处用 `mm://` URL 配合 `FilesLs`、`FilesStat`。
  上面的测试表明这些形式今天同样不判权，所以只修 `/.mm/<mid>` 会让泄露继续存在。读闸必须同时覆盖
  `mm://<mid>:<版本>/…`，并且不能拒绝对 Guest 可读弥媒的空会话分享校验。
- 备忘录说 `MMOpenUrl("files://…")` 不判权。本测试中，非所有者对 `files:///.mm/<mid>/…` 做 `MMOpenUrl` 返回了 `1101`，
  即字节路由在那里已经拒绝；没有保护的是目录列表和 `/ipfs/<cid>`，只给 `FilesLs`/`FilesStat` 加读闸并不能关闭后者。
  `/ipfs/<cid>` 是否应要求会话，需要问维护者。

尚未确认：

- 非系统会话能否提交（`MFSetCid`）它通过 `/.mm/<mid>` 能写入的文件型弥媒。
- V0.24.62 的这种读取行为在更早版本中是否存在。

gen8 的根目录还有一个名为 `mm:` 的目录，其子项是 `<mid>:cur`，哈希与 `/.mm/<mid>` 不同；
只有当时存在 `cur` 目录树的三个弥媒有这类子项。用途不明，看起来是已打开的 `mm://…:cur` 树的内部缓存。

测试用弥媒 `R8F75nZSJitgtGwPn952zk2bLye`（mark `claude-mm-probe-20261014`，权限 `72f6f04`）、
`xLAtmWH4DgvO4rBKFXoYEzPgQI7`（mark `claude-mm-probe-private-20261014`，权限 `0`）和
`JSj2uUxGAe0e8BcVVqLn3Q9cmLg`（mark `claude-mm-probe-share-20261014`，权限 `0x07276704`，作者是一个已删除密钥的临时身份）
仍留在 gen8 上；CLI 没有删除弥媒的命令。gen8 和 minipc 上的测试身份、授权和探针程序已清除。

### 5.3 主要文件操作

| 目的 | CLI 示例 | API 职责 |
|---|---|---|
| 创建容器 | `mimei create --data-type file` | 创建文件型弥媒 |
| 导入文件/目录 | `mimei add --mid <mid> <path>` | 把本地内容加入弥媒 |
| 建目录 | `mimei mkdir <path>` | 修改弥媒目录树 |
| 复制或关联内容 | `mimei cp <source> <dest>` | 树内复制或加入 IPFS 路径 |
| 设置内容 | `mimei setdata` / `mimei setcid` | 用文件或已有 CID 设置内容 |
| 读取和列表 | `mimei get` / `mimei ls` | 获取文件和枚举目录 |
| 提交 | `mimei backup --mid <mid>` | 把 `cur` 提交为新的 `last` |
| 检查 | `mimei show --mid <mid>` | 查看版本、权限和状态 |

节点级 `files` 命令族还提供 `cp`、`mkdir`、`mv`、`rm`、`flush`、`ls` 和 `stat`。大文件和断点续传可先写临时文件，完成校验后直接转为弥媒中的 Mac/IPFS 文件，或先接入节点 Files 目录树、flush 后把根 CID 赋给文件型弥媒。

仓库确认的关键 API 关系包括：

- `MMCreate(sid, appid, ext, mark, tp, right)` 返回稳定的 `mid`。
- `MMOpen(sid, mid, ver)` 返回所选版本的操作会话 `mmsid`。
- `MFOpenByPath` 和 `MFOpenMacFile` 返回文件资源会话 `fsid`。
- `MFSetData`/`MFGetData`、`MFSetObject`/`MFGetObject`、`MFGetSize`、`MFGetMimeType`、`MFStat` 和 `MFReaddir` 都操作已打开的资源会话。
- `MFOpenTempFile` 创建临时资源；`MFTemp2MacFile`、`MFTemp2Ipfs` 和 `MFTemp2Files` 对应不同的完成去向。
- `FilesFlush` 返回节点 Files 子树的根 CID；已验证的安装版本可用 `MFSetCid` 把该 CID 赋给文件型弥媒。
- `MMBackup` 是通用弥媒快照操作，生成不可变版本并推进 `last`。

`MMBackup` 的第一个参数存在版本差异：仓库 API 文档称为登录 `sid`，仓库可运行的数据库备份示例却传入 `MMOpen` 返回的 `mmsid`；现有 JavaScript 应用也可能使用 `BELoginAsAuthor()` 返回的作者会话。适配层必须以安装版本的实测签名为准。

同样，不能把 `MMBackup` 与 `MFSetCid` 固定地当作每次修改都要连续调用的两个提交。前者是通用快照；后者是安装版本中把完成的根 CID 赋给文件型弥媒的路径，并可能直接返回新版本。对上层应用只能暴露一个经过验证的逻辑提交。

### 5.4 句柄类型与关闭规则

- `sid`：已认证 API 会话。
- `mid`：稳定弥媒身份。
- `mmsid`：`MMOpen` 返回的弥媒/版本操作会话。
- `fsid`：已打开的文件、目录或临时文件资源会话。
- CID/Mac ID：不可变内容身份，不是会话。

这些值即使都以字符串表示，也不能互换。所有打开的弥媒和资源会话都要在成功与失败路径关闭。仓库 Markdown 对文件资源写作 `MFClose`，而当前 Go stub 和示例主要使用 `MMClose`；实际可调用名称以安装版本为准。

### 5.5 目录表达关系

文件化应用可以通过目录层级表达父子关系：文件夹的子目录和文件就是它的子内容，使目录树而不是数据库表成为权威关系图。应用可以采用用户可见文件名，也可以采用不透明 ID；这属于应用数据模型选择，不是 Leither 的强制要求。无论采用哪种方式，服务端都应规范化路径、拒绝路径穿越，并定义重名、移动和改名规则。

### 5.6 V0.24.23 上核实的提交开销与节点 Files 行为

以下内容于 2026-09-16 在 minipc（Leither V0.24.23，linux/amd64）上观测，并经 Leither 维护者确认。详细记录见英文版
[Leither and MiMei 第 5.6 节](../LEITHER_AND_MIMEI.md#56-commit-cost-and-node-files-behavior-verified-on-v02423)，
据此形成的 LifeDrive 方案见 [LifeDrive Backup Storage Plan](https://github.com/cfa532/LifeAlbum/blob/0876ac5a20e0b4f9225e63febe9426f8694795d9/docs/LIFEDRIVE_BACKUP_STORAGE_PLAN.md)。

**提交会钉住整棵树。** `MFSetCid(sid, mid, rootCid)` 通过 `backupWithFileId -> backupModifyInfo -> addRef` 记录版本，`addRef` 对新根执行递归 `IpfsPinAdd`，遍历根下全部块，而不只是新增块：13 块用时 31 ms，12,210 块的 LifeDrive 备份盘用时 51.037 s。提交开销与整个弥媒大小成正比，并随每个新文件增长。在同一个弥媒内把目录分得更深无济于事，只有独立的弥媒才是独立的钉住根。调用方 RPC 超时短于钉住时间时会看到超时，但提交仍可能成功，结果不确定；重试会再次钉住整棵树。

**引用、钉住与垃圾回收（维护者确认）：**

- `MMAddRef` 钉住被引用的 IPFS 内容，`MMDelRef` 解除钉住。对本节点已有的内容，`MMAddRef` 开销小且同步完成。
- 引用远程节点的内容需要时间同步过来，由 Leither 负责；本地路径访问需等待同步完成。
- 垃圾回收在磁盘空间不足时运行，只清除已无任何引用的 IPFS 块，从不删除磁盘上的普通文件。
- 对另一个弥媒的引用是弥媒 ID，子弥媒提交新版本时 ID 不变，因此提交子弥媒不给父弥媒带来任何开销；父弥媒仅更改引用的提交开销很小。
- 同步只传输已提交的 `last` 版本，不传输 `cur`。

**文件型弥媒的提交路径：** 目录树用 `FilesFlush` + `MFSetCid` 提交（与 TweetBackendApp 一致，且不要在 `MFSetCid` 之后调用 `MMBackup`）。对类型 1 弥媒，`MMBackup` 把 `cur` 作为字节快照，`last` 变成字节节点而非目录；以创建标志调用 `MFOpenByPath` 会在 `mmfs.(*MTreeFS).OpenFileForFind`（`MTreeNode2WebDav.go:281`）触发服务端 panic。普通用户会话调用 `MFOpenTempFile` 返回 `1101: Permission denied`，需要节点的 `mimei` 服务授权。

**命令行 `mimei` 文件命令：** `mimei mkdir` 默认刷新到 `last`；新弥媒上 `--flush=false` 不产生 `last`；`mimei add` 会用该文件替换整个根；CLI `ipfs add` 后 `mimei cp /ipfs/<cid>` 或 `files cp /ipfs/<cid>` 报 `file does not exist`，因为开头的斜杠表示节点 Files 树里的路径，写成 `ipfs/<cid>`（或 `ipfs://<cid>`）则添加后立即复制成功；`mimei cp mm://<other-mid>:last …` 挂起 219 s，客户端被终止后服务端 `MFCopy` 仍在运行。Leither 自带的 `files cp` 帮助把本地内容加入 MFS 的方式记为 `ipfs add --pin=false <file>` 后 `files cp /ipfs/<CID> <路径>`（只引用根的惰性复制），V0.24.23 上源路径须写成 `ipfs/<CID>`。

**`IpfsAdd` 与会话消息队列：** `IpfsAdd(sid, localPath)` 在内容存好后才返回。它向调用会话发送进度消息；不调用 `PullMsg` 的服务会话，队列（容量 10）满后每条消息等待 3 s（`sendmsg-timeout`），104 KB 的添加因此耗时 9 s，批量导入期间用 `PullMsg` 清空队列后恢复到亚秒级。这样添加的块未被钉住，下一次 `MFSetCid` 钉住整棵树时一并处理：19,616 块耗时 3 分 31 秒。

**节点 Files 性能：** 对目录调用 `FilesStat` 会遍历子树统计大小，在 12,210 块的盘上 45 s 内未返回，判断是否存在应改为列出父目录（`FilesLs`，约 350 ms）。当天稍后所有节点 Files 列表都挂起（`files ls /` 超过 60–120 s），而 `mm://` 列表 353 ms 返回；此时依赖 `FilesCopy`、`FilesFlush` 或工作树路径的操作在全节点范围停滞。同期 `Leither.log` 每 2–3 分钟轮转一个 270 MB 的文件。

## 6. 弥媒引用

弥媒引用连接不同弥媒或维持被引用资源的生命周期，它不同于一个文件型弥媒内部的目录层级。

以下情况需要显式弥媒引用：

- 主容器引用一个可以独立分享的文件型弥媒。
- 一个应用引用另一个独立保存的对象或集合。
- 一个容器需要保留外部不可变资源。
- 同步父弥媒时需要带上一个直接子弥媒。

Leither 只能自动跟随直接引用，不能无限递归整个对象图。因此应用必须定义自己的遍历规则，并在有意的数据边界使用跨弥媒引用。

## 7. 备份、发布、支撑与同步

| 操作 | 含义 |
|---|---|
| Backup | 把 `cur` 提交为不可变版本并推进 `last` |
| Publish | 具有发布权限者从持有最新版本数据的节点发布或重新发布弥媒状态 |
| Provide | 支撑节点声明可以提供弥媒副本的数据，角色类似 CDN |
| Find providers | 查找可提供该弥媒的节点 |
| Sync | 把指定版本和数据同步到本地节点 |
| Unprovide | 撤销本节点的支撑声明 |
| Unpublish | 撤销弥媒信息的网络发布 |

发布不局限于原作者或作者的根节点：任何拥有所需发布权限者，只要发布节点持有弥媒最新版本的数据，就可以执行 `publish` 或重新发布。TweetWeb 因而可以从 gen8 或 minipc 发布，但必须先确认所选节点上的弥媒数据是最新的；本地存在 `last` 标签不代表它已经与其他节点的最新版本一致。旧副本必须先更新，再重新发布。

发布会增加弥媒的版本号。发布后应确认新的版本号，不能继续把发布前的版本号视为最新版本。

支撑节点提供副本，角色类似 CDN。`provide` 声明提供数据的职责，本身不授予发布权限；支撑节点若同时具备发布权限和最新数据，也可以重新发布。文件型弥媒遵循相同规则，不能把 `publish -> provide` 当作发布时必须连续执行的两个步骤。

修改数据后先提交版本，再发布弥媒；重新发布已有版本前，先确认本节点持有最新数据。数据库弥媒使用 `MMBackup`；通过目录根 CID 提交的文件型弥媒使用 `FilesFlush -> MFSetCid`，其后不能再调用 `MMBackup`。副本支撑和副本健康检查属于分发职责。

如果一个节点可以查到弥媒，而其他节点返回 `no provider found`，应调查发布与跨节点发现过程；仅凭这个现象不能断定发布者缺少一次 `MiMeiProvide` 调用。

#### 未文档化的可变参数，以及 `delref=false` 的实际作用

安装版 `MMBackup` 的参数比公开文档多。`api/MiMei.md` 写的是
`MMBackup(sid, mid, memo string) (ver string, err error)`，而 V0.24.02 运行节点上的真实
签名是：

```text
MMBackup   func(sid, mid, memo string, opts ...string) (string, error)
MMAddRef   func(sid, mid string, fileids ...string) (int, error)
MMDelRef   func(sid, mid string, fileids ...string) (int, error)
```

TweetBackendApp 的 JavaScript entry 在**全部 39 处** `MMBackup` 调用中都传入
`"delref=false"`，包括纯新增的场景。Go 版本在每个写入路径上也这样做。

在运行节点上探测的结果是：该选项**没有任何可观察到的效果**。每一组结束后用 `MMGetRef`
读回引用列表，以下四种情况无法区分：

| 传入的选项 | backup 之后的引用表 |
|---|---|
| 不传 | 不变 |
| `delref=false` | 不变 |
| `delref=true` | 不变 |
| `zzz=nonsense` | 不变 |

测试覆盖两种场景：内容变更后的普通再次 backup，以及完整的
`MMAddRef` → backup → `MMDelRef` → backup 循环。在所有分组中，新增的引用都能在普通
backup 后保留，而 `MMDelRef` 都会将其移除，与是否传选项无关。

其中"无意义选项"这一组是关键对照：**无法识别的选项会被接受且不报错。** 因此"选项被接受"
完全不能说明它已被实现；传入选项的代码也不能因为"没有失败"就推断其行为。不要把
`delref=false` 当作对 backup 行为的描述。

以上结论是通过 `MMGetRef` 在类型 `2`（数据库）弥媒上观察得到的。如果存在在该处看不到的
影响——例如被引用对象数据的垃圾回收，或文件型弥媒上的不同处理——本次测试不会发现。

应用必须分别记录本地提交、网络发布、支撑声明和副本策略是否满足。Backup 成功不等于异地副本已经完成。

**不要同步本节点自己已经拥有的弥媒。** 在 V0.24.02 上，当弥媒存在但没有**其它**节点
提供它时，`SyncMiMei` 会以越界错误 panic：提供者列表解析为空，随后仍被按 `[0]` 取值。
节点把它作为已恢复的运行时错误报告，而不是普通失败：

```text
SyncMiMei:(runtime.boundsError{x:0, y:0, signed:false, code:0x0})
  Leither/pnet.(*MiMei).getSyncInfo   dhtMiMei.go:495
```

三种情况可以区分触发条件：有真实远程提供者的弥媒同步正常；不存在的弥媒会更早被拒绝，
返回干净的 `30000:No MiMei mid[...]`；而所有已公告提供者都是本节点自身地址的弥媒会
panic。节点自身的后台例程也会触发同一缺陷，应用无法完全规避；但应用可以避免自己触发
它——先解析属主节点，属主是本节点时跳过同步调用。这个判断本身也是正确的：无法从自己
拉取数据。

## 8. 身份、会话和权限

Leither 身份以非对称密钥为基础。节点密钥的公钥摘要标识节点；证书和通行证用于证明身份和委托访问；登录产生 API 使用的会话 ID。

弥媒权限控制谁能操作容器。应用可以根据自己的业务操作、分享规则和生命周期要求增加更细的策略。

常用 CLI 公共参数包括：

- `--node`：目标节点 URL、ID 或名称。
- `--key-file`：认证或签名密钥。
- `--ppt`：通行证文件。
- `--dht-ids`：指定 DHT 网络。
- `--api-version`：明确 API 合约版本。
- `--json`：机器可读的统一返回封装。

应用运行时应直接调用对应 API；CLI 用于运维、诊断、恢复和人工复现 API 操作。

## 9. 三种 API 边界

1. **CLI：** 完整的管理操作面，部分命令支持结构化 JSON。
2. **外部 API/RPC：** CLI 操作的程序化接口；应用需要完整节点能力时应优先使用。
3. **容器内 MApp API：** Leither 给容器内应用代码提供的较小接口集合。

公开仓库又提供了两种重要证据：

- `lapi/stub.go` 描述当前导出的 Go RPC 客户端接口以及句柄/返回类型。
- `example/rpc/` 展示登录、打开弥媒、使用返回会话、备份、临时文件、节点 Files 文件系统和读回验证的调用顺序。

这些内容不自动等于更新版安装程序的合约。仓库 Markdown、客户端 stub、示例、JavaScript MApp 绑定、CLI 帮助和二进制程序可能在不同时间更新。Go stub 中没有某个函数，不代表安装版 CLI 或 JavaScript 运行时一定没有；旧 Markdown 写有某个函数，也不能证明它在当前运行时可调用。

TweetBackendApp 展示了 MApp 模式，但某个容器内 Go 接口没有暴露全部外部能力，不代表 Leither 节点本身缺少这些能力；安装版 CLI 和外部 API 可能仍然提供。

容器内 JavaScript 可以在经过验证的 MApp 接口范围内直接实现文件型服务；外部 API/RPC 和 CLI 仍用于管理、诊断、恢复以及容器内未开放的能力。应用不能假定任意一个接口面暴露节点的全部能力。

### 9.1 Go MApp 解释器合约

TweetBackendApp 的 Go 版本验证了以下运行模型：

- Leither 使用 `ixgo` 解释器编译上传的源码，不运行预编译的原生二进制。
- 应用是一个名为 `lapp` 的扁平 package，不应依赖子 package。
- package 只暴露一个 `RunMApp`，并在内部完成 entry 分派。
- `go.mod` 只用于本地语法和类型检查；Leither 构建上传应用时不读取它。
- 运行时 import 必须使用 Leither 注册的准确路径，包括 `Leither/lapi`。
- 本地 CLI 模式可能通过位置参数 `Entry` 传入入口名；容器模式的已验证版本可能把入口名放在 `Request["entry"]`，同时令 `Entry` 为空。分派器应兼容两种来源。
- 不要保留与 Go 分派器入口同名的旧 JavaScript MApp 文件。V0.24.05 `darwin/arm64` 已验证在移除这些旧服务端入口后，Go LifeDrive 可以正确编译和分派。`lifedrive.js` 等浏览器对象使用独立名称，并由网页引导页加载，而不是作为服务端入口调用。Leither 升级后必须重新探测 Go 与浏览器对象混合发布的行为。
- 客户端应访问备份后的 `last` 或数字版本。已验证版本在 upload 后仍可能让 `cur` 运行旧编译结果，直到执行应用 backup。
- `--local` 可用于只读和分派探测，但已验证版本在该模式没有作者身份，`BELoginAsAuthor` 写入会失败；授权修改必须在容器模式验证。

Leither Go 程序使用的是允许列表，不是可以任意下载模块的 Go 环境，**而且这个列表会随版本
变化，并且已经缩小**。本文档早期修订列出的若干 package 现在已无法解析。

以下结果通过在安装节点上编译单 import 探测 MApp 得到，环境为 **V0.24.02, linux/amd64**。
不可用的 import 在编译期失败，错误为 `could not import <路径> (not found package)`。

| Import 路径 | V0.24.02 | 说明 |
|---|---|---|
| `errors` | 可用 | 生产环境已验证 |
| `fmt` | 可用 | 生产环境已验证 |
| `io` | 可用 | 生产环境已验证 |
| `os` | 可用 | 探测返回 `os.PathSeparator` |
| `sort` | 可用 | 生产环境已验证 |
| `strconv` | 可用 | 生产环境已验证 |
| `strings` | 可用 | 生产环境已验证 |
| `time` | 可用 | 探测返回 `time.Second` |
| `encoding/json` | 可用 | 完整 `Marshal`/`Unmarshal` 往返已验证，map 键有序 |
| `Leither/lapi` | 可用 | 容器内 Leither API |
| `bytes` | **未找到** | 早期修订曾列为可用 |
| `encoding/gob` | **未找到** | 早期修订曾列为可用 |
| `net/http` | **未找到** | 早期修订曾列为可用 |
| `crypto/sha256`、`crypto/ed25519` | **未找到** | 应用内无法做哈希或签名校验 |
| `unicode/utf16` | **未找到** | 代理对处理必须自行实现 |
| `github.com/shirou/gopsutil/*` | **未找到** | 已测 `mem` |
| `github.com/hprose/hprose-golang/v3/*` | **未找到** | 已测 `rpc/core` |

真正重要的是"缺失"而不是"存在"：

- 没有 `crypto/*` 和 `net/http`，应用无法自行哈希、校验签名或发起对外请求。这类能力只能
  来自 `lapi`；`lapi` 不提供，就等于该应用没有这个能力。
- `encoding/json` 可用且会对 map 键排序，因此不需要为确定性输出手写 JSON 编码器。

不要用本地 `go build` 成功来推断 package 集合，也不要把某个版本的清单沿用到升级之后。
请针对安装的二进制重新探测。

import 成功不等于运行行为与原生 Go 一致。在已验证的 Go MApp 运行时中，`init()` 和 package 级变量初始化器不会执行，变量会静默保留零值。因此分派表、map、错误哨兵和配置应在函数或显式的每次调用初始化路径中构造；依赖 package 初始化的库必须先在目标 Leither 版本做实际操作探测。本地 `go build` 只能证明语法和类型正确。

#### 访问未在公开接口上暴露的 lapi 能力

`lapi.GetLApi()` 返回的值静态类型是 `lapi.LApi` 接口，但其背后的对象是 `*frame.LApi`，
方法集**更大**。若干操作存在于该具体类型上，却不在 `lapi.LApi` 接口中，也不在
`github.com/3and4/Leither/lapi` 的任何已发布版本中。因此按公开接口编写的应用看不到它们，
并会得出"这些能力不存在"的错误结论。

在 V0.24.02 上确认存在，签名通过 `fmt.Sprintf("%T", ...)` 从运行节点读取：

| 方法 | 签名 |
|---|---|
| `MiMeiSync` | `func(sid, dhts, mid string, param map[string]string) error` |
| `MiMeiPublish` | `func(sid, dhts, mid string) ([]lapi.DhtReply, error)` |
| `MiMeiProvide` | `func(sid, dhts, mid string) ([]lapi.DhtReply, error)` |
| `MiMeiUnprovide` | `func(sid, dhts, mid string) ([]lapi.DhtReply, error)` |
| `MiMeiUnpublish` | `func(sid, dhts, mid string) ([]lapi.DhtReply, error)` |
| `MiMeiIsProvider` | `func(sid, mid string) (bool, error)` |
| `RunMApp` | `func(entry string, req map[string]string, args []any, opts ...string) (any, error)` |

`dhts` 参数选择要操作的 DHT，`""` 表示全部，JavaScript 运行时传的也是 `""`。

访问方式是把句柄断言到本地声明的接口，每个操作一个：

```go
type mimeiSyncer interface {
	MiMeiSync(sid, dhts, mid string, param map[string]string) error
}

if s, ok := api.(mimeiSyncer); ok {
	err := s.MiMeiSync(sid, "", mid, param)
}
```

这样应用仍然只 import `Leither/lapi`。内部的 `Leither/api` package 同样暴露这些方法，
但它的 `GetLApi()` 直接返回 `*frame.LApi`，会让应用依赖节点内部类型，没有任何收益。
类型断言还能安全降级：在没有该方法的版本上断言失败，调用方可以报告能力不可用，而不是
无法编译。

`Ed25519Verify` 通过所有已测试路径都**不可达**。需要校验签名的应用应把它视为不可用，
而不是降级成"检查格式是否正确"这种看起来像校验的替代品。

#### 每次请求的解释器内存开销

每次 Go MApp 请求都会保留解释器的编译状态，且不会释放。在两个独立节点上测量，使用同一个
最简单的 entry，连续 30 次请求，并以同节点上的 JavaScript MApp 作为对照：

| 节点 | Go MApp | JavaScript MApp |
|---|---|---|
| 节点 A | +609 MB（20.3 MB/请求） | +22.7 MB（0.76 MB/请求） |
| 节点 B | +242 MB（8.1 MB/请求） | +22.7 MB（0.76 MB/请求） |

对占用 610 MB 的节点做堆剖析，占比最高的是编译过程而不是应用数据：
`ixgo.(*function).regInstr`、`go/types.(*Checker).recordTypeAndValue`、
`go/ssa.createFunction`。goroutine 数量始终很低，因此这是内存滞留，不是并发泄漏。

运行影响严重且容易误判。超过约 3 GB 之后，节点对普通 HTTP 探测仍在毫秒级响应，而每个应用
调用都要几分钟——端口级监控会报告节点健康，客户端却在超时。实测某节点 `HEAD /` 用时 8 ms，
而最简单的 `health` entry 用时 2 分 3.8 秒。

V0.24.05 的表现明显改善，但还不能证明滞留已经消失。在 `darwin/arm64` 上，LifeDrive 使用
8 次 16 MiB `upload_write` 上传 128 MiB，随后移入垃圾箱并清除；完整 12 次 MApp 调用约用
2.44 秒，进程 RSS 从 257,984 KiB 上升到 300,800 KiB。这 42 MiB 上限包含上传和 IPFS
缓冲，未用堆剖析分离，不能全部归类为解释器滞留，但探测结束后也没有立即回落。

因此应使用较大的有界分块减少调用次数，监控**持有监听套接字的那个进程**的 RSS，并在每次
Leither 升级后重新测量。完成 V0.24.05 长时间堆剖析并证明编译状态释放之前，计划性重启仍是
合理保护措施。

#### JavaScript MApp 没有模块加载

JavaScript MApp **完全没有模块加载能力**。`require` 作为函数存在，但在 V0.24.02 上每次
调用都会失败并抛出 `GoError: Invalid module`——同一应用内的兄弟文件、`fs` 和 `path` 等
Node 内置模块、第三方包，全部如此。已通过探测验证。

对应用结构有两个后果：

- 每个 entry 文件必须完全自包含。公共辅助函数无法 import，只能在每个文件里重复，或在上传
  前由构建步骤生成进去。TweetBackendApp 的 JavaScript entry 在 68 个文件中有 65 个重复了
  各自的响应包装函数，原因正在于此——这是平台约束，不是疏漏。
- 调用了 `require` 的 entry 一定失败——每次都会在那一行抛异常。需要文件系统、解压或子进程
  的 JavaScript entry 无法实现；这类能力必须来自 `lapi`，或者移出 MApp。

JavaScript entry 文件还必须**调用自身**。文件是作为表达式求值的，末尾必须带调用，否则
entry 不返回任何内容，调用方看到的是 HTTP 200 加空响应体：

```javascript
((request, args) => {
    // ...
    return result
})(request, args)     // 末尾这个调用是必需的
```

以上探测结果的环境为安装的 **V0.24.02, linux/amd64** 节点（2026-08-25）。本文档与早期修订
冲突时，以探测结果为准；Go package 允许列表尤其**已经缩小**，任何 package 清单——包括这一
份——都应视为与版本绑定。

### 9.2 浏览器 Web 应用引导流程

Leither 已观察到两种网页入口。局域网直接访问格式为：

```text
http://<节点IP>:<Leither端口>/entry?aid=<应用ID>&ver=last
```

通过 `mimei setdomain` 绑定的域名会返回负责选择提供者的引导页；Tahoe V0.24.05 的局域网
直连 URL 则直接返回应用 `index.html`，既不定义 `window.getParam()`，也不重写外部资源地址。
两种方式都不需要节点根路径路由或应用名 URL 目录，但同时支持两种入口的网页包必须处理这一差异。

域名引导页依次完成：解析应用弥媒和提供者；通过 `window.getParam()` 暴露 `aid`、
`mid`、`ver`、`addrs`、`ips`、`CurNode` 和 `remote`；竞速访问候选提供者；从第一个可达节点
读取应用 `index.html`；把该页面注入引导文档；将链接和脚本对象重写到
`http://<提供者>/mm/<应用mid>:<版本>/`；最后按文档顺序加载并执行脚本。

TweetWeb 是已经运行验证的打包先例。应用弥媒中的网页对象位于应用根层，例如：

```text
index.html
lifedrive.js（或其他不冲突的浏览器脚本名）
内联 CSS（或使用归一化短名不同的外部样式表）
hprose.js 等显式引用的浏览器依赖
```

在域名引导模式中，`index.html` 可以使用 `lifedrive.js` 这样的裸对象名，不使用节点根路径，
也不使用应用名子目录。Leither 注入器会以 `/mm/<mid>:<ver>/` 为基准解析它们。Vite 8 在相对 base 下会生成 `./` 前缀，
TweetWeb 在构建后的 HTML 转换中移除该前缀；其他 Leither Web 应用也应采用这一做法，除非已经
对目标 Leither 版本验证过 `./` 对象名。

Tahoe 的局域网直连 `/entry` 不执行上述重写。裸 `lifedrive.js` 会解析成节点根路径
`/lifedrive.js` 并返回 404；真正发布的对象位于
`/mm/<应用ID>:<版本>/lifedrive.js`。LifeDrive 在 `index.html` 中加入一个很小的内联加载器：
域名模式优先读取 `window.getParam()`，直连模式则读取 URL 中的 `aid` 和 `ver`，随后构造
`/mm/` 模块地址。加载器不包含部署专用的应用 ID 或节点地址。

对象命名也与 Leither 版本有关。TweetWeb 使用 `index_entry.js`，但 Tahoe V0.24.05 会从
`index.html`、`index_entry.js` 和 `index.css` 派生出相同短名 `index`，并以
`entry shortName[index] ... exist` 拒绝整个应用包。因此 LifeDrive 使用独立的顶层名称
`lifedrive.js`，并把 CSS 内联到 `index.html`。应保留相同的引导布局，但必须避免目标运行时
的入口名归一化冲突。

域名注入后，应用脚本从 `window.getParam()` 读取运行时信息，而不是硬编码应用 ID 或节点地址。
应用 ID 主要取自 `aid`，并可用 `mid` 作兼容回退；当前 RPC 节点取自 `ips[CurNode]`。直连模式
从 `/entry` URL 及其 origin 取得同类信息。需要故障转移的客户端应保留按提供者分组的其他地址，
而不是只保存胜出的一个。

该网页引导与 MApp 入口调用是两个不同用途：带 `entry=<入口名>` 时调用服务端代码；省略
`entry` 时加载应用网页。

## 10. 文件化应用的运行模型

采用 JSON 对象树的文件化应用可以保持以下不变量：

- 其权威容器都是文件型弥媒。
- 每个逻辑对象都有稳定的应用对象 ID。
- 目录树是权威父子关系。
- 每个业务对象保存为版本化 JSON，必要时加密保存。
- 一个业务操作的内容和元数据共用一次备份边界。
- 普通读取只读 `last`，不读取未提交 `cur`。
- 提交、发布和副本状态分别可观测。
- 所有索引和缓存都能从文件及 JSON 完整重建。
- 任何敏感内容进入分布式存储前必须加密。

## 11. 参考定位

本参考采用以下证据优先级：

1. 目标安装版 Leither 二进制的实测行为和帮助。
2. 应用实际使用的安装版 API 绑定完成的成功往返测试。
3. 使用同一 Leither 版本并已正常工作的应用。
4. 当前仓库 stub 和可运行示例。
5. 仓库及网站 Markdown 文档。

较低优先级的资料仍然非常适合了解架构、设计意图和发现待验证能力，但不能替代版本固定和能力检测。

> Leither 是一个版本化的数据与应用容器，也是分布式服务平台。弥媒在不可变内容之上提供稳定身份，支持文件和数据库两种存储模型，并组合了版本、引用、权限、发布、支撑和同步能力。采用文件型弥媒的应用可以用普通文件或 JSON 表达对象、用目录表达关系；需要完整节点能力的应用应通过外部 API 集成。

这一模型比“Leither 只是一个 IPFS 服务节点”的描述更加完整。
