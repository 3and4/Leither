print("Test OpenByUrl")

--读取sid
sid =request['sid']
print("mimei sid", sid)

local mm  = require('mimei');
local url = "/mm/do-vfh8zWJenJEBpMvvwckrXH1t"
local mid = "do-vfh8zWJenJEBpMvvwckrXH1t"
local mmsid, err = mm.MMOpen(sid, mid, "last", "")
if (err ~= nil) then
	print('MMOpenByUrl err=',  err);
	return err
end

print('MMOpenByUrl ok mmsid =',  mmsid);

dirs, err = mm.MFReaddir(mmsid, 10)
if (err ~= nil) then
	print('MMOpen err=',  err);
	return err
end
for k,v in ipairs(dirs) do
        print(k,v)
end

print('MFReaddir ok');
