package common

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/lithammer/shortuuid"
)

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objType,
		shardID:    shardID,
	}
}

// Shard: 1, Object: 1, ID: 1 => 001 001 001
// 1 << 8 = 0001 0000 0000
// 1 << 4 =   	    1 0000
// 1 << 0 =              1
// => 0001 0001 0001

func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func genShortUUID() string {
	return shortuuid.New()
}
