sid =request['sid']
print("Test ZSet. sid", sid)
local mm = require('mimei');

mid, err = mm.MMCreate(sid, '', '', 'testZset', 2, 0x07276705);
if (err ~= nil) then
	print('MMCreate err=',  err);
	return err
end
print("create mimei mid", mid)

mmsid, err = mm.MMOpen(sid, mid, 'cur');
if (err ~= nil) then
	print('MMOpen err',  err);
	return err
end
print('MMOpen mmsid=', mmsid);

local sp1 = scorepair.new(-1, 'm1')
print("sp1", sp1:getScore(), sp1:getMember());

local sp2 = scorepair.new(-2, 'm2')
print("sp2", sp2:getScore(), sp2:getMember());


local sp3 = scorepair.new(-3, 'm3')
print("sp3", sp3:getScore(), sp3:getMember());

local sp4 = scorepair.new(-4, 'm4')
print("sp4", sp4:getScore(), sp4:getMember());

local sp5 = scorepair.new(-5, 'm5')
print("sp5", sp5:getScore(), sp5:getMember());

local n, err =mm.ZAdd(mmsid, 'zkey', sp1, sp2, sp3, sp4, sp5)
if (err ~= nil) then
	print('ZAdd err',  err);
	return err
end
print("ZAdd n=", n)

--	ZRange(key []byte, start int, stop int) ([]dbpub.ScorePair, error)
local ret, err = mm.ZRange(mmsid, 'zkey', 2, -1)
if (err ~= nil) then
	print('ZRange err',  err);
	return err
end
print("loop value")
for k,v in pairs(ret) do
	print(k, v)
end

--[[

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
--]]