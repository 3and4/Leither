--[[
local log = require('log');
print("testlog begin")
log.Debug("testlog")
print("testlog end")

str = "ver=2225"
reg = "(%a+)=(.*)"

_, _, key, value = string.find(str, reg)
print("key", key)
print("value",value)
]]--
sid =request['sid']
print("sid=", sid)
local mm = require('mimei');

-- ^h^{        ^r
mid='ilc_mDQ-vS9jRIRw2w70pyf8ASN'

-- ^i^s  ^`     ^r
mmsid, err = mm.MMOpen(sid, mid, 'cur');
if (err ~= nil) then
        print('MMOpen err=',  err);
        return err
end
print('MMOpen mmsid=', mmsid);

-- load image file
local inp = assert(io.open('1.jpg', 'rb'))
local t = inp:read(1)
--local t = "test"
--for i=1,1000 do t[i] = math.random(0,255) end

fsid, err = mm.MFOpenTempFile(sid);
count, err = mm.MFSetData(fsid, t, 0)
if (err ~= nil) then
        print('MMSetData err=',  err);
        return err
end
print('MMSetData count=', count);
