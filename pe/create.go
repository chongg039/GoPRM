package pe

import (
	"math/rand"
	"time"
)

/*
CreateProcess means create a process
*/
func (p *Process) CreateProcess(name string, level int, rname ...string) Process {
	const (
		status string = "ready"
		cpu    string = "notused"
		memory string = "notused"
		// files  string = "notused"
	)

	pid := generateRandomPID()
	// level := generateRandomLevel()

	p.Name = name
	p.PID = pid
	p.Status = status
	p.Priority = level
	p.CPUState = cpu
	p.Memory = memory

	var rs RequestResource

	// if len(rname) == 0 {
	// 	rs.Name = "no resource needed"
	// 	rs.OK = true
	// 	p.RequestResArr = append(p.RequestResArr, rs)
	// } else {
	for i := 0; i < len(rname); i++ {
		rs.Name = rname[i]
		rs.OK = false
		p.RequestResArr = append(p.RequestResArr, rs)
	}
	// }

	// append new created process to the ready queue
	// readyqueue = append(readyqueue, *p)
	return *p
}

// InsertSortedQueue replace method "append"
func (p *Process) InsertSortedQueue(rqa QueuesArr) QueuesArr {

	rqa[p.Priority] = append(rqa[p.Priority], *p)

	return rqa

	// 这里想使用 key\value 将所有进程有序插入一个队列，失败
	// for k, v := range readyqueue {
	// 	r := readyqueue[:k+1]
	// 	l := readyqueue[k+1:]
	// 	if p.Priority >= v.Priority {
	// 		readyqueue = append(r, *p)
	// 		readyqueue = append(readyqueue, l...)
	// 		break
	// 	}
	// }

	// 下面这段代码想将所有进程有序放入一个队列，但是存在问题
	// var min, max, mid int
	// var fmid, sum float64
	// min = 0
	// max = len(readyqueue) - 1
	// sum = float64(min) + float64(max)
	// fmid = math.Ceil(sum / 2)
	// mid = int(fmid)
	// log.Println("fmid is :", fmid)
	// log.Println("mid is :", mid)

	// for min <= max {
	// 	mid = (min + max) / 2
	// 	if readyqueue[mid].Priority > p.Priority {
	// 		max = mid - 1
	// 		//continue
	// 	} else if readyqueue[mid].Priority < p.Priority {
	// 		min = mid + 1
	// 		//continue
	// 	} else {
	// 		r := readyqueue[:mid]
	// 		l := readyqueue[mid:]
	// 		readyqueue = append(r, *p)
	// 		readyqueue = append(readyqueue, l...)
	// 		return readyqueue
	// 	}
	// }
	// r := readyqueue[:min]
	// l := readyqueue[min:]
	// readyqueue = append(r, *p)
	// readyqueue = append(readyqueue, l...)
	// return readyqueue
}

// generateRandomPID gets the PID in the range of 2000~15000
func generateRandomPID() (pid int) {
	min := 2000
	max := 15000

	rand.Seed(time.Now().Unix())
	pid = rand.Intn(max-min) + min
	return
}

// generateRandomLevel gets the level in the range of 0-2(init, user, system)
// func generateRandomLevel() (level int) {
// 	min := 0
// 	max := 2

// 	rand.Seed(time.Now().Unix())
// 	level = rand.Intn(max-min) + min
// 	return
// }
