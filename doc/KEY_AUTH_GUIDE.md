# Leither 密钥认证接入指南

> 版本：2026-09-06（T1 发布契约面，lapi v0.3.0 结构化签名请求已并入；F1/F3/F4a/b/c
> 修订已并入正文；线格式勘误：LoginReply 无 Isystem 字段——以 §3 结构为准）。
> 本文档自包含：外部应用/agent **仅凭本文档**
> 即可完成「密钥即身份」的用户认证接入。文档中所有命令均经真实 CLI 实测。

## 0. 承诺范围

**本期发布（对外承诺）：**

- **CertFor=Self 密钥登录**——应用用户持私钥自签通行证（PPT）登录节点；
- lapi `IAuth` 七个方法（`LoginWithPPT` / `Logout` / `SetUserInfo` / `SignPPT` /
  `Sign` / `PPTStr2Map` / `SignInfo2Map`）；
- `lpki` 五个命令（`genkey` / `genpk` / `gencert` / `signppt` / `verifyppt`）；
- 会话语义（sid、TTL）+ 身份读取 `GetVar(sid, "userid")`。

**范围限定（T0 标准措辞）：本功能在本地节点可用；跨节点同步属未承诺范围。**
认证数据存于节点本地 MM 数据库，不在节点间自动同步。

**正向事实（可依赖）：**

- **同一密钥可在任意节点登录**——认证是纯密码学验签，无需注册、无中心；
- **userid 全网一致**——userid 由公钥单向推导（见 §9），与节点无关。

**不在承诺面（内部/实验，勿依赖）：** Bind 型第三方身份绑定登录（`CertFor=Bind`）、
byname 用户名密码登录。相关说明见 §7 安全须知。

## 1. 核心概念

### 1.1 密钥即身份

- 算法：**Ed25519**。两种密钥类型：`sodiumv2`（默认，RIPEMD-160 推导 id）与
  `sodium`（旧型，SHA-256 截断 20 字节推导 id）。两型均已全网共识、不可变更；
  新应用一律用 `sodiumv2`（`genkey` 默认即 v2）。
- **userid = 密钥 id**（20 字节的 base64url 无填充编码，27 字符，
  形如 `e1P5pQnueXJZQSU6XdCt5IP1DuC`）。
- 无注册流程：持私钥即是用户。第一次登录时节点自动建立本地用户档案。

### 1.2 PPT（通行证）

PPT 是**有时效的签名 JSON 文档**，结构：

```json
{
  "Signature":  "<base64 签名>",
  "SignerID":   "<签名者 key id = userid>",
  "SignerInfo": "<base64 编码的签名者证书>",
  "KeyType":    "sodiumv2",
  "Version":    1,
  "Data":       "CertFor=Self;EndTime=2026-09-01T10:38:18Z;SignTime=2026-09-01T09:38:18Z;"
}
```

- `Data` 为 `k=v;` 形式的排序拼接（保留键见 §9）。
- **登录用 PPT 必须含 `CertFor=Self`**；服务端验签后 userid 即取 `SignerID`。
- `SignTime`/`EndTime` 为 **RFC3339 UTC**（如 `2026-09-01T10:38:18Z`），签发统一
  写 UTC，含显式 UTC 标识，跨节点无歧义；旧格式（`yyyyMMddHHmmss` + 时区缩写）
  仍可被校验端兼容读取（2026-09-02 前的存量 PPT）。有效期防重放。

### 1.3 会话（sid）

- 登录成功返回 `sid`——bearer token，之后所有 RPC 调用首参带 sid。
- 会话存于节点内存，**TTL 1 小时**，节点重启即全部失效（客户端需重新登录）。
- 读取当前身份：`GetVar(sid, "userid")`。
- `Logout` 为异步释放：返回后极短时间内 sid 可能仍可用，客户端不应依赖
  「Logout 后立即失效」做强时序断言。
- **会话与 PPT 有效期相互独立**：PPT 只在登录瞬间用于验签；会话建立后，
  PPT 过期**不影响**已建立的会话（会话由自身 1 小时 TTL 控制）。因此
  「PPT 短时效」不影响长会话的使用体验。

## 2. 五分钟接入（CLI 路径，语言无关）

以下命令用 2026-09-01 构建的二进制逐一实测通过；签发时间格式后随 F1（2026-09-02）
统一为 §1.2 的 RFC3339 UTC（命令与旗标不变，仅输出时间格式更新）：

