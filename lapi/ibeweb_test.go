package lapi

import (
	"testing"
)

// TestBEWebStubImplementsIBackEnd BackEndStub 满足 IBackEnd（含新 IBEWeb）。
func TestBEWebStubImplementsIBackEnd(t *testing.T) {
	var _ IBackEnd = (*BackEndStub)(nil)
}

// TestBEReadFileUnwiredReturnsError 字段未接线 → 返回错误，不 panic。
func TestBEReadFileUnwiredReturnsError(t *testing.T) {
	var s BackEndStub // BEWebStub 未接线
	out, err := s.BEReadFile("index.html", "web")
	if err == nil {
		t.Fatal("unwired BEReadFile should return error")
	}
	if out != nil {
		t.Errorf("got %q, want nil", out)
	}
}

// TestBEReadFileDelegatesToField 委托到 BEWebStub.BEReadFile 字段并透传参数（name + ops）。
func TestBEReadFileDelegatesToField(t *testing.T) {
	var gotName string
	var gotOps []string
	s := BackEndStub{
		BEWebStub: &BEWebStub{
			BEReadFile: func(name string, ops ...string) ([]byte, error) {
				gotName, gotOps = name, ops
				return []byte("web:" + name), nil
			},
		},
	}
	out, err := s.BEReadFile("index.html", "web", "minify")
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "web:index.html" {
		t.Errorf("got %q, want %q", out, "web:index.html")
	}
	if gotName != "index.html" || len(gotOps) != 2 || gotOps[0] != "web" || gotOps[1] != "minify" {
		t.Errorf("params: name=%q ops=%v", gotName, gotOps)
	}
}
