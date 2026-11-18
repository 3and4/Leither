# 身份、会话与跨节点调用

> **状态**：本文描述 Leither **当前实现**的身份 / 会话 / 跨节点调用边界，供外部应用（含 MApp）接入时判读。
> 标注「未承诺」的内容**不属于** [`doc/KEY_AUTH_GUIDE.md`](../doc/KEY_AUTH_GUIDE.md) 的对外承诺面；
> 两者冲突时以该指南为准。

## 1. 四类身份不要混

| 身份 | 载体 | 有效范围 | 建立方式 |
|---|---|---|---|
| 用户身份 | 会话里的 `userid` | **签发该会话的节点** | `LoginWithPPT` / `Login` 等；`GetVar(sid, "userid")` 读取 |
| 应用作者身份 | 会话里的 `userid`（= 应用记录的 `Author`） | **仅该节点** | MApp 后端 `BELoginAsAuthor()` |
| 应用身份 | 会话里的 `appid` | **仅该节点** | MApp 后端 `BELoginAsApp()` |
| 节点身份 | 节点间登录会话的 `userid`（= 调用节点的系统 key id） | 目标节点 | 节点间 `NLogin`（节点 PPT） |

- 认证数据与会话都在**节点本地**，不在节点间同步；会话是内存态（TTL 1 小时，节点重启即失效）。
- 「同一密钥可在任意节点登录、userid 由公钥单向推导且全网一致」是密钥认证的对外承诺；
  **会话本身不跨节点**，也不存在「把 sid 拿到另一节点继续用」。

## 2. 应用作者身份（仅 MApp 后端）

- `BELoginAsAuthor()` / `BELoginAsApp()` / `BESignPPT()` / `BESign()` / `BEOpenAppData*()` /
  `Session*()` 只存在于 MApp 进程内的 `lapi` 面；**不在** `/webapi/`、`/ws/`、`/ws3/` 的
  RPC 面上（RPC 面只暴露 `WebApi` 的具体方法）。
- `BELoginAsAuthor()` 返回**本节点内存会话**：uid = 应用记录里的 `Author`，不需要应用作者私钥；
  2026-09-03 起已冻结（**冻而不废**：对外不再演进、禁止新增依赖，删除走观测驱动）。
- `BELoginAsApp()` 返回的会话 uid = **appid**，不继承作者全权；`BEOpenAppDataApp()` 打开的
  应用数据区仍按应用记录里的作者归属。
- `BESignPPT()` / `BESign()` 由**节点系统钥**签名（不是应用作者钥）；`BESignPPT` 固定
  `CertFor=App`，而登录侧只接受 `Self` / `Bind` / `Gift` ⇒ 它签出的不是可登录凭证。

## 3. 跨节点 `RunMApp`

跨节点调用由**发起节点**判权（调用者 uid 必须是目标应用作者，或发起节点的系统身份），
然后转发到目标节点：

- 转发前**删除 `nid` 和 `sid`**；目标节点按 guest 执行（除非请求另带可验证凭证，如 `ppt`，
  由目标节点 `ParamAuth` 重新认证并换成本地会话）。
- 目标节点**不重复判权**，也拿不到调用方的 `sid`。
- 目标节点 `Request` 中各键的来源：

| 键 | 来源 | 可信性 |
|---|---|---|
| `author` | 节点按本地应用记录注入（覆盖调用方传值） | 以本地记录为准 |
| `mmroot` | 节点按 `Manifest.MMRoot` 注入（仅清单声明时覆盖） | 清单声明时可信 |
| `nodeappcode` | 平台向目标节点取的握手码（见 §4） | **bearer 码，不是节点证明** |
| `sidapp` | 仅本地 CLI `lpki runapp` 设置；容器 / 跨节点不注入 | 不可信 |
| `nid` | 调用方填的**目标节点**选择器，转发前删除 | 不是调用方身份 |

> 需要目标端确认调用方时，`nodeappcode` 只能证明「出示者知道这串码」，不能证明
> 「是哪个节点在调用」。

## 4. `nodeappcode`

- 铸造入口：`GetVar(sid, "nodeappcode", <appid>)`（`getvar?name=nodeappcode&arg0=<appid>`）。
- 语义：接收节点按 `(调用方会话 uid, appid)` 懒创建一个会话，有效期约 1 小时，写入
  `nodeid`（铸造时调用方会话的 uid，可能是普通用户 uid）与 `forapp`（appid）。
- 消费：`SessionGet(code, "nodeid")` / `SessionGet(code, "forapp")`。
- **不是权限凭证**：任何可解析会话（含普通用户）都能铸造；`forapp` 由消费方自行校验，
  平台不做消费侧强制；码泄露即可被出示。

## 5. 会话与凭证的边界

- 会话：节点内存、TTL 1 小时、重启失效；`Logout` 为异步释放。
- PPT：bearer 凭据，**不绑定签发 / 使用节点**；有效期以 `EndTime` 为准。
- byname（用户名密码）与 `Register`：**内部 / 实验面，未承诺**（注册接口无鉴权、无节流）；
  新应用请使用密钥 / PPT 方案。
- `mmsid`：数据区句柄，**永不作登录身份**。
