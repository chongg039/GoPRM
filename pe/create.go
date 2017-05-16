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
		Parent:    p,
		Children:  nil,
	}
}

// AppendPCBEle replace method "append"
func (pcbPool *PCBPool) AppendPCBEle(p *PCB) {

	var pcbele *PCBEle
	pcbele = new(PCBEle)
	pcbele.Data = *p

	if pcbPool[p.Priority].Head == nil {
		pcbPool[p.Priority].Head = pcbele
	} else {
		var h = pcbPool[p.Priority].Head
		for h.Next != nil {
			h = h.Next
		}
		h.Next = pcbele
	}
	pcbPool[p.Priority].Length++

}

// RemovePCBEle remove PCB ele from PCBPool
func (pcbPool *PCBPool) RemovePCBEle(p *PCB) bool {
	//n为当前节点，h为前一个节点，初始状态为同一个节点
	// var h, n = pcbPool[p.Priority].Head, pcbPool[p.Priority].Head

	// if n == pcbPool[p.Priority].Head && n != nil && n.Data.PID == p.PID {
	// 	pcbPool[p.Priority].Head = pcbPool[p.Priority].Head.Next
	// 	pcbPool[p.Priority].Length--
	// 	return true
	// }
	// for n != nil {
	// 	if n.Data.Name == p.Name {
	// 		//由于有垃圾回收，所以不用考虑释放内存
	// 		h.Next = n.Next
	// 		pcbPool[p.Priority].Length--
	// 		return true
	// 	}
	// 	h = n
	// 	n = n.Next
	// }
	// return false
	if pcbPool[p.Priority].Head == nil {
		return true
	}
	n := pcbPool[p.Priority].Head
	for n.Next != nil {
		if n.Next.Data.Name == p.Name {
			n.Next = n.Next.Next
			pcbPool[p.Priority].Length--
			return true
		}
		n = n.Next
	}
	return false
}

// generateRandomPID gets the PID in the range of 2000~15000
func generateRandomPID() (pid int) {
	min := 2000
	max := 15000

	rand.Seed(time.Now().Unix())
	pid = rand.Intn(max-min) + min
	return
}
