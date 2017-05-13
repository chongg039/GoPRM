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
func (p *PCB) CreatePCB(name string, level int) *PCB {
	const (
		status string = "ready"
		cpu    string = "notused"
		memory string = "notused"
	)

	pid := generateRandomPID()

	return &PCB{
		Name:      name,
		PID:       pid,
		Status:    status,
		Priority:  level,
		CPUState:  cpu,
		Memory:    memory,
		ReqResArr: []RequestResource{},
		parent:    p,
		children:  make([]interface{}, 5),
	}
}

// InsertSortedQueue replace method "append"
func (p *PCB) InsertSortedQueue(pcbPool *PCBPool) *PCBPool {

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
