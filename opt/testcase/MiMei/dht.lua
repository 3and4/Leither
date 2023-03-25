--读取sid
local sid =request['sid']
print("sid=", sid)

local mm = require('mimei');

--创建弥媒
local mid, err = mm.MMCreate(sid, '', '', 'testdht', 2, 0x07276705);
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

--备份
--	MMBackup(sid, mid, memo string)error
err = mm.MMBackup(sid, mid, '')
if (err ~= nil) then
	print('MMBackup err=%v',  err);
	return err
end


--异步接收备份消息
local ver
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
	
	if (key=="ver") then
		print("ver", value)
		ver = value
		--return value
		break
	end
	if (key=="err") then
		print("err", value)
		return value
	end	
	
	print(key, value)
end

print("ver", ver)

local mmsid, err = mm.MMOpen(sid, mid, 'last');
if (err ~= nil) then
	print('MMOpen err=',  err);
	return err
end


local value, err = mm.Get(mmsid, "sid")
if (err ~= nil) then
	print('HMGet err=',  err);
	return err
end
print("read sid", value)

--publish
local net = require('net');

err = net.MiMeiPublish(sid, "", mid)
if (err ~= nil) then
	print('MiMeiPublish err=',  err);
	return err
end
print("MiMeiPublish ok")

local sysvar = require('var')

--获取本地的弥媒信息
local mminfo, err  = sysvar.GetVar(sid, "mminfo", mid)
if (err ~= nil) then
	print('GetVar  mminfo err=',  err);
	return err
end

--[[
	Author  string    //作者
	AppType string    //应用类型
	Ext     string    //扩展信息
	Mark    string    //标识
	Create  time.Time //创建时间
	Right   uint64    //权限
--]]
print("Author ", mminfo['author'])
for k,v in pairs(mminfo) do
        print(k,v)
end

--获取版本信息
--获取本地的弥媒信息
print("readver")
local mmver, err  = sysvar.GetVar(sid, "mimeiversions", mid)
if (err ~= nil) then
	print('GetVar  mmver err=',  err);
	return err
end

--获取引用数
local refs, err = mm.MMGetRefs(sid, mid)
if (err ~= nil) then
	print('MMGetRefs err=',  err);
	return err
end

print("show versions")
for k,v in pairs(mmver['versions']) do
	local ver = v['version']
	print(ver, v['macRes'], v['macRef'])
--	for k1,v1 in pairs(v) do
--        print(k1, v1)
--	end
	ref = refs[ver]
	for k1,v1 in pairs(ref) do
        print(k1, v1)
	end
end

print("show specialVers")
for k,v in pairs(mmver['specialVers']) do
	print(v['verName'], v['version'])
--	for k1,v1 in pairs(v) do
--        print(k1, v1)
--	end
end



--获取网络信息
local sr, err = net.MiMeiShow(sid, "", mid)
if (err ~= nil) then
	print('MiMeiShow err=',  err);
	return err
end

print("from net")
local mmvers = sr:getVersions()
local vers = mmvers:getVersions()
local svers = mmvers:getSpecialVers()

--type MiMeiVersions struct {
--	Versions    []*MiMeiVersion   //所有的版本
--	SpecialVers []*SpecialVersion //特殊版本
--}

--print("vers", vers)
for k,v in pairs(vers) do
	--print(k, v)
	print(v:getVersion(), v:getMacRes(), v:getMacRef())
end


--print("svers",svers)
for k,v in pairs(svers) do
	print(v:getVerName(), v:getVersion())
end

return "success"
