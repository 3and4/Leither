// signreq_test.go — F4b L0 契约（G5 死线）：结构化签名请求 + Y1 修复测试。
//
// 裁决：D10（SignPPT 同改结构化，一次付清）、D11 变体 A'（lapi 面移除盲签裸 Sign，
// api 层保留 api-only SignRaw）、D14=B（契约步升审计记录 v2）。
package lapi

import (

	"testing"
)

// TestLApiStub_Sign_Unwired Y1 修复：nil 函数字段不再静默假成功，返回 errUnwired
// （对齐 BE 面先例 backend.go/backend_test.go——只断言非 nil）。
func TestLApiStub_Sign_Unwired(t *testing.T) {
	stub := &LApiStub{}
	if _, err := stub.Sign("s", &SignRequest{Content: []byte("m")}); err == nil {
		t.Error("Sign() unwired: want error, got nil（静默假成功）")
	}
	if _, err := stub.SignPPT("s", &SignPPTRequest{Info: map[string]string{"a": "b"}}); err == nil {
		t.Error("SignPPT() unwired: want error, got nil（静默假成功）")
	}
}

// TestLApiStub_Sign_Dispatch 结构化请求透传分发：stub 字段收到原样 req 指针内容。
func TestLApiStub_Sign_Dispatch(t *testing.T) {
	var gotSign *SignRequest
	var gotPPT *SignPPTRequest
	stub := &LApiStub{AuthStub: AuthStub{
		Sign: func(sid string, req *SignRequest) ([]byte, error) {
			gotSign = req
			return []byte("sig"), nil
		},
		SignPPT: func(sid string, req *SignPPTRequest) (string, error) {
			gotPPT = req
			return "ppt", nil
		},
	}}

	req := &SignRequest{Content: []byte("doc"), Purpose: "test", AidDecl: "app-1"}
	sig, err := stub.Sign("s1", req)
	if err != nil || string(sig) != "sig" {
		t.Fatalf("Sign dispatch: sig=%q err=%v", sig, err)
	}
	if gotSign != req {
		t.Error("SignRequest 未原样透传")
	}

	preq := &SignPPTRequest{Info: map[string]string{"name": "n"}, Period: 30, Purpose: "p", AidDecl: "app-1"}
	ppt, err := stub.SignPPT("s1", preq)
	if err != nil || ppt != "ppt" {
		t.Fatalf("SignPPT dispatch: ppt=%q err=%v", ppt, err)
	}
	if gotPPT != preq {
		t.Error("SignPPTRequest 未原样透传")
	}
}
