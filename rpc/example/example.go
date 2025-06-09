// example.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/rpc"
)

const testFile = "/testfiles.txt"

func main() {
	//以下是不需要认证身份的示例
	//节点的地址端口是127.0.0.1:4800
	stub := rpc.InitLApiStubByUrl("127.0.0.1:4800")
	ver, err := stub.GetVar("", "ver")
	if err != nil {
		panic(err)
	}

	fmt.Println("ver:", ver)

	//获取一个通行证
	//这是通过127.0.0.1:4800,向本地节点申请的
	//有效期是24小时
	strPPT, err := rpc.GetLocalPassport(4800, 24)
	if err != nil {
		panic(err)
	}
	fmt.Println("strPPT:", strPPT)
	fmt.Println("LoginWithPPT", stub.AuthStub.LoginWithPPT)

	reply, err := stub.LoginWithPPT(strPPT)
	fmt.Println("LoginWithPPT", reply, err)
	if err != nil {
		panic(err)
	}
	sid := reply.Sid
	defer stub.Logout(sid, "")

	fmt.Println("reply", sid, reply.Uid)

	//先删除files中的测试文件
	err = stub.FilesRm(sid, testFile, false, true)
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	//生成一个临时文件
	fsid, err := stub.MFOpenTempFile(sid)
	if err != nil {
		panic(err)
	}

	fmt.Println("fsid", fsid)

	//退出时关闭这个文件
	defer stub.MMClose(fsid)

	//写入数据
	count, err := stub.MFSetData(fsid, []byte("hello world"), 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("count", count)

	//把临时文件复制到files的根目录，路径为"/testfiles.txt"
	ipfsid, err := stub.MFTemp2Files(fsid, testFile)
	if err != nil {
		//这里可能会有重名
		//panic(err)
		fmt.Println("err:", err.Error())
	}
	fmt.Println("ipfsid", ipfsid)

	//这里测试读取
	// files/ps 		/ps	=	files/ps
	fsid2, err := stub.MMOpenUrl(sid, testFile)
	if err != nil {
		panic(err)
	}
	defer stub.MMClose(fsid2)

	data, err := stub.MFGetData(fsid2, 0, 100)
	if err != nil {
		panic(err)
	}
	fmt.Println("data:", string(data))

	//再删除files中的测试文件
	err = stub.FilesRm(sid, testFile, false, true)
	if err != nil {
		fmt.Println("err:", err.Error())
	}

}
