sid =request['sid']
print("mimei sid", sid)
local mm = require('mimei');
print('create mimei');

mid, err = mm.MMCreate(sid, '', '', 'testMiMei', 2, 0x07276705);
if (err ~= nil) then
	print('MMCreate err=',  err);
	return err
end
print("mid", mid)

mmsid, err = mm.MMOpen(sid, mid, 'cur');
if (err ~= nil) then
	print('MMOpen err',  err);
	return err
end
print('MMOpen mmsid=', mmsid);

local f = {}
f["name"] = "lua1";
f["key"] = "lua2";
f["3"] = 3;			--检查数据类型
f["4"] = "lua4";
f["6"] = "lua6";

--print("f[4]", f[4])
--print("f[5]", f[5])
print("f[6]", f["6"])
--print("maxn", table.maxn(f))
for k,v in ipairs(f) do
        print(k,v)
end

--hset hget
n, err = mm.HSet(mmsid, 'key', 'field10', f);
print("Hset result=", n, err)

value, err = mm.HGet(mmsid, 'key', 'field10');
print("loop value")
for k,v in pairs(value) do
	print(k, v)
end
print(value, err)
print("f[6]", f["6"])
print("f[4]", f["4"])
print("f[name]", f["name"])
print("f[key]", f["key"])
print("f[3]", f["3"])