```bash
# 1) 生成私钥（文件自动 0600 权限；JSON 两段式明文，妥善保管）
./Leither lpki genkey -o user.key

# 2) 生成自签名证书（绑定身份元信息）
./Leither lpki gencert -k user.key -m "name=alice" -o user.ca

# 3) 签发登录 PPT：-p 有效期（分钟，默认 4320=72h）；登录必须带 CertFor=Self
./Leither lpki signppt -c user.ca -p 60 -m "CertFor=Self" -o login.ppt

# 4) 本地验签（可选自检）
./Leither lpki verifyppt -c user.ca -i login.ppt   # 成功输出 "ppt is valid"，
                                                   # 失败以非零退出码 + 错误文本退出
```

导出公钥（分发/登记用，文件只含公钥）：

```bash
./Leither lpki genpk -i user.key -o user.pub
```

## 3. 登录调用（RPC）

Leither 节点对外暴露 hprose RPC（WebSocket）：`ws://<host>:<port>/ws/`。
**WS 与 HTTP 恒同端口**（均为节点的 `ServicePort`，默认 4800；节点改
`ServicePort` 时两者同步跟随）。方法名与 Go 接口一一对应。

**Go 示例**（`github.com/hprose/hprose-golang/v3`，实测于 **v3.0.16**——
注意 hprose-golang 的 v2 与 v3 API 差异很大，务必用 v3）：

```go
import (
    "github.com/hprose/hprose-golang/v3/rpc/core"
    "github.com/hprose/hprose-golang/v3/rpc/websocket" // 非 blank import
)

// LoginReply 是 RPC 返回的线格式（JSON 字段名首字母大写，与 Go 导出字段一致）：
//   {"Uid":"<userid>","Sid":"<40 位 hex 会话 id>"}
// （无 Isystem 字段——系统用户判定不通过登录返回，勿依赖。）
type LoginReply struct {
    Uid string // 当前用户 id（CertFor=Self 登录时 = 密钥 id）
    Sid string // 会话 id，后续所有 RPC 调用的首参
}

// nodeStub 用 v3 类型化代理：字段即方法名，签名与节点方法一一对应
type nodeStub struct {
    LoginWithPPT func(ppt string) (*LoginReply, error)
    GetVar       func(sid, name string) (string, error)
    Logout       func(sid, info string) error
}

func login(nodeWS, ppt string) (*nodeStub, *LoginReply, error) {
    // 关键：v3 的 websocket 包不自动注册传输，必须显式调 RegisterTransport()，
    // 否则 NewClient 报 "unsupported protocol: ws"。
    websocket.RegisterTransport()

    client := core.NewClient(nodeWS) // 形如 ws://127.0.0.1:4800/ws/
    var stub *nodeStub
    client.UseService(&stub)

    reply, err := stub.LoginWithPPT(ppt)
    if err != nil {
        return nil, nil, err
    }
    return stub, reply, nil
}
```

登录后用 sid 调业务 API，例如读取身份并登出：

```go
stub, reply, err := login("ws://127.0.0.1:4800/ws/", ppt)
uid, err := stub.GetVar(reply.Sid, "userid") // == 密钥 id
err = stub.Logout(reply.Sid, "")             // 异步释放（§1.3）
```

**PPT 字符串来源**：`lpki signppt -o login.ppt` 产出的文件内容即 JSON 原文，
原样读入字符串传入即可（容忍首尾空白/换行，建议 `strings.TrimSpace`）。

**完整可运行参考实现**：`example/key-auth-login/`（`main.go` + `go.mod`，
经真实节点实测：LoginWithPPT → GetVar("userid") → Logout 全链路通过）。

**等价的 `Login` 形式**：`Login(ppt, "", "byppt")` 与 `LoginWithPPT(ppt)` 同效。
三个参数依次为：凭证（byppt 时 = PPT 字符串）、口令（byppt 时留空占位 `""`）、
登录类型常量 `"byppt"`。

**错误语义**：认证失败以 RPC error 返回（无数字错误码，错误为文本）。
实测典型文本：
- PPT 非 JSON/格式损坏 → `readSignInfoFromString: ... invalid character ...`
- `CertFor` 缺失或非法 → `Login type[] invalid`
- 签名无效（Data 被篡改/密钥不匹配） → `checkSign false`
- PPT 过期 → 过期报错（含有效期信息）
客户端应**以「是否返回 error」判定成败**，错误文本仅供人读，勿做脆弱的
字符串匹配（文本未来可能调整）。

