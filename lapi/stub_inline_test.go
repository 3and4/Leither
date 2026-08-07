package lapi

import (
	"bytes"
	"testing"
)

// TestInlineAppCssUnwiredReturnsOriginal 字段未接线（nil）→ 返回原样，不 panic。
func TestInlineAppCssUnwiredReturnsOriginal(t *testing.T) {
	var s LApiStub // InlineAppCss 未接线
	in := []byte("<html><head></head></html>")
	out, err := s.InlineAppCss("", "", "", "", "", in)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !bytes.Equal(out, in) {
		t.Errorf("unwired InlineAppCss: got %q, want input unchanged", out)
	}
}

// TestInlineAppCssDelegatesToField 委托到 MiMeiStub.InlineAppCss 字段并透传参数。
func TestInlineAppCssDelegatesToField(t *testing.T) {
	var gotSid, gotAppid, gotMid, gotVer, gotRoot string
	s := LApiStub{
		MiMeiStub: MiMeiStub{
			InlineAppCss: func(sid, appid, mid, ver, root string, html []byte) ([]byte, error) {
				gotSid, gotAppid, gotMid, gotVer, gotRoot = sid, appid, mid, ver, root
				return append([]byte("X:"), html...), nil
			},
		},
	}
	out, err := s.InlineAppCss("s1", "a1", "m1", "v1", "r1", []byte("h"))
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "X:h" {
		t.Errorf("got %q, want %q", out, "X:h")
	}
	if gotSid != "s1" || gotAppid != "a1" || gotMid != "m1" || gotVer != "v1" || gotRoot != "r1" {
		t.Errorf("args not passed through: %q %q %q %q %q", gotSid, gotAppid, gotMid, gotVer, gotRoot)
	}
}
