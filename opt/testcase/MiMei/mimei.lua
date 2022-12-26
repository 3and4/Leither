--读取sid
sid =request['sid']
print("mimei sid", sid)

local mm = require('mimei');

print('create mimei');

--[[生成一个弥媒,
MMCreate(
	sid 	, //会话id
	appid 	, //应用id,目前空表示系统应用
	ext		, //扩展名，表示弥媒的类别
	mark	, //弥媒的标识，
	dataType, //存储类型，1文件，2数据库
	right	, //权限
) 
--]]
mid, err = mm.MMCreate(sid, '', '', 'testMiMei', 2, 0x07276705);
if (err ~= nil) then
	print('MMCreate err=',  err);
	return err
end
print("mid", mid)


--[[打开一个弥媒,
MMOpen(
	sid, 	//会话id
	mid, 	//弥媒id
	ver, 	//版本
	opt...	//可选项
)
返回值：会话id和错误
--]]

mmsid, err = mm.MMOpen(sid, mid, 'cur');
if (err ~= nil) then
	print('MMOpen err',  err);
	return err
end
print('MMOpen mmsid=', mmsid);

--[[
Set(mmsid, key, value) error
Get(mmsid, key) (value, error)
--]]

err = mm.Set(mmsid, 'key0', 'value0');
if (err ~= nil) then
	print('Set err',  err);
	return err
end

value, err = mm.Get(mmsid, 'key0');
if (err ~= nil) then
	print('Set err',  err);
	return err
end
print('Get key0 value=', value);



----[[
--set get table
local f = {}
f["name"] = "lua1";
f["key"] = "lua2";
f["3"] = 3;			--检查数据类型
f["4"] = "lua4";
--f[5] = "lua5"; 	--检查类型混用，这个是不可以的，有的话会出错
f["6"] = "lua6";

print("f[4]", f[4])
print("f[5]", f[5])
print("f[6]", f["6"])
print("maxn", table.maxn(f))
for k,v in ipairs(f) do
        print(k,v)
end
err = mm.Set(mmsid, 'key1', f);
if (err ~= nil) then
	print('Set err',  err);
	return err
end
print("set ok")

value, err = mm.Get(mmsid, 'key1');
if (err ~= nil) then
	print('Set err',  err);
	return err
end
--print('Get key1 value=', value["6"]);
print('Get key1 name=', value.name);
print('Get key1 key=', value["key"]);
--这个是测试name有没有变成大写
print('Get key1 Name=', value.Name);
print('Get key1 Key=', value["Key"]);
----]]


--hset hget
n, err = mm.HSet(mmsid, 'key', 'field0', 'value0');
print(n, err)

value, err = mm.HGet(mmsid, 'key', 'field0');
print(value, err)


--hmset hmget
local fv1 = fvpair.new('field1', 'value1')
print("fv1", fv1:getField(), fv1:getValue());

local fv2 = fvpair.new('field2', 'value2')
print("fv2", fv2:getField(), fv2:getValue());

err = mm.HMSet(mmsid, 'key', fv1, fv2);
print(err)

values, err = mm.HMGet(mmsid, 'key', 'field0', 'field1', 'field2');
print(values, err)
for k,v in ipairs(values) do
        print(k,v)
end


return mid;
