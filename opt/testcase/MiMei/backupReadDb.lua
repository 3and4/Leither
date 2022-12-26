--读取sid
local sid =request['sid']
print("sid=", sid)

local mm = require('mimei');

--创建弥媒
local mid, err = mm.MMCreate(sid, '', '', 'testBackup', 2, 0x07276705);
if (err ~= nil) then
	print('MMCreate err=',  err);
	return err
end
print("MMCreate mid=", mid)

local mmsid2, err = mm.MMOpen(sid, mid, 'last');


local ret, err = mm.HMGet(mmsid2, "hkey", "field99", "field0")
if (err ~= nil) then
	print('HMGet err=',  err);
	return err
end

print("HMGet ret")
for k,v in ipairs(ret) do
        print(k,v)
end

--取所有
local fvs, err = mm.HGetAll(mmsid2, "hkey")
print("HGetAll fvs")
for k,v in ipairs(fvs) do
        print(k, v:getField(), v:getValue())
end

return ver