**其他语言**：hprose 有 Python / JavaScript / Java / PHP / C# 客户端，
方法名、参数顺序同上。

## 4. lapi IAuth 参考

`lapi.IAuth`（`Leither/lapi` 模块；外部 RPC 同名方法语义一致）：

| 方法 | 说明 |
|---|---|
| `LoginWithPPT(strPPT string) (*LoginReply, error)` | PPT 登录。服务端验签 + 有效期检查；`CertFor=Self` 时身份=签名者 key id。PPT 无效/过期/类型非法均报错 |
| `Logout(sid, info string) error` | 登出（异步释放，见 §1.3） |
| `SetUserInfo(sid string, info map[string]string) error` | 更新用户档案字段。**已知行为：固定键名（如 `name` 以外的保留字段）可能报错**；仅写自定义业务字段 |
| `SignPPT(sid string, req *SignPPTRequest) (string, error)` | 以当前登录用户身份签 PPT（**F4b 起结构化请求**：`req.Info` 信息键值、`req.Period` 有效期——**分钟，合法区间 (0, 10080]（7 天）**、`req.Purpose` 用途声明入审计、`req.AidDecl` 应用身份声明——与服务端导出值不一致即拒绝）。普通用户会话中保留键（`Userid`/`BindID`/`BindType`/`SubType`/`CertPK`/`CertPKID`/`SignTime`/`EndTime`/`AppID`/`NodeId`）会被服务端剥离——客户端不得依赖经此通道铸造这些字段 |
| `Sign(sid string, req *SignRequest) ([]byte, error)` | 以当前用户私钥对消息签名（Ed25519；**F4b 起结构化请求**：`req.Content` 被签内容、`req.Purpose`/`req.AidDecl` 同上。**盲签裸 message 入口已移除**——api 层另留 api-only `SignRaw`，不属对外承诺面） |
| `PPTStr2Map(strPPT string) (map[string]string, error)` | 解析 PPT 的 Data 为键值 map（本地操作，不验签） |
| `SignInfo2Map(strSignInfo string) (map[string]string, error)` | 解析签名文档为 map |

**身份读取的补充说明**：§1.3/§3 用到的 `GetVar(sid, name string)` 属于节点的
**通用变量接口**（非 IAuth 七方法之一，但在对外承诺面内——本文档承诺的用法是
`GetVar(sid, "userid")` 读取当前登录身份，返回用户 id 字符串；对不存在的键
返回空串而非报错）。`GetVar` 的完整键清单见《Leither 应用开发指南》§6.2。

**CLI 与 RPC 铸造 PPT 的保留键差异（刻意设计，勿混淆）**：
- **CLI `signppt -m "..."`** 是**本地离线铸造**——你持有私钥，可写入任意字段
  （登录 PPT 必须显式带 `CertFor=Self`，正因 CLI 不自动注入）；
- **RPC `SignPPT`** 是**受约束的会话内铸造**——服务端为防已登录用户伪造身份
  声明，会剥离保留键（见上表 `SignPPT` 行）。
两者不矛盾：CLI 铸造 = 你就是私钥主人，RPC 铸造 = 服务端替你签且不信任你填的身份字段。

## 5. HTTP 便捷接口

节点所有者本机取登录 PPT（签的是**节点系统身份**，用于节点管理，不适合应用用户）：

```bash
curl "http://127.0.0.1:4800/getvar?name=ppt&arg0=1440&nojson"   # arg0=有效期分钟，缺省 1440=24h
```

应用用户的 PPT 应由用户私钥在客户端侧签（§2 或程序化），不走此接口。

## 6. 单位速查（2026-09-01 起全表面统一为**分钟**）

| 入口 | 参数 | 单位 |
|---|---|---|
| `lapi.SignPPT` / RPC `SignPPT` | `period` | 分钟，(0, 10080] |
| `lapi.BESignPPT`（后端） | `period` | 分钟 |
| `lpki signppt -p` | 有效期 | 分钟，默认 4320 |
| `getvar?name=ppt&arg0=N` | 有效期 | 分钟，缺省 1440 |

## 7. 安全须知

