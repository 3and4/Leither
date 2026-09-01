# key-auth-login — 密钥认证最小参考实现

《Leither 密钥认证接入指南》（`../../KEY_AUTH_GUIDE.md`）的配套可运行示例，
2026-09-01 经真实节点（V0.24.10）实测全链路通过：
genkey → gencert → signppt(CertFor=Self) → verifyppt → LoginWithPPT →
GetVar("userid") → Logout。

## 用法

```bash
# 1) 生成测试身份与登录 PPT（CLI 路径）
LEITHER_BIN=/path/to/Leither ./integrate.sh        # 产物在 out/

# 2) 起节点（或复用已有节点），拿到 ws 地址（默认端口 4800）

# 3) RPC 登录实测
NODE_WS=ws://127.0.0.1:4800/ws/ go run . -ppt out/login.ppt
```

成功输出形如：

```
LoginWithPPT OK: uid=GxeV... sid=0429... isystem=false
GetVar userid = GxeV...        # == uid == 密钥 id
Logout OK
```

## 文件

- `integrate.sh` — CLI 侧：生成密钥/证书/登录 PPT/公钥（对应指南 §2）
- `main.go` — RPC 侧：hprose v3 类型化代理登录（对应指南 §3，含
  `websocket.RegisterTransport()` 关键调用）
- `go.mod` — 依赖 `github.com/hprose/hprose-golang/v3`（实测 v3.0.16）
