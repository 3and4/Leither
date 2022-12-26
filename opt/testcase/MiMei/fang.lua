--读取sid
sid =request['sid']
print("mimei sid=%s", sid)


local mm = require('mimei');

print('create mimei');
mid, err = mm.MMCreate(sid, '', '', 'testMiMei', 2, 0x07276705);

if (err ~= nil) then
	print('MMCreate err=',  err);
	return err
end

--print("mid=%f", 1234)

log.Debug('MMCreate mid=%s', mid);

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
	log.Debug('MMOpen err=%v',  err);
	return err
end
log.Debug('MMOpen mmsid=%s', mmsid);

--[[
Set(mmsid, key, value) error
Get(mmsid, key) (value, error)
--]]
err = mm.Set(mmsid, 'key0', 'value0');
if (err ~= nil) then
	log.Debug('Set err=%v',  err);
	return err
end

value, err = mm.Get(mmsid, 'key0');
if (err ~= nil) then
	log.Debug('Set err=%v',  err);
	return err
end

log.Debug('Get key0 value=%s', value);

return mid;




local Fi = {}
Fi.__index = Fi
function Fi.new(name)
    local self = {}
    self.name=name
    setmetatable(self, Fi)
    return self
end
f=Fi.new('ergerefe')
--print(f:name)

mmsid = 'dc5bc5fd11423621bf011770f3acbead5a6cdbff'
ret, err = mm.HSet(mmsid, 'Test', 'E9BJr32Rn4kHG1VSJqVbjHE44ky',f);
if (err ~= nil) then
 print('Set err=',  err);
 return err
end
print ('Hset ret=', ret);