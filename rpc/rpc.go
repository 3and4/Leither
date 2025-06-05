package rpc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	. "github.com/hprose/hprose-golang/v3/rpc/websocket"
)

func init() {
	RegisterHandler()
	RegisterTransport()
}

func InitLApiStubByUrl(url string) *LApiStub {
	var stub LApiStub

	InitStubByUrl(&stub, url)
	return &stub
}

// 提取出来，可以指定api的Stub使用
func InitStubByUrl(stub any, url string) {
	hasWS := strings.HasPrefix(url, "ws")
	if !strings.HasPrefix(url, "http") && !hasWS {
		url = "http://" + url
	}

	//这里可以确定是http开的头
	t := strings.Replace(url, "http", "ws", 1) //现在支持
	if t[len(t)-1] != '/' {
		t = t + "/"
	}
	url = t + "ws3/" //待优化

	Client := core.NewClient(url)

	Client.UseService(stub)
	Client.Timeout = time.Hour * 24
}

// GetLocalPassport 从本地节点获取通行证
func GetLocalPassport(port int, validPeriod int) (string, error) {
	url := fmt.Sprintf("http://127.0.0.1:%d/getvar?name=ppt&arg0=%d&nojson", port, validPeriod)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("请求通行证失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("获取通行证失败，状态码: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	ppt := string(body)
	if ppt == "" {
		return "", fmt.Errorf("获取到的通行证为空")
	}

	return ppt, nil
}