1. **私钥即身份**：丢失 = 永久失去该身份；泄露 = 可被冒名。密钥文件为明文
   JSON（权限 0600），请按私钥材料同等标准保管、备份。
2. **PPT 是 bearer token**：有效期内任何持有者可用；**不绑定签发/使用节点**，
   有效期内可跨节点重放。请使用短时效（建议分钟级，如 15–60 分钟），
   过期后重新签发。
3. **时间字段**：`SignTime`/`EndTime` 自 2026-09-02 起签发统一为 **RFC3339 UTC**
   （如 `2026-09-01T10:38:18Z`），含显式 UTC 标识，无时区缩写歧义，跨节点可移植。
   校验端对旧格式（`yyyyMMddHHmmss` + 时区缩写，如 `20260901183818CST`）仍按
   兼容读取；旧格式按签发机时区表解析，跨时区场景仍建议把 PPT 有效期设短。
4. **byname（用户名密码）为遗留路径，不在本期承诺面**：口令存储自
   2026-09-01 起为 argon2id 哈希（存量明文记录在首次成功登录时透明迁移）；
   空口令不可注册也不可登录。新应用请使用密钥方案。
5. **临时用户与登录有节流保护（2026-09-02 起）**：临时用户（未注册身份首次登录
   自动创建）有并发配额上限（节点内 4096 个，满额淘汰最久者，TTL 30 分钟惰性回收）；
   登录接口按身份指纹（10 次/分钟）、节点全局（120 次/分钟）与新建临时用户
   （30 次/分钟）三档限流，超限返回 `login rate limit exceeded`。会话仍为内存态
   无持久化。
6. **签名审计（F4c，2026-09-04 起）**：节点对 RPC `Sign`/`SignPPT` 全量审计
   （成功/失败均记；记录含被签内容全文 + sha256 摘要 + 会话来源标注），JSONL 落
   `logs/signaudit.log`（独立轮转，缺省 64MB×90 天）。线上调用 `aid` 恒空
   （仅容器内 MApp 调用经遮蔽注入导出值）。审计查询：

   ```bash
   Leither lpki signaudit --uid <uid> --limit 20          # 文本模式
   Leither lpki signaudit --op signppt --since 24h --json # JSON envelope
   ```

   普通用户只能查自己的记录（服务端强制）；系统用户可查全部。
   审计写入失败会降级进主日志（`[SIGNAUDIT]`）而不阻断签名。
7. **aid 级签名授权（F4b，2026-09-04 起）**：节点对每个「用户 × 应用」的代签做
   名单判定。本期缺省**姿态 b（deny 名单读法）**——无记录=放行（零迁移 grandfather）；
   应用会话（app 会话）代签**本期一律拒绝**（`uid=appid` 歧义防绕过），容器内 MApp
   以调用方用户会话执行、不受影响。声明身份与导出不符即拒绝
   （审计 `rule=signacl:aid-spoof`）。管理面（每用户自助域）：

   ```bash
   Leither lpki signacl deny  <aid> -k user.key   # 拒绝该应用代我签名
   Leither lpki signacl unset <aid> -k user.key   # 撤销
   Leither lpki signacl list  --json -k user.key  # 查看我的名单
   ```

   deny 诊断：`Leither lpki signaudit --decision deny [--rule signacl:deny-list]`。
   收紧姿态（`SystemVars.json` 的 `SignACL.Mode="a"` + `TightenSince` cutoff）与
   安装期勾选 UX 待 GrantServices 授权体系设计定稿后开放，本期勿切。

## 8. addkey 私钥驻留边界（F4a，2026-09-03 起）

`lpki addkey` 用于把**私钥**上传给节点（落库为用户私钥，供节点代签）。为收窄私钥驻留
边界，节点端按**调用来源**分层受理（裁决：`T1/f4a-implementation-plan.md` §2.1）：

| 来源 | 处置 |
|---|---|
| 进程内直调（容器 MApp/宿主，无 RPC 上下文） | 放行（与节点同机同信任域） |
| 回环/本机（127.0.0.1、::1、本机网卡地址） | 放行，**明文直传**（现状不变） |
| 同网段（客户端 IP ∈ 本机任一接口子网） | 默认放行，但**只受理信封形态**（`addkeysealed`）；`SystemVars.AddkeyAllowLAN=false` 即关闭（偏执模式） |
| 其余（跨网段私网/公网/不可分类，即 WAN 远程） | **拒绝**（保持关闭） |

