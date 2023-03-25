--读取sid
sid =request['sid']
print("sid=", sid)
local mm = require('mimei');

--创建弥媒
mid, err = mm.MMCreate(sid, '', '', 'testMiMei', 2, 0x07276705);
if (err ~= nil) then
	print('MMCreate err=',  err);
	return err
end
print("mid", mid)

mmsid, err = mm.MMOpen(sid, mid, 'cur');
if (err ~= nil) then
	print('MMOpen err=',  err);
	return err
end
print('MMOpen mmsid=', mmsid);

-- load image file
local inp = assert(io.open('1.jpg', 'rb'))
local img = inp:read("*a")
print(#img)

fsid, err = mm.MFOpenTempFile(sid);
print("fsid=", fsid)
err = mm.MFSetObject(fsid, img)
if (err ~= nil) then
	print('MMSetObject err=',  err);
	return err
end

macid, err = mm.MFTemp2MacFile(fsid, mid);
if (err ~= nil) then
	print('MFTemp2MacFile err=',  err);
	return err
end
print('MFTemp2MacFile macid=', macid);
local dmacid = macid
-- 
fsid, err = mm.MFOpenTempFile(sid);
err = mm.MFSetObject(fsid, macid);
print("fsid=", fsid, err)
macid, err = mm.MFTemp2MacFile(fsid, mid);
ret, err = mm.Hset(mmsid, "Test", macid, dmacid);
print("Hset ret=",ret, macid);

local sp = scorepair.new(1666868524589, macid)
print("sp", sp:getScore(), sp:getMember());

ret,err  = mm.Zadd(mmsid, "Test", sp)
print("Zadd ret", ret, err) 


err = mm.MMBackup(sid, mid, '')
if (err ~= nil) then
	print('MMBackup err=%v',  err);
	return err
end