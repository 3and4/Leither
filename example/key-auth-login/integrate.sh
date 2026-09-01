#!/usr/bin/env bash
# integrate.sh — 端到端生成测试用户密钥与登录 PPT（依据《Leither 密钥认证接入指南》§2）。
#
# 输入环境变量：
#   LEITHER_BIN  Leither CLI 二进制路径（必填）
#   WORKDIR      产物输出目录（默认：脚本所在目录下的 out/）
#   PPT_MINUTES  PPT 有效期（分钟，默认 60）
#
# 产物（$WORKDIR 下）：
#   user.key  私钥（0600，JSON 两段式明文，妥善保管）
#   user.ca   自签名证书
#   login.ppt 登录 PPT（CertFor=Self）
#   user.pub  公钥（分发/登记用）
#
# 之后用 login_client 实测登录：
#   NODE_WS=ws://127.0.0.1:4800/ws/ go run ./login_client -ppt "$WORKDIR/login.ppt"

set -euo pipefail

if [[ -z "${LEITHER_BIN:-}" ]]; then
    echo "ERROR: 环境变量 LEITHER_BIN 未设置（Leither CLI 二进制路径）" >&2
    exit 2
fi
if [[ ! -x "$LEITHER_BIN" ]]; then
    echo "ERROR: LEITHER_BIN 不可执行: $LEITHER_BIN" >&2
    exit 2
fi

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
WORKDIR="${WORKDIR:-$SCRIPT_DIR/out}"
PPT_MINUTES="${PPT_MINUTES:-60}"

mkdir -p "$WORKDIR"
cd "$WORKDIR"

echo "==> [1/5] 生成私钥 (Ed25519/sodiumv2, genkey 默认即 v2)"
"$LEITHER_BIN" lpki genkey -o user.key

echo "==> [2/5] 生成自签名证书"
"$LEITHER_BIN" lpki gencert -k user.key -m "name=t1-rehearsal-user" -o user.ca

echo "==> [3/5] 签发登录 PPT (CertFor=Self, 有效期 ${PPT_MINUTES} 分钟)"
"$LEITHER_BIN" lpki signppt -c user.ca -p "$PPT_MINUTES" -m "CertFor=Self" -o login.ppt

echo "==> [4/5] 本地验签自检（期望输出 \"ppt is valid\"）"
"$LEITHER_BIN" lpki verifyppt -c user.ca -i login.ppt

echo "==> [5/5] 导出公钥"
"$LEITHER_BIN" lpki genpk -i user.key -o user.pub

echo
echo "完成。产物位于 $WORKDIR:"
ls -l user.key user.ca login.ppt user.pub
echo
echo "下一步：NODE_WS=ws://127.0.0.1:4800/ws/ go run $SCRIPT_DIR/login_client -ppt $WORKDIR/login.ppt"
