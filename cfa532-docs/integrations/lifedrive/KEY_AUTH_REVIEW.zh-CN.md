# LifeDrive 的 Leither 设备密钥认证设计

**Canonical home:** `cfa532/Leither/cfa532-docs/`
**Consolidated:** 2026-10-16 from [LifeAlbum/docs/zh-CN/LIFEDRIVE_LEITHER_AUTH_REVIEW.md](https://github.com/cfa532/LifeAlbum/blob/ade508c37bec3cd2d772e00f28d004eef5cdfde4/docs/zh-CN/LIFEDRIVE_LEITHER_AUTH_REVIEW.md).

Historical LifeDrive native-key design and V0.24.12 observations, not the current household authentication contract. Conflicting verification rows and old installer commands are retained as evidence, not reconciled by this move.

---

**状态：** 已在 `codex/lifedrive-key-auth` 实现

**适用版本：** Leither V0.24.11 或更高版本

**上游依据：** [KEY_AUTH_GUIDE](https://github.com/3and4/Leither/blob/main/doc/KEY_AUTH_GUIDE.md) 与 [key-auth-login 示例](https://github.com/3and4/Leither/tree/main/example/key-auth-login)

**目的：** 说明 LifeDrive 如何在不保留用户名/密码兼容层的前提下，使用 `sodiumv2`、PPT、SID 与弥媒原生权限提供方便的无密码登录。

## 1. 最终选择

| 设计问题 | 当前选择 | 原因 |
|---|---|---|
| 日常登录凭证 | 每台设备独立的 Ed25519 `sodiumv2` 密钥 | 不共享密码；可以单独撤销遗失设备 |
| 首次配对 | 导入一次由安装程序生成的身份文件 | 不需要手抄密钥或在浏览器中配置 Leither 内部参数 |
| 浏览器存储 | IndexedDB 中不可导出的 WebCrypto `CryptoKey` | 原始私钥文本只在导入期间短暂出现，不写入 Web Storage |
| 登录证明 | 浏览器本地签发 15 分钟 `CertFor=Self` PPT | 与 Leither 新的 `LoginWithPPT` 原生流程一致 |
| 会话 | Leither 一小时 SID；LifeDrive 在 45 分钟后使用设备密钥提前续期 | 日常使用无重复输入，同时缩短旧 SID 的实际使用时间 |
| 授权 | 读放行；写在每次 MApp 请求检查设备清单 | Leither 未向 MApp 提供可区分设备的调用，写入授权只能由应用判定 |
| 多设备 | 每台设备不同 UID、不同身份文件、相同 owner profile | 保持用户资料连续，又能精确撤销设备 |
| 公开名称 | 只作为 `<name>.lepan.org` 地址和显示名称 | 名称不是秘密，也不参与认证 |
| 旧密码机制 | 不迁移、不兼容，检测到旧状态立即停止 | LifeDrive 尚无真实用户，不值得永久保留两套认证面 |

没有采用以下方案：

- 不让 Leither 节点替浏览器签名；这只能证明服务器持有密钥，不能证明当前访问者是谁。
- 不把发布者或 `hostkey` 复制到浏览器；它们保留为节点侧发布与恢复权限。
- 不让多台设备复制同一个私钥；共享密钥无法只撤销其中一台设备。
- 当前不直接使用 WebAuthn/passkey；Leither 的登录合约要求 `sodiumv2` 证书和 PPT。以后可以用 passkey 作为解锁设备密钥的本地门槛，但不改变 Leither 身份格式。

## 2. 完整流程

```text
安装 LifeDrive
    ↓
Leither genkey 生成设备专用 sodiumv2 密钥
    ↓
Leither gencert 创建自签证书
    ↓
Leither signppt + verifyppt 生成并验证第一张 CertFor=Self PPT
    ↓
安装程序读取 PPT 的 SignerID（设备 UID）和 SignerInfo
    ↓
MMSetRight(authorSID, driveMID, deviceUID, 0x07276704)
    ↓
写出权限为 0600、绑定 App MID 的设备身份文件
    ↓
用户在目标浏览器导入一次
    ↓
浏览器验证 App MID、字段和公私钥配对关系
    ↓
私钥种子转换为不可导出的 WebCrypto CryptoKey
    ↓
每次认证在本地签发 15 分钟 PPT
    ↓
LoginWithPPT → 核对返回 UID → GetVar(SID, "userid")
    ↓
RunMApp → access_check → 读放行 / 写比对 lifeDrive.devices 设备清单
```

用户名密码注册、`auth.SetUserInfo(newpass)` 和 `Login(username, password)` 已从 LifeDrive 删除。

## 3. 设备签发与身份文件

### 登录凭证是哪一个文件

只有安装程序生成的 `*.identity.json` 是浏览器登录凭证。节点目录下的其他密钥文件都不是，导入它们会被拒绝。

| 文件 | 身份 | 用途 | 能否用于浏览器登录 |
|---|---|---|---|
| `<app>-<label>-<时间>.identity.json` | 每台设备独立的 `sodiumv2` UID | 浏览器登录 | **是，唯一凭证** |
| `hostkey.cfg` | 节点身份（`sodium` v1） | 发布应用、授予/撤销权限、以作者身份执行数据操作 | 否 |
| `gen8.key` | 某个用户密钥（`sodium` v1） | LifeDrive 不使用 | 否 |

设备 UID 与节点 UID 无关。`lifeDrive-owner.sh` 为每台设备用 `lpki genkey` 生成全新密钥，UID 取自该密钥自签 PPT 的 `SignerID`；`hostkey.cfg` 只是执行 `MMSetRight` 授权的那一方，不参与 UID 推导。例如同一节点上：

```text
hostkey.cfg → 9k81K7ksnPEDwIx4Z_2rwtpsbTM   节点，sodium v1
设备身份    → KE-BUXZaAPjywHz9dxDz04J1sx3   sodiumv2
```

浏览器登录合约要求 `sodiumv2`，而 `hostkey.cfg` 与 `gen8.key` 都是 `sodium` v1，因此即使有人想用它们登录也不可能成功。把节点密钥放进浏览器等于交出整个节点，绝不可这样做。

身份文件导入时按 `app_id` 校验，只能用于生成它的那一套 LifeDrive；换用其他应用的身份文件会提示 “That identity belongs to a different LifeDrive.”。

`lifeDrive-owner.sh` 使用 Leither 官方 `genkey`、`gencert`、`signppt` 和 `verifyppt`。LifeDrive 不自行编码证书；它复用 Leither 已验证 PPT 中的 `SignerInfo`，从而避免复制一套容易分叉的证书实现。

身份文件格式如下；示例值全部为占位符：

```json
{
  "type": "lifedrive-device-identity",
  "version": 1,
  "app_id": "<LifeDrive App MID>",
  "uid": "<27-character Leither key UID>",
  "label": "Alice's MacBook",
  "public_username": "alice",
  "key_type": "sodiumv2",
  "public_key": "<base64url Ed25519 public key>",
  "private_key": "<sensitive base64url Ed25519 secret>",
  "signer_info": "<Leither certificate encoding>",
  "created_at": "<UTC timestamp>"
}
```

安装程序的处理规则：

- 最终文件以 `0600` 权限写在 LifeDrive 应用目录之外；目标目录必须预先存在并且可写。
- 已存在的目标文件绝不覆盖。
- 临时密钥、证书和 PPT 在授权完成后删除。
- 身份文件不进入 Git、npm/GitHub 发布包、应用目录、AV1 请求或安装日志。
- 文件只有在弥媒授权成功后才移动到最终位置，避免交付一个尚未获准访问的身份。
- 身份文件本身包含明文私钥，`0600` 不是加密。它必须像 SSH 私钥一样安全传输并离线保管。

## 4. 浏览器配对与本地存储

导入时，浏览器按顺序执行：

1. 限制文件最大为 64 KiB，并验证 schema、字段长度、UID 和 base64url 格式。
2. 要求 `app_id` 等于当前页面的 LifeDrive App MID，避免把同一节点上另一套 LifeDrive 的身份导错位置。
3. 从 secret 的前 32 字节建立 RFC 8410 PKCS#8 Ed25519 私钥。
4. 导入文件内的公钥，随机生成挑战，执行一次签名和验证，确认公私钥配对。
5. 只把不可导出的私钥 `CryptoKey`、`SignerInfo` 和公开元数据写入 IndexedDB。
6. 以 App MID 作为 IndexedDB 记录键，所以同一浏览器来源可安全保存多套 LifeDrive 身份。

原始 JSON 私钥不写入 `localStorage`、`sessionStorage` 或应用状态。JavaScript 字符串无法可靠清零，但解码后的私钥字节、PKCS#8 缓冲区和验证挑战会在导入完成后清零。

“Sign out”和“Forget this device”的含义不同：

- **Sign out：** 删除并注销 SID，保留不可导出的设备密钥；下次可以直接重新认证。
- **Forget this device：** 先注销 SID，再删除当前 App MID 对应的浏览器密钥；以后必须重新导入离线身份文件。

## 5. PPT 与 SID 生命周期

浏览器使用 UTC 时间生成按键名排序的 PPT Data：

```text
CertFor=Self;EndTime=<UTC>;SignTime=<UTC>;
```

私钥只对该 Data 串执行 Ed25519 签名。生成的 PPT 包含 `Signature`、`SignerID`、`SignerInfo`、`KeyType: "sodiumv2"`、`Version: 1` 和 `Data`。

PPT 的有效期是 15 分钟，只在登录调用期间存在于内存。它不写入浏览器存储。LifeDrive 通过 Leither WebSocket `/ws/` 调用 `LoginWithPPT`，并要求：

```text
LoginWithPPT 返回的 Uid
    == 身份文件 uid
    == GetVar(sid, "userid")
```

任一值不一致即注销该 SID 并拒绝访问。成功后，浏览器在按 App MID 隔离的 `sessionStorage` 中保存 SID、UID、设备标签、公开名称和会话创建时间。

Leither SID 的内存有效期是一小时。LifeDrive 在会话达到 45 分钟后、下一次受保护请求开始前签发新 PPT 并换取新 SID，然后注销旧 SID。并发请求共享同一个续期过程，避免同时创建多个替代会话。页面重开且没有 SID 时，只要浏览器密钥仍在，也会自动认证。

## 6. 授权与共享 owner profile

每个设备使用不同 key UID。安装程序以应用发布者或节点身份创建固定 Drive 弥媒，并分别授权：

```text
MMCreate(authorSID, appMID, "lifedrive.app-drive/v1", "lifeDrive", 1, 0x07276704)
MMSetRight(authorSID, driveMID, deviceUID, 0x07276704)
```

MApp 不因 PPT 登录成功而跳过授权。每个请求仍然：

1. 用应用作者身份解析固定 Drive MID。
2. 用浏览器 SID 执行 `MMOpen`。
3. 读取检查 `last`；修改检查 `cur`。
4. 从同一 SID 读取 `userid`。

设备注册表只是节点侧的管理清单。

**已在 Leither V0.24.12 验证的实际行为，与最初设计意图不同，必须如实记录：**

- `MMOpen` 不构成授权边界。未授权 UID 同样可以打开 `last` 与 `cur`。
- `MMGetRight` 无法区分设备。节点侧 `mimei show` 能显示 `0x07276704` 的授权，但 MApp 可调用的 `MMGetRight` 对已授权与未授权 UID 一律返回默认的 `0x4`。
- `MMBackup` 仅限作者。持有 `0x07276704` 的设备与未授权设备都会收到 `1101:Permission denied`，因此设备无法自行提交。
- MApp 的全部文件操作都使用应用作者会话（`BELoginAsAuthor`）。Leither 只看到作者写入，看不到设备。

结论：Leither 当前没有向 MApp 提供任何可区分已授权设备的调用，所以**写入授权必须由 LifeDrive 自己判定**。

当前实现（`device_registry.go`）：

```text
读取（view）  → 不检查，任何通过认证的 UID 都可以读
写入（其他）  → 调用 deviceMayModify(uid)，比对安装程序维护的设备清单
```

设备清单是 Leither 服务目录下的 `lifeDrive.devices`，由 `lifeDrive-owner.sh` 追加、`lifeDrive-device.sh` 删除。检查 fail-closed：清单缺失、不可读或为空时，任何写入都被拒绝。

这是应用层强制、以作者密钥执行的授权，属于 [Leither technical foundation](../../reference/LEITHER_AND_MIMEI.md) 第 8 节描述的临时兼容层。持有作者密钥或能写入该文件的人可以绕过它，必须按 [LifeDrive security design](https://github.com/cfa532/LifeAlbum/blob/ade508c37bec3cd2d772e00f28d004eef5cdfde4/docs/LIFEDRIVE_INTERIM_SECURITY.md) 如实披露，不能声称由弥媒原生权限强制执行。

所有设备共享固定标记为 `owner-profile` 的 profile 弥媒。Profile 不再由设备 UID 派生，因此新增或轮换设备不会建立另一个用户资料。第一次访问 profile 时，当前已授权设备会获得该 profile 的权限。

## 7. 日常操作

首次安装：

```bash
npx --yes @inoku/lifedrive
```

指定身份文件输出位置：

```bash
npx --yes @inoku/lifedrive --identity-out /secure/offline/alice-mac.identity.json
```

增加独立设备：

```bash
npx --yes @inoku/lifedrive --add-device "Alice's iPhone"
```

列出或撤销设备：

```bash
npx --yes @inoku/lifedrive --list-devices
npx --yes @inoku/lifedrive --revoke-device <device-uid>
```

管理命令直接使用下载包内的新脚本，但不会替换或发布 LifeDrive 应用文件。安装程序拒绝同时指定 `--upgrade`、`--add-device`、`--list-devices` 或 `--revoke-device` 中的多个操作。

撤销规则：

- 至少必须保留另一台已登记设备，防止误撤销最后的日常入口。
- 目标 UID 的 Drive 和 `owner-profile` 权限都会设置为 `0`。
- 撤销不删除目标浏览器或离线文件里的私钥。

**已验证（V0.24.12）：撤销立即阻止写入，但不阻止读取。** `lifeDrive-device.sh revoke` 从 `lifeDrive.devices` 移除该 UID，之后该设备的任何修改请求都在 `access_check` 被拒绝。读取不受影响——Leither 给任何通过认证的 UID 保留默认读权限，这是本产品接受的设计。`MMSetRight(uid, 0)` 对读写都没有实际作用（`MMGetRight` 仍返回 `0x4`），保留它只是为了在原生权限可用时不必重建登记流程。

## 8. 明确的非兼容变更

这是一次有意的 clean break：

- Owner 状态格式改为 `lifedrive-key-auth-v1`。
- 浏览器会话格式改为按 App MID 隔离的 v3 格式。
- Owner profile schema 改为 v4，并使用固定的 `owner-profile` 标记。
- 安装和升级检测到密码时代的 owner 状态时会停止，不创建兼容账号，也不自动迁移。
- 旧的用户名/密码页面、`Login` 调用、密码重置和 `--recover` 路径已删除。

这样可以避免长期承担双认证入口、迁移歧义和降级攻击面。因为 LifeDrive 目前没有真实用户，旧原型状态应作为新安装处理。

## 9. 恢复与安全边界

- SID 丢失或达到续期时间：浏览器用现有设备密钥签发新 PPT。
- 浏览器网站数据丢失：重新导入该设备的离线身份文件。
- 设备与其身份文件同时丢失：在物理 Leither 节点新增替代设备，然后撤销旧 UID。
- 发布者或节点密钥只用于应用发布和物理节点恢复，不作为日常浏览器身份。
- 身份文件被复制：复制者拥有该设备私钥；应立即签发替代设备并撤销旧 UID。
- 时钟严重不准：PPT 可能尚未生效或已经过期；浏览器和 Leither 节点都应保持 UTC 时间同步。

生产访问必须使用 HTTPS/WSS。WebCrypto Ed25519 只在安全浏览器来源（或 localhost）可用；使用 `--skip-domain` 的 LAN 部署必须自行提供可信 HTTPS 反向代理。

PPT 在 15 分钟有效期内是 bearer credential，必须防止日志记录、URL 拼接和网络窃听。不可导出的 `CryptoKey` 可以阻止普通导出，但不能抵御已在 LifeDrive 来源中执行的恶意 JavaScript；CSP、依赖供应链和发布完整性仍然属于认证边界。

当前不承诺硬件设备证明、passkey 用户验证、跨节点身份同步或远程设备邀请。这些能力可以在不恢复密码登录的情况下作为后续层加入。

## 10. 实现位置与验证范围

| 位置 | 职责 |
|---|---|
| `apps/lifedrive-server/lifeDrive-owner.sh` | 生成 Leither 密钥/证书/PPT、授权 UID、输出身份文件 |
| `apps/lifedrive-server/lifeDrive-device.sh` | 列出和撤销设备 |
| `apps/lifedrive-server/lifeDrive-setup.sh` | 初始安装、设备管理和 clean-break 检查 |
| `apps/inoku-web-demo/src/deviceIdentity.ts` | 身份导入、KeyPair 验证、IndexedDB 和 PPT 签名 |
| `apps/inoku-web-demo/src/leitherClient.ts` | `LoginWithPPT`、UID 核对、SID 续期与注销 |
| `apps/lifedrive-server/access_entries.go` | 每次请求的授权边界：读放行，写调用设备清单检查 |
| `apps/lifedrive-server/device_registry.go` | 读取 `lifeDrive.devices`，fail-closed 判定设备可否写入 |
| `apps/lifedrive-server/profile_entries.go` | 跨设备稳定的 owner profile |

当前实现已经通过 TypeScript 静态检查、Vue 生产构建、Go 编译、Shell 语法检查和 Git diff 检查。

已在 Leither V0.24.12 节点完成的运行时验证：

| 项目 | 结果 |
|---|---|
| `lifeDrive-owner.sh` 端到端签发身份文件 | 通过（修复密钥文件解析后） |
| 浏览器本地签发 UTC PPT 并 `LoginWithPPT` | 通过 |
| UID 与 `GetVar(sid,"userid")` 核对 | 通过 |
| SID 丢失后使用浏览器密钥静默重新认证 | 通过 |
| `MMCreate` 确定性解析同一 Drive MID | 通过 |
| 提交后读取可见（`MFSetCid` + `MMBackup`） | 通过 |
| 未授权设备被拒绝写入 | 通过：登录即被拒绝，提示“未授权” |
| `--revoke-device` 阻止写入 | 通过：移出清单后修改请求被拒绝 |
| 提交后读取可见（`MFSetCid`） | 通过：暂存树 CID 与弥媒新版本一致 |

未验证：真实撤销后的拒绝行为、多设备并发续期、上传分片路径。
