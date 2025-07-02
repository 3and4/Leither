package rpc

type FVPair struct {
	Field string
	Value any
}

type ScorePair struct {
	Score  int64
	Member string
}

// 原来两处引用，现在只剩下一处了
func ScorePairs2Members(sps []ScorePair) []string {
	ret := make([]string, 0, len(sps))
	for _, sp := range sps {
		ret = append(ret, sp.Member)
	}
	return ret
}

type MDbStub struct {
	Begin            func(mmsid string, timeout int) error
	Commit           func(mmsid string) error
	Rollback         func(mmsid string) error
	GetLastSeq       func(mmsid string) (uint64, error)
	Set              func(mmsid, key string, value any) error
	Get              func(mmsid, key string) (any, error)
	Del              func(mmsid string, key ...string) (int64, error)
	Incr             func(mmsid, key string) (int64, error)
	IncrBy           func(mmsid, key string, delta int64) (int64, error)
	Strlen           func(mmsid, key string) (int64, error)
	Hmclear          func(mmsid string, key ...string) (int64, error)
	Hclear           func(mmsid string, key string) (int64, error)
	Hdel             func(mmsid, key string, field ...string) (int64, error)
	Hlen             func(mmsid, key string) (int64, error)
	Hset             func(mmsid, key, field string, value any) (int64, error)
	Hget             func(mmsid, key, field string) (any, error)
	Hmget            func(mmsid, key string, fields ...string) ([]any, error)
	Hmset            func(mmsid, key string, values ...FVPair) error
	Hgetall          func(mmsid, key string) ([]FVPair, error)
	Hkeys            func(mmsid, key string) ([]string, error)
	Hscan            func(mmsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)
	Hrevscan         func(mmsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)
	HincrBy          func(mmsid, key, field string, delta int64) (ret int64, err error)
	Lpush            func(mmsid, key string, value ...any) (int64, error)
	Lpop             func(mmsid, key string) (any, error)
	Rpush            func(mmsid, key string, value ...any) (int64, error)
	Rpop             func(mmsid, key string) (any, error)
	Lrange           func(mmsid, key string, start, stop int32) ([]any, error)
	Lclear           func(mmsid, key string) (int64, error)
	Lmclear          func(mmsid string, keys ...string) (int64, error)
	Lindex           func(mmsid, key string, index int32) (any, error)
	Llen             func(mmsid, key string) (int64, error)
	Lset             func(mmsid, key string, index int32, value any) error
	Zadd             func(mmsid, key string, args ...ScorePair) (int64, error)
	Zaddwithseq      func(mmsid, key string, members ...string) (ret int64, err error)
	Zcard            func(mmsid, key string) (int64, error)
	Zcount           func(mmsid, key string, mins, maxs int64) (int64, error)
	Zrem             func(mmsid, key string, members ...string) (int64, error)
	Zscore           func(mmsid, key, member string) (int64, error)
	Zrank            func(mmsid, key, member string) (int64, error)
	Zrange           func(mmsid, key string, start, stop int) ([]ScorePair, error)
	Zrangebyscore    func(mmsid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error)
	Zremrangebyscore func(mmsid, key string, mins, maxs int64) (int64, error)
	Zrevrange        func(mmsid, key string, start, stop int) (ret []ScorePair, err error)
	Zrevrangebyscore func(mmsid, key string, mins, maxs int64, offset int, count int) (ret []ScorePair, err error)
	Zmclear          func(mmsid string, key ...string) (int64, error)
	Zclear           func(mmsid, key string) (int64, error)
	ZincrBy          func(mmsid, key string, delta int64, member string) (ret int64, err error)
	Sadd             func(mmsid, key string, args ...string) (int64, error)
	Scard            func(mmsid, key string) (int64, error)
	Sclear           func(mmsid, key string) (int64, error)
	Sdiff            func(mmsid string, keys ...string) ([]string, error)
	Sinter           func(mmsid string, keys ...string) ([]string, error)
	Smclear          func(mmsid string, key ...string) (int64, error)
	Smembers         func(mmsid, key string) ([]string, error)
	Srem             func(mmsid, key string, m string) (int64, error)
	Sunion           func(mmsid string, keys ...string) ([]string, error)
	Scan             func(mmsid string, begin, match string, count int, inclusive bool, tp byte) (keys []string, err error)
	MMSyncData       func(sid, mid, tp string, mark []byte, count uint64, verseq ...uint64) ([]byte, error)
	CheckIntegrity   func(mmsid string, dataType uint8, key string, bRepair bool) error //校验数据
}
