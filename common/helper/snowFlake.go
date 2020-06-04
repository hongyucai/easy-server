package helper

import (
	"log"
	"sync"
	"time"
)
/*
雪花算法
0 - 0000000000000000000000000000000000000 - 00000 00000 -000000000000
1   41									    5     5      12

1位，不用。二进制中最高位为 1 的都是负数，但是我们生成的 id 一般都使用整数，所以这个最高位固定是 0
41位，用来记录时间戳（毫秒）。
	如果只用来表示正整数（计算机中正数包含 0），可以表示的数值范围是：0 至 241−1，减 1 是因为可表示的数值范围是从 0 开始算的，而不是 1。
	也就是说 41 位可以表示 241−1 个毫秒的值，转化成单位年则是 (241−1)/(1000∗60∗60∗24∗365)=69 年
10位，用来记录工作机器 id。
	可以部署在 210=1024 个节点，包括 5 位 datacenterId 和 5 位 workerId,5 位（bit）可以表示的最大正整数是 25−1=31，即可以用 0、1、2、3、....31 这 32 个数字，来表示不同的 datecenterId 或 workerId
12位，序列号，用来记录同毫秒内产生的不同 id。
	12 位（bit）可以表示的最大正整数是 212−1=4095，即可以用 0、1、2、3、....4094 这 4095 个数字，来表示同一机器同一时间截（毫秒) 内产生的 4095 个 ID 序号
*/
const (
	workerIdBits int64  = 5
	datacenterIdBits int64 = 5
	sequenceBits int64 = 12

	maxWorkerId int64 = -1 ^ (-1 << uint64(workerIdBits))
	maxDatacenterId int64 = -1 ^ (-1 << uint64(datacenterIdBits))
	maxSequence int64 = -1 ^ (-1 << uint64(sequenceBits))

	timeLeft uint8 = 22
	dataLeft uint8 = 17
	workLeft uint8 = 12

	twepoch int64 = 1525705533000
)

type  worker struct {
	mu        sync.Mutex
	laststamp int64
	workerid int64
	datacenterid int64
	sequence int64
}

func(w *worker) getCurrentTime() int64 {
	return time.Now().UnixNano() / 1e6
}
//var i int = 1
func(w *worker) nextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	timestamp := w.getCurrentTime()
	if timestamp < w.laststamp {
		log.Fatal("can not generate id")
	}
	if w.laststamp == timestamp {
		// 这其实和 <==>
		// w.sequence++
		// if w.sequence++ > maxSequence  等价
		w.sequence = (w.sequence + 1) & maxSequence
		if w.sequence == 0 {
			// 之前使用 if, 只是没想到 GO 可以在一毫秒以内能生成到最大的 Sequence, 那样就会导致很多重复的
			// 这个地方使用 for 来等待下一毫秒
			for timestamp <= w.laststamp {
				//i++
				//fmt.Println(i)
				timestamp = w.getCurrentTime()
			}
		}
	} else {
		w.sequence = 0
	}
	w.laststamp = timestamp

	return ((timestamp - twepoch) << timeLeft) | (w.datacenterid << dataLeft)  | (w.workerid << workLeft) | w.sequence
}
func (w *worker) tilNextMillis() int64 {
	timestamp := w.getCurrentTime()
	if (timestamp <= w.laststamp) {
		timestamp = w.getCurrentTime()
	}
	return timestamp
}

func GenSnowFlakeId() int64{
	w := new(worker)
	// 上一次时间
	w.laststamp = -1
	w.workerid  = 1
	w.datacenterid = 1
	w.sequence = 14

	/*i := 0
	r := make([]int64, 0)
	for {
		id := w.nextId()
		fmt.Printf("id: %d\n",id )
		r = append(r, id)
		i++
		if  i > 10000000 {
			break
		}
	}*/

	return w.nextId()
}