(function(global) {
    global.openDB = (name, version) => {
        return new Promise((resolve, reject)=>{
            var version = version || 1;
            var request = indexedDB.open(name, version)
            request.onerror = function(e) {
                console.log("openDB error!");
                reject(e)
            };
            request.onsuccess = function(e) {
                resolve(e.target.result);
            };
            request.onupgradeneeded = function(e) {
                console.log("indexDB version change to " + version);
                var db = e.target.result;
                if(!db.objectStoreNames.contains('students')) {
                    db.createObjectStore('students',{keyPath:'id'});
                }
            }
        })
    }

    global.updateDataByKey = (db, storeName, data) =>{
        return db.then(db=>{
            var transaction = db.transaction(storeName, 'readwrite');
            var store = transaction.objectStore(storeName);
            var request=store.get(data.id);
            request.onsuccess = e => {
                console.log("store.put")
                //var student=e.target.result;
                store.put(data);
            };
        })
    }

    global.getDataByKey = (db, storeName, id) =>{
        return new Promise((resolve, reject)=>{ 
            db.then(db=>{
                var transaction = db.transaction(storeName);
                var store = transaction.objectStore(storeName);
                var request=store.get(id);
                request.onerror = function(e) {
                    console.log("getDataByKey error!");
                    reject(e)
                };
                request.onsuccess= e => {
                    console.log("store.get", e.target.result)                    
                    resolve(e.target.result);
                };
            })
        })
    }

    global.db = openDB("test", "1");
    console.log("db=", db)

    var student =  { id: "102", name: "Bill", age: 35, email: "bill@company.com" }
    updateDataByKey(db, "students", student).then(()=>{
        return getDataByKey(db, "students", "102")
    }).then(d=>{
        console.log("getdata ret=", d)
    })

})(this || [eval][0]('this'));
