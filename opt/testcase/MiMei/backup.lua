--读取sid
local sid =request['sid']
print("sid=", sid)

local mm = require('mimei');

--创建弥媒
local mid, err = mm.MMCreate(sid, '', '', 'test Backup', 2, 0x07276705);
if (err ~= nil) then
	print('MMCreate err=',  err);
	return err
end
print("MMCreate mid=", mid)

--打开弥媒
local mmsid, err = mm.MMOpen(sid, mid, 'cur');
if (err ~= nil) then
	print('MMOpen err=',  err);
	return err
end
print('MMOpen mmsid=', mmsid);

--制造一些变动差异
err = mm.Set(mmsid, 'sid', sid);
if (err ~= nil) then
	print('Set err',  err);
	return err
end

--测试的基准数量
local TestCount = 100 

--生成一个随机值
local idxDel = os.time() % TestCount 
local count
--写基准数量个值
for i=0, TestCount, 1 do
	local key = string.format("key%d", i)
	if (i == idxDel) then
		count, err = mm.Del(mmsid, key)
		print("del", key)
	else 
		local value =string.format("value%d", i)
		print("set", key, value)
		err = mm.Set(mmsid, key, value);
	end
	if (err ~= nil) then
		print('err=%v',  err);
		return err
	end
end



--写一下hashtable
for i=0, TestCount, 1 do
	local field = string.format("field%d", i)
	local value =string.format("value%d", i)
	print("hset", "hkey", field, value)
	n, err = mm.HSet(mmsid, "hkey", field, value);
	if (err ~= nil) then
		print('HSet err=%v',  err);
		return err
	end
end

--制造一些变动差异
err = mm.Set(mmsid, 'sidend', sid);
if (err ~= nil) then
	print('Set err',  err);
	return err
end


--显示值
for i=0, TestCount, 1 do
	local key = string.format("key%d", i)
	value, err = mm.Get(mmsid, key);
	if (err ~= nil) then
		print('Get err=%v',  err);
		return err
	end

	print("Get", key, value);
end

for i=0, TestCount, 1 do
	local field = string.format("field%d", i)
	value, err = mm.HGet(mmsid, "hkey", field);
	if (err ~= nil) then
		print('Set err=%v',  err);
		return err
	end

	print("HGet", "hkey", field, value);
end

local sp = scorepair.new(1, 'm1')
print("sp", sp:getScore(), sp:getMember());

local n, err =mm.ZAdd(mmsid, 'zkey', sp)
if (err ~= nil) then
	print('ZAdd err',  err);
	return err
end
print("ZAdd n=", n)
--[[
local ret, err = mm.Zrem(mmsid, "zkey", "m1")
if (err ~= nil) then
        print("Zrem err", err);
        return err
end
print("Zrem ret=", ret)
--]]

--备份
--	MMBackup(sid, mid, memo string)error
local ver
ver, err = mm.MMBackup(sid, mid, '')
if (err ~= nil) then
	print('MMBackup err=',  err);
	return err
end


--[[异步接收备份消息
local msg = require('msg');
while true
do
	repeat
		strmsg, err = msg.PullMsg(sid, 3)
		if (err ~= nil) then
			print('PullMsg err=%v',  err);
			return err
		end
	until(strmsg ~= nil)
	info = strmsg:getMsg()
	reg = "(%a+)=(.*)"
	_, _, key, value = string.find(info, reg)
	
	if (key=="result") then
		print("ver", value)
		ver = value
		--return value
		break
	end
	if (key=="error") then
		print("err", value)
		return value
	end	
	
	print(key, value)
end
--]]
print("ver", ver)
--[[

--打开备份的版本
--打开弥媒

--local mmsid2, err = mm.MMOpen(sid, mid, ver);
if (err ~= nil) then
	print('MMOpen err=',  err);
	return err
end
print('MMOpen mmsid2=', mmsid2, ver);
for i=0, 100, 1 do
	local field = string.format("field%d", i)
	value, err = mm.HGet(mmsid2, "hkey", field);
	if (err ~= nil) then
		print('Set err=%v',  err);
		return err
	end

	print("HGet", "hkey", field, value);
end
--]]
local mmsid2, err = mm.MMOpen(sid, mid, 'last');
--[[
local v, err = mm.HGet(mmsid2, "hkey", "field99")
if (err ~= nil) then
	print('HGet err=',  err);
	return err
end
print("HGet v", v)
--]]

local fvs, err = mm.HMGet(mmsid2, "hkey", "field99")
if (err ~= nil) then
	print('HMGet err=',  err);
	return err
end

print("read fvs")
for k,v in ipairs(fvs) do
        print(k,v)
end

local value, err = mm.Get(mmsid, "sid")
if (err ~= nil) then
	print('Get sid err=',  err);
	return err
end
print("read sid", value)

local value, err = mm.Get(mmsid, "sidend")
if (err ~= nil) then
	print('Get sidend err=',  err);
	return err
end
print("read sidend", value)

--[[
fvs, err = mm.HMGet(mmsid, "hkey", "field0", "field1")
if (err ~= nil) then
	print('HMGet err=',  err);
	return err
end

print("read fvs cur")
for k,v in ipairs(fvs) do
        print(k,v)
end
--]]
return ver
