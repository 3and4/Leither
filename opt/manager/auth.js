(function(global) {
console.log("auth.js");  
var ayApi = ["GetVarByContext", "Act", "Login", "Getvar","Getnodeip", "SwarmLocal","DhtGetAllKeys", "DhtGet", "DhtGets", "SignPPT", "SwarmAddrs"]

function getcurips(){
    //缺省的地址，用于本地调试程序
    var ips = "127.0.0.1:4800"
    //获取节点链接
    if (window.getParam != null){
        p=window.getParam()
        console.log("p=", p)
        ips = p["ips"][p.CurNode]
        // hosturl = "ws://" + p["ips"][p.CurNode] + "/ws/"
        // baseurl = "http://" + p["ips"][p.CurNode] + "/"
    } else if (window.location.host != ""){
        ips = window.location.host
        // hosturl = "ws://" + window.location.host + "/ws/"
        // baseurl = "http://" +  window.location.host + "/"
    }
    { //for test
        //ips = "192.168.3.7:8000"
    }
    global.vue.ips = ips
    return ips
}

global.getLocalApiHandler = ()=>{
    var ips = getcurips()
    var hosturl = "ws://" + ips +"/ws/"
    var baseurl = "http://" + ips + "/"

    //global.baseUrl = baseurl

    var apiHandler = {}
    //生成操作句柄
    apiHandler.client = hprose.Client.create(hosturl, ayApi);

    // console.log("client=", apiHandler.client);
    // console.log("baseurl=", baseurl);

    //以上部分可以提取公用代码
    return apiHandler.client.GetVarByContext("", "ppt").then(ppt=>{
        showPPT(ppt)

        return apiHandler.client.Login(ppt)
    }).then(reply=>{
        vue.isLogined = true
        console.log("sid=", reply.sid)
        console.log("uid=", reply.uid)
        apiHandler.baseUrl = baseurl

        apiHandler.sid = reply.sid
        apiHandler.leitherid = reply.uid
       //vue.mynodedata.nodeid = reply.uid
        vue.systemdata.localnodeid = reply.uid
        // vue.baseUrl = baseurl
        // vue.mynodedata.leitherid= reply.uid
        // vue.sid = reply.sid
        //global.sid = reply.sid
        //console.log("apiHandler=", apiHandler)
        return apiHandler
        //查询应用            
        //showapps(sid)
    }/*, e=>{
        //如果不能登录，需要进入用户注册登录页面
        console.log("getppt fail e=", e)
       //这里如果截取错误，流程上会出问题
    }*/)

}
global.doAll = (api)=>{
    console.log("doAll")
    console.log("api=", api)
    global.api = api
    
    var initReply = (api)=>{
        vue.sid = api.sid
        vue.baseUrl = api.baseUrl
        vue.mynodedata.nodeid= api.leitherid
    }

    initReply(api)


    function getappvers(api, iapp) {
        console.log("getappvers", iapp)
        var app = vue.appsdata.apps[iapp]
        console.log("app=", app, "api.sid=", api.sid)
    
        return api.client.Getvar(api.sid, "appversions", app.iD)
            .then((appvers)=>{
            console.log("appvers=", appvers)
            vue.appsdata.apps[iapp].appvers = appvers
            vue.appsdata.apps = vue.appsdata.apps
            return appvers
        })
    }

    function showapps(api){
        console.log("showapps")
        return api.client.Getvar(api.sid, "appinfos").then(apps=>{
            console.log("apps=", apps)
            vue.appsdata.apps = apps
            //请求版本信息
            var ret = []
            for (iapp in apps){
                ret[iapp] =getappvers(api, iapp)
            }
            return hprose.Future.all(ret)
        })
    }

    showapps(api).then(()=>{
        console.log("all=")
        var apps = vue.appsdata.apps
        vue.appsdata.apps = []
        vue.appsdata.apps = apps
    })

    api.client.GetVarByContext("", "ip").then(
        function(ip){
            console.log("ip=", ip)
        }
    )    



    api.client.GetVarByContext("", "mac").then(
        function(mac){
            console.log("mac=", mac)
        }
    )
    
    
    
    api.client.Getvar("", "nodepk", "", "raw").then(
        function(hostpk){
            console.log("hostpk=", hostpk)
            vue.mynodedata.hostpk = hostpk
        }
    )
    
    api.client.Getvar("", "nodeaddrs").then(
        function(ips){
            console.log("ips=", ips, ips.length)
            //vue.mynodedata.hostpk = hostpk
            if (ips.length > 0){
                vue.mynodedata.ips= ips[0]
                if (ips.length>1){
                    vue.mynodedata.gateway= ips[1]
                    vue.swarmdata.gateway = ips[1]
                } 
            }
        }, err=>{
            console.log(err)
        }
    )
    
    api.client.Getvar("", "peerid").then(
        function(peerid){
            console.log("peerid=", peerid)
            vue.mynodedata.peerid = peerid
        }
    )
    api.client.Getvar("", "nodeid").then(
        function(nodeid){
            console.log("nodeid=", nodeid)
            vue.mynodedata.nodeid = nodeid
        }
    )
    api.client.SwarmLocal("").then(
        function(sl){
            console.log("SwarmLocal =", sl)
            vue.mynodedata.swarmlocal = sl
        }
    )
    
    // api.client.SwarmPeers("").then(
    //     function(peers){
    //         console.log("SwarmPeers =", peers)
    //         vue.swarmdata.peers = peers
    //     }
    // )
    
    api.client.DhtGetAllKeys("").then((keys)=>{
        console.log("DhtGetAllKeys keys=", keys)
    })

    api.client.DhtGets("", "IPFS").then((ret)=>{
        console.log("DhtGets IPFS ret=", ret)
    }) 
    
    api.client.DhtGet("", "LAN").then((ret)=>{
        console.log("DhtGet LAN keys=", ret)
    })
    
    api.client.DhtGets("").then((ret)=>{
        console.log("vue.swarmdata.dhts =", ret)
        vue.swarmdata.dhts = ret    
    })
    
    api.client.DhtGets("", "WAN").then((ret)=>{
        console.log("DhtGets WAN ret=", ret)
    }) 

    api.client.SwarmAddrs(api.sid).then((addrs)=>{
        console.log("SwarmAddrs keys=", addrs)
        vue.swarmdata.addrs = addrs
        //for (k in addrs){
            //var addr = addrs[k] 
            //console.log("k=", k)
            //console.log("addr=", addr)
        //}
    })
    api.client.Getvar(api.sid, "systemvars").then(
        function(cfg){
            console.log("cfg=", cfg)
            vue.systemdata.oldVars = cfg
            vue.systemdata.systemvars = cfg
 //           vue.mynodedata.nodeid = nodeid
        }
    )   
}

var apiHandler = getLocalApiHandler()
global.apiHandler = apiHandler

apiHandler.then(doAll, e=>{
    console.log("e=", e)
})


global.getApiHandler = (addr, ppt)=>{
    console.log("getApiHandler", addr)
    //
    hosturl = "ws://" + addr + "/ws/"
    var apiHandler = {}
    //生成操作句柄
    apiHandler.client = hprose.Client.create(hosturl, ayApi);

    //获取一个通行证

    //以上部分可以提取公用代码
    return apiHandler.client.GetVarByContext("", "ppt", {
        period:60*24*7,
        subtype:"mac",
        ppt:ppt
    }).then(
        apiHandler.client.Login
    ).then(reply=>{
        console.log("reply=", reply)
        console.log("sid=", reply.sid)
        console.log("uid=", reply.uid)
        apiHandler.baseUrl = "http//" + addr + "/"
        apiHandler.sid = reply.sid
        apiHandler.leitherid = reply.uid

        return apiHandler
        //查询应用            
        //showapps(sid)
    }, err=>{
        console.log(err)
    })
}

global.onswitch = pid =>{
    console.log("global.onswitch", pid)
    var addrs =vue.swarmdata.addrs[pid]

    console.log("global.onswitch", addrs)
    var laddr = maddrs2laddr(addrs)
    loginPeer(laddr)
    // for (id in vue.swarmdata.peers){
    //     var addr = vue.swarmdata.peers[id]
    //     if (addr.lastIndexOf(pid) >=0) {
    //         console.log("addr=", addr)
    //         var paddr=maddr2ips(addr)
    //         console.log(paddr)
    //         loginPeer(paddr)
    //     }

    // }
}

function maddrs2laddr(addrs){
 //   var laddripv4, laddripv6
    for (i in addrs){
        var addr = addrs[i]
        console.log(addr)

        var ipv6, ipv4
        var port = 80
        var ay = addr.split("/")
        console.log(ay)
        for (var i=1; i<ay.length; i++){
            console.log(ay[i], ay[i+1])
            switch(ay[i++]){
            case "ip6":
                ipv6 = "[" + ay[i] + "]"
                break
            case "ip4":
                ipv4 = ay[i]
                break
            case "tcp":
                port = ay[i]
                break
            default:
            }
        }
        console.log("ipv6", ipv6, "port", port, "ipv4", ipv4)
        
        //优先选ipv6
        if (ipv6) {
            return ipv6 + ":" + port
        }
        if (ipv4) {
            if (ipv4.indexOf("192")==0){
                return ipv4 + ":" + port
            }
        }        
    }

    return ""
}

//地址转换
function maddr2ips(addr){
    var ipv6, ipv4
    var port = 80
    var ay = addr.split("/")
    console.log(ay)
    for (var i=1; i<ay.length; i++){
        console.log(ay[i], ay[i+1])
        switch(ay[i++]){
        case "ip6":
            ipv6 = "[" + ay[i] + "]"
            break
        case "ip4":
            ipv4 = ay[i]
            break
        case "tcp":
            port = ay[i]
            break
        default:
        }
    }
    console.log("ipv6", ipv6, "port", port, "ipv4", ipv4)
    if (ipv6) {
        return ipv6 + ":" + port
    }
    if (ipv4) {
        return ipv4 + ":" + port
    }
    return null
}

function loginPeer(piurl){
    apiHandler.then((api)=>{
        console.log("api.Client.SignPPT")
        //生成ppt
        return api.client.SignPPT(api.sid, {
            CertFor:"Self",
            Userid:api.Userid
        }, 1)
    }).then(ppt=>{
        console.log("ppt=", ppt)
        //知道对方的url
        return getApiHandler(piurl, ppt)
    }).then(doAll, e=>{
        console.log("e=", e)
    })
}

//读出通行证中的userid
function showPPT(strppt){
    console.log("ppt=", strppt)
    var ppt = JSON.parse(strppt)
    // for (k in ppt){
    //     console.log("k=", k)        
    // }
    console.log("data=", ppt.Data)

}

// global.ontest = () =>{
//     console.log("ontest")
//     loginPeer('192.168.3.7:8000')   
// }

})(this || [eval][0]('this'));

            