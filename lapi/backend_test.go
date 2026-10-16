package lapi

import "testing"

// BackEndStub 的 Session 方法应委托给内嵌 SessionStub 的真实实现。
// 曾因方法遮蔽内嵌字段导致无限自递归（fatal stack overflow）。
func TestBackEndStub_CreateSession_DelegatesToSessionStub(t *testing.T) {
	stub := &BackEndStub{SessionStub: &SessionStub{}}
	const want = "sid-123"
	stub.SessionStub.CreateSession = func() string { return want }

	if got := stub.CreateSession(); got != want {
		t.Fatalf("CreateSession() = %q, want %q", got, want)
	}
}

func TestBackEndStub_SessionSet_DelegatesToSessionStub(t *testing.T) {
	stub := &BackEndStub{SessionStub: &SessionStub{}}
	const sid, key = "sid-1", "k"
	val := any("v")
	stub.SessionStub.SessionSet = func(gotSid, gotKey string, gotVal any) error {
		if gotSid != sid || gotKey != key || gotVal != val {
			t.Fatalf("SessionSet args = (%q, %q, %v), want (%q, %q, %v)", gotSid, gotKey, gotVal, sid, key, val)
		}
		return nil
	}

	if err := stub.SessionSet(sid, key, val); err != nil {
		t.Fatalf("SessionSet() error: %v", err)
	}
}

func TestBackEndStub_SessionGet_DelegatesToSessionStub(t *testing.T) {
	stub := &BackEndStub{SessionStub: &SessionStub{}}
	const sid, key = "sid-1", "k"
	want := any("stored-value")
	stub.SessionStub.SessionGet = func(gotSid, gotKey string) (any, error) {
		if gotSid != sid || gotKey != key {
			t.Fatalf("SessionGet args = (%q, %q), want (%q, %q)", gotSid, gotKey, sid, key)
		}
		return want, nil
	}

	got, err := stub.SessionGet(sid, key)
	if err != nil {
		t.Fatalf("SessionGet() error: %v", err)
	}
	if got != want {
		t.Fatalf("SessionGet() = %v, want %v", got, want)
	}
}

func TestBackEndStub_SessionDelete_DelegatesToSessionStub(t *testing.T) {
	stub := &BackEndStub{SessionStub: &SessionStub{}}
	const sid, key = "sid-1", "k"
	stub.SessionStub.SessionDelete = func(gotSid, gotKey string) error {
		if gotSid != sid || gotKey != key {
			t.Fatalf("SessionDelete args = (%q, %q), want (%q, %q)", gotSid, gotKey, sid, key)
		}
		return nil
	}

	if err := stub.SessionDelete(sid, key); err != nil {
		t.Fatalf("SessionDelete() error: %v", err)
	}
}

func TestBackEndStub_ReleaseSession_DelegatesToSessionStub(t *testing.T) {
	stub := &BackEndStub{SessionStub: &SessionStub{}}
	const sid = "sid-1"
	stub.SessionStub.ReleaseSession = func(gotSid string) error {
		if gotSid != sid {
			t.Fatalf("ReleaseSession(%q), want %q", gotSid, sid)
		}
		return nil
	}

	if err := stub.ReleaseSession(sid); err != nil {
		t.Fatalf("ReleaseSession() error: %v", err)
	}
}

// 未接线时不应 panic，且应显式报"未接线"错误（D1：不允许静默零值成功）。
func TestBackEndStub_BEOpenAppDataNode_UnwiredReturnsError(t *testing.T) {
	stub := &BackEndStub{} // 不含 BEAppDataStub

	got, err := stub.BEOpenAppDataNode("cur", "mark")
	if err == nil {
		t.Fatal("BEOpenAppDataNode() unwired: want explicit error, got nil")
	}
	if got != "" {
		t.Fatalf("BEOpenAppDataNode() = %q, want empty", got)
	}
}

// 未接线的后端方法应统一显式报错（D1），而非静默返回零值。
func TestBackEndStub_UnwiredMethodsReturnError(t *testing.T) {
	stub := &BackEndStub{}

	if _, err := stub.BEOpenAppDataApp("cur", "mark"); err == nil {
		t.Error("BEOpenAppDataApp() unwired: want error")
	}
	if err := stub.BEMMSync("dhts", "mid", nil); err == nil {
		t.Error("BEMMSync() unwired: want error")
	}
	if _, err := stub.BELoginAsAuthor(); err == nil {
		t.Error("BELoginAsAuthor() unwired: want error")
	}
	if _, err := stub.BELoginAsApp(); err == nil {
		t.Error("BELoginAsApp() unwired: want error")
	}
	if _, err := stub.BESignPPT(nil, 1); err == nil {
		t.Error("BESignPPT() unwired: want error")
	}
	if _, err := stub.BESign(nil); err == nil {
		t.Error("BESign() unwired: want error")
	}
	if err := stub.SessionSet("s", "k", nil); err == nil {
		t.Error("SessionSet() unwired: want error")
	}
	if _, err := stub.SessionGet("s", "k"); err == nil {
		t.Error("SessionGet() unwired: want error")
	}
	if err := stub.SessionDelete("s", "k"); err == nil {
		t.Error("SessionDelete() unwired: want error")
	}
	if err := stub.ReleaseSession("s"); err == nil {
		t.Error("ReleaseSession() unwired: want error")
	}
	// CreateSession 无 error 返回值：未接线返回空 sid（调用方判 sid==""）
	if sid := stub.CreateSession(); sid != "" {
		t.Errorf("CreateSession() unwired = %q, want empty", sid)
	}
}

func TestLApiStub_MMOpen_NoPanicWhenUnset(t *testing.T) {
	stub := &LApiStub{} // MiMeiStub.MMOpen 未设置

	got, err := stub.MMOpen("sid", "mid", "cur")
	if err != nil {
		t.Fatalf("MMOpen() error: %v", err)
	}
	if got != "" {
		t.Fatalf("MMOpen() = %q, want empty", got)
	}
}

func TestLApiStub_FilesRm_NoPanicWhenUnset(t *testing.T) {
	stub := &LApiStub{} // NetStub.FilesRm 未设置

	if err := stub.FilesRm("sid", "/ps", false, false); err != nil {
		t.Fatalf("FilesRm() error: %v", err)
	}
}