**信封（LAK1）**：目标节点的 X25519 加密子钥由节点 Ed25519 身份私钥纯 stdlib 派生
（SHA-512(seed) 钳制，`crypto/ecdh`），经 `GetVar(sid,"hostpk","","x25519")` 发布；
信封 = ECDH 临时钥 + HKDF-SHA256（info/AAD 绑定 hostid）+ chacha20poly1305。
**误投他节点必然解密失败**（hostid 进 AAD）。CLI 侧 TOFU 指纹钉住加密公钥
（`<WorkPath>/known_nodes.json`，0600）：首连展示指纹交互确认（脚本用 `--yes` 免确认），
公钥变更即拒绝（需显式删除记录行再信任）。

CLI 行为一览：

```bash
Leither lpki addkey keyfile                      # 本机回环：明文直传（现状）
Leither lpki addkey keyfile --local              # 显式声明最严档：锁死本机回环目标
Leither lpki addkey keyfile -n 192.168.1.10:4800 # 同网段：TOFU 首连确认 → 信封上传
Leither lpki addkey keyfile -n <公网节点>        # 拒绝（WAN 远程 addkey 关闭）
```

兼容性注意：

- **旧 CLI → 新节点**：同网段明文上传被节点端拒绝（信封强制），报
  `plaintext addkey from LAN rejected`——需升级 CLI。
- **新 CLI → 旧节点**：`hostpk x25519` 不被识别，报"目标节点版本过旧"。
- **身份白名单**（`SystemVars.AddkeyAllowUsers`，会话 uid 列表）：裁决为"仅入文档、
  默认空、有需求再实现"——字段已定义，判定**未实现**。
- `/ws3/`（hprose v3）上 `SetUserInfo` 因注入参数改变方法表（树内无 v3 调用方）；
  `/ws/`、`/webapi/` 线签名不变。
- 残余风险：节点前有反代/隧道时 `RemoteAddr` 是代理地址（来源判定语义需重估）；
  docker bridge 等容器网络下"同机子网"可能把同宿主容器算同网段——偏执模式可关闭。
- 观测：受理/拒绝审计日志打 `[F4A-ADDKEY]` 标签（来源 IP、会话 uid、层级、明文/信封）。

## 9. 附录：PPT 线格式（供非 Go 实现者）

自行实现 PPT 签发/校验时的精确规则：

- **密钥**：Ed25519（RFC 8032）。私钥 64 字节（32B 种子 + 32B 公钥），
  公钥 32 字节。
- **key id（=userid）**：`sodiumv2` = RIPEMD-160(公钥) 的 20 字节，
  base64url **无填充**编码（27 字符）；`sodium`（旧型）= SHA-256(公钥)
  前 20 字节同法编码。
- **Data 序列化**：全部键值按键名字典序排序，拼接为 `k1=v1;k2=v2;…`
  （每个键值对以 `;` 结尾；**键与值均不得含 `=` `;`**，无转义机制；
  重复键解析即报错）。
- **时间字段**：`SignTime`/`EndTime` 为 **RFC3339 UTC**（如 `2026-09-02T04:00:00Z`），
  签发时按 UTC 归一化写入；校验方须同时兼容读取旧格式 `yyyyMMddHHmmss` + 时区缩写
  （2026-09-02 前的存量 PPT，按签发机时区表解析）。
- **签名**：Ed25519 对「Data 序列化串」签名（服务端实现细节：`Signature`
  覆盖排序后的 Data 字符串）。
- **PPT JSON 字段**：`Signature`（base64 标准编码）、`SignerID`、
  `SignerInfo`（签名者证书的 gob+base64 编码，含证书公钥与 `CertPK`/
  `CertPKID` 声明，服务端从中取验签公钥）、`KeyType`、`Version`、`Data`。
- **保留键**：`CertFor`（`Self`/`Gift`/`Bind`/`App`）、`SignTime`、`EndTime`、
  `Userid`、`BindID`、`BindType`、`SubType`、`CertPK`、`CertPKID`、`AppID`、`NodeId`。

> 多数接入方无需自行实现：CLI（§2）+ RPC（§3）已覆盖全部主流语言。
> 本附录供需要在受限环境（如纯前端/无 CLI）内实现的少数场景参考。
