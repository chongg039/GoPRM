package pe

import (
	"math/rand"
	"time"
)

/*
CreateProcess means create a process
*/
func (p *PCB) CreateProcess(name string, level int, rname ...string) PCB {
	const (
		status string = "ready"
		cpu    string = "notused"
		memory string = "notused"
		// files  string = "notused"
	)

	pid := generateRandomPID()

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
func (p *PCB) InsertSortedQueue(pcbPool PCBPool) PCBPool {

	pcb[p.Priority] = append(pcb[p.Priority], *p)

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
