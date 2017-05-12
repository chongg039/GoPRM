package pe

import (
	"math/rand"
	"time"
)

// CreateRCB initialize resource
func CreateRCB(name string, t int) *RCB {
	return &RCB{
		Name:        name,
		Total:       t,
		Available:   t,
		BlockedList: Queue{},
	}
}

/*
CreatePCB means create a process
*/
func CreatePCB(name string, level int) *PCB {
	const (
		status string = "ready"
		cpu    string = "notused"
		memory string = "notused"
		// files  string = "notused"
	)

	pid := generateRandomPID()

	// var rs RequestResource

	// for i := 0; i < len(rname); i++ {
	// 	rs.Name = rname[i]
	// 	rs.OK = false
	// 	p.RequestResArr = append(p.RequestResArr, rs) // 添加请求数组应该在request函数中进行
	// }

	// append new created process to the ready queue
	// readyqueue = append(readyqueue, *p)
	return &PCB{
		Name:      name,
		PID:       pid,
		Status:    status,
		Priority:  level,
		CPUState:  cpu,
		Memory:    memory,
		ReqResArr: []RequestResource{},
	}
}

// InsertSortedQueue replace method "append"
func (p *PCB) InsertSortedQueue(pcbPool PCBPool) PCBPool {

	pcbPool[p.Priority] = append(pcbPool[p.Priority], *p)

	return pcbPool
}

// generateRandomPID gets the PID in the range of 2000~15000
func generateRandomPID() (pid int) {
	min := 2000
	max := 15000

	rand.Seed(time.Now().Unix())
	pid = rand.Intn(max-min) + min
	return
}
