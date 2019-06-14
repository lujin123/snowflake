// `snowflake`简单实现，用32位标识时间戳，8位标识机器，13位标识自增序列，
// 所以最多可以有256台机器，一台机器一秒最多可以生成8192个ID
// 总的长度只有53位，因为`js`整型最多可以标识53位不丢失精度，所以为了避免后端处理成字符串，这样简单的处理下

package snowflake

import (
	"sync"
	"time"
)

const (
	StartStamp = 1480166465631 //起始时间戳

	// 占用位数
	StampBit    = 32 //秒级时间戳位数
	MachineBit  = 8  //机器标识位数
	SequenceBit = 13 //自增序列位数

	// 左移位数
	MachineLeft = SequenceBit
	StampLeft   = MachineLeft + MachineBit

	// 最大值
	MaxMachine  = ^(int64(-1) << MachineBit)
	MaxSequence = ^(int64(-1) << SequenceBit)
)

type Snowflake struct {
	sync.Mutex
	sequence  int64
	machineId int64
	lastStamp int64
}

func New(machineId int64) *Snowflake {
	if machineId > MaxMachine {
		panic("Machine id greater than maximum.")
	}
	return &Snowflake{
		machineId: machineId,
		lastStamp: time.Now().Unix(),
	}
}

func (s *Snowflake) NextID() int64 {
	s.Lock()
	defer s.Unlock()

	nowStamp := s.getNowStamp()
	if nowStamp < s.lastStamp {
		panic("Clock moved backwards, Refusing to generate id")
	}
	if nowStamp == s.lastStamp {
		// 同一秒内
		s.sequence = (s.sequence + 1) & MaxSequence
		if s.sequence == 0 {
			// 序列用完，取下一秒
			nowStamp = s.getNextStamp()
		}
	} else {
		// 不是同一秒序列重置
		s.sequence = 0
	}

	s.lastStamp = nowStamp
	return (nowStamp-StartStamp)<<StampLeft | s.machineId<<MachineLeft | s.sequence
}

func (s *Snowflake) getNowStamp() int64 {
	return time.Now().Unix()
}

func (s *Snowflake) getNextStamp() int64 {
	stamp := s.getNowStamp()
	if stamp <= s.lastStamp {
		stamp = s.getNowStamp()
	}
	return stamp
}
