package lapi

// IBEWeb 定义后端 web 内容服务 API（mmroot 文件读取 + 网页预处理）。
type IBEWeb interface {
	// BEReadFile 读取应用 mmroot 下文件并返回最终内容。
	// name: mmroot 下相对路径（如 index.html、css/theme.css）。
	// ops: 预处理操作（可多个，节点按内部固定顺序应用）；空=原始读；
	//      含 "web" = 网页可服务化流水线（css-inline + link-rewrite，仅 html 生效）；
	//      未知操作名返回错误。
	BEReadFile(name string, ops ...string) ([]byte, error)
}

// BEWebStub 是 IBEWeb 的宿主接线结构（函数字段，容器进程内直调）。
type BEWebStub struct {
	BEReadFile func(name string, ops ...string) ([]byte, error)
}
