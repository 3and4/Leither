(function(global) {
    console.log("manager.js")  
    
    var c = localStorage.getItem("selected")
    if (c == null){
        c = "system"
    }


    //定义一个组件
    Vue.component('headnodestate', {
        props: ['localid', "nodeid"],
        template: `<div class="main_header">
        <p class="main_header_title">当前状态</p>
        <div class="main_header_hint2">
        本地节点&nbsp;<span class="theme_color">{{localid}}</span>&nbsp;&nbsp;&nbsp;
        当前浏览&nbsp;&nbsp;<span class="theme_color">{{nodeid}}</span>
        
        </div>
    </div>`
      })
    global.vue = new Vue({
    el: '#leitherbody',
    data: {
        ips:"",
        isLogined:true,
        selected: c,
        sid:"",
        mynodedata:{
            nodeid:" ",
            hostpk:" ",
            ips:" ",
            gateway:"  ",
            peerid:" ",
            swarmlocal:""
        },
        swarmdata:{
            dhts:[],
            //peers:[],
            addrs:[],
            gateway:"  "
        },
        appsdata:{
            apps:[]
        },
        systemdata:{
            localnodeid:"",
            oldVars:"",
            systemvars:"",
            bChanged: false        
        }
        
    },
    // 在 `methods` 对象中定义方法
    methods: {
        ontest:  event =>{
            console.log("vue.ontest", event)
            global.ontest()
        },
        
        onswitch:pid =>{
            console.log("vue.onswitch", pid)
            global.onswitch(pid)
        },

        select: name=> {
            global.vue.selected = name
            localStorage.setItem("selected", name)
        },

        getAppUrl: (appid, ver)=>{
            console.log("getAppUrl")
            return global.vue.baseUrl + "entry?aid="+ appid +"&ver="+ver
        },
        printvars: e=>{
            global.vue.systemdata.bChanged = (e !== global.vue.systemdata.oldVars)
            console.log(e)
            console.log("systemvars=", global.vue.systemdata.systemvars)
            global.vue.systemdata.systemvars =  e
        },
        resetSystemdata:()=>{
            global.vue.systemdata.systemvars = global.vue.systemdata.oldVars
            global.vue.systemdata.bChanged = false
        },
        saveSystemdata:()=>{
            console.log("saveSystemdata ", global.vue.systemdata.systemvars)
            global.apiHandler.then(api=>{
                console.log("api=", api)
                api.client.Act(api.sid, "savesystemvars", global.vue.systemdata.systemvars).then(()=>{
                    console.log("save ok")
                    global.vue.systemdata.oldVars = global.vue.systemdata.systemvars
                },e=>{
                    console.log(e)
                })                
            })
           // global.vue.systemdata.systemvars = global.vue.systemdata.oldVars
        }
    }
  })

  //刷新

//  global.vue.greet()
})(this || [eval][0]('this'));